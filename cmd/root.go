package cmd


import (
	"log"

	"github.com/spf13/cobra"
)

var (
	BuildDate  = "None"
	CommitHash = "None"
)

var rootCmd = &cobra.Command{
	Use: "Report Chaser Service",
}

// Execute : runs registered cobra commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
