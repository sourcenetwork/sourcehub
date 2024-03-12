package policy_cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/go-jose/go-jose/v3"

	"github.com/sourcenetwork/sourcehub/x/acp/did"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func newJWSVerifier(resolver did.Resolver) jwsVerifier {
	return jwsVerifier{
		resolver: resolver,
	}
}

// jwsVerifier verifies the Signature of a JWS which contains a PolicyCmd
type jwsVerifier struct {
	resolver did.Resolver
}

// Verify verifies the integrity of the JWS payload, returns the Payload if OK
//
// The verification extracts a VerificationMethod from the resolved Actor DID in the PolicyCmd.
// The JOSE header attributes are ignored and only the key derived from the Actor DID is accepted.
// This is done to assure no impersonation happens by thinkering the JOSE header in order to produce a valid
// JWS, signed by key different than that of the DID owner.
func (s *jwsVerifier) Verify(ctx context.Context, jwsStr string) (*types.PolicyCmdPayload, error) {
	jws, err := jose.ParseSigned(jwsStr)
	if err != nil {
		return nil, fmt.Errorf("failed parsing jws: %v", err)
	}

	payloadBytes := jws.UnsafePayloadWithoutVerification()
	payload := &types.PolicyCmdPayload{}
	err = jsonpb.UnmarshalString(string(payloadBytes), payload)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshaling PolcyCmd payload: %v", err)
	}

	did := payload.Actor
	doc, err := s.resolver.Resolve(ctx, did)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve actor did: %v", err)
	}
	// TODO this should technically be Authentication
	if len(doc.VerificationMethod) == 0 {
		return nil, fmt.Errorf("resolved actor did does not contain any verification methods")
	}

	method := doc.VerificationMethod[0]
	jwkRaw, err := json.Marshal(method.PublicKeyJWK)
	if err != nil {
		return nil, fmt.Errorf("error verifying signature: jwk marshal: %v", err)
	}

	jwk := jose.JSONWebKey{}
	err = jwk.UnmarshalJSON(jwkRaw)
	if err != nil {
		return nil, fmt.Errorf("error verifying signature: jwk unmarshal: %v", err)
	}

	_, err = jws.Verify(jwk)
	if err != nil {
		return nil, fmt.Errorf("could not verify actor signature for jwk: %v", err)
	}

	return payload, nil
}
