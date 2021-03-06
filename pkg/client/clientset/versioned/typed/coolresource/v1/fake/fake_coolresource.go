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
	coolresourcev1 "github.com/mattgartman/k8s-testcrd/pkg/apis/coolresource/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCoolResources implements CoolResourceInterface
type FakeCoolResources struct {
	Fake *FakeMattgartmanV1
	ns   string
}

var coolresourcesResource = schema.GroupVersionResource{Group: "mattgartman.com", Version: "v1", Resource: "coolresources"}

var coolresourcesKind = schema.GroupVersionKind{Group: "mattgartman.com", Version: "v1", Kind: "CoolResource"}

// Get takes name of the coolResource, and returns the corresponding coolResource object, and an error if there is any.
func (c *FakeCoolResources) Get(name string, options v1.GetOptions) (result *coolresourcev1.CoolResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(coolresourcesResource, c.ns, name), &coolresourcev1.CoolResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coolresourcev1.CoolResource), err
}

// List takes label and field selectors, and returns the list of CoolResources that match those selectors.
func (c *FakeCoolResources) List(opts v1.ListOptions) (result *coolresourcev1.CoolResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(coolresourcesResource, coolresourcesKind, c.ns, opts), &coolresourcev1.CoolResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &coolresourcev1.CoolResourceList{ListMeta: obj.(*coolresourcev1.CoolResourceList).ListMeta}
	for _, item := range obj.(*coolresourcev1.CoolResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested coolResources.
func (c *FakeCoolResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(coolresourcesResource, c.ns, opts))

}

// Create takes the representation of a coolResource and creates it.  Returns the server's representation of the coolResource, and an error, if there is any.
func (c *FakeCoolResources) Create(coolResource *coolresourcev1.CoolResource) (result *coolresourcev1.CoolResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(coolresourcesResource, c.ns, coolResource), &coolresourcev1.CoolResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coolresourcev1.CoolResource), err
}

// Update takes the representation of a coolResource and updates it. Returns the server's representation of the coolResource, and an error, if there is any.
func (c *FakeCoolResources) Update(coolResource *coolresourcev1.CoolResource) (result *coolresourcev1.CoolResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(coolresourcesResource, c.ns, coolResource), &coolresourcev1.CoolResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coolresourcev1.CoolResource), err
}

// Delete takes name of the coolResource and deletes it. Returns an error if one occurs.
func (c *FakeCoolResources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(coolresourcesResource, c.ns, name), &coolresourcev1.CoolResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCoolResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(coolresourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &coolresourcev1.CoolResourceList{})
	return err
}

// Patch applies the patch and returns the patched coolResource.
func (c *FakeCoolResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *coolresourcev1.CoolResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(coolresourcesResource, c.ns, name, pt, data, subresources...), &coolresourcev1.CoolResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*coolresourcev1.CoolResource), err
}
