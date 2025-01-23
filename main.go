package main

import(
	"fmt"
	"github.com/austinwilson1296/cli_chat/cmd"
	"github.com/austinwilson1296/cli_chat/utils"

)

func main(){
	fmt.Println("Starting CLI Application!")
	utils.State = &utils.UserState{}

	cmd.Execute()
}