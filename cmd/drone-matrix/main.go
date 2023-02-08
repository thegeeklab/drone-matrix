// Copyright (c) 2020, the Drone Plugins project authors.
// Copyright (c) 2021, Robert Kaussow <mail@thegeeklab.de>

// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/drone-plugins/drone-matrix/plugin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/thegeeklab/drone-plugin-lib/v2/urfave"
	"github.com/urfave/cli/v2"
)

//nolint:gochecknoglobals
var (
	BuildVersion = "devel"
	BuildDate    = "00000000"
)

func main() {
	settings := &plugin.Settings{}

	if _, err := os.Stat("/run/drone/env"); err == nil {
		_ = godotenv.Overload("/run/drone/env")
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version=%s date=%s\n", c.App.Name, c.App.Version, BuildDate)
	}

	app := &cli.App{
		Name:    "drone-matrix",
		Usage:   "build notifications for matrix",
		Version: BuildVersion,
		Flags:   append(settingsFlags(settings, urfave.FlagsPluginCategory), urfave.Flags()...),
		Action:  run(settings),
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(settings *plugin.Settings) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		urfave.LoggingFromContext(ctx)

		plugin := plugin.New(
			*settings,
			urfave.PipelineFromContext(ctx),
			urfave.NetworkFromContext(ctx),
		)

		if err := plugin.Validate(); err != nil {
			return fmt.Errorf("validation failed: %w", err)
		}

		if err := plugin.Execute(); err != nil {
			return fmt.Errorf("execution failed: %w", err)
		}

		return nil
	}
}
