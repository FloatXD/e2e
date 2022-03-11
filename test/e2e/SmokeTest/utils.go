package SmokeTest

import (
	"bytes"
	"context"
	ldapis "github.com/hwameistor/local-disk-manager/pkg/apis"
	ldv1 "github.com/hwameistor/local-disk-manager/pkg/apis/hwameistor/v1alpha1"
	"github.com/sirupsen/logrus"

	"github.com/niulechuan/e2e/test/e2e/framework"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"os/exec"
	"regexp"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
	"time"
)

func int32Ptr(i int32) *int32 { return &i }

func boolPter(i bool) *bool { return &i }

func runInLinux(cmd string) string {
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		logrus.Printf("ERROR:%+v \n", err)
	}
	return string(result)
}

func nodeList() *apiv1.NodeList {
	logrus.Printf("get node list\n")
	f := framework.NewDefaultFramework(ldapis.AddToScheme)
	client := f.GetClient()
	nodelist := &apiv1.NodeList{}
	err := client.List(context.TODO(), nodelist)
	if err != nil {
		f.ExpectNoError(err)
		logrus.Printf("%+v \n", err)
	}
	return nodelist
}

func addLabels() {
	logrus.Printf("add node labels\n")
	f := framework.NewDefaultFramework(ldapis.AddToScheme)
	client := f.GetClient()
	nodelist := &apiv1.NodeList{}
	err := client.List(context.TODO(), nodelist)
	if err != nil {
		f.ExpectNoError(err)
		logrus.Printf("%+v \n", err)
	}
	for _, nodes := range nodelist.Items {
		node := &apiv1.Node{}
		nodeKey := k8sclient.ObjectKey{
			Name: nodes.Name,
		}
		err := client.Get(context.TODO(), nodeKey, node)
		if err != nil {
			logrus.Printf("%+v \n", err)
			f.ExpectNoError(err)
		}
		_, boolLabel := node.Labels["localstorage.hwameistor.io/local-storage"]
		if !boolLabel {
			node.Labels["localstorage.hwameistor.io/local-storage"] = "true"
			logrus.Printf("adding labels \n")
			err := client.Update(context.TODO(), node)
			if err != nil {
				logrus.Printf("%+v \n", err)
				f.ExpectNoError(err)
			}
			time.Sleep(20 * time.Second)
		}
		_, boolLabel = node.Labels["csi.driver.hwameistor.io/local-storage"]
		if !boolLabel {
			node.Labels["csi.driver.hwameistor.io/local-storage"] = "true"
			logrus.Printf("adding labels \n")
			err := client.Update(context.TODO(), node)
			if err != nil {
				logrus.Printf("%+v \n", err)
				f.ExpectNoError(err)
			}
			time.Sleep(20 * time.Second)
		}
		_, boolLabel = node.Labels["csi.driver.hwameistor.io/local-storage"]
		if !boolLabel {
			node.Labels["localstorage.hwameistor.io/local-storage-topology-node"] = nodes.Name
			logrus.Printf("adding labels \n")
			err := client.Update(context.TODO(), node)
			if err != nil {
				logrus.Printf("%+v \n", err)
				f.ExpectNoError(err)
			}
			time.Sleep(20 * time.Second)
		}
	}
}

func installHelm() {
	logrus.Printf("helm install hwameistor\n")
	_ = runInLinux("cd /root/helm-charts-main/charts && helm install hwameistor -n hwameistor --create-namespace --generate-name")
	logrus.Printf("waiting for intall hwameistor\n")
	time.Sleep(1 * time.Minute)
}
func uninstallHelm() {
	logrus.Printf("helm uninstall hwameistor\n")
	_ = runInLinux("helm list -A | grep 'hwameistor' | awk '{print $1}' | xargs helm uninstall -n hwameistor")
	logrus.Printf("clean all hwameistor crd\n")
	f := framework.NewDefaultFramework(extv1.AddToScheme)
	client := f.GetClient()
	crdList := extv1.CustomResourceDefinitionList{}
	err := client.List(context.TODO(), &crdList)
	if err != nil {
		logrus.Printf("%+v \n", err)
		f.ExpectNoError(err)
	}
	for _, crd := range crdList.Items {
		myBool, _ := regexp.MatchString(".*hwameistor.*", crd.Name)
		if myBool {
			err := client.Delete(context.TODO(), &crd)
			if err != nil {
				logrus.Printf("%+v \n", err)
				f.ExpectNoError(err)
			}
		}

	}
	logrus.Printf("waiting for uninstall hwameistor\n")
	time.Sleep(1 * time.Minute)

}

func createLdc() {
	logrus.Printf("create ldc for each node\n")
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
			logrus.Printf("Create LDC failed ：%+v \n", err)
			f.ExpectNoError(err)
		}
	}
	logrus.Printf("wait 1 minutes for create ldc\n")
	time.Sleep(1 * time.Minute)

}

func deleteAllPVC() {
	logrus.Printf("delete All PVC\n")
	f := framework.NewDefaultFramework(ldapis.AddToScheme)
	client := f.GetClient()
	pvcList := &apiv1.PersistentVolumeClaimList{}
	err := client.List(context.TODO(), pvcList)
	if err != nil {
		logrus.Printf("get pvc list error:%+v \n", err)
		f.ExpectNoError(err)
	}

	for _, pvc := range pvcList.Items {
		logrus.Printf("delete pvc:%+v \n", pvc.Name)
		ctx, _ := context.WithTimeout(context.TODO(), time.Minute)
		err := client.Delete(ctx, &pvc)
		if err != nil {
			logrus.Printf("delete pvc error:%+v \n", err)
			f.ExpectNoError(err)
		}
		time.Sleep(30 * time.Second)
	}

}

func deleteAllSC() {
	logrus.Printf("delete All SC\n")
	f := framework.NewDefaultFramework(ldapis.AddToScheme)
	client := f.GetClient()
	scList := &storagev1.StorageClassList{}
	err := client.List(context.TODO(), scList)
	if err != nil {
		logrus.Printf("get sc list error:%+v \n", err)
		f.ExpectNoError(err)
	}

	for _, sc := range scList.Items {
		logrus.Printf("delete sc:%+v \n", sc.Name)
		ctx, _ := context.WithTimeout(context.TODO(), time.Minute)
		err := client.Delete(ctx, &sc)
		if err != nil {
			logrus.Printf("delete sc error:%+v \n", err)
			f.ExpectNoError(err)
		}
		time.Sleep(30 * time.Second)
	}

}
func ExecInPod(config *rest.Config, namespace, podName, command, containerName string) (string, string, error) {
	k8sCli, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", "", err
	}
	cmd := []string{
		"sh",
		"-c",
		command,
	}
	const tty = false
	req := k8sCli.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).SubResource("exec").Param("container", containerName)
	req.VersionedParams(
		&v1.PodExecOptions{
			Command: cmd,
			Stdin:   false,
			Stdout:  true,
			Stderr:  true,
			TTY:     tty,
		},
		scheme.ParameterCodec,
	)

	var stdout, stderr bytes.Buffer
	myExec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return "", "", err
	}
	err = myExec.Stream(remotecommand.StreamOptions{
		Stdin:  nil,
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if err != nil {
		return "", "", err
	}
	return strings.TrimSpace(stdout.String()), strings.TrimSpace(stderr.String()), err
}
