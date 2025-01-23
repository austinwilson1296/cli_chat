package cmd

import (
	"fmt"
	"github.com/austinwilson1296/cli_chat/utils"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)



func init() {
	rootCmd.AddCommand(stateCmd)
	
	
}

var stateCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Current authenticated user",
	Run: func(cmd *cobra.Command, args []string) {
		user := utils.State.GetCurrentUser()
		fmt.Printf("You are :%s\n",user.Username)
	},
}