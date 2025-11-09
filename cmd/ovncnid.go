package cmd

import (
	ovncnid "github.com/cybercoder/histack/pkg/daemon/ovn-cni-server"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ovncnidCMD)
}

var ovncnidCMD = &cobra.Command{
	Use:   "ovncnid",
	Short: "OVN CNI Daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := ovncnid.Start("/var/run/histack-ovn-cni.sock"); err != nil {
			// klog.Fatalf("Error on starting ovn cni daemon: %v", err)
			return err
		}
		return nil
	},
}
