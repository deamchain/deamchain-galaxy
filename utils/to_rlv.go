package utils

import "math/big"

// ToDeam number of DEAM to Wei
func ToDeam(deam uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(deam), big.NewInt(1e18))
}
