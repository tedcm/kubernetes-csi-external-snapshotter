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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/kubernetes-csi/external-snapshotter/client/v4/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// VolumeSnapshots returns a VolumeSnapshotInformer.
	VolumeSnapshots() VolumeSnapshotInformer
	// VolumeSnapshotClasses returns a VolumeSnapshotClassInformer.
	VolumeSnapshotClasses() VolumeSnapshotClassInformer
	// VolumeSnapshotContents returns a VolumeSnapshotContentInformer.
	VolumeSnapshotContents() VolumeSnapshotContentInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// VolumeSnapshots returns a VolumeSnapshotInformer.
func (v *version) VolumeSnapshots() VolumeSnapshotInformer {
	return &volumeSnapshotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VolumeSnapshotClasses returns a VolumeSnapshotClassInformer.
func (v *version) VolumeSnapshotClasses() VolumeSnapshotClassInformer {
	return &volumeSnapshotClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// VolumeSnapshotContents returns a VolumeSnapshotContentInformer.
func (v *version) VolumeSnapshotContents() VolumeSnapshotContentInformer {
	return &volumeSnapshotContentInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
