package resource

import (
	"errors"
	"time"

	"github.com/benbjohnson/clock"
)

func VersionGenerator(c clock.Clock, timezone string, layout string) (string, error) {
	if timezone == "" {
		timezone = SourceDefaults.Timezone
	}

	if layout == "" {
		layout = SourceDefaults.Format
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}

	version := c.Now().In(location).Format(layout)

	if version == layout {
		err = errors.New("invalid date format")
		return "", err
	}

	return version, err
}
