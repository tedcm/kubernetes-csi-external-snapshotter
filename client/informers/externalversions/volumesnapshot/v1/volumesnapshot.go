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

package v1

import (
	"context"
	time "time"

	volumesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	versioned "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned"
	internalinterfaces "github.com/kubernetes-csi/external-snapshotter/client/v4/informers/externalversions/internalinterfaces"
	v1 "github.com/kubernetes-csi/external-snapshotter/client/v4/listers/volumesnapshot/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VolumeSnapshotInformer provides access to a shared informer and lister for
// VolumeSnapshots.
type VolumeSnapshotInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.VolumeSnapshotLister
}

type volumeSnapshotInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVolumeSnapshotInformer constructs a new informer for VolumeSnapshot type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVolumeSnapshotInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVolumeSnapshotInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVolumeSnapshotInformer constructs a new informer for VolumeSnapshot type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVolumeSnapshotInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SnapshotV1().VolumeSnapshots(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SnapshotV1().VolumeSnapshots(namespace).Watch(context.TODO(), options)
			},
		},
		&volumesnapshotv1.VolumeSnapshot{},
		resyncPeriod,
		indexers,
	)
}

func (f *volumeSnapshotInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVolumeSnapshotInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *volumeSnapshotInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&volumesnapshotv1.VolumeSnapshot{}, f.defaultInformer)
}

func (f *volumeSnapshotInformer) Lister() v1.VolumeSnapshotLister {
	return v1.NewVolumeSnapshotLister(f.Informer().GetIndexer())
}
