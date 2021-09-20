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
// GENERATED BY gen_go_data.go
// gen_go_data -package beta -var YAML_key_ring blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudkms/beta/key_ring.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudkms/beta/key_ring.yaml
var YAML_key_ring = []byte("info:\n  title: Cloudkms/KeyRing\n  description: DCL Specification for the Cloudkms KeyRing resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a KeyRing\n    parameters:\n    - name: KeyRing\n      required: true\n      description: A full instance of a KeyRing\n  apply:\n    description: The function used to apply information about a KeyRing\n    parameters:\n    - name: KeyRing\n      required: true\n      description: A full instance of a KeyRing\n  list:\n    description: The function used to list information about many KeyRing\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    KeyRing:\n      title: KeyRing\n      x-dcl-id: projects/{{project}}/locations/{{location}}/keyRings/{{name}}\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time at which this KeyRing was created.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          readOnly: true\n          description: Output only. The resource name for the KeyRing in the format\n            `projects/*/locations/*/keyRings/*`.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n")

// 2043 bytes
// MD5: 1b1277bc399eacb268579297cc47fc2c
