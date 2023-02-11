package todo

import (
	"github.com/spf13/cobra"
	"todo/pkg"
)

var all bool
var listCmd = &cobra.Command{
	Use: "list",
	Aliases: []string{"l"},
	Short: "Show all todos",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string){
		todo.ShowToDoList(all)
	},
}

func init(){
	listCmd.Flags().BoolVarP(&all,"all","a",false,"Return all todos")
	rootCmd.AddCommand(listCmd)
}