/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package framework contains provider-independent helper code for
// building and running E2E tests with Ginkgo. The actual Ginkgo test
// suites gets assembled by combining this framework, the optional
// provider support code and specific tests via a separate .go file
// like Kubernetes' test/e2e.go.
package framework

import (
	"fmt"
	"strings"

	"github.com/onsi/ginkgo"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	cacheddiscovery "k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/dynamic"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
)

// Framework supports common operations used by e2e tests; it will keep a client & a namespace for you.
// Eventual goal is to merge this with integration test framework.
type Framework struct {
	clientConfig                     *rest.Config
	ClientSet                        clientset.Interface
	KubemarkExternalClusterClientSet clientset.Interface
	DynamicClient                    dynamic.Interface

	// configuration for framework's client
	Options Options

	// Timeouts contains the custom timeouts used during the test execution.
	Timeouts *TimeoutContext
}

// AfterEachActionFunc is a function that can be called after each test
type AfterEachActionFunc func(f *Framework, failed bool)

// Options is a struct for managing test framework options.
type Options struct {
	ClientQPS   float32
	ClientBurst int
}

// NewFrameworkWithCustomTimeouts makes a framework with with custom timeouts.
func NewFrameworkWithCustomTimeouts(baseName string, timeouts *TimeoutContext) *Framework {
	f := NewDefaultFramework(baseName)
	f.Timeouts = timeouts
	return f
}

// NewDefaultFramework makes a new framework and sets up a BeforeEach/AfterEach for
// you (you can write additional before/after each functions).
func NewDefaultFramework(baseName string) *Framework {
	options := Options{
		ClientQPS:   20,
		ClientBurst: 50,
	}
	return NewFramework(baseName, options, nil)
}

// NewFramework creates a test framework.
func NewFramework(baseName string, options Options, client clientset.Interface) *Framework {
	f := &Framework{
		Options:   options,
		ClientSet: client,
		Timeouts:  NewTimeoutContextWithDefaults(),
	}

	ginkgo.BeforeEach(f.BeforeEach)

	return f
}

// LoadConfig returns a config for a rest client with the UserAgent set to include the current test name.
func LoadConfig() (config *restclient.Config, err error) {
	defer func() {
		if err == nil && config != nil {
			testDesc := ginkgo.CurrentGinkgoTestDescription()
			if len(testDesc.ComponentTexts) > 0 {
				componentTexts := strings.Join(testDesc.ComponentTexts, " ")
				config.UserAgent = fmt.Sprintf("%s -- %s", rest.DefaultKubernetesUserAgent(), componentTexts)
			}
		}
	}()

	return restclient.InClusterConfig()
}

// BeforeEach gets a client and makes a namespace.
func (f *Framework) BeforeEach() {
	if f.ClientSet == nil {
		ginkgo.By("Creating a kubernetes client")
		config, err := LoadConfig()
		ExpectNoError(err)

		config.QPS = f.Options.ClientQPS
		config.Burst = f.Options.ClientBurst
		f.clientConfig = rest.CopyConfig(config)
		f.ClientSet, err = clientset.NewForConfig(config)
		ExpectNoError(err)
		f.DynamicClient, err = dynamic.NewForConfig(config)
		ExpectNoError(err)

		// create scales getter, set GroupVersion and NegotiatedSerializer to default values
		// as they are required when creating a REST client.
		if config.GroupVersion == nil {
			config.GroupVersion = &schema.GroupVersion{}
		}
		if config.NegotiatedSerializer == nil {
			config.NegotiatedSerializer = scheme.Codecs
		}
		ExpectNoError(err)
		discoClient, err := discovery.NewDiscoveryClientForConfig(config)
		ExpectNoError(err)
		cachedDiscoClient := cacheddiscovery.NewMemCacheClient(discoClient)
		restMapper := restmapper.NewDeferredDiscoveryRESTMapper(cachedDiscoClient)
		restMapper.Reset()
	}

}

// ClientConfig an externally accessible method for reading the kube client config.
func (f *Framework) ClientConfig() *rest.Config {
	ret := rest.CopyConfig(f.clientConfig)
	// json is least common denominator
	ret.ContentType = runtime.ContentTypeJSON
	ret.AcceptContentTypes = runtime.ContentTypeJSON
	return ret
}
