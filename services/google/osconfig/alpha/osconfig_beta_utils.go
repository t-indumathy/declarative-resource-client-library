// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package alpha

import (
	"errors"
	"strings"
)

// expandGuestPolicyInstances truncates the instances field to just the part from "zones" forward.
func expandGuestPolicyInstances(_ *GuestPolicyAssignment, i interface{}) ([]string, error) {
	instances, ok := i.([]string)
	if !ok {
		return nil, errors.New("expand instances encountered unknown instances type")
	}
	for j, instance := range instances {
		parts := strings.Split(instance, "/")
		for k, part := range parts {
			if part == "projects" || part == "zones" {
				instances[j] = strings.Join(parts[k:], "/")
				break
			}
		}
	}
	return instances, nil
}

// flattenGuestPolicyInstances returns the instances field unaltered.
func flattenGuestPolicyInstances(i interface{}) []string {
	if items, ok := i.([]interface{}); ok {
		instances := make([]string, len(items))
		for j, item := range items {
			instances[j] = item.(string)
		}
		return instances
	}
	return nil
}
