/*
Copyright The Kubernetes Authors.

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
	v2alpha1 "github.com/solo-io/solo-kit/test/mocks/v2alpha1/kube/apis/testing.solo.io/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFrequentlyChangingAnnotationsResources implements FrequentlyChangingAnnotationsResourceInterface
type FakeFrequentlyChangingAnnotationsResources struct {
	Fake *FakeTestingV2alpha1
	ns   string
}

var frequentlychangingannotationsresourcesResource = schema.GroupVersionResource{Group: "testing.solo.io", Version: "v2alpha1", Resource: "fcars"}

var frequentlychangingannotationsresourcesKind = schema.GroupVersionKind{Group: "testing.solo.io", Version: "v2alpha1", Kind: "FrequentlyChangingAnnotationsResource"}

// Get takes name of the frequentlyChangingAnnotationsResource, and returns the corresponding frequentlyChangingAnnotationsResource object, and an error if there is any.
func (c *FakeFrequentlyChangingAnnotationsResources) Get(name string, options v1.GetOptions) (result *v2alpha1.FrequentlyChangingAnnotationsResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(frequentlychangingannotationsresourcesResource, c.ns, name), &v2alpha1.FrequentlyChangingAnnotationsResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.FrequentlyChangingAnnotationsResource), err
}

// List takes label and field selectors, and returns the list of FrequentlyChangingAnnotationsResources that match those selectors.
func (c *FakeFrequentlyChangingAnnotationsResources) List(opts v1.ListOptions) (result *v2alpha1.FrequentlyChangingAnnotationsResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(frequentlychangingannotationsresourcesResource, frequentlychangingannotationsresourcesKind, c.ns, opts), &v2alpha1.FrequentlyChangingAnnotationsResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.FrequentlyChangingAnnotationsResourceList{ListMeta: obj.(*v2alpha1.FrequentlyChangingAnnotationsResourceList).ListMeta}
	for _, item := range obj.(*v2alpha1.FrequentlyChangingAnnotationsResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested frequentlyChangingAnnotationsResources.
func (c *FakeFrequentlyChangingAnnotationsResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(frequentlychangingannotationsresourcesResource, c.ns, opts))

}

// Create takes the representation of a frequentlyChangingAnnotationsResource and creates it.  Returns the server's representation of the frequentlyChangingAnnotationsResource, and an error, if there is any.
func (c *FakeFrequentlyChangingAnnotationsResources) Create(frequentlyChangingAnnotationsResource *v2alpha1.FrequentlyChangingAnnotationsResource) (result *v2alpha1.FrequentlyChangingAnnotationsResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(frequentlychangingannotationsresourcesResource, c.ns, frequentlyChangingAnnotationsResource), &v2alpha1.FrequentlyChangingAnnotationsResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.FrequentlyChangingAnnotationsResource), err
}

// Update takes the representation of a frequentlyChangingAnnotationsResource and updates it. Returns the server's representation of the frequentlyChangingAnnotationsResource, and an error, if there is any.
func (c *FakeFrequentlyChangingAnnotationsResources) Update(frequentlyChangingAnnotationsResource *v2alpha1.FrequentlyChangingAnnotationsResource) (result *v2alpha1.FrequentlyChangingAnnotationsResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(frequentlychangingannotationsresourcesResource, c.ns, frequentlyChangingAnnotationsResource), &v2alpha1.FrequentlyChangingAnnotationsResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.FrequentlyChangingAnnotationsResource), err
}

// Delete takes name of the frequentlyChangingAnnotationsResource and deletes it. Returns an error if one occurs.
func (c *FakeFrequentlyChangingAnnotationsResources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(frequentlychangingannotationsresourcesResource, c.ns, name), &v2alpha1.FrequentlyChangingAnnotationsResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFrequentlyChangingAnnotationsResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(frequentlychangingannotationsresourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v2alpha1.FrequentlyChangingAnnotationsResourceList{})
	return err
}

// Patch applies the patch and returns the patched frequentlyChangingAnnotationsResource.
func (c *FakeFrequentlyChangingAnnotationsResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2alpha1.FrequentlyChangingAnnotationsResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(frequentlychangingannotationsresourcesResource, c.ns, name, pt, data, subresources...), &v2alpha1.FrequentlyChangingAnnotationsResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.FrequentlyChangingAnnotationsResource), err
}