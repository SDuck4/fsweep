package cmd

import (
	"fmt"
	"os"

	"github.com/SDuck4/fsweep/internal"
	"github.com/spf13/cobra"
)

var cfgFile string
var name string

var rootCmd = &cobra.Command{
	Use:   "fsweep <path> <number-of-days>",
	Short: "Sweep old files",
	Long: `Sweep old files based on file modified time. For example:

./fsweep /var/log/httpd 30
`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		internal.Sweep(args, flags)
	},
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", ".*", "file name pattern to delete in regexp")
	rootCmd.Flags().BoolP("assumeyes", "y", false, "assume that the answer to any question which would be asked is yes")
}
