// Copyright © 2017 Ray Johnson <ray.johnson@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"github.com/docker/docker/pkg/term"
	"github.com/spf13/cobra"
	"os"
)

const usageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{ wrappedFlagUsages . | trimRightSpace}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

const filesManInfo = `Many of the options for this command can be put in a config file.
You can create a config file at $HOME/.zap.toml.  Configs found in the config file will override built-in
defaults but can be overridden by explict command-line options.

The format of the config file is written in Toml.  Sections in brackets (e.g. [broker]) can be
referenced with the --broker flag.  Values should be of the form qos = "1" and the keys will
have the same name as the option values listed above.`

var cfgFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string, revision string) {
	rootCmd := SetupRootCommand(version, revision)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// SetupRootCommand sets of the cobra data structures for command line processing
func SetupRootCommand(version string, revision string) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var versionFlag bool
	var rootCmd = &cobra.Command{
		Use:   "zap",
		Args:  cobra.NoArgs,
		Short: "Listen or publish to a MQTT broker",
		Long: `zap - what happens when technology meets mosquito

zap is a little utility for publishing or subscribing to events for the
MQTT message bus`,
		Run: func(cmd *cobra.Command, args []string) {
			if versionFlag {
				fmt.Println("zap version " + version + ", Revision: " + revision)
			} else {
				fmt.Println(cmd.Help())
			}
		},
	}

	rootCmd.Flags().BoolVar(&versionFlag, "version", false, "Display version information")

	cobra.AddTemplateFunc("wrappedFlagUsages", wrappedFlagUsages)

	rootCmd.AddCommand(
		newSubscribeCommand(),
		newPublishCommand(),
		newStatsCommand(),
	)
	rootCmd.SetUsageTemplate(usageTemplate)

	setUpLogging()

	return rootCmd
}

func wrappedFlagUsages(cmd *cobra.Command) string {
	width := 80
	if ws, err := term.GetWinsize(0); err == nil {
		width = int(ws.Width)
	}
	return cmd.Flags().FlagUsagesWrapped(width - 1)
}

func setUpLogging() {

	// TODO: need to set up my own log handlers and get rid of
	// all the fmt.Print statements.  Can also handle verbose this way

	// standard
	// stats
	// verbose
	// debug

}
