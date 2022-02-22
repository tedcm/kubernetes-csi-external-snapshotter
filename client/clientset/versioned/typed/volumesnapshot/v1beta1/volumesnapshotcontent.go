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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1beta1"
	scheme "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumeSnapshotContentsGetter has a method to return a VolumeSnapshotContentInterface.
// A group's client should implement this interface.
type VolumeSnapshotContentsGetter interface {
	VolumeSnapshotContents() VolumeSnapshotContentInterface
}

// VolumeSnapshotContentInterface has methods to work with VolumeSnapshotContent resources.
type VolumeSnapshotContentInterface interface {
	Create(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.CreateOptions) (*v1beta1.VolumeSnapshotContent, error)
	Update(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.UpdateOptions) (*v1beta1.VolumeSnapshotContent, error)
	UpdateStatus(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.UpdateOptions) (*v1beta1.VolumeSnapshotContent, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VolumeSnapshotContent, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VolumeSnapshotContentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeSnapshotContent, err error)
	VolumeSnapshotContentExpansion
}

// volumeSnapshotContents implements VolumeSnapshotContentInterface
type volumeSnapshotContents struct {
	client rest.Interface
}

// newVolumeSnapshotContents returns a VolumeSnapshotContents
func newVolumeSnapshotContents(c *SnapshotV1beta1Client) *volumeSnapshotContents {
	return &volumeSnapshotContents{
		client: c.RESTClient(),
	}
}

// Get takes name of the volumeSnapshotContent, and returns the corresponding volumeSnapshotContent object, and an error if there is any.
func (c *volumeSnapshotContents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeSnapshotContent, err error) {
	result = &v1beta1.VolumeSnapshotContent{}
	err = c.client.Get().
		Resource("volumesnapshotcontents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeSnapshotContents that match those selectors.
func (c *volumeSnapshotContents) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeSnapshotContentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VolumeSnapshotContentList{}
	err = c.client.Get().
		Resource("volumesnapshotcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeSnapshotContents.
func (c *volumeSnapshotContents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("volumesnapshotcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a volumeSnapshotContent and creates it.  Returns the server's representation of the volumeSnapshotContent, and an error, if there is any.
func (c *volumeSnapshotContents) Create(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.CreateOptions) (result *v1beta1.VolumeSnapshotContent, err error) {
	result = &v1beta1.VolumeSnapshotContent{}
	err = c.client.Post().
		Resource("volumesnapshotcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshotContent).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a volumeSnapshotContent and updates it. Returns the server's representation of the volumeSnapshotContent, and an error, if there is any.
func (c *volumeSnapshotContents) Update(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.UpdateOptions) (result *v1beta1.VolumeSnapshotContent, err error) {
	result = &v1beta1.VolumeSnapshotContent{}
	err = c.client.Put().
		Resource("volumesnapshotcontents").
		Name(volumeSnapshotContent.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshotContent).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *volumeSnapshotContents) UpdateStatus(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.UpdateOptions) (result *v1beta1.VolumeSnapshotContent, err error) {
	result = &v1beta1.VolumeSnapshotContent{}
	err = c.client.Put().
		Resource("volumesnapshotcontents").
		Name(volumeSnapshotContent.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshotContent).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the volumeSnapshotContent and deletes it. Returns an error if one occurs.
func (c *volumeSnapshotContents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("volumesnapshotcontents").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeSnapshotContents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("volumesnapshotcontents").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched volumeSnapshotContent.
func (c *volumeSnapshotContents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeSnapshotContent, err error) {
	result = &v1beta1.VolumeSnapshotContent{}
	err = c.client.Patch(pt).
		Resource("volumesnapshotcontents").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
