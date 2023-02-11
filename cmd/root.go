package todo

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use: "todo",
	Version: version,
	Short: "todo - A simple todo CLI tool",
	Long: "todo - A simple todo CLI tool",
	Run: func(cmd *cobra.Command, args []string){},
}

func Execute(){
	if err := rootCmd.Execute(); err !=nil{
		fmt.Fprintf(os.Stderr,"Something went wrong! - '%s'",err)
		os.Exit(1)
	}
}
