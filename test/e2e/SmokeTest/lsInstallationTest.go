package SmokeTest

import (
	"context"
	ldapis "github.com/hwameistor/local-disk-manager/pkg/apis"
	"github.com/niulechuan/e2e/test/e2e/framework"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = ginkgo.Describe("volume", func() {
	f := framework.NewDefaultFramework(ldapis.AddToScheme)
	client := f.GetClient()

	ginkgo.It("Configure the base environment", func() {
		installHelm()
		addLabels()
	})
	ginkgo.Context("test localstorage", func() {
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
	ginkgo.Context("test hwameistor-csi-controller", func() {
		ginkgo.It("check status", func() {
			deployment := &appsv1.Deployment{}
			deploymentKey := k8sclient.ObjectKey{
				Name:      "hwameistor-csi-controller",
				Namespace: "hwameistor",
			}

			err := client.Get(context.TODO(), deploymentKey, deployment)
			if err != nil {
				f.ExpectNoError(err)
				logrus.Printf("%+v \n", err)
			}
			gomega.Expect(deployment.Status.AvailableReplicas).To(gomega.Equal(int32(1)))
		})
	})
	ginkgo.Context("test hwameistor-scheduler", func() {
		ginkgo.It("check status", func() {
			deployment := &appsv1.Deployment{}
			deploymentKey := k8sclient.ObjectKey{
				Name:      "hwameistor-scheduler",
				Namespace: "hwameistor",
			}

			err := client.Get(context.TODO(), deploymentKey, deployment)
			if err != nil {
				f.ExpectNoError(err)
				logrus.Printf("%+v \n", err)
			}
			gomega.Expect(deployment.Status.AvailableReplicas).To(gomega.Equal(int32(1)))
		})
	})
	ginkgo.It("Clean up the environment", func() {
		uninstallHelm()

	})
})
