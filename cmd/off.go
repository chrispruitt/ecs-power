package cmd

import (
	lib "github.com/chrispruitt/ecs-power/lib"

	"github.com/spf13/cobra"
)

// Run Command ./pentaho-cli run
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Scale cluster down",
	Run: func(cmd *cobra.Command, args []string) {
		lib.PowerOff(clusterName)
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
	offCmd.PersistentFlags().StringVarP(&clusterName, "cluster", "c", "", "cluster name")
	offCmd.MarkPersistentFlagRequired("cluster")
}
