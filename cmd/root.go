package cmd

import (
	"fmt"
	"os"

	"github.com/chevaliersoft/pswdgen/pkg/gen"

	"github.com/spf13/cobra"
)

const (
  MaxLength = 2048
)

var info gen.PswdInfo = gen.PswdInfo{}

var rootCmd = &cobra.Command{
	Use:   "pswdgen",
	Short: "pswdgen generates a random password",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
    if info.Len > 2048 {
      fmt.Printf("error: length must be lower than %d.\n", MaxLength)
      os.Exit(1)
    }
		gen.Gen(info)
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
