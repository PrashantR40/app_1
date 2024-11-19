/*
Copyright k0s authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1beta1 "github.com/k0sproject/k0s/pkg/apis/k0s/v1beta1"
	scheme "github.com/k0sproject/k0s/pkg/client/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ClusterConfigsGetter has a method to return a ClusterConfigInterface.
// A group's client should implement this interface.
type ClusterConfigsGetter interface {
	ClusterConfigs(namespace string) ClusterConfigInterface
}

// ClusterConfigInterface has methods to work with ClusterConfig resources.
type ClusterConfigInterface interface {
	Create(ctx context.Context, clusterConfig *v1beta1.ClusterConfig, opts v1.CreateOptions) (*v1beta1.ClusterConfig, error)
	Update(ctx context.Context, clusterConfig *v1beta1.ClusterConfig, opts v1.UpdateOptions) (*v1beta1.ClusterConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ClusterConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ClusterConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	ClusterConfigExpansion
}

// clusterConfigs implements ClusterConfigInterface
type clusterConfigs struct {
	*gentype.ClientWithList[*v1beta1.ClusterConfig, *v1beta1.ClusterConfigList]
}

// newClusterConfigs returns a ClusterConfigs
func newClusterConfigs(c *K0sV1beta1Client, namespace string) *clusterConfigs {
	return &clusterConfigs{
		gentype.NewClientWithList[*v1beta1.ClusterConfig, *v1beta1.ClusterConfigList](
			"clusterconfigs",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.ClusterConfig { return &v1beta1.ClusterConfig{} },
			func() *v1beta1.ClusterConfigList { return &v1beta1.ClusterConfigList{} }),
	}
}
