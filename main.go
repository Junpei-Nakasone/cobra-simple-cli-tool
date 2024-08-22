package main

import (
    "fmt"
    "os"
    "time"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "cobra-simple-cli",
    Short: "Cobra simple CLI tool",
    Long:  `This is just simple CLI tool for demo use.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello from Cobra Simple CLI!")
    },
}

var dayOfWeekCmd = &cobra.Command{
    Use:   "dayofweek",
    Short: "Prints the current day of the week",
    Long:  `This command prints the current day of the week.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Today is", time.Now().Weekday())
    },
}

func main() {
    rootCmd.AddCommand(dayOfWeekCmd)
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
