package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var descriptionCmd = &cobra.Command{
	Use:  "0.0.1",
	Long: "A simple test with Cobra cmd lib",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run test...")
	},
}

func Execute() {
	if err := descriptionCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
