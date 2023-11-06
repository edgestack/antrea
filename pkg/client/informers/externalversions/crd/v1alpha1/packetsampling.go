// Copyright 2023 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	crdv1alpha1 "antrea.io/antrea/pkg/apis/crd/v1alpha1"
	versioned "antrea.io/antrea/pkg/client/clientset/versioned"
	internalinterfaces "antrea.io/antrea/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "antrea.io/antrea/pkg/client/listers/crd/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PacketSamplingInformer provides access to a shared informer and lister for
// PacketSamplings.
type PacketSamplingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PacketSamplingLister
}

type packetSamplingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPacketSamplingInformer constructs a new informer for PacketSampling type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPacketSamplingInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPacketSamplingInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPacketSamplingInformer constructs a new informer for PacketSampling type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPacketSamplingInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1alpha1().PacketSamplings().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1alpha1().PacketSamplings().Watch(context.TODO(), options)
			},
		},
		&crdv1alpha1.PacketSampling{},
		resyncPeriod,
		indexers,
	)
}

func (f *packetSamplingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPacketSamplingInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *packetSamplingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&crdv1alpha1.PacketSampling{}, f.defaultInformer)
}

func (f *packetSamplingInformer) Lister() v1alpha1.PacketSamplingLister {
	return v1alpha1.NewPacketSamplingLister(f.Informer().GetIndexer())
}
