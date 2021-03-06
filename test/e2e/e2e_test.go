package e2e

import (
	"math/rand"
	"os"
	"testing"
	"time"

	_ "github.com/niulechuan/e2e/test/e2e/SmokeTest"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func TestE2E(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "hwameistor e2e test")
}
