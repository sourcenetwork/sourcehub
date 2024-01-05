package pkg

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	cmtjson "github.com/cometbft/cometbft/libs/json"
	rpcclient "github.com/cometbft/cometbft/rpc/jsonrpc/client"
	cmttypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc"

	rpctypes "github.com/cometbft/cometbft/rpc/core/types"

	"github.com/sourcenetwork/sourcehub/app"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

var ErrPromiseCanceled = errors.New("promise canceled")

const defaultChanSize = 100
const websocketEndpoint = "/websocket"
const txEventsQuery = "tm.event='Tx'"
const txHashEvent = "tx.hash"

type Executed bool

func NewACPQueryClient(addr string) (types.QueryClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return types.NewQueryClient(conn), nil
}

// TxListener connects to a CometBFT node RPC serv
// and subscribes as a liestener to Txs
// Multiplexes Txs to different listeners - only one listener per tx
type TxListener struct {
	subscribers sync.Map
	defaultChan chan rpctypes.ResultEvent
	errChan     chan error
	client      *rpcclient.WSClient
}

func NewTxListenerFromClient(client *rpcclient.WSClient) (*TxListener, error) {
	ch := make(chan rpctypes.ResultEvent, defaultChanSize)
	return &TxListener{
		subscribers: sync.Map{},
		defaultChan: ch,
		client:      client,
	}, nil
}

// NewTxListener creates a WebSocket client for CometBFT rpc
func NewTxListener(addr string) (*TxListener, error) {
	client, err := rpcclient.NewWS(addr, websocketEndpoint)
	if err != nil {
		return nil, err
	}
	return NewTxListenerFromClient(client)
}

func processResult(result rpctypes.ResultEvent) (cmttypes.EventDataTx, error) {
	eventData := result.Data.(cmttypes.EventDataTx)
	if eventData.Result.Code != 0 {
		txHash := result.Events[txHashEvent][0]
		return cmttypes.EventDataTx{}, NewTxError(txHash,
			eventData.Result.Code,
			eventData.Result.Log,
			eventData.Result.Codespace)
	}
	return eventData, nil
}

type TxError struct {
	code      uint32
	txHash    string
	msg       string
	codespace string
}

func NewTxError(txHash string, code uint32, msg string, codespace string) *TxError {
	return &TxError{
		code:      code,
		msg:       msg,
		codespace: codespace,
		txHash:    txHash,
	}
}

func (e *TxError) Error() string {
	return fmt.Sprintf("tx %v: code %v codespace %v: %v", e.txHash, e.code, e.codespace, e.msg)
}

// Listen runs the listener in a thread and returns
// a default reponse channel which returns responses to unregistered txs
// If ctx is terminated, the listener stops and cleans up
func (l *TxListener) Listen(ctx context.Context) error {

	err := l.client.OnStart()
	if err != nil {
		return err
	}

	err = l.client.Subscribe(ctx, txEventsQuery)
	if err != nil {
		return err
	}

	listen := func() {
		log.Printf("TxListener waiting")
		for {
			select {
			case resp := <-l.client.ResponsesCh:
				if resp.Error != nil {
					l.errChan <- resp.Error
					log.Printf("TxListener err: %v", resp.Error)
					continue
				}

				event := rpctypes.ResultEvent{}
				err := cmtjson.Unmarshal(resp.Result, &event)
				if err != nil {
					log.Printf("TxListener err: %v", err)
					l.errChan <- err
					continue
				}

				log.Printf("TxListener Event received")

				txHashList, found := event.Events[txHashEvent]
				if !found {
					log.Printf("event with no hash")
					continue
				}
				txHash := txHashList[0]
				any, found := l.subscribers.LoadAndDelete(txHash)
				fulfiller := any.(Fulfiller[rpctypes.ResultEvent])
				if found {
					_, err := processResult(event)
					if err != nil {
						fulfiller.ProduceErr(err)
					} else {
						fulfiller.Produce(event)
					}
				} else {
					//FIXME thsi may block, maybe check if chan is full
					l.defaultChan <- event
				}

			case <-ctx.Done():
				l.client.Stop()
				l.subscribers.Range(func(key any, value any) bool {
					fulfiller := value.(Fulfiller[rpctypes.ResultEvent])
					fulfiller.Cancel()
					return true
				})
				return
			}
		}
	}

	go listen()

	return nil
}

