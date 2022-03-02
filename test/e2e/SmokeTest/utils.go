package SmokeTest

import (
	"context"
	"fmt"
	ldapis "github.com/hwameistor/local-disk-manager/pkg/apis"
	_ "github.com/niulechuan/e2e/pkg/apis"
	"github.com/niulechuan/e2e/test/e2e/framework"
	apiv1 "k8s.io/api/core/v1"
	"os/exec"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"time"
)

func int32Ptr(i int32) *int32 { return &i }

func boolPter(i bool) *bool { return &i }

func runInLinux(cmd string) string {
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("ERROR:%+v \n", err)
	}
	return string(result)
}

func nodeList() *apiv1.NodeList {
	f := framework.NewDefaultFramework(ldapis.AddToScheme)
	client := f.GetClient()
	nodelist := &apiv1.NodeList{}
	err := client.List(context.TODO(), nodelist)
	if err != nil {
		f.ExpectNoError(err)
		fmt.Printf("%+v \n", err)
	}
	return nodelist
}

func addLabels() {
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
}

func uninstallHelm() {
	_ = runInLinux("helm list -A | grep 'hwameistor' | awk '{print $1}' | xargs helm uninstall -n hwameistor")
	_ = runInLinux("kubectl get crd | grep 'hwameistor' | awk '{print $1}' | xargs -n1 kubectl delete crd")
}
