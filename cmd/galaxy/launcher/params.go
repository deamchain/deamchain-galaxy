package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// testnet
		"enode://48d4a5e5cce3891f92a78b75aced365460de5d8a547c945bfbb31f7b9bd16e68ccfcc61d8fb58f33fa638e9d06ec769998fde8fb119dda307be7465a79a55913@65.108.43.238:18888",
		"enode://2843ce606c5a350a093b40b4824ce1dfec932b05976995b2e370681e751250e10f149cbf1c4ce65fb4153b227ca62246b656778fc7b142cbafea1735ef68ea28@65.108.124.81:18888",
		"enode://c9fabacf90ebffebaa88729ab02efc8d18df8fcfd60764044a3e93a1370241495436a82c4f2b75c6dea6f7c5b9e1eb668a9b37f54a9e5f674059e94daaaa6025@65.108.2.89:18888",
		// mainnet
		"enode://11980fc6fadfe551d33de498f67aeb1b58a0994003e2c7d5a180b812ae061f1dbfdac934c073960d243141ceb7eb783af3372d3142042537a271925a6d598766@65.108.141.55:18888",
		"enode://8fe1ab5291a0042da486da112bec79085c8828da26972a5b9eb38e66eff17a8f6723cce6595d545e88d76b41db58872755ab1abe9e51971a2200e2fb97fa6d8e@65.21.89.25:18888",
		"enode://eb251b62eb9c4d3ca04258516cdfb44da94ccc4b8495ce2c81a7fd4d815fd0513e14e85276f0fbc6bd3da0de0780245ce8b0fc317a0adebcc44e97d1ee7b30b7@95.217.205.113:18888",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
