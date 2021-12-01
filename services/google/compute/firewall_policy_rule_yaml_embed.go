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
// gen_go_data -package compute -var YAML_firewall_policy_rule blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/firewall_policy_rule.yaml

package compute

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/firewall_policy_rule.yaml
var YAML_firewall_policy_rule = []byte("info:\n  title: Compute/FirewallPolicyRule\n  description: The Compute FirewallPolicyRule resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a FirewallPolicyRule\n    parameters:\n    - name: FirewallPolicyRule\n      required: true\n      description: A full instance of a FirewallPolicyRule\n  apply:\n    description: The function used to apply information about a FirewallPolicyRule\n    parameters:\n    - name: FirewallPolicyRule\n      required: true\n      description: A full instance of a FirewallPolicyRule\n  delete:\n    description: The function used to delete a FirewallPolicyRule\n    parameters:\n    - name: FirewallPolicyRule\n      required: true\n      description: A full instance of a FirewallPolicyRule\n  deleteAll:\n    description: The function used to delete all FirewallPolicyRule\n    parameters:\n    - name: firewallpolicy\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many FirewallPolicyRule\n    parameters:\n    - name: firewallpolicy\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    FirewallPolicyRule:\n      title: FirewallPolicyRule\n      x-dcl-id: locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}\n      x-dcl-locations:\n      - global\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - priority\n      - match\n      - action\n      - direction\n      - firewallPolicy\n      properties:\n        action:\n          type: string\n          x-dcl-go-name: Action\n          description: The Action to perform when the client connection triggers the\n            rule. Can currently be either \"allow\" or \"deny()\" where valid values for\n            status are 403, 404, and 502.\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description for this resource.\n        direction:\n          type: string\n          x-dcl-go-name: Direction\n          x-dcl-go-type: FirewallPolicyRuleDirectionEnum\n          description: 'The direction in which this rule applies. Possible values:\n            INGRESS, EGRESS'\n          enum:\n          - INGRESS\n          - EGRESS\n        disabled:\n          type: boolean\n          x-dcl-go-name: Disabled\n          description: Denotes whether the firewall policy rule is disabled. When\n            set to true, the firewall policy rule is not enforced and traffic behaves\n            as if it did not exist. If this is unspecified, the firewall policy rule\n            will be enabled.\n        enableLogging:\n          type: boolean\n          x-dcl-go-name: EnableLogging\n          description: 'Denotes whether to enable logging for a particular rule. If\n            logging is enabled, logs will be exported to the configured export destination\n            in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you\n            cannot enable logging on \"goto_next\" rules.'\n        firewallPolicy:\n          type: string\n          x-dcl-go-name: FirewallPolicy\n          description: The firewall policy of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/FirewallPolicy\n            field: name\n            parent: true\n        kind:\n          type: string\n          x-dcl-go-name: Kind\n          readOnly: true\n          description: Type of the resource. Always `compute#firewallPolicyRule` for\n            firewall policy rules\n          x-kubernetes-immutable: true\n        match:\n          type: object\n          x-dcl-go-name: Match\n          x-dcl-go-type: FirewallPolicyRuleMatch\n          description: A match condition that incoming traffic is evaluated against.\n            If it evaluates to true, the corresponding 'action' is enforced.\n          required:\n          - layer4Configs\n          properties:\n            destIPRanges:\n              type: array\n              x-dcl-go-name: DestIPRanges\n              description: CIDR IP address range. Maximum number of destination CIDR\n                IP ranges allowed is 256.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            layer4Configs:\n              type: array\n              x-dcl-go-name: Layer4Configs\n              description: Pairs of IP protocols and ports that the rule should match.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: FirewallPolicyRuleMatchLayer4Configs\n                required:\n                - ipProtocol\n                properties:\n                  ipProtocol:\n                    type: string\n                    x-dcl-go-name: IPProtocol\n                    description: The IP protocol to which this rule applies. The protocol\n                      type is required when creating a firewall rule. This value can\n                      either be one of the following well known protocol strings (`tcp`,\n                      `udp`, `icmp`, `esp`, `ah`, `ipip`, `sctp`), or the IP protocol\n                      number.\n                  ports:\n                    type: array\n                    x-dcl-go-name: Ports\n                    description: 'An optional list of ports to which this rule applies.\n                      This field is only applicable for UDP or TCP protocol. Each\n                      entry must be either an integer or a range. If not specified,\n                      this rule applies to connections through any port. Example inputs\n                      include: ``.'\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: string\n                      x-dcl-go-type: string\n            srcIPRanges:\n              type: array\n              x-dcl-go-name: SrcIPRanges\n              description: CIDR IP address range. Maximum number of source CIDR IP\n                ranges allowed is 256.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n        priority:\n          type: integer\n          format: int64\n          x-dcl-go-name: Priority\n          description: An integer indicating the priority of a rule in the list. The\n            priority must be a positive value between 0 and 2147483647. Rules are\n            evaluated from highest to lowest priority where 0 is the highest priority\n            and 2147483647 is the lowest prority.\n          x-kubernetes-immutable: true\n        ruleTupleCount:\n          type: integer\n          format: int64\n          x-dcl-go-name: RuleTupleCount\n          readOnly: true\n          description: Calculation of the complexity of a single firewall policy rule.\n        targetResources:\n          type: array\n          x-dcl-go-name: TargetResources\n          description: A list of network resource URLs to which this rule applies.\n            This field allows you to control which network's VMs get this rule. If\n            this field is left blank, all VMs within the organization will receive\n            the rule.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Compute/Network\n              field: selfLink\n        targetServiceAccounts:\n          type: array\n          x-dcl-go-name: TargetServiceAccounts\n          description: A list of service accounts indicating the sets of instances\n            that are applied with this rule.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Iam/ServiceAccount\n              field: name\n")

// 8009 bytes
// MD5: 09d815541f3b42bf80ade81b2f61f26c
