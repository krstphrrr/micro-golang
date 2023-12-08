package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var generatePermissions = &cobra.Command{
	Use: "test",
	Short: "Generate permissions for a user",
	Long: `Generate permissions for a user registered in a cognito pool`,
	Run: runPermission,

}
func init(){
	rootCmd.AddCommand(generatePermissions)
}

func runPermission(cmd *cobra.Command, args []string){
	// var
	fmt.Println("hey")
}