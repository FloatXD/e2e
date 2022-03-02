package SmokeTest

import (
	"context"
	"fmt"
	ldapis "github.com/hwameistor/local-disk-manager/pkg/apis"
	_ "github.com/niulechuan/e2e/pkg/apis"
	"github.com/niulechuan/e2e/test/e2e/framework"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"time"
)

var _ = Describe("volume", func() {
	f := framework.NewDefaultFramework(ldapis.AddToScheme)
	client := f.GetClient()
	nodelist := &apiv1.NodeList{}
	err := client.List(context.TODO(), nodelist)
	if err != nil {
		f.ExpectNoError(err)
		fmt.Printf("%+v \n", err)
	}
	for _, nodes := range nodelist.Items {
		node := &apiv1.Node{}
		nodeKey := k8sclient.ObjectKey{
			Name: nodes.Name,
		}
		err := client.Get(context.TODO(), nodeKey, node)
		if err != nil {
			fmt.Printf("%+v \n", err)
			f.ExpectNoError(err)
		}
		_, boolLabel := node.Labels["csi.driver.uds.dce.daocloud.io/local.storage.daocloud.io"]
		if !boolLabel {
			node.Labels["localstorage.hwameistor.io/local-storage"] = "true"
			node.Labels["localstorage.hwameistor.io/local-storage-topology-node"] = nodes.Name
			fmt.Printf("adding labels \n")
			err := client.Update(context.TODO(), node)
			if err != nil {
				fmt.Printf("%+v \n", err)
				f.ExpectNoError(err)
			}
			fmt.Printf("wait 1 minute\n")
			time.Sleep(1 * time.Minute)
		}
	}
	Describe("Ls test", func() {
		Context("check hwameistor", func() {
			It("check status", func() {
				daemonset := &appsv1.DaemonSet{}
				daemonsetKey := k8sclient.ObjectKey{
					Name:      "hwameistor",
					Namespace: "hwameistor",
				}

				err := client.Get(context.TODO(), daemonsetKey, daemonset)
				if err != nil {
					f.ExpectNoError(err)

				}
				Expect(daemonset.Status.DesiredNumberScheduled).To(Equal(daemonset.Status.NumberAvailable))
			})
		})

		Context("check hwameistor", func() {
			It("check status", func() {
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
				Expect(deployment.Status.AvailableReplicas).To(Equal(int32(1)))
			})
		})
		Context("check hwameistor", func() {
			It("check status", func() {
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
				Expect(deployment.Status.AvailableReplicas).To(Equal(int32(1)))
			})
		})
	})
})
