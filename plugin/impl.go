// Copyright (c) 2020, the Drone Plugins project authors.
// Copyright (c) 2021, Robert Kaussow <mail@thegeeklab.de>

// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"errors"
	"fmt"
	"strings"

	"github.com/matrix-org/gomatrix"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"github.com/sirupsen/logrus"
	"github.com/thegeeklab/drone-plugin-lib/v2/template"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

// Settings for the plugin.
type Settings struct {
	Username    string
	Password    string
	UserID      string
	AccessToken string
	Homeserver  string
	RoomID      string
	Template    string
}

var ErrAuthSourceNotSet = errors.New("either username and password or userid and accesstoken are required")

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	if (p.settings.Username == "" || p.settings.Password == "") &&
		(p.settings.UserID == "" || p.settings.AccessToken == "") {
		return ErrAuthSourceNotSet
	}

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	muid := id.NewUserID(prepend("@", p.settings.UserID), p.settings.Homeserver)

	client, err := mautrix.NewClient(p.settings.Homeserver, muid, p.settings.AccessToken)
	if err != nil {
		return fmt.Errorf("failed to initialize client: %w", err)
	}

	if p.settings.UserID == "" || p.settings.AccessToken == "" {
		_, err := client.Login(&mautrix.ReqLogin{
			Type:                     "m.login.password",
			Identifier:               mautrix.UserIdentifier{Type: mautrix.IdentifierTypeUser, User: p.settings.Username},
			Password:                 p.settings.Password,
			InitialDeviceDisplayName: "Drone",
			StoreCredentials:         true,
		})
		if err != nil {
			return fmt.Errorf("failed to authenticate user: %w", err)
		}
	}

	logrus.Info("logged in successfully")

	joined, err := client.JoinRoom(prepend("!", p.settings.RoomID), "", nil)
	if err != nil {
		return fmt.Errorf("failed to join room: %w", err)
	}

	message, err := template.RenderTrim(p.network.Context, *p.network.Client, p.settings.Template, p.pipeline)
	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	formatted := bluemonday.UGCPolicy().SanitizeBytes(
		blackfriday.Run([]byte(message)),
	)

	content := gomatrix.HTMLMessage{
		Body:          message,
		MsgType:       "m.notice",
		Format:        "org.matrix.custom.html",
		FormattedBody: string(formatted),
	}

	if _, err := client.SendMessageEvent(joined.RoomID, event.EventMessage, content); err != nil {
		return fmt.Errorf("failed to submit message: %w", err)
	}

	logrus.Info("message sent successfully")

	return nil
}

func prepend(prefix, input string) string {
	if strings.TrimSpace(input) == "" {
		return input
	}

	if strings.HasPrefix(input, prefix) {
		return input
	}

	return prefix + input
}
