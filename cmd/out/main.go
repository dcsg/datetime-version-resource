package main

import (
	"encoding/json"
	"os"

	"github.com/benbjohnson/clock"
	resource "github.com/dcsg/datetime-version-resource"
)

func main() {
	outRequest := resource.OutRequest{
		Source: resource.SourceDefaults,
	}

	err := json.NewDecoder(os.Stdin).Decode(&outRequest)
	if err != nil {
		panic(err)
	}

	version, err := resource.VersionGenerator(
		clock.New(),
		outRequest.Source.Timezone,
		outRequest.Source.Format,
	)
	if err != nil {
		panic(err)
	}

	outResponse := resource.OutResponse{
		Version: resource.Version{
			Version: version,
		},
		Metadata: []resource.Metadata{
			{
				Name:  "version",
				Value: version,
			},
			{
				Name:  "timezone",
				Value: outRequest.Source.Timezone,
			},
			{
				Name:  "format",
				Value: outRequest.Source.Format,
			},
		},
	}

	err = json.NewEncoder(os.Stdout).Encode(outResponse)
	if err != nil {
		panic(err)
	}
}
