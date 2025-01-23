package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/austinwilson1296/cli_chat/internal/database"
	"github.com/austinwilson1296/cli_chat/utils"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var roomName string
var roomDescription string

var roomCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new room",
	Run: func(cmd *cobra.Command, args []string) {
		roomName, _ := cmd.Flags().GetString("name")
		roomDescription, _ := cmd.Flags().GetString("description")
		ownerID := utils.State.GetCurrentUser()

		db, err := sql.Open("sqlite3", File)
		if err != nil {
			log.Fatalf("Unable to connect to the database: %v", err)
		}
		defer db.Close()

		// Initialize queries
		DbQuery = database.New(db)

		roomID := uuid.New()
		roomIdString := roomID.String()

		room := database.CreateRoomParams{
			RoomID: roomIdString,
			RoomName: roomName,
			Description: sql.NullString{String:roomDescription,Valid: true},
			RoomOwner: ownerID.UserID,
		}

		err = DbQuery.CreateRoom(context.Background(),room)
		if err != nil{
			log.Fatal("error creating room")
			return
		}

		fmt.Printf("Creating room: %s\nDescription: %s\n", roomName, roomDescription)
		
	},
}

func init() {
	roomCreateCmd.Flags().StringVarP(&roomName, "name", "n", "", "Name of the room (required)")
	roomCreateCmd.Flags().StringVarP(&roomDescription, "description", "d", "", "Description of the room")
	roomCreateCmd.MarkFlagRequired("name") // Make "name" a required flag
	roomCmd.AddCommand(roomCreateCmd)
}
