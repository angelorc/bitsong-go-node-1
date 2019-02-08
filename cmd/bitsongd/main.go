package main

import (
	"fmt"
	"github.com/BitSongOfficial/bitsong-go-node/api"
	"github.com/BitSongOfficial/bitsong-go-node/cmd/utils"
	"github.com/BitSongOfficial/bitsong-go-node/config"
	"github.com/BitSongOfficial/bitsong-go-node/core/bitsong"
	"github.com/BitSongOfficial/bitsong-go-node/genesis"
	"github.com/BitSongOfficial/bitsong-go-node/gui"
	"github.com/BitSongOfficial/bitsong-go-node/log"
	"github.com/tendermint/tendermint/libs/common"
	tmNode "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
	rpc "github.com/tendermint/tendermint/rpc/client"
	"os"
)

var cfg = config.GetConfig()

func main() {
	err := common.EnsureDir(utils.GetBitsongHome()+"/config", 0777)

	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	app := bitsong.NewBitsongBlockchain()
	node := startTendermintNode(app)

	client := rpc.NewLocal(node)
	status, _ := client.Status()
	if status.NodeInfo.Network != genesis.Network {
		log.Error("Different networks")
		os.Exit(1)
	}

	app.SetTmNode(node)

	if !cfg.ValidatorMode {
		go api.RunApi(app, client)
		go gui.Run(cfg.GUIListenAddress)
	}

	// Wait forever
	common.TrapSignal(func() {
		// Cleanup
		node.Stop()
		app.Stop()
	})
}

func startTendermintNode(app *bitsong.Blockchain) *tmNode.Node {
	cfg := config.GetTmConfig()
	nodeKey, err := p2p.LoadOrGenNodeKey(cfg.NodeKeyFile())

	if err != nil {
		panic(err)
	}

	node, err := tmNode.NewNode(
		cfg,
		privval.LoadOrGenFilePV(cfg.PrivValidatorKeyFile(), cfg.PrivValidatorStateFile()),
		nodeKey,
		proxy.NewLocalClientCreator(app),
		genesis.GetTestnetGenesis,
		tmNode.DefaultDBProvider,
		tmNode.DefaultMetricsProvider(cfg.Instrumentation),
		log.With("module", "tendermint"),
	)

	if err != nil {
		log.Error(fmt.Sprintf("Failed to create a node: %v", err))
		os.Exit(1)
	}

	if err = node.Start(); err != nil {
		log.Error(fmt.Sprintf("Failed to start node: %v", err))
		os.Exit(1)
	}

	log.Info("Started node", "nodeInfo", node.Switch().NodeInfo())

	return node
}
