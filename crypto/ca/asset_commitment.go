package ca

import "chain/crypto/ed25519/ecmath"

// AssetCommitment is a point pair representing an ElGamal commitment to an AssetPoint.
type AssetCommitment PointPair

// Create an asset commitment. Nil aek means make it nonblinded (and
// the returned scalar blinding factor is nil).
func CreateAssetCommitment(assetID AssetID, aek AssetKey) (*AssetCommitment, *ecmath.Scalar) {
	A := pointHash([]byte("AssetID"), assetID[:])
	if aek == nil {
		return &AssetCommitment{A, ecmath.ZeroPoint}, nil
	}
	c := scalarHash([]byte("AC.c"), assetID[:], aek)
	var H, C ecmath.Point
	H.ScMulAdd(&A, &ecmath.One, &c)
	C.ScMul(&J, &c)
	return &AssetCommitment{H, C}, &c
}

func (ac *AssetCommitment) Bytes() []byte {
	return (*PointPair)(ac).Bytes()
}
