package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"github.com/austinwilson1296/cli_chat/internal/auth"
	"github.com/austinwilson1296/cli_chat/internal/database"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)



func init() {
	rootCmd.AddCommand(addUserCmd)
	addUserCmd.Flags().String("u", "", "Username to register")
	addUserCmd.Flags().String("p", "", "Password to register")
}

var addUserCmd = &cobra.Command{
	Use:   "add-user",
	Short: "Register a user for CLI Chat App",
	Run: func(cmd *cobra.Command, args []string) {
		// Get flags
		username, _ := cmd.Flags().GetString("u")
		password, _ := cmd.Flags().GetString("p")

		// Validate inputs
		if username == "" || password == "" {
			fmt.Println("Error: Username and Password cannot be empty.")
			return
		}

		hashPass,err := auth.HashPassword(password)
		if err != nil{
			fmt.Println("unable to hash password")
		}

		userID := uuid.New()
		// Initialize database connection
		db, err := sql.Open("sqlite3", File)
		if err != nil {
			log.Fatalf("Unable to connect to the database: %v", err)
		}
		defer db.Close()

		// Initialize queries
		DbQuery = database.New(db)

		// Create user
		user := database.CreateUserParams{
			ID: userID,
			Username: username,
			Password: hashPass,
		}

		err = DbQuery.CreateUser(context.Background(), user)
		if err != nil {
			log.Fatalf("Unable to create user: %v", err)
		}

		fmt.Printf("User '%s' created successfully.\n", username)
	},
}
