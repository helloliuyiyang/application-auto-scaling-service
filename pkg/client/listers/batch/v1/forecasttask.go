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

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "nanto.io/application-auto-scaling-service/pkg/apis/batch/v1"
)

// ForecastTaskLister helps list ForecastTasks.
// All objects returned here must be treated as read-only.
type ForecastTaskLister interface {
	// List lists all ForecastTasks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ForecastTask, err error)
	// ForecastTasks returns an object that can list and get ForecastTasks.
	ForecastTasks(namespace string) ForecastTaskNamespaceLister
	ForecastTaskListerExpansion
}

// forecastTaskLister implements the ForecastTaskLister interface.
type forecastTaskLister struct {
	indexer cache.Indexer
}

// NewForecastTaskLister returns a new ForecastTaskLister.
func NewForecastTaskLister(indexer cache.Indexer) ForecastTaskLister {
	return &forecastTaskLister{indexer: indexer}
}

// List lists all ForecastTasks in the indexer.
func (s *forecastTaskLister) List(selector labels.Selector) (ret []*v1.ForecastTask, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ForecastTask))
	})
	return ret, err
}

// ForecastTasks returns an object that can list and get ForecastTasks.
func (s *forecastTaskLister) ForecastTasks(namespace string) ForecastTaskNamespaceLister {
	return forecastTaskNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ForecastTaskNamespaceLister helps list and get ForecastTasks.
// All objects returned here must be treated as read-only.
type ForecastTaskNamespaceLister interface {
	// List lists all ForecastTasks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ForecastTask, err error)
	// Get retrieves the ForecastTask from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ForecastTask, error)
	ForecastTaskNamespaceListerExpansion
}

// forecastTaskNamespaceLister implements the ForecastTaskNamespaceLister
// interface.
type forecastTaskNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ForecastTasks in the indexer for a given namespace.
func (s forecastTaskNamespaceLister) List(selector labels.Selector) (ret []*v1.ForecastTask, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ForecastTask))
	})
	return ret, err
}

// Get retrieves the ForecastTask from the indexer for a given namespace and name.
func (s forecastTaskNamespaceLister) Get(name string) (*v1.ForecastTask, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("forecasttask"), name)
	}
	return obj.(*v1.ForecastTask), nil
}