// Subscribe configures the Listener to return the RPC Response
// when the response for txHash is received.
//
// Note: The listener does not logs responses.
// If the response for txHash is received before being registered
// the response will be sent on the default channel.
func (l *TxListener) Subscribe(txHash string) Promise[rpctypes.ResultEvent] {
	promise, fulfiller := NewPromise[rpctypes.ResultEvent](txHash)
	l.subscribers.Store(txHash, fulfiller)
	log.Printf("subscribed to tx: %v", txHash)
	return promise
}

func (l *TxListener) ErrChan() <-chan error {
	return l.errChan
}

func (l *TxListener) DefaultChan() <-chan rpctypes.ResultEvent {
	return l.defaultChan
}

// Fulfiller fullfils a Promise
type Fulfiller[T any] struct {
	resultChan chan<- T
	errChan    chan<- error
}

func NewFulfiller[T any](resultChan chan<- T, errChan chan<- error) Fulfiller[T] {
	return Fulfiller[T]{
		resultChan: resultChan,
		errChan:    errChan,
	}
}

// Produce produces the result value, sends it to the correct channel
// and marks it as fullfilled
func (f *Fulfiller[T]) Produce(value T) {
	f.resultChan <- value
	f.cleanup()
}

func (f *Fulfiller[T]) cleanup() {
	close(f.resultChan)
	close(f.errChan)
}

// ProduceErr produces an error as the Promise result and fulfills it
func (f *Fulfiller[T]) ProduceErr(err error) {
	f.errChan <- err
	f.cleanup()
}

// Cancel cancels a promise, meaning it wont produce an error
func (f *Fulfiller[T]) Cancel() {
	f.ProduceErr(ErrPromiseCanceled)
}

// Promise represents a value which will be produced
// The result will be either an error or a result.
// Callers should select between either one of the channels.
type Promise[T any] struct {
	resultChan <-chan T
	errChan    <-chan error
	fulfilled  bool
	txHash     string
	val        T
	err        error
}

// NewPromise creates new Promise
func NewPromise[T any](txHash string) (Promise[T], Fulfiller[T]) {
	resultChan := make(chan T, 1)
	errChan := make(chan error, 1)

	promise := Promise[T]{
		resultChan: resultChan,
		errChan:    errChan,
		fulfilled:  false,
		txHash:     txHash,
	}

	fulfiller := NewFulfiller(resultChan, errChan)

	return promise, fulfiller
}

// Returns Hash of the Tx being awaited
func (p *Promise[T]) GetTxHash() string {
	return p.txHash
}

func (p *Promise[T]) Terminated() bool {
	return p.fulfilled
}

// Error returns a channel which produces an error
func (p *Promise[T]) Error() <-chan error {
	return p.errChan
}

// Result returns a channel which produces the result
func (p *Promise[T]) Result() <-chan T {
	return p.resultChan
}

// GetResult returns the stored result from a Promise.
// If Promise hasn't been fulfilled return err
func (p *Promise[T]) GetResult() (T, error) {
	var zero T
	if !p.Terminated() {
		return zero, errors.New("Promise not terminated")
	}

	return p.val, p.err
}

// Await blocks untils the Promisse is resolved
func (p *Promise[T]) Await() (T, error) {
	var zero T
	for {
		select {
		case val, valid := <-p.Result():
			if !valid {
				continue
			}
			p.val = val
			p.fulfilled = true

			return val, nil
		case err, valid := <-p.Error():
			if !valid {
				continue
			}

			p.err = err
			p.fulfilled = true
			return zero, err
		}
	}
}

func MapPromise[T any, U any](promise Promise[T], mapper func(T) U) Promise[U] {
	pipe := NewPipe(mapper, promise.resultChan)

	return Promise[U]{
		resultChan: pipe.ReceiveEnd(),
		errChan:    promise.errChan,
		fulfilled:  promise.fulfilled,
		txHash:     promise.txHash,
	}
}

type PipeChan[T any, U any] struct {
	in     <-chan T
	out    chan U
	mapper func(T) U
}

func NewPipe[T any, U any](mapper func(T) U, producer <-chan T) PipeChan[T, U] {
	out := make(chan U, 1)
	pipe := PipeChan[T, U]{
		in:     producer,
		out:    out,
		mapper: mapper,
	}

	go pipe.run()

	return pipe
}

