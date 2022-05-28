// Copyright (c) 2020, the Drone Plugins project authors.
// Copyright (c) 2021, Robert Kaussow <mail@thegeeklab.de>

// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package main

import (
	"github.com/drone-plugins/drone-matrix/plugin"
	"github.com/urfave/cli/v2"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "username",
			EnvVars:     []string{"PLUGIN_USERNAME", "MATRIX_USERNAME"},
			Usage:       "sets username for authentication",
			Destination: &settings.Username,
		},
		&cli.StringFlag{
			Name:        "password",
			EnvVars:     []string{"PLUGIN_PASSWORD", "MATRIX_PASSWORD"},
			Usage:       "sets password for authentication",
			Destination: &settings.Password,
		},
		&cli.StringFlag{
			Name:        "userid",
			EnvVars:     []string{"PLUGIN_USERID,PLUGIN_USER_ID", "MATRIX_USERID", "MATRIX_USER_ID"},
			Usage:       "sets userid for authentication",
			Destination: &settings.UserID,
		},
		&cli.StringFlag{
			Name:        "accesstoken",
			EnvVars:     []string{"PLUGIN_ACCESSTOKEN,PLUGIN_ACCESS_TOKEN", "MATRIX_ACCESSTOKEN", "MATRIX_ACCESS_TOKEN"},
			Usage:       "sets access token for authentication",
			Destination: &settings.AccessToken,
		},
		&cli.StringFlag{
			Name:        "homeserver",
			EnvVars:     []string{"PLUGIN_HOMESERVER", "MATRIX_HOMESERVER"},
			Usage:       "sets matrix home server url to use",
			Value:       "https://matrix.org",
			Destination: &settings.Homeserver,
		},
		&cli.StringFlag{
			Name:        "roomid",
			EnvVars:     []string{"PLUGIN_ROOMID", "MATRIX_ROOMID"},
			Usage:       "sets roomid to send messages to",
			Destination: &settings.RoomID,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "template",
			EnvVars:     []string{"PLUGIN_TEMPLATE", "MATRIX_TEMPLATE"},
			Usage:       "sets message template",
			Value:       "Build {{ build.Status }} [{{ repo.Owner }}/{{ repo.Name }}#{{ truncate commit.SHA 8 }}]({{ build.Link }}) ({{ build.Branch }}) by {{ commit.Author }}",
			Destination: &settings.Template,
		},
	}
}
