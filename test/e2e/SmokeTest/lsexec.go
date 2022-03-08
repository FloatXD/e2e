package SmokeTest

import (
	"context"
	"fmt"
	lsapi "github.com/hwameistor/local-storage/pkg/apis/client/clientset/versioned/scheme"
	_ "github.com/niulechuan/e2e/pkg/apis"
	"github.com/niulechuan/e2e/test/e2e/framework"
	"github.com/onsi/ginkgo"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var _ = ginkgo.Describe("volume", func() {
	f := framework.NewDefaultFramework(lsapi.AddToScheme)
	client := f.GetClient()
	ginkgo.Describe("LDM test", func() {

		ginkgo.It("get ready", func() {
			config, err := config.GetConfig()
			if err != nil {
				return
			}

			deployment := &appsv1.Deployment{}
			deployKey := k8sclient.ObjectKey{
				Name:      "demo-2048",
				Namespace: "default",
			}
			err = client.Get(context.TODO(), deployKey, deployment)
			if err != nil {
				fmt.Printf("%+v \n", err)
				f.ExpectNoError(err)
			}

			apps, err := labels.NewRequirement("app", selection.In, []string{"demo"})
			selector := labels.NewSelector()
			selector = selector.Add(*apps)
			listOption := k8sclient.ListOptions{
				LabelSelector: selector,
			}
			podlist := &v1.PodList{}
			err = client.List(context.TODO(), podlist, &listOption)

			if err != nil {
				fmt.Printf("%+v \n", err)
				f.ExpectNoError(err)
			}

			containers := deployment.Spec.Template.Spec.Containers
			for _, container := range containers {
				for _, pod := range podlist.Items {
					a, _, err := ExecInPod(config, deployment.Namespace, pod.Name, "cd /data && echo it-is-a-test >test", container.Name)
					fmt.Printf("%+v \n", a)
					a, _, err = ExecInPod(config, deployment.Namespace, pod.Name, "cd /data && cat test", container.Name)
					fmt.Printf("%+v \n", a)
					if err != nil {
						fmt.Printf("%+v \n", err)
						f.ExpectNoError(err)
					}
				}
			}

		})
	})
})
