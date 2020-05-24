package main_test

import (
	"encoding/json"
	"os/exec"

	resource "github.com/dcsg/datetime-version-resource"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"io/ioutil"
	"testing"

	"fmt"
	"os"

	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/types"
)

func TestOut(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "In Suite")
}

var (
	binPath string
	err     error

	tmpDir string
)

var _ = BeforeEach(func() {
	if _, err = os.Stat("/opt/resource/in"); err == nil {
		binPath = "/opt/resource/in"
	} else {
		binPath, err = gexec.Build("github.com/dcsg/datetime-version-resource/cmd/in")
		Expect(err).NotTo(HaveOccurred())
	}

	tmpDir, err = ioutil.TempDir("", "datetime_version_resource_in")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterEach(func() {
	err := os.RemoveAll(tmpDir)
	Expect(err).NotTo(HaveOccurred())
})

func RunIn(s resource.Source, v resource.Version, matchers ...types.GomegaMatcher) *gexec.Session {
	payload := resource.InRequest{
		Source: s,
		Version: v,
	}

	b, err := json.Marshal(&payload)
	Expect(err).NotTo(HaveOccurred())

	c := exec.Command(binPath, tmpDir)
	c.Stdin = bytes.NewBuffer(b)
	sess, err := gexec.Start(c, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	<-sess.Exited
	if len(matchers) == 0 {
		Expect(sess).To(gexec.Exit(0), fmt.Sprintf("Expected session to exit 0, exited with %d.\n\nStdout: %s\n\nStderr: %s", sess.ExitCode(), sess.Out.Contents(), sess.Err.Contents()))
	} else {
		for _, matcher := range matchers {
			Expect(sess).To(matcher)
		}
	}

	return sess
}
