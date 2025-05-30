package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	release   = "UNKNOWN"
	buildDate = "UNKNOWN"
	gitHash   = "UNKNOWN"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Показать информацию о версии сборки",
	Run: func(cmd *cobra.Command, args []string) {
		if err := json.NewEncoder(os.Stdout).Encode(struct {
			Release   string `json:"release"`
			BuildDate string `json:"buildDate"`
			GitHash   string `json:"gitHash"`
		}{
			Release:   release,
			BuildDate: buildDate,
			GitHash:   gitHash,
		}); err != nil {
			fmt.Printf("error while encoding version info: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
