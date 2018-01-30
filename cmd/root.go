package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a CLI task manager",
}
