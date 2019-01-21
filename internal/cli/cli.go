package cli

import (
	"github.com/igarciaolaizola/h2-forward/internal/server"
	"github.com/spf13/cobra"
)

// NewCommand create and returns the root cli command
func NewCommand() *cobra.Command {
	var addr string
	var port int

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Launches the http2 proxy server",
		RunE: func(c *cobra.Command, args []string) error {
			return server.Run(addr, port)
		},
	}

	cmd.Flags().StringVar(&addr, "addr", "localhost:8080", "listening address")
	cmd.Flags().IntVar(&port, "port", 8081, "port to forward")
	return cmd
}
