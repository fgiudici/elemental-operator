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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"github.com/rancher/wrangler/v2/pkg/generic"
	v1 "k8s.io/api/rbac/v1"
)

// RoleBindingController interface for managing RoleBinding resources.
type RoleBindingController interface {
	generic.ControllerInterface[*v1.RoleBinding, *v1.RoleBindingList]
}

// RoleBindingClient interface for managing RoleBinding resources in Kubernetes.
type RoleBindingClient interface {
	generic.ClientInterface[*v1.RoleBinding, *v1.RoleBindingList]
}

// RoleBindingCache interface for retrieving RoleBinding resources in memory.
type RoleBindingCache interface {
	generic.CacheInterface[*v1.RoleBinding]
}
