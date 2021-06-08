package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/go-api/artisan/crud"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "create crud api from model.json",
	Long:  `create crud api from model.json`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("model name not null")
			return
		}
		app := &crud.Command{
			Name: args[0],
		}
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
