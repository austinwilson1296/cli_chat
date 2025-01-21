package cmd

import(
	
	"github.com/spf13/cobra"
	
)

var (
	rootCmd = &cobra.Command{
		Use: "CLI Chat App",
		Short: "A downloadable cli chat app",
		Long: "Allows you and your friends to communicate privately from your computer!",
		}
)

func Execute(){
	cobra.CheckErr(rootCmd.Execute())
}

