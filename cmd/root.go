
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ldc",
	Short: "CLI application to admin AWS Cognito Userpool groups",
	Long: `CLI application to admin AWS Cognito Userpool groups. For example:

	Example usage: 
	ldc list 
	ldc groups
	ldc modify -u d107d39c-df84-4c5f-98ba-97a778507df3 -g NWERN -i=true (short-hand flags)
	ldc modify --username d107d39c-df84-4c5f-98ba-97a778507df3 --group NWERN --include=true (long-form flags)`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cog-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


