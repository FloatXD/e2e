package SmokeTest

import (
	"context"
	"fmt"
	ldapis "github.com/hwameistor/local-disk-manager/pkg/apis"
	ldv1 "github.com/hwameistor/local-disk-manager/pkg/apis/hwameistor/v1alpha1"
	_ "github.com/niulechuan/e2e/pkg/apis"
	"github.com/niulechuan/e2e/test/e2e/framework"
	apiv1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	fmt.Printf("get node list\n")
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
	fmt.Printf("add node labels\n")
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
		_, boolLabel := node.Labels["localstorage.hwameistor.io/local-storage"]
		if !boolLabel {
			node.Labels["localstorage.hwameistor.io/local-storage"] = "true"
			node.Labels["csi.driver.hwameistor.io/localstorage"] = "true"
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

func installHelm() {
	fmt.Printf("helm install hwameistor\n")
	_ = runInLinux("cd /root/helm-charts-hwameistor-0.2.3/charts && helm install hwameistor -n hwameistor --create-namespace --generate-name")
	fmt.Printf("waiting for intall hwameistor\n")
	time.Sleep(1 * time.Minute)
}
func uninstallHelm() {
	fmt.Printf("helm uninstall hwameistor\n")
	_ = runInLinux("helm list -A | grep 'hwameistor' | awk '{print $1}' | xargs helm uninstall -n hwameistor")
	fmt.Printf("clean all hwameistor crd\n")
	_ = runInLinux("kubectl get crd | grep 'hwameistor' | awk '{print $1}' | xargs -n1 kubectl delete crd")
	fmt.Printf("waiting for uninstall hwameistor\n")
	time.Sleep(1 * time.Minute)

}

func createLdc() {
	fmt.Printf("create ldc for each node\n")
	nodelist := nodeList()
	for _, nodes := range nodelist.Items {
		f := framework.NewDefaultFramework(ldapis.AddToScheme)
		client := f.GetClient()
		exmlocalDiskClaim := &ldv1.LocalDiskClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "localdiskclaim-" + nodes.Name,
				Namespace: "kube-system",
			},
			Spec: ldv1.LocalDiskClaimSpec{
				NodeName: nodes.Name,
				Description: ldv1.DiskClaimDescription{
					DiskType: "HDD",
				},
			},
		}
		err := client.Create(context.TODO(), exmlocalDiskClaim)
		if err != nil {
			fmt.Printf("Create LDC failed ：%+v \n", err)
			f.ExpectNoError(err)
		}
	}
	fmt.Printf("wait 1 minutes for create ldc\n")
	time.Sleep(1 * time.Minute)

}

func deleteAllPVC() bool {
	fmt.Printf("delete All PVC\n")
	f := framework.NewDefaultFramework(ldapis.AddToScheme)
	client := f.GetClient()
	pvcList := &apiv1.PersistentVolumeClaimList{}
	err := client.List(context.TODO(), pvcList)
	if err != nil {
		fmt.Printf("get pvc list error:%+v \n", err)
		f.ExpectNoError(err)
	}
	waitTime := 0

	for len(pvcList.Items) != 0 {
		if waitTime < 90 {
			for _, pvc := range pvcList.Items {
				fmt.Printf("delete pvc:%+v \n", pvc.Name)
				err := client.Delete(context.TODO(), &pvc)
				if err != nil {
					fmt.Printf("delete pvc error:%+v \n", err)
					f.ExpectNoError(err)
				}
				time.Sleep(30 * time.Second)
			}
			pvcList = &apiv1.PersistentVolumeClaimList{}
			waitTime = waitTime + 5
			time.Sleep(5 * time.Second)

		} else {
			fmt.Printf("delete PVC out of time")
			return false
		}

	}
	return true

}

func deleteAllSC() bool {
	fmt.Printf("delete All SC\n")
	f := framework.NewDefaultFramework(ldapis.AddToScheme)
	client := f.GetClient()
	scList := &storagev1.StorageClassList{}
	err := client.List(context.TODO(), scList)
	if err != nil {
		fmt.Printf("get sc list error:%+v \n", err)
		f.ExpectNoError(err)
	}
	waitTime := 0

	for len(scList.Items) != 0 {
		if waitTime < 90 {
			for _, sc := range scList.Items {
				fmt.Printf("delete sc:%+v \n", sc.Name)
				err := client.Delete(context.TODO(), &sc)
				if err != nil {
					fmt.Printf("delete sc error:%+v \n", err)
					f.ExpectNoError(err)
				}
				time.Sleep(30 * time.Second)
			}
			scList = &storagev1.StorageClassList{}
			waitTime = waitTime + 5
			time.Sleep(5 * time.Second)

		} else {
			fmt.Printf("delete sc out of time")
			return false
		}

	}
	return true

}
