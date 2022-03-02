package SmokeTest

import (
	"context"
	"fmt"
	"github.com/niulechuan/e2e/pkg/apis"
	"github.com/niulechuan/e2e/test/e2e/framework"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
	"time"
)

var _ = Describe("volume", func() {

	f := framework.NewDefaultFramework(apis.AddToScheme)
	client := f.GetClient()

	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//flag.Parse()
	//
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	panic(err)
	//}
	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err)
	//}
	//
	//req := clientset.CoreV1().RESTClient().Post().
	//	Resource("pods").
	//	Name(podName).
	//	Namespace(namespace).SubResource("exec").Param("container", containerName)

	Describe("ha-dlocal test", func() {
		Context("create a HA-SC", func() {
			It("SC", func() {
				//create sc
				deleteObj := apiv1.PersistentVolumeReclaimDelete
				waitForFirstConsumerObj := storagev1.VolumeBindingWaitForFirstConsumer
				examplesc := &storagev1.StorageClass{
					ObjectMeta: metav1.ObjectMeta{
						Name: "local-storage-hdd-lvm-ha",
					},
					Provisioner: "local.storage.daocloud.io",
					Parameters: map[string]string{
						"replicaNumber":             "2",
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
		Context("create a HA-PVC", func() {
			It("PVC STATUS should be Pending", func() {
				//create PVC
				storageClassName := "local-storage-hdd-lvm-ha"
				examplePvc := &apiv1.PersistentVolumeClaim{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pvc-lvm-ha",
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
					Name:      "pvc-lvm-ha",
					Namespace: "default",
				}
				err = client.Get(context.TODO(), pvcKey, pvc)
				if err != nil {
					fmt.Printf("Failed to find pvc ：%+v \n", err)
					f.ExpectNoError(err)
				}
				Expect(pvc.Status.Phase).To(Equal(apiv1.ClaimPending))
			})

		})
		Context("create a deployment", func() {

			It("PVC STATUS should be Bound", func() {
				//create deployment
				exampleDeployment := &appsv1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "demo-2048-ha",
						Namespace: "default",
					},
					Spec: appsv1.DeploymentSpec{
						Replicas: int32Ptr(1),
						Strategy: appsv1.DeploymentStrategy{
							Type: appsv1.RecreateDeploymentStrategyType,
						},
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
								SchedulerName: "cherry-io-scheduler",
								Affinity: &apiv1.Affinity{
									NodeAffinity: &apiv1.NodeAffinity{
										RequiredDuringSchedulingIgnoredDuringExecution: &apiv1.NodeSelector{
											NodeSelectorTerms: []apiv1.NodeSelectorTerm{
												{
													[]apiv1.NodeSelectorRequirement{
														{
															Key:      "kubernetes.io/hostname",
															Operator: apiv1.NodeSelectorOpIn,
															Values: []string{
																"k8s-node1",
															},
														},
													},
													[]apiv1.NodeSelectorRequirement{},
												},
											},
										},
										PreferredDuringSchedulingIgnoredDuringExecution: nil,
									},
									PodAffinity:     nil,
									PodAntiAffinity: nil,
								},
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
												Name:      "2048-volume-lvm-ha",
												MountPath: "/data",
											},
										},
									},
								},
								Volumes: []apiv1.Volume{
									{
										Name: "2048-volume-lvm-ha",
										VolumeSource: apiv1.VolumeSource{
											PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
												ClaimName: "pvc-lvm-ha",
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
					Name:      "pvc-lvm-ha",
					Namespace: "default",
				}
				err = client.Get(context.TODO(), pvcKey, pvc)
				if err != nil {
					fmt.Printf("%+v \n", err)
					f.ExpectNoError(err)
				}
				Expect(pvc.Status.Phase).To(Equal(apiv1.ClaimBound))
			})
			It("deploy STATUS should be AVAILABLE", func() {
				deployment := &appsv1.Deployment{}
				deployKey := k8sclient.ObjectKey{
					Name:      "demo-2048-ha",
					Namespace: "default",
				}
				err := client.Get(context.TODO(), deployKey, deployment)
				if err != nil {
					fmt.Printf("%+v \n", err)
					f.ExpectNoError(err)
				}
				Expect(deployment.Status.AvailableReplicas).To(Equal(int32(1)))
			})

		})
		Context("Using volumes", func() {
			It("Write", func() {
				//create a request
				output := runInLinux("kubectl get pod |grep demo-2048")
				containerId := strings.Split(output, "   ")[0]
				output = runInLinux("kubectl exec " + containerId + " -- sh -c \"cd /data && echo it-is-a-test >test\"")
				output = runInLinux("kubectl exec " + containerId + " -- sh -c \"cd /data && cat test\"")
				Expect(output).To(Equal("it-is-a-test\n"))
			})
			It("Delete", func() {
				output := runInLinux("kubectl get pod |grep demo-2048")
				containerId := strings.Split(output, "   ")[0]
				output = runInLinux("kubectl exec " + containerId + " -- sh -c \"cd /data && rm -rf test\"")
				output = runInLinux("kubectl exec " + containerId + " -- sh -c \"cd /data && ls \"")
				Expect(output).To(Equal(""))
			})
		})
		Context("HA test", func() {
			It("Write test file", func() {

				output := runInLinux("kubectl get pod |grep demo-2048")
				containerId := strings.Split(output, "   ")[0]
				output = runInLinux("kubectl exec " + containerId + " -- sh -c \"cd /data && echo it-is-a-test >test\"")
				output = runInLinux("kubectl exec " + containerId + " -- sh -c \"cd /data && cat test\"")
				Expect(output).To(Equal("it-is-a-test\n"))
			})
			It("update deploy", func() {
				//delete deploy
				deployment := &appsv1.Deployment{}
				deployKey := k8sclient.ObjectKey{
					Name:      "demo-2048",
					Namespace: "default",
				}
				err := client.Get(context.TODO(), deployKey, deployment)
				if err != nil {
					f.ExpectNoError(err)
				}

				newAffinity := []apiv1.NodeSelectorTerm{
					{
						[]apiv1.NodeSelectorRequirement{
							{
								Key:      "kubernetes.io/hostname",
								Operator: apiv1.NodeSelectorOpIn,
								Values: []string{
									"k8s-master",
								},
							},
						},
						[]apiv1.NodeSelectorRequirement{},
					},
				}
				deployment.Spec.Template.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms = newAffinity

				err = client.Update(context.TODO(), deployment)
				fmt.Printf("wait 1 minute")
				time.Sleep(1 * time.Minute)
				err = client.Get(context.TODO(), deployKey, deployment)
				if err != nil {
					fmt.Printf("%+v \n", err)
					f.ExpectNoError(err)
				}
				Expect(deployment.Status.AvailableReplicas).To(Equal(int32(1)))
			})
			It("check test file", func() {
				//delete deploy
				output := runInLinux("kubectl get pod |grep demo-2048")
				containerId := strings.Split(output, "   ")[0]
				output = runInLinux("kubectl exec " + containerId + " -- sh -c \"cd /data && cat test\"")
				Expect(output).To(Equal("it-is-a-test\n"))
			})
		})
		Context("Delete test object", func() {
			It("Delete test object", func() {
				//delete deploy
				deployment := &appsv1.Deployment{}
				deployKey := k8sclient.ObjectKey{
					Name:      "demo-2048-ha",
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

				//delete PVC
				pvc := &apiv1.PersistentVolumeClaim{}
				pvcKey := k8sclient.ObjectKey{
					Name:      "pvc-lvm-ha",
					Namespace: "default",
				}
				err = client.Get(context.TODO(), pvcKey, pvc)
				if err != nil {
					fmt.Printf("Failed to find pvc ：%+v \n", err)
					f.ExpectNoError(err)
				}
				err = client.Delete(context.TODO(), pvc)
				if err != nil {
					fmt.Printf("Failed to delete pvc ：%+v \n", err)
					f.ExpectNoError(err)
				}

				//delete SC
				sc := &storagev1.StorageClass{}
				scKey := k8sclient.ObjectKey{
					Name: "local-storage-hdd-lvm-ha",
				}
				err = client.Get(context.TODO(), scKey, sc)
				if err != nil {
					f.ExpectNoError(err)
				}
				err = client.Delete(context.TODO(), sc)
				if err != nil {
					f.ExpectNoError(err)
				}

				time.Sleep(1 * time.Minute)

			})

		})
	})
})