func (c *PipeChan[T, U]) ReceiveEnd() <-chan U {
	return c.out
}

// run
func (c *PipeChan[T, U]) run() {
	for {
		select {
		case t, ok := <-c.in:
			if !ok {
				close(c.out)
				return
			}
			u := c.mapper(t)
			c.out <- u
		}
	}
}

type ACPClient struct {
	client   cosmosClient
	listener *TxListener
}

func NewACPClient(nodeGRPCAddr string, listener *TxListener) (ACPClient, error) {
	client, err := newCosmosClient(nodeGRPCAddr)
	if err != nil {
		return ACPClient{}, err
	}

	return ACPClient{
		client:   client,
		listener: listener,
	}, nil
}

func (c *ACPClient) TxCreatePolicy(ctx context.Context, session Session, msg *types.MsgCreatePolicy) (*Promise[string], error) {
	resp, err := c.buildAndBroadcast(ctx, session, msg)
	if err != nil {
		return nil, err
	}

	promise := c.listener.Subscribe(resp.TxHash)
	promiseEx := MapPromise(promise, func(result rpctypes.ResultEvent) string {
		id := result.Events["sourcehub.acp.EventPolicyCreated.policy_id"][0]
		return strings.ReplaceAll(id, "\"", "")
	})
	return &promiseEx, nil
}

