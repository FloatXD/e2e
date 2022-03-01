package SmokeTest

import (
	"context"
	"fmt"
	ldapis "github.com/hwameistor/local-disk-manager/pkg/apis"
	ldv1 "github.com/hwameistor/local-disk-manager/pkg/apis/hwameistor/v1alpha1"
	_ "github.com/niulechuan/e2e/pkg/apis"
	"github.com/niulechuan/e2e/test/e2e/framework"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("volume", func() {
	f := framework.NewDefaultFramework(ldapis.AddToScheme)
	client := f.GetClient()

	//LocalDisk := &ldv1.LocalDisk{}
	//LocalDiskKey := k8sclient.ObjectKey{
	//	Name:      "k8s-node1-sdb",
	//	Namespace: "kube-system",
	//}
	//
	//err := client.Get(context.TODO(), LocalDiskKey, LocalDisk)
	//if err != nil {
	//	f.ExpectNoError(err)
	//	fmt.Printf("%+v \n", err)
	//}
	localDiskList := &ldv1.LocalDiskList{}
	err := client.List(context.TODO(), localDiskList)
	if err != nil {
		f.ExpectNoError(err)
		fmt.Printf("%+v \n", err)
	}

	Describe("LDM test", func() {
		//Context("LD test", func() {
		//	It("Check existing LD", func() {
		//		localDiskNumber := 0
		//		for i, localDisk := range localDiskList.Items {
		//			fmt.Printf("%+v \n", localDisk.Name)
		//			localDiskNumber = i
		//		}
		//		fmt.Printf("There are %d local volumes", localDiskNumber)
		//	})
		//	It("Manage new disks", func() {
		//		localDiskNumber := 0
		//		output := runInLinux("cd /data && sh adddisk.sh")
		//		fmt.Printf("%+v \n", output)
		//		for i, localDisk := range localDiskList.Items {
		//			fmt.Printf("%+v \n", localDisk.Name)
		//			localDiskNumber = i
		//		}
		//		fmt.Printf("There are %d local volumes", localDiskNumber)
		//	})
		//})
		Context("LDC test", func() {
			It("Create new LDC", func() {
				exmlocalDiskClaim := &ldv1.LocalDiskClaim{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "localdiskclaim-node-1",
						Namespace: "kube-system",
					},
					Spec: ldv1.LocalDiskClaimSpec{
						NodeName: " k8s-node1",
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
				localDiskClaim := &ldv1.LocalDiskClaim{}
				localDiskClaimKey := k8sclient.ObjectKey{
					Name:      "localdiskclaim-node-1",
					Namespace: "kube-system",
				}
				err = client.Get(context.TODO(), localDiskClaimKey, localDiskClaim)
				if err != nil {
					fmt.Printf("%+v \n", err)
					f.ExpectNoError(err)
				}
				Expect(localDiskClaim.Status.Status).To(Equal(ldv1.LocalDiskClaimStatusBound))
			})
		})

	})

})
