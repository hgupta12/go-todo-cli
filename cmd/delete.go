package todo

import (
	todo "todo/pkg"

	"github.com/spf13/cobra"
)

var dltCmd = &cobra.Command{
	Use: "delete",
	Aliases: []string{"d"},
	Short: "Delete the todo with the given id",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string){
		todo.DeleteToDo(args[0])
	},
}

func init(){
	rootCmd.AddCommand(dltCmd)
}