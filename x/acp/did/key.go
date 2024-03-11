package did

import (
	"context"

	"github.com/cyware/ssi-sdk/did/key"
)

var _ Resolver = (*KeyResolver)(nil)

type KeyResolver struct {
}

func (r *KeyResolver) Resolve(ctx context.Context, did string) (DIDDocument, error) {
	didkey := key.DIDKey(did)
	doc, err := didkey.Expand()
	if err != nil {
		return DIDDocument{}, err
	}

	return DIDDocument{
		Document: doc,
	}, nil
}
