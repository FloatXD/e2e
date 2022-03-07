package SmokeTest

import (
	"context"
	"fmt"
	ldapis "github.com/hwameistor/local-disk-manager/pkg/apis"
	_ "github.com/niulechuan/e2e/pkg/apis"
	"github.com/niulechuan/e2e/test/e2e/framework"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = ginkgo.Describe("volume", func() {
	f := framework.NewDefaultFramework(ldapis.AddToScheme)
	client := f.GetClient()

	ginkgo.Describe("LDM test", func() {
		ginkgo.It("get ready", func() {
			installHelm()
		})
		ginkgo.Context("check local-disk-manager", func() {
			ginkgo.It("check status", func() {
				daemonset := &appsv1.DaemonSet{}
				daemonsetKey := k8sclient.ObjectKey{
					Name:      "local-disk-manager",
					Namespace: "hwameistor",
				}

				err := client.Get(context.TODO(), daemonsetKey, daemonset)
				if err != nil {
					f.ExpectNoError(err)
					fmt.Printf("%+v \n", err)
				}
				gomega.Expect(daemonset.Status.DesiredNumberScheduled).To(gomega.Equal(daemonset.Status.NumberAvailable))
			})
		})
		ginkgo.It("delete helm", func() {
			uninstallHelm()

		})
	})
})
