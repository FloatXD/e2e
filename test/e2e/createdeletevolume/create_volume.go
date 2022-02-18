package createdeletevolume

import (
	"context"

	"github.com/niulechuan/e2e/pkg/apis"
	"github.com/niulechuan/e2e/pkg/apis/cache/v1alpha1"
	"github.com/niulechuan/e2e/test/e2e/framework"
	"github.com/onsi/ginkgo"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = ginkgo.Describe("Memcached", func() {
	f := framework.NewDefaultFramework(apis.AddToScheme)
	memcached := &v1alpha1.Memcached{}
	client := f.GetClient()
	backupKey := k8sclient.ObjectKey{
		Name:      "test-memcached",
		Namespace: "default",
	}

	if err := client.Get(context.TODO(), backupKey, memcached); err != nil {
		f.ExpectNoError(err)
	}
})
