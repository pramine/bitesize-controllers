/*
Copyright 2018 Pearson Technology

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
package v1

import (
	v1 "github.com/pearsontechnology/bitesize-controllers/vault-controller/pkg/apis/vaultpolicy/v1"
	scheme "github.com/pearsontechnology/bitesize-controllers/vault-controller/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PoliciesGetter has a method to return a PolicyInterface.
// A group's client should implement this interface.
type PoliciesGetter interface {
	Policies() PolicyInterface
}

// PolicyInterface has methods to work with Policy resources.
type PolicyInterface interface {
	Create(*v1.Policy) (*v1.Policy, error)
	Update(*v1.Policy) (*v1.Policy, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Policy, error)
	List(opts meta_v1.ListOptions) (*v1.PolicyList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Policy, err error)
	PolicyExpansion
}

// policies implements PolicyInterface
type policies struct {
	client rest.Interface
}

// newPolicies returns a Policies
func newPolicies(c *VaultpolicyV1Client) *policies {
	return &policies{
		client: c.RESTClient(),
	}
}

// Get takes name of the policy, and returns the corresponding policy object, and an error if there is any.
func (c *policies) Get(name string, options meta_v1.GetOptions) (result *v1.Policy, err error) {
	result = &v1.Policy{}
	err = c.client.Get().
		Resource("policies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Policies that match those selectors.
func (c *policies) List(opts meta_v1.ListOptions) (result *v1.PolicyList, err error) {
	result = &v1.PolicyList{}
	err = c.client.Get().
		Resource("policies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested policies.
func (c *policies) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("policies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a policy and creates it.  Returns the server's representation of the policy, and an error, if there is any.
func (c *policies) Create(policy *v1.Policy) (result *v1.Policy, err error) {
	result = &v1.Policy{}
	err = c.client.Post().
		Resource("policies").
		Body(policy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a policy and updates it. Returns the server's representation of the policy, and an error, if there is any.
func (c *policies) Update(policy *v1.Policy) (result *v1.Policy, err error) {
	result = &v1.Policy{}
	err = c.client.Put().
		Resource("policies").
		Name(policy.Name).
		Body(policy).
		Do().
		Into(result)
	return
}

// Delete takes name of the policy and deletes it. Returns an error if one occurs.
func (c *policies) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("policies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *policies) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Resource("policies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched policy.
func (c *policies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Policy, err error) {
	result = &v1.Policy{}
	err = c.client.Patch(pt).
		Resource("policies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
