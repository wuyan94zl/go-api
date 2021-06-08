package cmd

import (
	"fmt"
	"github.com/wuyan94zl/go-api/artisan/queue"

	"github.com/spf13/cobra"
)

// queueCmd represents the queue command
var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "create queue code",
	Long:  `create queue code`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("model name not null")
			return
		}
		app := &queue.Command{
			Name: args[0],
		}
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(queueCmd)
}
