package cmd

import (
	lib "github.com/chrispruitt/ecs-power/lib"

	"github.com/spf13/cobra"
)

// Run Command ./pentaho-cli run
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Scale ecs cluster up based on preset ssm parameter auto scale values",
	Run: func(cmd *cobra.Command, args []string) {
		lib.PowerOn(clusterName)
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
	onCmd.PersistentFlags().StringVarP(&clusterName, "cluster", "c", "", "cluster name")
	onCmd.MarkPersistentFlagRequired("cluster")
}
