package SmokeTest

import (
	_ "github.com/niulechuan/e2e/pkg/apis"
	"github.com/onsi/ginkgo"
)

var _ = ginkgo.Describe("volume", func() {
	ginkgo.Describe("LDM test", func() {
		ginkgo.It("get ready", func() {
			uninstallHelm()
		})

	})
})
