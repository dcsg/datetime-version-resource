package main_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	resource "github.com/dcsg/datetime-version-resource"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("In", func() {
	var (
		session *gexec.Session
		inResponse resource.InResponse
		version = resource.Version{
			Version: "20200101-123004",
		}
	)

	It("writes an `version` file", func() {
		session = RunIn(resource.SourceDefaults, version)

		b, err := ioutil.ReadFile(filepath.Join(tmpDir, "version"))
		Expect(err).NotTo(HaveOccurred())

		Expect(version.Version).To(Equal(string(b)))
	})

	It("emits metadata", func() {
		session = RunIn(resource.SourceDefaults, version)

		err = json.NewDecoder(bytes.NewBuffer(session.Out.Contents())).Decode(&inResponse)
		Expect(err).NotTo(HaveOccurred())

		Expect(inResponse.Metadata).To(ConsistOf(
			resource.Metadata{Name: "version", Value: version.Version},
			resource.Metadata{Name: "timezone", Value: resource.SourceDefaults.Timezone},
			resource.Metadata{Name: "format", Value: resource.SourceDefaults.Format},
		))
	})
})
