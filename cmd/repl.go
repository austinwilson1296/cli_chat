package cmd

import (
    "bufio"
    "fmt"
    "os"
    "strings"
	"os/exec"
    "github.com/spf13/cobra"
)

func StartRepl(rootCmd *cobra.Command) {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("\nWelcome to the Party 😀!\n\n Type 'exit' to quit.🚪")
    for {
        fmt.Print(">> ")
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            continue
        }

        input = strings.TrimSpace(input)
        if input == "exit" {
            fmt.Println("Goodbye!")
            break
        }

        // Split the input into command and arguments
        args := strings.Split(input, " ")

        // Find and execute the command
        foundCmd, _, err := rootCmd.Find(args)
        if err != nil || foundCmd == nil {
            fmt.Printf("Unknown command: %s\n", args[0])
            continue
        }

        // Execute the found command with the arguments
        rootCmd.SetArgs(args)
        if err := rootCmd.Execute(); err != nil {
            fmt.Printf("Command execution error: %v\n", err)
        }
    }
}
