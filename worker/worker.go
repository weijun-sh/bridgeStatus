package worker

import (
	//"time"
	"github.com/weijun-sh/bridgeStatus/rpc"
	"github.com/weijun-sh/bridgeStatus/tokens"
)

// StartWork start get balance
func StartWork() {
	rpc.InitClient()

	tokens.GetBalanceOfToken()
}
