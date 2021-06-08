# Copyright 2021 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from connector import channel
from google3.cloud.graphite.mmv2.services.google.compute import forwarding_rule_pb2
from google3.cloud.graphite.mmv2.services.google.compute import forwarding_rule_pb2_grpc

from typing import List


class ForwardingRule(object):
    def __init__(
        self,
        all_ports: bool = None,
        allow_global_access: bool = None,
        backend_service: str = None,
        creation_timestamp: str = None,
        description: str = None,
        ip_address: str = None,
        ip_protocol: str = None,
        ip_version: str = None,
        is_mirroring_collector: bool = None,
        load_balancing_scheme: str = None,
        metadata_filter: list = None,
        name: str = None,
        network: str = None,
        network_tier: str = None,
        port_range: str = None,
        ports: list = None,
        region: str = None,
        self_link: str = None,
        service_label: str = None,
        service_name: str = None,
        subnetwork: str = None,
        target: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.all_ports = all_ports
        self.allow_global_access = allow_global_access
        self.backend_service = backend_service
        self.creation_timestamp = creation_timestamp
        self.description = description
        self.ip_address = ip_address
        self.ip_protocol = ip_protocol
        self.ip_version = ip_version
        self.is_mirroring_collector = is_mirroring_collector
        self.load_balancing_scheme = load_balancing_scheme
        self.metadata_filter = metadata_filter
        self.name = name
        self.network = network
        self.network_tier = network_tier
        self.port_range = port_range
        self.ports = ports
        self.region = region
        self.self_link = self_link
        self.service_label = service_label
        self.service_name = service_name
        self.subnetwork = subnetwork
        self.target = target
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = forwarding_rule_pb2_grpc.ComputeBetaForwardingRuleServiceStub(
            channel.Channel()
        )
        request = forwarding_rule_pb2.ApplyComputeBetaForwardingRuleRequest()
        if Primitive.to_proto(self.all_ports):
            request.resource.all_ports = Primitive.to_proto(self.all_ports)

        if Primitive.to_proto(self.allow_global_access):
            request.resource.allow_global_access = Primitive.to_proto(
                self.allow_global_access
            )

        if Primitive.to_proto(self.backend_service):
            request.resource.backend_service = Primitive.to_proto(self.backend_service)

        if Primitive.to_proto(self.creation_timestamp):
            request.resource.creation_timestamp = Primitive.to_proto(
                self.creation_timestamp
            )

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.ip_address):
            request.resource.ip_address = Primitive.to_proto(self.ip_address)

        if ForwardingRuleIPProtocolEnum.to_proto(self.ip_protocol):
            request.resource.ip_protocol = ForwardingRuleIPProtocolEnum.to_proto(
                self.ip_protocol
            )

        if ForwardingRuleIPVersionEnum.to_proto(self.ip_version):
            request.resource.ip_version = ForwardingRuleIPVersionEnum.to_proto(
                self.ip_version
            )

        if Primitive.to_proto(self.is_mirroring_collector):
            request.resource.is_mirroring_collector = Primitive.to_proto(
                self.is_mirroring_collector
            )

        if ForwardingRuleLoadBalancingSchemeEnum.to_proto(self.load_balancing_scheme):
            request.resource.load_balancing_scheme = ForwardingRuleLoadBalancingSchemeEnum.to_proto(
                self.load_balancing_scheme
            )

        if ForwardingRuleMetadataFilterArray.to_proto(self.metadata_filter):
            request.resource.metadata_filter.extend(
                ForwardingRuleMetadataFilterArray.to_proto(self.metadata_filter)
            )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if ForwardingRuleNetworkTierEnum.to_proto(self.network_tier):
            request.resource.network_tier = ForwardingRuleNetworkTierEnum.to_proto(
                self.network_tier
            )

        if Primitive.to_proto(self.port_range):
            request.resource.port_range = Primitive.to_proto(self.port_range)

        if Primitive.to_proto(self.ports):
            request.resource.ports.extend(Primitive.to_proto(self.ports))
        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.self_link):
            request.resource.self_link = Primitive.to_proto(self.self_link)

        if Primitive.to_proto(self.service_label):
            request.resource.service_label = Primitive.to_proto(self.service_label)

        if Primitive.to_proto(self.service_name):
            request.resource.service_name = Primitive.to_proto(self.service_name)

        if Primitive.to_proto(self.subnetwork):
            request.resource.subnetwork = Primitive.to_proto(self.subnetwork)

        if Primitive.to_proto(self.target):
            request.resource.target = Primitive.to_proto(self.target)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaForwardingRule(request)
        self.all_ports = Primitive.from_proto(response.all_ports)
        self.allow_global_access = Primitive.from_proto(response.allow_global_access)
        self.backend_service = Primitive.from_proto(response.backend_service)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.description = Primitive.from_proto(response.description)
        self.ip_address = Primitive.from_proto(response.ip_address)
        self.ip_protocol = ForwardingRuleIPProtocolEnum.from_proto(response.ip_protocol)
        self.ip_version = ForwardingRuleIPVersionEnum.from_proto(response.ip_version)
        self.is_mirroring_collector = Primitive.from_proto(
            response.is_mirroring_collector
        )
        self.load_balancing_scheme = ForwardingRuleLoadBalancingSchemeEnum.from_proto(
            response.load_balancing_scheme
        )
        self.metadata_filter = ForwardingRuleMetadataFilterArray.from_proto(
            response.metadata_filter
        )
        self.name = Primitive.from_proto(response.name)
        self.network = Primitive.from_proto(response.network)
        self.network_tier = ForwardingRuleNetworkTierEnum.from_proto(
            response.network_tier
        )
        self.port_range = Primitive.from_proto(response.port_range)
        self.ports = Primitive.from_proto(response.ports)
        self.region = Primitive.from_proto(response.region)
        self.self_link = Primitive.from_proto(response.self_link)
        self.service_label = Primitive.from_proto(response.service_label)
        self.service_name = Primitive.from_proto(response.service_name)
        self.subnetwork = Primitive.from_proto(response.subnetwork)
        self.target = Primitive.from_proto(response.target)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = forwarding_rule_pb2_grpc.ComputeBetaForwardingRuleServiceStub(
            channel.Channel()
        )
        request = forwarding_rule_pb2.DeleteComputeBetaForwardingRuleRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.all_ports):
            request.resource.all_ports = Primitive.to_proto(self.all_ports)

        if Primitive.to_proto(self.allow_global_access):
            request.resource.allow_global_access = Primitive.to_proto(
                self.allow_global_access
            )

        if Primitive.to_proto(self.backend_service):
            request.resource.backend_service = Primitive.to_proto(self.backend_service)

        if Primitive.to_proto(self.creation_timestamp):
            request.resource.creation_timestamp = Primitive.to_proto(
                self.creation_timestamp
            )

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.ip_address):
            request.resource.ip_address = Primitive.to_proto(self.ip_address)

        if ForwardingRuleIPProtocolEnum.to_proto(self.ip_protocol):
            request.resource.ip_protocol = ForwardingRuleIPProtocolEnum.to_proto(
                self.ip_protocol
            )

        if ForwardingRuleIPVersionEnum.to_proto(self.ip_version):
            request.resource.ip_version = ForwardingRuleIPVersionEnum.to_proto(
                self.ip_version
            )

        if Primitive.to_proto(self.is_mirroring_collector):
            request.resource.is_mirroring_collector = Primitive.to_proto(
                self.is_mirroring_collector
            )

        if ForwardingRuleLoadBalancingSchemeEnum.to_proto(self.load_balancing_scheme):
            request.resource.load_balancing_scheme = ForwardingRuleLoadBalancingSchemeEnum.to_proto(
                self.load_balancing_scheme
            )

        if ForwardingRuleMetadataFilterArray.to_proto(self.metadata_filter):
            request.resource.metadata_filter.extend(
                ForwardingRuleMetadataFilterArray.to_proto(self.metadata_filter)
            )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if ForwardingRuleNetworkTierEnum.to_proto(self.network_tier):
            request.resource.network_tier = ForwardingRuleNetworkTierEnum.to_proto(
                self.network_tier
            )

        if Primitive.to_proto(self.port_range):
            request.resource.port_range = Primitive.to_proto(self.port_range)

        if Primitive.to_proto(self.ports):
            request.resource.ports.extend(Primitive.to_proto(self.ports))
        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.self_link):
            request.resource.self_link = Primitive.to_proto(self.self_link)

        if Primitive.to_proto(self.service_label):
            request.resource.service_label = Primitive.to_proto(self.service_label)

        if Primitive.to_proto(self.service_name):
            request.resource.service_name = Primitive.to_proto(self.service_name)

        if Primitive.to_proto(self.subnetwork):
            request.resource.subnetwork = Primitive.to_proto(self.subnetwork)

        if Primitive.to_proto(self.target):
            request.resource.target = Primitive.to_proto(self.target)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteComputeBetaForwardingRule(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = forwarding_rule_pb2_grpc.ComputeBetaForwardingRuleServiceStub(
            channel.Channel()
        )
        request = forwarding_rule_pb2.ListComputeBetaForwardingRuleRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeBetaForwardingRule(request).items

    def to_proto(self):
        resource = forwarding_rule_pb2.ComputeBetaForwardingRule()
        if Primitive.to_proto(self.all_ports):
            resource.all_ports = Primitive.to_proto(self.all_ports)
        if Primitive.to_proto(self.allow_global_access):
            resource.allow_global_access = Primitive.to_proto(self.allow_global_access)
        if Primitive.to_proto(self.backend_service):
            resource.backend_service = Primitive.to_proto(self.backend_service)
        if Primitive.to_proto(self.creation_timestamp):
            resource.creation_timestamp = Primitive.to_proto(self.creation_timestamp)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.ip_address):
            resource.ip_address = Primitive.to_proto(self.ip_address)
        if ForwardingRuleIPProtocolEnum.to_proto(self.ip_protocol):
            resource.ip_protocol = ForwardingRuleIPProtocolEnum.to_proto(
                self.ip_protocol
            )
        if ForwardingRuleIPVersionEnum.to_proto(self.ip_version):
            resource.ip_version = ForwardingRuleIPVersionEnum.to_proto(self.ip_version)
        if Primitive.to_proto(self.is_mirroring_collector):
            resource.is_mirroring_collector = Primitive.to_proto(
                self.is_mirroring_collector
            )
        if ForwardingRuleLoadBalancingSchemeEnum.to_proto(self.load_balancing_scheme):
            resource.load_balancing_scheme = ForwardingRuleLoadBalancingSchemeEnum.to_proto(
                self.load_balancing_scheme
            )
        if ForwardingRuleMetadataFilterArray.to_proto(self.metadata_filter):
            resource.metadata_filter.extend(
                ForwardingRuleMetadataFilterArray.to_proto(self.metadata_filter)
            )
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if ForwardingRuleNetworkTierEnum.to_proto(self.network_tier):
            resource.network_tier = ForwardingRuleNetworkTierEnum.to_proto(
                self.network_tier
            )
        if Primitive.to_proto(self.port_range):
            resource.port_range = Primitive.to_proto(self.port_range)
        if Primitive.to_proto(self.ports):
            resource.ports.extend(Primitive.to_proto(self.ports))
        if Primitive.to_proto(self.region):
            resource.region = Primitive.to_proto(self.region)
        if Primitive.to_proto(self.self_link):
            resource.self_link = Primitive.to_proto(self.self_link)
        if Primitive.to_proto(self.service_label):
            resource.service_label = Primitive.to_proto(self.service_label)
        if Primitive.to_proto(self.service_name):
            resource.service_name = Primitive.to_proto(self.service_name)
        if Primitive.to_proto(self.subnetwork):
            resource.subnetwork = Primitive.to_proto(self.subnetwork)
        if Primitive.to_proto(self.target):
            resource.target = Primitive.to_proto(self.target)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class ForwardingRuleMetadataFilter(object):
    def __init__(self, filter_match_criteria: str = None, filter_label: list = None):
        self.filter_match_criteria = filter_match_criteria
        self.filter_label = filter_label

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = forwarding_rule_pb2.ComputeBetaForwardingRuleMetadataFilter()
        if ForwardingRuleMetadataFilterFilterMatchCriteriaEnum.to_proto(
            resource.filter_match_criteria
        ):
            res.filter_match_criteria = ForwardingRuleMetadataFilterFilterMatchCriteriaEnum.to_proto(
                resource.filter_match_criteria
            )
        if ForwardingRuleMetadataFilterFilterLabelArray.to_proto(resource.filter_label):
            res.filter_label.extend(
                ForwardingRuleMetadataFilterFilterLabelArray.to_proto(
                    resource.filter_label
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ForwardingRuleMetadataFilter(
            filter_match_criteria=ForwardingRuleMetadataFilterFilterMatchCriteriaEnum.from_proto(
                resource.filter_match_criteria
            ),
            filter_label=ForwardingRuleMetadataFilterFilterLabelArray.from_proto(
                resource.filter_label
            ),
        )


class ForwardingRuleMetadataFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ForwardingRuleMetadataFilter.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ForwardingRuleMetadataFilter.from_proto(i) for i in resources]


class ForwardingRuleMetadataFilterFilterLabel(object):
    def __init__(self, name: str = None, value: str = None):
        self.name = name
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = forwarding_rule_pb2.ComputeBetaForwardingRuleMetadataFilterFilterLabel()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ForwardingRuleMetadataFilterFilterLabel(
            name=Primitive.from_proto(resource.name),
            value=Primitive.from_proto(resource.value),
        )


class ForwardingRuleMetadataFilterFilterLabelArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ForwardingRuleMetadataFilterFilterLabel.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            ForwardingRuleMetadataFilterFilterLabel.from_proto(i) for i in resources
        ]


class ForwardingRuleIPProtocolEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return forwarding_rule_pb2.ComputeBetaForwardingRuleIPProtocolEnum.Value(
            "ComputeBetaForwardingRuleIPProtocolEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return forwarding_rule_pb2.ComputeBetaForwardingRuleIPProtocolEnum.Name(
            resource
        )[len("ComputeBetaForwardingRuleIPProtocolEnum") :]


class ForwardingRuleIPVersionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return forwarding_rule_pb2.ComputeBetaForwardingRuleIPVersionEnum.Value(
            "ComputeBetaForwardingRuleIPVersionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return forwarding_rule_pb2.ComputeBetaForwardingRuleIPVersionEnum.Name(
            resource
        )[len("ComputeBetaForwardingRuleIPVersionEnum") :]


class ForwardingRuleLoadBalancingSchemeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return forwarding_rule_pb2.ComputeBetaForwardingRuleLoadBalancingSchemeEnum.Value(
            "ComputeBetaForwardingRuleLoadBalancingSchemeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return forwarding_rule_pb2.ComputeBetaForwardingRuleLoadBalancingSchemeEnum.Name(
            resource
        )[
            len("ComputeBetaForwardingRuleLoadBalancingSchemeEnum") :
        ]


class ForwardingRuleMetadataFilterFilterMatchCriteriaEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return forwarding_rule_pb2.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum.Value(
            "ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return forwarding_rule_pb2.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum.Name(
            resource
        )[
            len("ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum") :
        ]


class ForwardingRuleNetworkTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return forwarding_rule_pb2.ComputeBetaForwardingRuleNetworkTierEnum.Value(
            "ComputeBetaForwardingRuleNetworkTierEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return forwarding_rule_pb2.ComputeBetaForwardingRuleNetworkTierEnum.Name(
            resource
        )[len("ComputeBetaForwardingRuleNetworkTierEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
