package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/austinwilson1296/cli_chat/internal/auth"
	"github.com/austinwilson1296/cli_chat/internal/database"
	"github.com/austinwilson1296/cli_chat/utils"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)



func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().String("u", "", "Username to register")
	loginCmd.Flags().String("p", "", "Password to register")
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Chat Cli app",
	Run: func(cmd *cobra.Command, args []string) {
		// Get flags
		username, _ := cmd.Flags().GetString("u")
		password, _ := cmd.Flags().GetString("p")

		// Validate inputs
		if username == "" || password == "" {
			fmt.Println("Error: Username and Password cannot be empty.")
			return
		}

		// Initialize database connection
		db, err := sql.Open("sqlite3", File)
		if err != nil {
			log.Fatalf("Unable to connect to the database: %v", err)
		}
		defer db.Close()

		// Initialize queries
		DbQuery = database.New(db)

		user,err := DbQuery.GetUser(context.Background(),username)
		if err != nil {
			fmt.Printf("unable to locate username with provided name: %s\n",username)
			os.Exit(1)
		}
		err = auth.CheckPasswordHash(password,user.Password)
		if err != nil{
			fmt.Printf("incorrect password for user: %s\n",username)
			os.Exit(1)
		}
		


		stateUser := utils.User{
			Username: username,
			UserID: user.ID,
		}

		fmt.Printf("User '%s' logged in successfully.\n\n", username)
		utils.State.SetUser(&stateUser)
		utils.State.UpdateLastActive()
		StartRepl(rootCmd)
	},
}