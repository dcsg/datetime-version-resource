package main_test

import (
	"bytes"
	"encoding/json"
	"time"

	resource "github.com/dcsg/datetime-version-resource"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Out", func() {
	var (
		session *gexec.Session

		outResponse resource.OutResponse
	)

	Describe("Generate a date time version", func() {
		Context("With default timezone and format", func() {
			It("generates a version", func() {
				session = RunOut(resource.Source{})

				err = json.NewDecoder(bytes.NewBuffer(session.Out.Contents())).Decode(&outResponse)
				Expect(err).NotTo(HaveOccurred())

				versionTime, err := time.Parse("20060102-150405", outResponse.Version.Version)
				Expect(err).NotTo(HaveOccurred())

				Expect(versionTime.Hour()).To(Equal(time.Now().UTC().Hour()))
				Expect(versionTime.Minute()).To(Equal(time.Now().UTC().Minute()))
				Expect(versionTime.Year()).To(Equal(time.Now().UTC().Year()))
				Expect(versionTime.Month()).To(Equal(time.Now().UTC().Month()))
				Expect(versionTime.Day()).To(Equal(time.Now().UTC().Day()))
			})

			It("emits metadata", func() {
				session = RunOut(resource.Source{})

				err = json.NewDecoder(bytes.NewBuffer(session.Out.Contents())).Decode(&outResponse)
				Expect(err).NotTo(HaveOccurred())

				Expect(outResponse.Metadata).To(ConsistOf(
					resource.Metadata{Name: "version", Value: outResponse.Version.Version},
					resource.Metadata{Name: "timezone", Value: "UTC"},
					resource.Metadata{Name: "format", Value: "20060102-150405"},
				))
			})
		})
	})

	Context("when timezone is defined", func() {
		It("emits metadata", func() {
			session = RunOut(resource.Source{
				Timezone: "Europe/Lisbon",
			})

			err = json.NewDecoder(bytes.NewBuffer(session.Out.Contents())).Decode(&outResponse)
			Expect(err).NotTo(HaveOccurred())

			Expect(outResponse.Metadata).To(ConsistOf(
				resource.Metadata{Name: "version", Value: outResponse.Version.Version},
				resource.Metadata{Name: "timezone", Value: "Europe/Lisbon"},
				resource.Metadata{Name: "format", Value: "20060102-150405"},
			))
		})

		It("errors when an invalid timezone is provided", func() {
			source := resource.Source{
				Timezone: "invalid",
			}

			session = RunOut(source, Not(gexec.Exit(0)))

			Expect(session.Err).To(gbytes.Say("unknown time zone invalid"))
		})
	})

	Context("when datetime format is defined", func() {
		It("emits metadata", func() {
			session = RunOut(resource.Source{
				Format: "20060102.150405",
			})

			err = json.NewDecoder(bytes.NewBuffer(session.Out.Contents())).Decode(&outResponse)
			Expect(err).NotTo(HaveOccurred())

			Expect(outResponse.Metadata).To(ConsistOf(
				resource.Metadata{Name: "version", Value: outResponse.Version.Version},
				resource.Metadata{Name: "timezone", Value: "UTC"},
				resource.Metadata{Name: "format", Value: "20060102.150405"},
			))
		})

		It("errors when an invalid date format is provided", func() {
			source := resource.Source{
				Format: "invalid",
			}

			session = RunOut(source, Not(gexec.Exit(0)))

			Expect(session.Err).To(gbytes.Say("invalid date format"))
		})
	})
})
