/*
Copyright 2022 The Kubernetes Authors.

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

	volumesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVolumeSnapshotContents implements VolumeSnapshotContentInterface
type FakeVolumeSnapshotContents struct {
	Fake *FakeSnapshotV1
}

var volumesnapshotcontentsResource = schema.GroupVersionResource{Group: "snapshot.storage.k8s.io", Version: "v1", Resource: "volumesnapshotcontents"}

var volumesnapshotcontentsKind = schema.GroupVersionKind{Group: "snapshot.storage.k8s.io", Version: "v1", Kind: "VolumeSnapshotContent"}

// Get takes name of the volumeSnapshotContent, and returns the corresponding volumeSnapshotContent object, and an error if there is any.
func (c *FakeVolumeSnapshotContents) Get(ctx context.Context, name string, options v1.GetOptions) (result *volumesnapshotv1.VolumeSnapshotContent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(volumesnapshotcontentsResource, name), &volumesnapshotv1.VolumeSnapshotContent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshotContent), err
}

// List takes label and field selectors, and returns the list of VolumeSnapshotContents that match those selectors.
func (c *FakeVolumeSnapshotContents) List(ctx context.Context, opts v1.ListOptions) (result *volumesnapshotv1.VolumeSnapshotContentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(volumesnapshotcontentsResource, volumesnapshotcontentsKind, opts), &volumesnapshotv1.VolumeSnapshotContentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &volumesnapshotv1.VolumeSnapshotContentList{ListMeta: obj.(*volumesnapshotv1.VolumeSnapshotContentList).ListMeta}
	for _, item := range obj.(*volumesnapshotv1.VolumeSnapshotContentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeSnapshotContents.
func (c *FakeVolumeSnapshotContents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(volumesnapshotcontentsResource, opts))
}

// Create takes the representation of a volumeSnapshotContent and creates it.  Returns the server's representation of the volumeSnapshotContent, and an error, if there is any.
func (c *FakeVolumeSnapshotContents) Create(ctx context.Context, volumeSnapshotContent *volumesnapshotv1.VolumeSnapshotContent, opts v1.CreateOptions) (result *volumesnapshotv1.VolumeSnapshotContent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(volumesnapshotcontentsResource, volumeSnapshotContent), &volumesnapshotv1.VolumeSnapshotContent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshotContent), err
}

// Update takes the representation of a volumeSnapshotContent and updates it. Returns the server's representation of the volumeSnapshotContent, and an error, if there is any.
func (c *FakeVolumeSnapshotContents) Update(ctx context.Context, volumeSnapshotContent *volumesnapshotv1.VolumeSnapshotContent, opts v1.UpdateOptions) (result *volumesnapshotv1.VolumeSnapshotContent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(volumesnapshotcontentsResource, volumeSnapshotContent), &volumesnapshotv1.VolumeSnapshotContent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshotContent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVolumeSnapshotContents) UpdateStatus(ctx context.Context, volumeSnapshotContent *volumesnapshotv1.VolumeSnapshotContent, opts v1.UpdateOptions) (*volumesnapshotv1.VolumeSnapshotContent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(volumesnapshotcontentsResource, "status", volumeSnapshotContent), &volumesnapshotv1.VolumeSnapshotContent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshotContent), err
}

// Delete takes name of the volumeSnapshotContent and deletes it. Returns an error if one occurs.
func (c *FakeVolumeSnapshotContents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(volumesnapshotcontentsResource, name, opts), &volumesnapshotv1.VolumeSnapshotContent{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeSnapshotContents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(volumesnapshotcontentsResource, listOpts)

	_, err := c.Fake.Invokes(action, &volumesnapshotv1.VolumeSnapshotContentList{})
	return err
}

// Patch applies the patch and returns the patched volumeSnapshotContent.
func (c *FakeVolumeSnapshotContents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *volumesnapshotv1.VolumeSnapshotContent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(volumesnapshotcontentsResource, name, pt, data, subresources...), &volumesnapshotv1.VolumeSnapshotContent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshotContent), err
}
