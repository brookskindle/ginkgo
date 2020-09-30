package remote_test

import (
	. "github.com/brookskindle/ginkgo"
	. "github.com/brookskindle/gomega"

	"testing"
)

func TestRemote(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Remote Spec Forwarding Suite")
}
