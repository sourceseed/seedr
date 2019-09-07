package seedr

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "seedr",
	Short: "Seedr helps you to get started quickly",
	Long: `Seedr helps you to get started quickly.
	Seedr can setup a project skeleton for you in seconds.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
