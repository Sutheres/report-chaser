package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(workerCmd)
}

var workerCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts Report Chaser server",
	Run:   startServer,
}

func startServer(cmd *cobra.Command, args []string) {

	defer func() {
		log.Println("shutting down...")
	}()

	log.Println("starting server...")

}