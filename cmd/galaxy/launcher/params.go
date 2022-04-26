package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://8c7ab07ba269364af68bc957ad32a7e2e01b5f19bbd9c1da4eedab2c4c140b4d95dae799913775b809e6d51a7662af04f0ff9ca934fb88e5188848234b4e2a36@65.108.43.238:18888",
		"enode://8c7ab07ba269364af68bc957ad32a7e2e01b5f19bbd9c1da4eedab2c4c140b4d95dae799913775b809e6d51a7662af04f0ff9ca934fb88e5188848234b4e2a36@65.108.124.81:18888",
		"enode://8c7ab07ba269364af68bc957ad32a7e2e01b5f19bbd9c1da4eedab2c4c140b4d95dae799913775b809e6d51a7662af04f0ff9ca934fb88e5188848234b4e2a36@65.108.2.89:18888",

		"enode://8c7ab07ba269364af68bc957ad32a7e2e01b5f19bbd9c1da4eedab2c4c140b4d95dae799913775b809e6d51a7662af04f0ff9ca934fb88e5188848234b4e2a36@65.108.141.55:18888",
		"enode://8c7ab07ba269364af68bc957ad32a7e2e01b5f19bbd9c1da4eedab2c4c140b4d95dae799913775b809e6d51a7662af04f0ff9ca934fb88e5188848234b4e2a36@65.21.89.25:18888",
		"enode://8c7ab07ba269364af68bc957ad32a7e2e01b5f19bbd9c1da4eedab2c4c140b4d95dae799913775b809e6d51a7662af04f0ff9ca934fb88e5188848234b4e2a36@95.217.205.113:18888",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
