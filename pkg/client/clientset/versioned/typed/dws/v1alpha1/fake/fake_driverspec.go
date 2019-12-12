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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "stash.us.cray.com/dpm/dws-operator/pkg/apis/dws/v1alpha1"
)

// FakeDriverSpecs implements DriverSpecInterface
type FakeDriverSpecs struct {
	Fake *FakeDwsV1alpha1
	ns   string
}

var driverspecsResource = schema.GroupVersionResource{Group: "dws.cray.com", Version: "v1alpha1", Resource: "driverspecs"}

var driverspecsKind = schema.GroupVersionKind{Group: "dws.cray.com", Version: "v1alpha1", Kind: "DriverSpec"}

// Get takes name of the driverSpec, and returns the corresponding driverSpec object, and an error if there is any.
func (c *FakeDriverSpecs) Get(name string, options v1.GetOptions) (result *v1alpha1.DriverSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(driverspecsResource, c.ns, name), &v1alpha1.DriverSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DriverSpec), err
}

// List takes label and field selectors, and returns the list of DriverSpecs that match those selectors.
func (c *FakeDriverSpecs) List(opts v1.ListOptions) (result *v1alpha1.DriverSpecList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(driverspecsResource, driverspecsKind, c.ns, opts), &v1alpha1.DriverSpecList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DriverSpecList), err
}

// Watch returns a watch.Interface that watches the requested driverSpecs.
func (c *FakeDriverSpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(driverspecsResource, c.ns, opts))

}

// Create takes the representation of a driverSpec and creates it.  Returns the server's representation of the driverSpec, and an error, if there is any.
func (c *FakeDriverSpecs) Create(driverSpec *v1alpha1.DriverSpec) (result *v1alpha1.DriverSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(driverspecsResource, c.ns, driverSpec), &v1alpha1.DriverSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DriverSpec), err
}

// Update takes the representation of a driverSpec and updates it. Returns the server's representation of the driverSpec, and an error, if there is any.
func (c *FakeDriverSpecs) Update(driverSpec *v1alpha1.DriverSpec) (result *v1alpha1.DriverSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(driverspecsResource, c.ns, driverSpec), &v1alpha1.DriverSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DriverSpec), err
}

// Delete takes name of the driverSpec and deletes it. Returns an error if one occurs.
func (c *FakeDriverSpecs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(driverspecsResource, c.ns, name), &v1alpha1.DriverSpec{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDriverSpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(driverspecsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DriverSpecList{})
	return err
}

// Patch applies the patch and returns the patched driverSpec.
func (c *FakeDriverSpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DriverSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(driverspecsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DriverSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DriverSpec), err
}
