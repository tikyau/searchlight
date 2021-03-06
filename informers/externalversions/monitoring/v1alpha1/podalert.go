/*
Copyright 2017 The Searchlight Authors.

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	monitoring_v1alpha1 "github.com/appscode/searchlight/apis/monitoring/v1alpha1"
	client "github.com/appscode/searchlight/client"
	internalinterfaces "github.com/appscode/searchlight/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/appscode/searchlight/listers/monitoring/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// PodAlertInformer provides access to a shared informer and lister for
// PodAlerts.
type PodAlertInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PodAlertLister
}

type podAlertInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newPodAlertInformer(client client.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.MonitoringV1alpha1().PodAlerts(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.MonitoringV1alpha1().PodAlerts(v1.NamespaceAll).Watch(options)
			},
		},
		&monitoring_v1alpha1.PodAlert{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *podAlertInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoring_v1alpha1.PodAlert{}, newPodAlertInformer)
}

func (f *podAlertInformer) Lister() v1alpha1.PodAlertLister {
	return v1alpha1.NewPodAlertLister(f.Informer().GetIndexer())
}
