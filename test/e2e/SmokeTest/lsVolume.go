package SmokeTest

import (
	"context"
	"fmt"
	"github.com/niulechuan/e2e/pkg/apis"
	"github.com/niulechuan/e2e/test/e2e/framework"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"time"
)

var _ = ginkgo.Describe("volume", func() {

	ginkgo.Describe("dlocal test", func() {
		f := framework.NewDefaultFramework(apis.AddToScheme)
		client := f.GetClient()

		ginkgo.It("get ready", func() {
			installHelm()
			createLdc()
			addLabels()
		})
		ginkgo.Context("create a SC", func() {
			ginkgo.It("SC", func() {
				//create sc
				deleteObj := apiv1.PersistentVolumeReclaimDelete
				waitForFirstConsumerObj := storagev1.VolumeBindingWaitForFirstConsumer
				examplesc := &storagev1.StorageClass{
					ObjectMeta: metav1.ObjectMeta{
						Name: "local-storage-hdd-lvm",
					},
					Provisioner: "localstorage.hwameistor.io",
					Parameters: map[string]string{
						"replicaNumber":             "1",
						"poolClass":                 "HDD",
						"poolType":                  "REGULAR",
						"volumeKind":                "LVM",
						"striped":                   "true",
						"csi.storage.k8s.io/fstype": "xfs",
					},
					ReclaimPolicy:        &deleteObj,
					AllowVolumeExpansion: boolPter(true),
					VolumeBindingMode:    &waitForFirstConsumerObj,
				}
				err := client.Create(context.TODO(), examplesc)
				if err != nil {
					fmt.Printf("Create SC failed ：%+v \n", err)
					f.ExpectNoError(err)
				}
			})
		})
		ginkgo.Context("create a PVC", func() {
			ginkgo.It("PVC STATUS should be Pending", func() {
				//create PVC
				storageClassName := "local-storage-hdd-lvm"
				examplePvc := &apiv1.PersistentVolumeClaim{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pvc-lvm",
						Namespace: "default",
					},
					Spec: apiv1.PersistentVolumeClaimSpec{
						AccessModes:      []apiv1.PersistentVolumeAccessMode{apiv1.ReadWriteOnce},
						StorageClassName: &storageClassName,
						Resources: apiv1.ResourceRequirements{
							Requests: apiv1.ResourceList{
								apiv1.ResourceStorage: resource.MustParse("1Gi"),
							},
						},
					},
				}
				err := client.Create(context.TODO(), examplePvc)
				if err != nil {
					fmt.Printf("Create PVC failed ：%+v \n", err)
					f.ExpectNoError(err)
				}

				pvc := &apiv1.PersistentVolumeClaim{}
				pvcKey := k8sclient.ObjectKey{
					Name:      "pvc-lvm",
					Namespace: "default",
				}
				err = client.Get(context.TODO(), pvcKey, pvc)
				if err != nil {
					fmt.Printf("Failed to find pvc ：%+v \n", err)
					f.ExpectNoError(err)
				}
				gomega.Expect(pvc.Status.Phase).To(gomega.Equal(apiv1.ClaimPending))
			})

		})
		ginkgo.Context("create a deployment", func() {

			ginkgo.It("PVC STATUS should be Bound", func() {
				//create deployment
				exampleDeployment := &appsv1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "demo-2048",
						Namespace: "default",
					},
					Spec: appsv1.DeploymentSpec{
						Replicas: int32Ptr(1),
						Selector: &metav1.LabelSelector{
							MatchLabels: map[string]string{
								"app": "demo",
							},
						},
						Template: apiv1.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"app": "demo",
								},
							},
							Spec: apiv1.PodSpec{
								SchedulerName: "hwameistor-scheduler",
								Containers: []apiv1.Container{
									{
										Name:  "web",
										Image: "daocloud.io/daocloud/dao-2048:latest",
										Ports: []apiv1.ContainerPort{
											{
												Name:          "http",
												Protocol:      apiv1.ProtocolTCP,
												ContainerPort: 80,
											},
										},
										VolumeMounts: []apiv1.VolumeMount{
											{
												Name:      "2048-volume-lvm",
												MountPath: "/data",
											},
										},
									},
								},
								Volumes: []apiv1.Volume{
									{
										Name: "2048-volume-lvm",
										VolumeSource: apiv1.VolumeSource{
											PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
												ClaimName: "pvc-lvm",
											},
										},
									},
								},
							},
						},
					},
				}
				err := client.Create(context.TODO(), exampleDeployment)
				if err != nil {
					fmt.Printf("%+v \n", err)
					f.ExpectNoError(err)
				}
				time.Sleep(1 * time.Minute)
				pvc := &apiv1.PersistentVolumeClaim{}
				pvcKey := k8sclient.ObjectKey{
					Name:      "pvc-lvm",
					Namespace: "default",
				}
				err = client.Get(context.TODO(), pvcKey, pvc)
				if err != nil {
					fmt.Printf("%+v \n", err)
					f.ExpectNoError(err)
				}
				gomega.Expect(pvc.Status.Phase).To(gomega.Equal(apiv1.ClaimBound))
			})
			ginkgo.It("deploy STATUS should be AVAILABLE", func() {
				deployment := &appsv1.Deployment{}
				deployKey := k8sclient.ObjectKey{
					Name:      "demo-2048",
					Namespace: "default",
				}
				err := client.Get(context.TODO(), deployKey, deployment)
				if err != nil {
					fmt.Printf("%+v \n", err)
					f.ExpectNoError(err)
				}
				gomega.Expect(deployment.Status.AvailableReplicas).To(gomega.Equal(int32(1)))
			})

		})
		ginkgo.Context("Using volumes", func() {
			ginkgo.It("Write", func() {

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
						_, _, err := ExecInPod(config, deployment.Namespace, pod.Name, "cd /data && echo it-is-a-test >test", container.Name)
						if err != nil {
							fmt.Printf("%+v \n", err)
							f.ExpectNoError(err)
						}
						output, _, err := ExecInPod(config, deployment.Namespace, pod.Name, "cd /data && cat test", container.Name)
						if err != nil {
							fmt.Printf("%+v \n", err)
							f.ExpectNoError(err)
						}
						gomega.Expect(output).To(gomega.Equal("it-is-a-test"))
					}
				}
			})
			ginkgo.It("Delete", func() {
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
						_, _, err := ExecInPod(config, deployment.Namespace, pod.Name, "cd /data && rm -rf test", container.Name)
						if err != nil {
							fmt.Printf("%+v \n", err)
							f.ExpectNoError(err)
						}
						output, _, err := ExecInPod(config, deployment.Namespace, pod.Name, "cd /data && ls", container.Name)
						if err != nil {
							fmt.Printf("%+v \n", err)
							f.ExpectNoError(err)
						}
						gomega.Expect(output).To(gomega.Equal(""))
					}
				}
			})
		})
		ginkgo.Context("Delete test object", func() {
			ginkgo.It("Delete test object", func() {
				//delete deploy
				deployment := &appsv1.Deployment{}
				deployKey := k8sclient.ObjectKey{
					Name:      "demo-2048",
					Namespace: "default",
				}
				err := client.Get(context.TODO(), deployKey, deployment)
				if err != nil {
					fmt.Printf("%+v \n", err)
					f.ExpectNoError(err)
				}
				err = client.Delete(context.TODO(), deployment)
				if err != nil {
					fmt.Printf("%+v \n", err)
					f.ExpectNoError(err)
				}

			})
			ginkgo.It("delete all pvc ", func() {
				r := deleteAllPVC()
				gomega.Expect(r).To(gomega.Equal(true))
			})
			ginkgo.It("delete all sc", func() {
				r := deleteAllSC()
				gomega.Expect(r).To(gomega.Equal(true))
			})
			ginkgo.It("delete helm", func() {
				uninstallHelm()

			})
		})
	})
})
