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

	ginkgo.Describe("Ls test", func() {
		ginkgo.It("get ready", func() {
			installHelm()
			addLabels()
		})
		ginkgo.Context("check hwameistor", func() {
			ginkgo.It("check status", func() {
				daemonset := &appsv1.DaemonSet{}
				daemonsetKey := k8sclient.ObjectKey{
					Name:      "hwameistor",
					Namespace: "hwameistor",
				}

				err := client.Get(context.TODO(), daemonsetKey, daemonset)
				if err != nil {
					f.ExpectNoError(err)

				}

				gomega.Expect(daemonset.Status.DesiredNumberScheduled).To(gomega.Equal(daemonset.Status.NumberAvailable))
			})
		})
		ginkgo.Context("check hwameistor-csi-controller", func() {
			ginkgo.It("check status", func() {
				deployment := &appsv1.Deployment{}
				deploymentKey := k8sclient.ObjectKey{
					Name:      "hwameistor-csi-controller",
					Namespace: "hwameistor",
				}

				err := client.Get(context.TODO(), deploymentKey, deployment)
				if err != nil {
					f.ExpectNoError(err)
					fmt.Printf("%+v \n", err)
				}
				gomega.Expect(deployment.Status.AvailableReplicas).To(gomega.Equal(int32(1)))
			})
		})
		ginkgo.Context("check hwameistor-scheduler", func() {
			ginkgo.It("check status", func() {
				deployment := &appsv1.Deployment{}
				deploymentKey := k8sclient.ObjectKey{
					Name:      "hwameistor-scheduler",
					Namespace: "hwameistor",
				}

				err := client.Get(context.TODO(), deploymentKey, deployment)
				if err != nil {
					f.ExpectNoError(err)
					fmt.Printf("%+v \n", err)
				}
				gomega.Expect(deployment.Status.AvailableReplicas).To(gomega.Equal(int32(1)))
			})
		})
		ginkgo.It("delete helm", func() {
			uninstallHelm()

		})
	})
})
