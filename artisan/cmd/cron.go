package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/go-api/artisan/command"
)

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "create cron code",
	Long: `create cron code`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("cron name not null")
			return
		}
		app := &command.Command{
			Name: args[0],
		}
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(cronCmd)
}
