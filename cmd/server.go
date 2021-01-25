package cmd

import (
	"github.com/Sutheres/report-chaser/internal/sec"
	"github.com/Sutheres/report-chaser/service"
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

	s := sec.NewClient(
		"https://sec.gov",
	)

	_ = service.NewService(
		"", "",
		service.WithSEC(s),
	)

	defer func() {
		log.Println("shutting down...")
	}()

	log.Println("starting server...")

}