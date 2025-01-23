package cmd

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/austinwilson1296/cli_chat/utils"
	"fmt"
)



func init() {
	rootCmd.AddCommand(roomCmd)
	
	
}

var roomCmd = &cobra.Command{
	Use:   "room",
	Short: "commands to interact with chatrooms",
	Long: "primary command to interact with rooms. \n1.List Rooms('list') \n2.Join Room('join' <room_name>) \n3.Delete Room('delete' <room_name> ) \n4.Create Room('create' <room_name>)",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
        if !utils.State.IsLoggedIn() {
            fmt.Println("Error: You must be logged in to use room commands.")
            // Exit the command execution
            cmd.Help() // Show the help text for the command
            return
        }
	},
}