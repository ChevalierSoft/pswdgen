package cmd

import (
	"fmt"
	"os"

	"pswdgen/pkg/gen"

	"github.com/spf13/cobra"
)

var info gen.PswdInfo = gen.PswdInfo{}

var rootCmd = &cobra.Command{
	Use:   "pswdgen",
	Short: "pswdgen generates a password for you",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		gen.Gen(&info)
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&info.Len, "length", "l", 16, "password length")
	rootCmd.PersistentFlags().BoolVarP(&info.Lower, "lower", "w", false, "use lower characters")
	rootCmd.PersistentFlags().BoolVarP(&info.Upper, "upper", "u", false, "use upper characters")
	rootCmd.PersistentFlags().BoolVarP(&info.Numbers, "numbers", "n", false, "use numbers")
	rootCmd.PersistentFlags().BoolVarP(&info.Special, "specials", "s", false, "use special characters")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
