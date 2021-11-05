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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "nanto.io/application-auto-scaling-service/pkg/apis/autoscaling/v1alpha1"
)

// CustomedHorizontalPodAutoscalerLister helps list CustomedHorizontalPodAutoscalers.
// All objects returned here must be treated as read-only.
type CustomedHorizontalPodAutoscalerLister interface {
	// List lists all CustomedHorizontalPodAutoscalers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CustomedHorizontalPodAutoscaler, err error)
	// CustomedHorizontalPodAutoscalers returns an object that can list and get CustomedHorizontalPodAutoscalers.
	CustomedHorizontalPodAutoscalers(namespace string) CustomedHorizontalPodAutoscalerNamespaceLister
	CustomedHorizontalPodAutoscalerListerExpansion
}

// customedHorizontalPodAutoscalerLister implements the CustomedHorizontalPodAutoscalerLister interface.
type customedHorizontalPodAutoscalerLister struct {
	indexer cache.Indexer
}

// NewCustomedHorizontalPodAutoscalerLister returns a new CustomedHorizontalPodAutoscalerLister.
func NewCustomedHorizontalPodAutoscalerLister(indexer cache.Indexer) CustomedHorizontalPodAutoscalerLister {
	return &customedHorizontalPodAutoscalerLister{indexer: indexer}
}

// List lists all CustomedHorizontalPodAutoscalers in the indexer.
func (s *customedHorizontalPodAutoscalerLister) List(selector labels.Selector) (ret []*v1alpha1.CustomedHorizontalPodAutoscaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CustomedHorizontalPodAutoscaler))
	})
	return ret, err
}

// CustomedHorizontalPodAutoscalers returns an object that can list and get CustomedHorizontalPodAutoscalers.
func (s *customedHorizontalPodAutoscalerLister) CustomedHorizontalPodAutoscalers(namespace string) CustomedHorizontalPodAutoscalerNamespaceLister {
	return customedHorizontalPodAutoscalerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CustomedHorizontalPodAutoscalerNamespaceLister helps list and get CustomedHorizontalPodAutoscalers.
// All objects returned here must be treated as read-only.
type CustomedHorizontalPodAutoscalerNamespaceLister interface {
	// List lists all CustomedHorizontalPodAutoscalers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CustomedHorizontalPodAutoscaler, err error)
	// Get retrieves the CustomedHorizontalPodAutoscaler from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CustomedHorizontalPodAutoscaler, error)
	CustomedHorizontalPodAutoscalerNamespaceListerExpansion
}

// customedHorizontalPodAutoscalerNamespaceLister implements the CustomedHorizontalPodAutoscalerNamespaceLister
// interface.
type customedHorizontalPodAutoscalerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CustomedHorizontalPodAutoscalers in the indexer for a given namespace.
func (s customedHorizontalPodAutoscalerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CustomedHorizontalPodAutoscaler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CustomedHorizontalPodAutoscaler))
	})
	return ret, err
}

// Get retrieves the CustomedHorizontalPodAutoscaler from the indexer for a given namespace and name.
func (s customedHorizontalPodAutoscalerNamespaceLister) Get(name string) (*v1alpha1.CustomedHorizontalPodAutoscaler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("customedhorizontalpodautoscaler"), name)
	}
	return obj.(*v1alpha1.CustomedHorizontalPodAutoscaler), nil
}
