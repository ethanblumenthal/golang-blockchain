package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethanblumenthal/golang-blockchain/database"
	"github.com/ethanblumenthal/golang-blockchain/node"
	"github.com/spf13/cobra"
)
func runCmd() *cobra.Command {
    var runCmd = &cobra.Command{
        Use: "run",
        Short: "Launches the gochain node and its HTTP API.",
        Run: func(cmd *cobra.Command, args []string) {
            ip, _ := cmd.Flags().GetString(flagIP)
            port, _ := cmd.Flags().GetUint64(flagPort)
            miner, _ := cmd.Flags().GetString(flagMiner)
            bootstrapIp, _ := cmd.Flags().GetString(flagBootstrapIp)
			bootstrapPort, _ := cmd.Flags().GetUint64(flagBootstrapPort)
			bootstrapAcc, _ := cmd.Flags().GetString(flagBootstrapAcc)
            
            fmt.Println("Launching gochain node and its HTTP API...")

            bootstrap := node.NewPeerNode(bootstrapIp, bootstrapPort, true, database.NewAccount(bootstrapAcc), false)
            n := node.New(getDataDirFromCmd(cmd), ip, port, database.NewAccount(miner), bootstrap)

            err := n.Run(context.Background())
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
        },
    }
    addDefaultRequiredFlags(runCmd)
    runCmd.Flags().String(flagIP, node.DefaultIP, "exposed IP for communication with peers")
    runCmd.Flags().Uint64(flagPort, node.DefaultHTTPPort, "exposed HTTP port for communication with peers")
    runCmd.Flags().String(flagMiner, node.DefaultMiner, "miner account of this node to receive block rewards")
    runCmd.Flags().String(flagBootstrapIp, node.DefaultBootstrapIp, "default bootstrap server to interconnect peers")
	runCmd.Flags().Uint64(flagBootstrapPort, node.DefaultBootstrapPort, "default bootstrap server port to interconnect peers")
	runCmd.Flags().String(flagBootstrapAcc, node.DefaultBootstrapAcc, "default bootstrap account to interconnect peers")

    return runCmd
}