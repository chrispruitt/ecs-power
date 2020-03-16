package cmd

import (
	"fmt"
	"os"

	lib "github.com/chrispruitt/ecs-power/lib"

	"github.com/spf13/cobra"
)

var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Scale ecs cluster up based on preset ssm parameter auto scale values",
	Run: func(cmd *cobra.Command, args []string) {
		err := lib.PowerOn(clusterName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("Done.")
		}
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
	onCmd.PersistentFlags().StringVarP(&clusterName, "cluster", "c", "", "cluster name")
	onCmd.MarkPersistentFlagRequired("cluster")
}
