package cmd

import (
	"os"

	"github.com/narumiruna/wolframalpha/pkg/simple"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "wolframalpha",
	Short: "WolframAlpha Simple API",

	RunE: run,
}

func init() {
	RootCmd.Flags().StringP("input", "i", "", "input value")
	RootCmd.Flags().StringP("output", "o", "output.png", "filename of output image")
	RootCmd.Flags().String("appid", "", "App ID")

	RootCmd.Flags().Int("width", 0, "width of output image")
	RootCmd.Flags().Int("fontsize", 0, "font size of output image")
	RootCmd.Flags().String("units", "", "units to use for measurements and quantities")
	RootCmd.Flags().Int("timeout", 0, "maximum amount of time (in seconds) allowed to process a query")
}

func run(cmd *cobra.Command, args []string) error {
	input, err := cmd.Flags().GetString("input")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	appID, err := cmd.Flags().GetString("appid")
	if err != nil {
		return err
	}

	if appID == "" {
		appID = os.Getenv("WOLFRAMALPHA_APP_ID")
	}

	width, err := cmd.Flags().GetInt("width")
	if err != nil {
		return err
	}

	fontsize, err := cmd.Flags().GetInt("fontsize")
	if err != nil {
		return err
	}

	units, err := cmd.Flags().GetString("units")
	if err != nil {
		return err
	}

	timeout, err := cmd.Flags().GetInt("timeout")
	if err != nil {
		return err
	}

	options := &simple.QueryOptions{}

	if width != 0 {
		options.Width = width
	}

	if fontsize != 0 {
		options.Fontsize = fontsize
	}

	if units != "" {
		options.Units = units
	}

	if timeout != 0 {
		options.Timeout = timeout
	}

	client := simple.New(appID)

	return client.QueryFile(input, output, options)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("failed to execute root command")
	}
}
