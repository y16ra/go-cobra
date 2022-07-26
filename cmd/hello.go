/*
Copyright © 2022 y16ra

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var rFlag bool

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello command greets to ginven args.",
	Long:  `'hello' prints a greeting message to the given name in the arguments.`,
	Example: `
	go-cobra hello foo bar           // Hello, foo and bar!
	`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := "Hello, World!"
		if len(args) > 0 {
			msg = "Hello, " + strings.Join(args, " and ") + "!"
		}
		if rFlag {
			msg = reverseString(msg)
		}
		fmt.Println(msg)
	},
}

func reverseString(input string) string {
	var msgTmp string

	for _, v := range input {
		msgTmp = string(v) + msgTmp
	}

	return msgTmp
}

func init() {
	rootCmd.AddCommand(helloCmd)

	// Define flags
	helloCmd.Flags().BoolVarP(&rFlag, "reverse", "r", false, "Reverses the output.")
}
