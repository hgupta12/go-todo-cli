package todo

import ( 
	"github.com/spf13/cobra"
	"todo/pkg"
)

var addCmd = &cobra.Command{
	Use: "add",
	Aliases: []string{"a"},
	Short: "Add a todo",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string){
		todo.AddToDo(args)
	},
}

func init(){
	rootCmd.AddCommand(addCmd)
}