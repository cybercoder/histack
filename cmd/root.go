package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ik8s",
	Short: "ik8s Stack - Multi-command dispatcher",
}

func Execute() error {
	return RootCmd.Execute()
}
