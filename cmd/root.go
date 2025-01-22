package cmd

import(
	
	"github.com/spf13/cobra"
	"github.com/austinwilson1296/cli_chat/internal/database"
)

const File string = "chatcli.db"
var DbQuery *database.Queries

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

