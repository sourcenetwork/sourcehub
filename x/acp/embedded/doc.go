// package embedded exposes a minimal ACP module without consensus
//
// The goal of this package is to provide a minimal implementation
// of the ACP module which can be used indepedent of a cosmos-sdk
// consensus engine.
// This useful for testing purposes only as it wouldn't be
// part of any deployment or real chain.
//
// Example usage:
// ```go
// acp, _ := NewLocalACP()
// ctx := acp.GetCtx()
// msgServer := acp.GetMsgServer()
// resp, err := msgServer.CreatePolicy(ctx, ...)
// ...
// /
// ````
package embedded
