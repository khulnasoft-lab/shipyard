package env

import (
	"net/http"

	"github.com/khulnasoft-lab/shipyard/constants"
	"github.com/khulnasoft-lab/shipyard/pkg/client"
	"github.com/khulnasoft-lab/shipyard/pkg/display"
	"github.com/khulnasoft-lab/shipyard/pkg/requests/uri"
	"github.com/spf13/cobra"
)

func NewRestartCmd(c client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "restart",
		GroupID: constants.GroupEnvironments,
		Short:   "Restart a stopped environment",
		Example: `  # Restart environment ID 12345
  shipyard restart environment 12345`,
		SilenceUsage: true,
	}

	cmd.AddCommand(newRestartEnvironmentCmd(c))

	return cmd
}

func newRestartEnvironmentCmd(c client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Aliases:      []string{"env"},
		Use:          "environment [environment ID]",
		SilenceUsage: true,
		Short:        "Restart a stopped environment",
		Example: `  # Restart environment ID 12345
  shipyard restart environment 12345`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return restartEnvironmentByID(c, args[0])
			}
			return errNoEnvironment
		},
	}

	return cmd
}

func restartEnvironmentByID(c client.Client, id string) error {
	params := make(map[string]string)
	if c.Org != "" {
		params["org"] = c.Org
	}

	_, err := c.Requester.Do(http.MethodPost, uri.CreateResourceURI("restart", "environment", id, "", params), nil)
	if err != nil {
		return err
	}

	display.Println("Environment queued for a restart.")
	return nil
}
