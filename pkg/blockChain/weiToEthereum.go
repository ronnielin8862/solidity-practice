package blockChain

import (
	"math/big"
)

func WeiToEther(val *big.Int) *big.Float {

	x, y := new(big.Float).SetInt(val), big.NewFloat(1000000000000000000)
	return new(big.Float).Quo(x, y)
}