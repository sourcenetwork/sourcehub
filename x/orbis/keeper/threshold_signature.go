package keeper

import (
	"encoding/hex"
	"fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
	decaf377 "github.com/mizufinance/decaf377-go"
	"github.com/mizufinance/decaf377-go/orbisfrost"
	blst "github.com/supranational/blst/bindings/go"

	"github.com/sourcenetwork/vera/x/orbis/types"
)

const (
	// ThresholdSignatureSchemeBLS12381G1PKG2SigAugV1 is the public-key-augmented
	// BLS scheme: the ring public key is prepended to the message before
	// hash-to-curve (IETF "message augmentation"), so a signature cannot be
	// scaled from one publicly-related derived key onto another. Must stay in
	// lockstep with orbis-rs `crypto::THRESHOLD_SIGNATURE_SCHEME`. The
	// pre-augmentation scheme string was "bls12_381_g1_pk_g2_sig_nul".
	ThresholdSignatureSchemeBLS12381G1PKG2SigAugV1 = "bls12_381_g1_pk_g2_sig_aug_v1"
	ThresholdSignatureSchemeDecaf377FROST          = "decaf377_frost"

	bls12381PublicKeySize = 48
	bls12381SignatureSize = 96
	// AUG_ ciphersuite DST (matches orbis-rs `bls12_381::sign::BLS_SIG_DOMAIN`).
	bls12381G2SignatureAugDST = "BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_AUG_"
	decaf377PublicKeySize     = decaf377.ElementSize
	decaf377SignatureSize     = decaf377.ElementSize + decaf377.ScalarSize
)

func verifyThresholdSignatureForRingUpdate(ring *types.Ring, message []byte, scheme string, signature []byte) error {
	return verifyThresholdSignature(scheme, ring.RingPk, message, signature)
}

func verifyThresholdSignature(scheme string, ringPK string, message []byte, signature []byte) error {
	switch normalizeThresholdSignatureScheme(scheme) {
	case ThresholdSignatureSchemeBLS12381G1PKG2SigAugV1:
		return verifyBLS12381ThresholdSignature(ringPK, message, signature)
	case ThresholdSignatureSchemeDecaf377FROST:
		return verifyDecaf377FROSTThresholdSignature(ringPK, message, signature)
	default:
		return errorsmod.Wrapf(types.ErrInvalidThresholdSignature, "unsupported threshold signature scheme %q", scheme)
	}
}

func normalizeThresholdSignatureScheme(scheme string) string {
	return strings.ToLower(strings.TrimSpace(scheme))
}

func verifyBLS12381ThresholdSignature(ringPK string, message []byte, signature []byte) error {
	publicKey, err := decodeHexBytes(ringPK)
	if err != nil {
		return errorsmod.Wrapf(types.ErrInvalidThresholdSignature, "invalid bls12_381 public key encoding: %s", err)
	}
	if len(publicKey) != bls12381PublicKeySize {
		return errorsmod.Wrapf(types.ErrInvalidThresholdSignature, "invalid bls12_381 public key length %d", len(publicKey))
	}
	if len(signature) != bls12381SignatureSize {
		return errorsmod.Wrapf(types.ErrInvalidThresholdSignature, "invalid bls12_381 signature length %d", len(signature))
	}

	// Augmented BLS: verify against H(ringPK || message). The two trailing
	// bools are (useHash, usePksAsAugs) — usePksAsAugs makes blst prepend the
	// exact `publicKey` bytes to the message, matching orbis-rs which prepends
	// the compressed ring public key before hash-to-curve.
	if !new(blst.P2Affine).VerifyCompressed(
		signature,
		true,
		publicKey,
		true,
		message,
		[]byte(bls12381G2SignatureAugDST),
		true,
		true,
	) {
		return types.ErrInvalidThresholdSignature
	}

	return nil
}

func verifyDecaf377FROSTThresholdSignature(ringPK string, message []byte, signature []byte) error {
	publicKey, err := decodeHexBytes(ringPK)
	if err != nil {
		return errorsmod.Wrapf(types.ErrInvalidThresholdSignature, "invalid decaf377 public key encoding: %s", err)
	}
	if len(publicKey) != decaf377PublicKeySize {
		return errorsmod.Wrapf(types.ErrInvalidThresholdSignature, "invalid decaf377 public key length %d", len(publicKey))
	}
	if len(signature) != decaf377SignatureSize {
		return errorsmod.Wrapf(types.ErrInvalidThresholdSignature, "invalid decaf377 signature length %d", len(signature))
	}

	point, err := decaf377.Decode(publicKey)
	if err != nil {
		return errorsmod.Wrapf(types.ErrInvalidThresholdSignature, "invalid decaf377 public key: %s", err)
	}
	if decaf377.Equivalent(point, decaf377.Identity()) {
		return errorsmod.Wrap(types.ErrInvalidThresholdSignature, "decaf377 public key is the identity")
	}

	ok, err := orbisfrost.Verify(publicKey, message, signature)
	if err != nil {
		return errorsmod.Wrapf(types.ErrInvalidThresholdSignature, "invalid decaf377 signature: %s", err)
	}
	if !ok {
		return types.ErrInvalidThresholdSignature
	}

	return nil
}

// rejectIdentityRingPublicKey prevents finalizing a ring with the one public
// key that makes Schnorr verification vacuous. FinalizeRing predates an
// explicit curve/scheme field and its tests use opaque placeholder keys, so
// other encoding validation remains in the scheme-specific consumers.
func rejectIdentityRingPublicKey(ringPK string) error {
	publicKey, err := decodeHexBytes(ringPK)
	if err != nil {
		return nil
	}

	switch len(publicKey) {
	case decaf377PublicKeySize:
		point, err := decaf377.Decode(publicKey)
		if err == nil && decaf377.Equivalent(point, decaf377.Identity()) {
			return errorsmod.Wrap(types.ErrInvalidRing, "decaf377 ring public key is the identity")
		}
	case bls12381PublicKeySize:
		point := new(blst.P1Affine).Uncompress(publicKey)
		if point != nil && !point.KeyValidate() {
			return errorsmod.Wrap(types.ErrInvalidRing, "bls12_381 ring public key is the identity or outside the prime-order subgroup")
		}
	}

	return nil
}

func decodeHexBytes(value string) ([]byte, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil, fmt.Errorf("empty value")
	}

	decoded, err := hex.DecodeString(value)
	if err != nil {
		return nil, fmt.Errorf("expected plain hex: %w", err)
	}

	return decoded, nil
}
