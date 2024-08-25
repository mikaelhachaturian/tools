package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decodeString)
}

var decodeString = &cobra.Command{
	Use:     "decode",
	Short:   "Decode base64 to string.",
	Aliases: []string{"d"},
	Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		decoded, err := base64.StdEncoding.DecodeString(args[0])
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(string(decoded))
	},
}
