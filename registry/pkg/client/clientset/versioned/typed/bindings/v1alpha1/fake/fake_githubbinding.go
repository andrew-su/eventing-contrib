/*
Copyright 2020 The Knative Authors

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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/eventing-contrib/github/pkg/apis/bindings/v1alpha1"
)

// FakeGitHubBindings implements GitHubBindingInterface
type FakeGitHubBindings struct {
	Fake *FakeBindingsV1alpha1
	ns   string
}

var githubbindingsResource = schema.GroupVersionResource{Group: "bindings.knative.dev", Version: "v1alpha1", Resource: "githubbindings"}

var githubbindingsKind = schema.GroupVersionKind{Group: "bindings.knative.dev", Version: "v1alpha1", Kind: "RegistryBinding"}

// Get takes name of the gitHubBinding, and returns the corresponding gitHubBinding object, and an error if there is any.
func (c *FakeGitHubBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.GitHubBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(githubbindingsResource, c.ns, name), &v1alpha1.GitHubBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubBinding), err
}

// List takes label and field selectors, and returns the list of GitHubBindings that match those selectors.
func (c *FakeGitHubBindings) List(opts v1.ListOptions) (result *v1alpha1.GitHubBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(githubbindingsResource, githubbindingsKind, c.ns, opts), &v1alpha1.GitHubBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GitHubBindingList{ListMeta: obj.(*v1alpha1.GitHubBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.GitHubBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gitHubBindings.
func (c *FakeGitHubBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(githubbindingsResource, c.ns, opts))

}

// Create takes the representation of a gitHubBinding and creates it.  Returns the server's representation of the gitHubBinding, and an error, if there is any.
func (c *FakeGitHubBindings) Create(gitHubBinding *v1alpha1.GitHubBinding) (result *v1alpha1.GitHubBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(githubbindingsResource, c.ns, gitHubBinding), &v1alpha1.GitHubBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubBinding), err
}

// Update takes the representation of a gitHubBinding and updates it. Returns the server's representation of the gitHubBinding, and an error, if there is any.
func (c *FakeGitHubBindings) Update(gitHubBinding *v1alpha1.GitHubBinding) (result *v1alpha1.GitHubBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(githubbindingsResource, c.ns, gitHubBinding), &v1alpha1.GitHubBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGitHubBindings) UpdateStatus(gitHubBinding *v1alpha1.GitHubBinding) (*v1alpha1.GitHubBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(githubbindingsResource, "status", c.ns, gitHubBinding), &v1alpha1.GitHubBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubBinding), err
}

// Delete takes name of the gitHubBinding and deletes it. Returns an error if one occurs.
func (c *FakeGitHubBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(githubbindingsResource, c.ns, name), &v1alpha1.GitHubBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGitHubBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(githubbindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GitHubBindingList{})
	return err
}

// Patch applies the patch and returns the patched gitHubBinding.
func (c *FakeGitHubBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GitHubBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(githubbindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GitHubBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubBinding), err
}
