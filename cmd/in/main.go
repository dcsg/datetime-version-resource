package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	resource "github.com/dcsg/datetime-version-resource"
)

func main() {
	artifactDirectory := os.Args[1]

	var (
		inRequest resource.InRequest
		payload   resource.InResponse
	)

	inRequest.Source = resource.SourceDefaults

	err := json.NewDecoder(os.Stdin).Decode(&inRequest)
	if err != nil {
		panic(err)
	}

	payload = resource.InResponse{
		Version: resource.Version{
			Version: inRequest.Version.Version,
		},
		Metadata: []resource.Metadata{
			{
				Name:  "version",
				Value: inRequest.Version.Version,
			},
			{
				Name:  "timezone",
				Value: inRequest.Source.Timezone,
			},
			{
				Name:  "format",
				Value: inRequest.Source.Format,
			},
		},
	}

	fp := filepath.Join(artifactDirectory, "version")

	err = ioutil.WriteFile(fp, []byte(payload.Version.Version), 0644)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(os.Stdout).Encode(payload)
	if err != nil {
		panic(err)
	}
}
