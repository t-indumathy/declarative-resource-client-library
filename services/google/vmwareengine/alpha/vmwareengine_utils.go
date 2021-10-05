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
// Package vmwareengine contains methods and types for handling vmwareengine GCP resources.
package alpha

import (
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func encodePrivateCloudUpdateClusterRequest(m map[string]interface{}) map[string]interface{} {
	if req, ok := m["managementCluster"].(map[string]interface{}); ok {
		delete(req, "clusterId")
		return req
	}
	return nil
}

func (op *updatePrivateCloudUpdateClusterOperation) UpdateMask() string {
	for _, fieldDiff := range op.FieldDiffs {
		fieldDiff.FieldName = strings.TrimPrefix(fieldDiff.FieldName, "ManagementCluster.")
	}
	return dcl.UpdateMask(op.FieldDiffs)
}
