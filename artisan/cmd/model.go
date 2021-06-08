package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyan94zl/go-api/artisan/model"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "model code to json",
	Long: `model code to json`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			fmt.Println("model name not null")
			return
		}
		model := &model.Command{
			Name: args[0],
		}
		model.Run()

	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
