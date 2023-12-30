package cli

import (
	"github.com/wcygan/fs/internal/client"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A CLI to interact with the file server",
}

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload [filename]",
	Short: "Upload a file to the server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		err := client.Upload(filename)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(uploadCmd)
}
