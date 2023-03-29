package config

import (
	"os"

	"github.com/benbjohnson/clock"
	"github.com/joho/godotenv"
)

type GlobalConfigInfo struct {
	Clock    clock.Clock
	Testnet  bool
	Port     string
	Eth_Node string
	Bsc_Node string
}

var GlobalConfig = &GlobalConfigInfo{}

func (g *GlobalConfigInfo) Init() {
	g.Clock = clock.New()
	envFile := os.Getenv("ENV")
	if envFile != "" {
		godotenv.Load(envFile)
	} else {
		godotenv.Load(".env")
	}
	g.Port = os.Getenv("PORT")
	if g.Port == "" {
		g.Port = "8080"
	}
	g.Testnet = os.Getenv("TESTNET") == "true"
	g.Eth_Node = os.Getenv("ETH_NODE")
	g.Bsc_Node = os.Getenv("BSC_NODE")
}
