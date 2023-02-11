package todo

import (
	todo "todo/pkg"

	"github.com/spf13/cobra"
)

var f bool
var markCmd = &cobra.Command{
	Use:"mark",
	Aliases: []string{"m"},
	Short:"Mark a todo as complete",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string){
		todo.MarkToDo(args[0],f)
	},
}

func init(){
	markCmd.Flags().BoolVarP(&f,"false","f",false,"Marks a todo as incomplete")
	rootCmd.AddCommand(markCmd)
}