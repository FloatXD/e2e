package createdeletevolume

import (
	"context"
	"fmt"
	ldv1 "github.com/hwameistor/local-disk-manager/pkg/apis/hwameistor/v1alpha1"
	"github.com/niulechuan/e2e/pkg/apis"
	"github.com/niulechuan/e2e/test/e2e/framework"
	. "github.com/onsi/ginkgo"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("volume", func() {
	f := framework.NewDefaultFramework(apis.AddToScheme)

	client := f.GetClient()

	LocalDisk := &ldv1.LocalDisk{}
	LocalDiskKey := k8sclient.ObjectKey{
		Name:      "k8s-node1-sdb",
		Namespace: "kube-system",
	}

	err := client.Get(context.TODO(), LocalDiskKey, LocalDisk)
	if err != nil {
		//f.ExpectNoError(err)
		fmt.Printf("%+v \n", err)
	}
	fmt.Printf("%+v \n", LocalDisk)

})
