package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var rootFlags struct {
	verbose bool
	port    int
	static  string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "vite-echo-monorepo server",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()
		e.Static("/", rootFlags.static)
		e.GET("/api/hello", func(c echo.Context) error {
			return c.String(http.StatusOK, "world!")
		})
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", rootFlags.port)))
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
	rootCmd.PersistentFlags().BoolVarP(&rootFlags.verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().IntVarP(&rootFlags.port, "port", "p", 8080, "Port to run the server on")
	rootCmd.Flags().StringVarP(&rootFlags.static, "static", "s", "./static", "Path to static files")
}
