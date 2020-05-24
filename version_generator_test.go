package resource_test

import (
	"github.com/benbjohnson/clock"
	resource "github.com/dcsg/datetime-version-resource"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VersionGenerator", func() {
	var (
		c clock.Clock
	)

	BeforeEach(func() {
		c = clock.NewMock()
	})

	Describe("with default timezone and format", func() {
		It("Generates a UTC datetime version", func() {
			version, err := resource.VersionGenerator(c, "", "")
			Expect(err).NotTo(HaveOccurred())

			Expect(version).To(Equal("19700101-000000"))
		})
	})

	Describe("with params timezone and format", func() {
		It("Generates a version in specified timezone and format", func() {
			version, err := resource.VersionGenerator(c, "Europe/Lisbon", "20060102.150405")
			Expect(err).NotTo(HaveOccurred())

			Expect(version).To(Equal("19700101.010000"))
		})

		It("Errors when an invalid date format is provided", func() {
			_, err := resource.VersionGenerator(c, "", "invalid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("invalid date format"))
		})

		It("Errors when an invalid timezone is provided", func() {
			_, err := resource.VersionGenerator(c, "invalid", "")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("unknown time zone invalid"))
		})
	})
})
