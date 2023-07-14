package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"server/cmd/api"
)

var rootCmd = &cobra.Command{
	Use:               "drawer",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `drawer`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + `drawer-admin ` + ` 可以使用 ` + `-h` + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
