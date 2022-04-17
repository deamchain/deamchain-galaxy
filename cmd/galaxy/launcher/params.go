package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://d1ea8768f9adbd0858a0291d3bbeee874429d8034be3e30b12c57f49655cf4d859606bef66d5f9d56def8eba7a6eab0b25668395ce72e19c4841fcfa5e5afdaa@185.64.104.17:15060",
		"enode://e8514c480b21d722ba99a74e290bb69a7b333ce8cac86f99f8897e678190b7f445fea0925bfac232cf31d5d5115d30f2cf8a06abfc728119fff6d794d129d223@185.25.50.199:15060",
		"enode://af9ad92f3004d220c2271b3cbf7e4095c08e07e6ad9d5edb725efbae9e4069a027ea0ed53343f5d280adcc43d8b92f30620d66f7b6d9b0df0a767df7ea34b3a2@185.25.50.202:15060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
