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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	appawarecontrollerv1 "k8s.io/application-aware-controller/pkg/apis/appawarecontroller/v1"
	testing "k8s.io/client-go/testing"
)

// FakeResourceRecommendations implements ResourceRecommendationInterface
type FakeResourceRecommendations struct {
	Fake *FakeAppawarecontrollerV1
	ns   string
}

var resourcerecommendationsResource = schema.GroupVersionResource{Group: "appawarecontroller.k8s.io", Version: "v1", Resource: "resourcerecommendations"}

var resourcerecommendationsKind = schema.GroupVersionKind{Group: "appawarecontroller.k8s.io", Version: "v1", Kind: "ResourceRecommendation"}

// Get takes name of the resourceRecommendation, and returns the corresponding resourceRecommendation object, and an error if there is any.
func (c *FakeResourceRecommendations) Get(ctx context.Context, name string, options v1.GetOptions) (result *appawarecontrollerv1.ResourceRecommendation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourcerecommendationsResource, c.ns, name), &appawarecontrollerv1.ResourceRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appawarecontrollerv1.ResourceRecommendation), err
}

// List takes label and field selectors, and returns the list of ResourceRecommendations that match those selectors.
func (c *FakeResourceRecommendations) List(ctx context.Context, opts v1.ListOptions) (result *appawarecontrollerv1.ResourceRecommendationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourcerecommendationsResource, resourcerecommendationsKind, c.ns, opts), &appawarecontrollerv1.ResourceRecommendationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appawarecontrollerv1.ResourceRecommendationList{ListMeta: obj.(*appawarecontrollerv1.ResourceRecommendationList).ListMeta}
	for _, item := range obj.(*appawarecontrollerv1.ResourceRecommendationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceRecommendations.
func (c *FakeResourceRecommendations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourcerecommendationsResource, c.ns, opts))

}

// Create takes the representation of a resourceRecommendation and creates it.  Returns the server's representation of the resourceRecommendation, and an error, if there is any.
func (c *FakeResourceRecommendations) Create(ctx context.Context, resourceRecommendation *appawarecontrollerv1.ResourceRecommendation, opts v1.CreateOptions) (result *appawarecontrollerv1.ResourceRecommendation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourcerecommendationsResource, c.ns, resourceRecommendation), &appawarecontrollerv1.ResourceRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appawarecontrollerv1.ResourceRecommendation), err
}

// Update takes the representation of a resourceRecommendation and updates it. Returns the server's representation of the resourceRecommendation, and an error, if there is any.
func (c *FakeResourceRecommendations) Update(ctx context.Context, resourceRecommendation *appawarecontrollerv1.ResourceRecommendation, opts v1.UpdateOptions) (result *appawarecontrollerv1.ResourceRecommendation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourcerecommendationsResource, c.ns, resourceRecommendation), &appawarecontrollerv1.ResourceRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appawarecontrollerv1.ResourceRecommendation), err
}

// Delete takes name of the resourceRecommendation and deletes it. Returns an error if one occurs.
func (c *FakeResourceRecommendations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resourcerecommendationsResource, c.ns, name), &appawarecontrollerv1.ResourceRecommendation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceRecommendations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourcerecommendationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &appawarecontrollerv1.ResourceRecommendationList{})
	return err
}

// Patch applies the patch and returns the patched resourceRecommendation.
func (c *FakeResourceRecommendations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *appawarecontrollerv1.ResourceRecommendation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcerecommendationsResource, c.ns, name, pt, data, subresources...), &appawarecontrollerv1.ResourceRecommendation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appawarecontrollerv1.ResourceRecommendation), err
}
