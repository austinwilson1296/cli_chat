package cmd

import (
  "fmt"
  "database/sql"
  "github.com/austinwilson1296/cli_chat/internal/auth"
	"github.com/austinwilson1296/cli_chat/internal/database"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(addUserCmd)
  addUserCmd.PersistentFlags().String("username","","username to register")
  addUserCmd.PersistentFlags().String("password","","password to register")
}

var addUserCmd = &cobra.Command{
  Use:   "Create User",
  Short: "Register a user for CLI Chat App",
  Run: func(cmd *cobra.Command, args []string) {
	
    username,_ := cmd.Flags().GetString("username")
	password,_ := cmd.Flags().GetString("password")

	if username == "" || password == ""{
		fmt.Printf("Username and Password cannot be empty")
		return
	}

	userID := auth.GenerateRandomId(10)

	user := database.CreateUserParams{
		ID: int64(userID),
		Username: sql.NullString{String: username,Valid: true},
		Password: sql.NullString{String: password,Valid: true},
	}

	fmt.Printf("%v",user)
  },


}