func (c *ACPClient) buildAndBroadcast(ctx context.Context, session Session, msg sdk.Msg) (*sdk.TxResponse, error) {

	addr, err := sdk.AccAddressFromBech32(session.Actor)
	if err != nil {
		return nil, err
	}

	log.Printf("broadcasting msg: %v", msg)

	tx, err := c.client.BuildTx(ctx, addr, session.PrivKey, msg)

	if err != nil {
		return nil, fmt.Errorf("buildtx err: %v", err)
	}

	resp, err := c.client.BroadcastTx(ctx, tx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ACPClient) TxRegisterObject(ctx context.Context, session Session, msg *types.MsgRegisterObject) (*Promise[Executed], error) {
	resp, err := c.buildAndBroadcast(ctx, session, msg)
	if err != nil {
		return nil, err
	}

	promise := c.listener.Subscribe(resp.TxHash)
	promiseEx := MapPromise(promise, func(_ rpctypes.ResultEvent) Executed {
		return true
	})
	return &promiseEx, nil
}

func (c *ACPClient) TxSetRelationship(ctx context.Context, session Session, msg *types.MsgSetRelationship) (*Promise[Executed], error) {
	resp, err := c.buildAndBroadcast(ctx, session, msg)
	if err != nil {
		return nil, err
	}

	promise := c.listener.Subscribe(resp.TxHash)
	promiseEx := MapPromise(promise, func(_ rpctypes.ResultEvent) Executed {
		return true
	})
	return &promiseEx, nil
}

func (c *ACPClient) TxDeleteRelationship(ctx context.Context, session Session, msg *types.MsgDeleteRelationship) (*Promise[Executed], error) {
	resp, err := c.buildAndBroadcast(ctx, session, msg)
	if err != nil {
		return nil, err
	}

	promise := c.listener.Subscribe(resp.TxHash)
	promiseEx := MapPromise(promise, func(_ rpctypes.ResultEvent) Executed {
		return true
	})
	return &promiseEx, nil
}

// TxCheckAccess builds builds a Tx and Broadcats a MsgCheckAccess
// Return a Promise containing the Decision ID
func (c *ACPClient) TxCheckAccess(ctx context.Context, session Session, msg *types.MsgCheckAccess) (*Promise[string], error) {
	resp, err := c.buildAndBroadcast(ctx, session, msg)
	if err != nil {
		return nil, err
	}

	promise := c.listener.Subscribe(resp.TxHash)
	promiseEx := MapPromise(promise, func(events rpctypes.ResultEvent) string {
		id := events.Events["sourcehub.acp.EventAccessDecisionCreated.decision_id"][0]
		return strings.ReplaceAll(id, "\"", "")
	})
	return &promiseEx, nil
}

// TODO add remaining txs methods

type cosmosClient struct {
	//encCfg     params.EncodingConfig
	txClient   txtypes.ServiceClient
	authClient authtypes.QueryClient
}

func newCosmosClient(grpcNodeAddr string) (cosmosClient, error) {
	//encCfg := app.MakeEncodingConfig()

	queryConn, err := grpc.Dial(grpcNodeAddr, grpc.WithInsecure())
	if err != nil {
		return cosmosClient{}, err
	}
	authClient := authtypes.NewQueryClient(queryConn)

	txConn, err := grpc.Dial(grpcNodeAddr, grpc.WithInsecure())
	if err != nil {
		return cosmosClient{}, err
	}
	txClient := txtypes.NewServiceClient(txConn)

	return cosmosClient{
		//encCfg:     encCfg,
		authClient: authClient,
		txClient:   txClient,
	}, nil
}

/*
func (c *cosmosClient) getKey(addr sdk.AccAddress) (cryptotypes.PrivKey, error) {
	record, err := c.keyring.KeyByAddress(addr)
	if err != nil {
		return nil, err
	}

	privBytes := record.GetLocal().PrivKey.Value
	// TODO use any unmarhsal
	priv := &secp256k1.PrivKey{}
	err = priv.Unmarshal(privBytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}
*/

func (c *cosmosClient) BroadcastTx(ctx context.Context, tx xauthsigning.Tx) (*sdk.TxResponse, error) {
	encoder := authtx.DefaultTxEncoder()
	txBytes, err := encoder(tx)
	if err != nil {
		return nil, fmt.Errorf("marshaling tx: %w", err)
	}

	grpcRes, err := c.txClient.BroadcastTx(
		ctx,
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		log.Printf("broadcasting err: %v", grpcRes)
		return nil, err
	}

	log.Printf("broadcast tx: %v", grpcRes)

	return grpcRes.TxResponse, nil
}

func (c *cosmosClient) BuildTx(ctx context.Context, addr sdk.AccAddress, privKey cryptotypes.PrivKey, msgs ...sdk.Msg) (xauthsigning.Tx, error) {
	registry := cdctypes.NewInterfaceRegistry()
	types.RegisterInterfaces(registry)
	cfg := authtx.NewTxConfig(
		codec.NewProtoCodec(registry),
		[]signing.SignMode{
			signing.SignMode_SIGN_MODE_DIRECT,
		},
	)

	txBuilder := cfg.NewTxBuilder()

	err := txBuilder.SetMsgs(msgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to set msgs: %v", err)
	}

	// FIXME
	txBuilder.SetGasLimit(200000)
	// TODO fetch coin name from somewhere
	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewInt64Coin(TokenName, 10)))

	acc, err := c.getAccount(ctx, addr.String())
	if err != nil {
		return nil, err
	}

	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	sigV2 := signing.SignatureV2{
		PubKey: privKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
			Signature: nil,
		},
		Sequence: acc.GetSequence(),
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, fmt.Errorf("set signatures: %v", err)
	}

	// TODO set chain id somewhere
	signerData := xauthsigning.SignerData{
		ChainID:       ChainID,
		AccountNumber: acc.GetAccountNumber(),
		Sequence:      acc.GetSequence(),
		PubKey:        privKey.PubKey(),
	}

	sigV2, err = tx.SignWithPrivKey(
		context.Background(),
		signing.SignMode_SIGN_MODE_DIRECT,
		signerData,
		txBuilder,
		privKey,
		cfg,
		acc.GetSequence())
	if err != nil {
		return nil, fmt.Errorf("sign tx: %v", err)
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, fmt.Errorf("set signatures: %v", err)
	}

	return txBuilder.GetTx(), nil

}

func (c *cosmosClient) getAccount(ctx context.Context, addr string) (*authtypes.BaseAccount, error) {
	msg := authtypes.QueryAccountRequest{
		Address: addr,
	}

	resp, err := c.authClient.Account(ctx, &msg)
	if err != nil {
		return nil, fmt.Errorf("fetching account: %v", err)
	}

	acc := authtypes.BaseAccount{}
	err = acc.Unmarshal(resp.Account.Value)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling account: %v", err)
	}

	return &acc, nil
}

// FIXME what's a better way of doing this for a lib?
func init() {
	// Set prefixes
	accountPubKeyPrefix := app.AccountAddressPrefix + "pub"
	validatorAddressPrefix := app.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := app.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := app.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := app.AccountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()
}
