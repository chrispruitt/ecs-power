package cmd

import (
	"fmt"
	"os"

	lib "github.com/chrispruitt/ecs-power/lib"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the current status of the ecs-cluster.",
	Run: func(cmd *cobra.Command, args []string) {
		groups, err := lib.Status(clusterName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			for _, group := range groups {
				fmt.Println(fmt.Sprintf("Group: %s - Desired: %v - Min: %v - Max: %v", *group.AutoScalingGroupName, *group.DesiredCapacity, *group.MinSize, *group.MaxSize))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.PersistentFlags().StringVarP(&clusterName, "cluster", "c", "", "cluster name")
}
