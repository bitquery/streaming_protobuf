package util

import "math/big"

var (
	Big1 = big.NewInt(1)
)

func BytesToBigInt(bytes []byte, signed bool) *big.Int {
	bits := len(bytes) * 8
	ret := new(big.Int).SetBytes(bytes)
	if signed && ret.Bit(bits-1) == 1 {
		ret.Sub(ret, new(big.Int).Lsh(Big1, uint(bits)))
	}
	return ret
}
