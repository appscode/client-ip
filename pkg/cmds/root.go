package cmds

import (
	"flag"
	"log"

	"github.com/appscode/analytics/pkg/dockerhub/printer"
	v "github.com/appscode/go/version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewRootCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "analytics [command]",
		Short:             `Analytics by AppsCode - Essential analytics for OSS`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			c.Flags().VisitAll(func(flag *pflag.Flag) {
				log.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
			})
		},
	}
	cmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	flag.CommandLine.Parse([]string{})

	cmd.AddCommand(NewCmdServer(version))
	cmd.AddCommand(v.NewCmdVersion())
	cmd.AddCommand(printer.NewCmdDockerHub())
	return cmd
}
