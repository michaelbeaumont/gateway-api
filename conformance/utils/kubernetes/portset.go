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

package kubernetes

import (
	"sigs.k8s.io/gateway-api/apis/v1beta1"
)

// PortSet is a wrapper for handling sets of ports.
type PortSet []v1beta1.PortNumber

func (s *PortSet) PopPort() (v1beta1.PortNumber, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	newPort := v1beta1.PortNumber((*s)[0])
	*s = (*s)[1:]
	return newPort, true
}