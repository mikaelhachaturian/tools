package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encodeString)
}

var encodeString = &cobra.Command{
	Use:     "encode",
	Short:   "Encode string to base64.",
	Aliases: []string{"e"},
	Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		encoded := base64.StdEncoding.EncodeToString([]byte(args[0]))
		fmt.Println(encoded)
	},
}
