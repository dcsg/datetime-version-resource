package main_test

import (
	"bytes"
	"encoding/json"

	resource "github.com/dcsg/datetime-version-resource"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Check", func() {
	var (
		session *gexec.Session

		err error

		checkResponse []string
	)

	Context("when invoked", func() {
		It("returns empty digest", func() {
			session = RunCheck(resource.Source{})

			err = json.NewDecoder(bytes.NewBuffer(session.Out.Contents())).Decode(&checkResponse)
			Expect(err).NotTo(HaveOccurred())

			Expect(checkResponse).To(Equal(make([]string, 0)))
		})
	})
})
