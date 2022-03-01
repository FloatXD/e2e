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
	"time"
)

var _ = Describe("volume", func() {
	f := framework.NewDefaultFramework(ldapis.AddToScheme)
	client := f.GetClient()

	localDiskList := &ldv1.LocalDiskList{}
	err := client.List(context.TODO(), localDiskList)
	if err != nil {
		f.ExpectNoError(err)
		fmt.Printf("%+v \n", err)
	}

	Describe("LDM test", func() {
		Context("LD test", func() {
			localDiskNumber := 0
			It("Check existing LD", func() {
				for i, localDisk := range localDiskList.Items {
					fmt.Printf("%+v \n", localDisk.Name)
					localDiskNumber = i + 1
				}
				fmt.Printf("There are %d local volumes \n", localDiskNumber)
				Expect(localDiskNumber).ToNot(Equal(0))
			})
			It("Manage new disks", func() {

				newlocalDiskNumber := 0
				output := runInLinux("cd /root && sh adddisk.sh")
				fmt.Printf("wait 1 minute \n")
				time.Sleep(2 * time.Minute)
				localDiskList := &ldv1.LocalDiskList{}
				err := client.List(context.TODO(), localDiskList)
				if err != nil {
					f.ExpectNoError(err)
					fmt.Printf("%+v \n", err)
				}
				fmt.Printf("%+v \n", output)
				for i, localDisk := range localDiskList.Items {
					fmt.Printf("%+v \n", localDisk.Name)
					newlocalDiskNumber = i + 1
				}
				fmt.Printf("There are %d local volumes", newlocalDiskNumber)
				Expect(newlocalDiskNumber).ToNot(Equal(localDiskNumber))
			})
		})
		Context("LDC test", func() {
			It("Create new LDC", func() {
				exmlocalDiskClaim := &ldv1.LocalDiskClaim{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "localdiskclaim-node-1",
						Namespace: "kube-system",
					},
					Spec: ldv1.LocalDiskClaimSpec{
						NodeName: "k8s-master",
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
				time.Sleep(1 * time.Minute)
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

				fmt.Printf("%+v", localDiskClaim)
				Expect(localDiskClaim.Status.Status).To(Equal(ldv1.LocalDiskClaimStatusBound))
			})
		})

	})

})
