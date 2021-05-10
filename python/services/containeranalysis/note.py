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
from google3.cloud.graphite.mmv2.services.google.container_analysis import note_pb2
from google3.cloud.graphite.mmv2.services.google.container_analysis import note_pb2_grpc

from typing import List


class Note(object):
    def __init__(
        self,
        name: str = None,
        short_description: str = None,
        long_description: str = None,
        related_url: list = None,
        expiration_time: str = None,
        create_time: str = None,
        update_time: str = None,
        related_note_names: list = None,
        vulnerability: dict = None,
        build: dict = None,
        image: dict = None,
        package: dict = None,
        discovery: dict = None,
        deployment: dict = None,
        attestation: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.short_description = short_description
        self.long_description = long_description
        self.related_url = related_url
        self.expiration_time = expiration_time
        self.related_note_names = related_note_names
        self.vulnerability = vulnerability
        self.build = build
        self.image = image
        self.package = package
        self.discovery = discovery
        self.deployment = deployment
        self.attestation = attestation
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = note_pb2_grpc.ContaineranalysisNoteServiceStub(channel.Channel())
        request = note_pb2.ApplyContaineranalysisNoteRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.short_description):
            request.resource.short_description = Primitive.to_proto(
                self.short_description
            )

        if Primitive.to_proto(self.long_description):
            request.resource.long_description = Primitive.to_proto(
                self.long_description
            )

        if NoteRelatedUrlArray.to_proto(self.related_url):
            request.resource.related_url.extend(
                NoteRelatedUrlArray.to_proto(self.related_url)
            )
        if Primitive.to_proto(self.expiration_time):
            request.resource.expiration_time = Primitive.to_proto(self.expiration_time)

        if Primitive.to_proto(self.related_note_names):
            request.resource.related_note_names.extend(
                Primitive.to_proto(self.related_note_names)
            )
        if NoteVulnerability.to_proto(self.vulnerability):
            request.resource.vulnerability.CopyFrom(
                NoteVulnerability.to_proto(self.vulnerability)
            )
        else:
            request.resource.ClearField("vulnerability")
        if NoteBuild.to_proto(self.build):
            request.resource.build.CopyFrom(NoteBuild.to_proto(self.build))
        else:
            request.resource.ClearField("build")
        if NoteImage.to_proto(self.image):
            request.resource.image.CopyFrom(NoteImage.to_proto(self.image))
        else:
            request.resource.ClearField("image")
        if NotePackage.to_proto(self.package):
            request.resource.package.CopyFrom(NotePackage.to_proto(self.package))
        else:
            request.resource.ClearField("package")
        if NoteDiscovery.to_proto(self.discovery):
            request.resource.discovery.CopyFrom(NoteDiscovery.to_proto(self.discovery))
        else:
            request.resource.ClearField("discovery")
        if NoteDeployment.to_proto(self.deployment):
            request.resource.deployment.CopyFrom(
                NoteDeployment.to_proto(self.deployment)
            )
        else:
            request.resource.ClearField("deployment")
        if NoteAttestation.to_proto(self.attestation):
            request.resource.attestation.CopyFrom(
                NoteAttestation.to_proto(self.attestation)
            )
        else:
            request.resource.ClearField("attestation")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyContaineranalysisNote(request)
        self.name = Primitive.from_proto(response.name)
        self.short_description = Primitive.from_proto(response.short_description)
        self.long_description = Primitive.from_proto(response.long_description)
        self.related_url = NoteRelatedUrlArray.from_proto(response.related_url)
        self.expiration_time = Primitive.from_proto(response.expiration_time)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.related_note_names = Primitive.from_proto(response.related_note_names)
        self.vulnerability = NoteVulnerability.from_proto(response.vulnerability)
        self.build = NoteBuild.from_proto(response.build)
        self.image = NoteImage.from_proto(response.image)
        self.package = NotePackage.from_proto(response.package)
        self.discovery = NoteDiscovery.from_proto(response.discovery)
        self.deployment = NoteDeployment.from_proto(response.deployment)
        self.attestation = NoteAttestation.from_proto(response.attestation)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = note_pb2_grpc.ContaineranalysisNoteServiceStub(channel.Channel())
        request = note_pb2.DeleteContaineranalysisNoteRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.short_description):
            request.resource.short_description = Primitive.to_proto(
                self.short_description
            )

        if Primitive.to_proto(self.long_description):
            request.resource.long_description = Primitive.to_proto(
                self.long_description
            )

        if NoteRelatedUrlArray.to_proto(self.related_url):
            request.resource.related_url.extend(
                NoteRelatedUrlArray.to_proto(self.related_url)
            )
        if Primitive.to_proto(self.expiration_time):
            request.resource.expiration_time = Primitive.to_proto(self.expiration_time)

        if Primitive.to_proto(self.related_note_names):
            request.resource.related_note_names.extend(
                Primitive.to_proto(self.related_note_names)
            )
        if NoteVulnerability.to_proto(self.vulnerability):
            request.resource.vulnerability.CopyFrom(
                NoteVulnerability.to_proto(self.vulnerability)
            )
        else:
            request.resource.ClearField("vulnerability")
        if NoteBuild.to_proto(self.build):
            request.resource.build.CopyFrom(NoteBuild.to_proto(self.build))
        else:
            request.resource.ClearField("build")
        if NoteImage.to_proto(self.image):
            request.resource.image.CopyFrom(NoteImage.to_proto(self.image))
        else:
            request.resource.ClearField("image")
        if NotePackage.to_proto(self.package):
            request.resource.package.CopyFrom(NotePackage.to_proto(self.package))
        else:
            request.resource.ClearField("package")
        if NoteDiscovery.to_proto(self.discovery):
            request.resource.discovery.CopyFrom(NoteDiscovery.to_proto(self.discovery))
        else:
            request.resource.ClearField("discovery")
        if NoteDeployment.to_proto(self.deployment):
            request.resource.deployment.CopyFrom(
                NoteDeployment.to_proto(self.deployment)
            )
        else:
            request.resource.ClearField("deployment")
        if NoteAttestation.to_proto(self.attestation):
            request.resource.attestation.CopyFrom(
                NoteAttestation.to_proto(self.attestation)
            )
        else:
            request.resource.ClearField("attestation")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteContaineranalysisNote(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = note_pb2_grpc.ContaineranalysisNoteServiceStub(channel.Channel())
        request = note_pb2.ListContaineranalysisNoteRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListContaineranalysisNote(request).items

    def to_proto(self):
        resource = note_pb2.ContaineranalysisNote()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.short_description):
            resource.short_description = Primitive.to_proto(self.short_description)
        if Primitive.to_proto(self.long_description):
            resource.long_description = Primitive.to_proto(self.long_description)
        if NoteRelatedUrlArray.to_proto(self.related_url):
            resource.related_url.extend(NoteRelatedUrlArray.to_proto(self.related_url))
        if Primitive.to_proto(self.expiration_time):
            resource.expiration_time = Primitive.to_proto(self.expiration_time)
        if Primitive.to_proto(self.related_note_names):
            resource.related_note_names.extend(
                Primitive.to_proto(self.related_note_names)
            )
        if NoteVulnerability.to_proto(self.vulnerability):
            resource.vulnerability.CopyFrom(
                NoteVulnerability.to_proto(self.vulnerability)
            )
        else:
            resource.ClearField("vulnerability")
        if NoteBuild.to_proto(self.build):
            resource.build.CopyFrom(NoteBuild.to_proto(self.build))
        else:
            resource.ClearField("build")
        if NoteImage.to_proto(self.image):
            resource.image.CopyFrom(NoteImage.to_proto(self.image))
        else:
            resource.ClearField("image")
        if NotePackage.to_proto(self.package):
            resource.package.CopyFrom(NotePackage.to_proto(self.package))
        else:
            resource.ClearField("package")
        if NoteDiscovery.to_proto(self.discovery):
            resource.discovery.CopyFrom(NoteDiscovery.to_proto(self.discovery))
        else:
            resource.ClearField("discovery")
        if NoteDeployment.to_proto(self.deployment):
            resource.deployment.CopyFrom(NoteDeployment.to_proto(self.deployment))
        else:
            resource.ClearField("deployment")
        if NoteAttestation.to_proto(self.attestation):
            resource.attestation.CopyFrom(NoteAttestation.to_proto(self.attestation))
        else:
            resource.ClearField("attestation")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class NoteRelatedUrl(object):
    def __init__(self, url: str = None, label: str = None):
        self.url = url
        self.label = label

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteRelatedUrl()
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteRelatedUrl(url=resource.url, label=resource.label,)


class NoteRelatedUrlArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteRelatedUrl.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteRelatedUrl.from_proto(i) for i in resources]


class NoteVulnerability(object):
    def __init__(
        self,
        cvss_score: float = None,
        severity: str = None,
        details: list = None,
        cvss_v3: dict = None,
        windows_details: list = None,
        source_update_time: str = None,
    ):
        self.cvss_score = cvss_score
        self.severity = severity
        self.details = details
        self.cvss_v3 = cvss_v3
        self.windows_details = windows_details
        self.source_update_time = source_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteVulnerability()
        if Primitive.to_proto(resource.cvss_score):
            res.cvss_score = Primitive.to_proto(resource.cvss_score)
        if NoteVulnerabilitySeverityEnum.to_proto(resource.severity):
            res.severity = NoteVulnerabilitySeverityEnum.to_proto(resource.severity)
        if NoteVulnerabilityDetailsArray.to_proto(resource.details):
            res.details.extend(NoteVulnerabilityDetailsArray.to_proto(resource.details))
        if NoteVulnerabilityCvssV3.to_proto(resource.cvss_v3):
            res.cvss_v3.CopyFrom(NoteVulnerabilityCvssV3.to_proto(resource.cvss_v3))
        else:
            res.ClearField("cvss_v3")
        if NoteVulnerabilityWindowsDetailsArray.to_proto(resource.windows_details):
            res.windows_details.extend(
                NoteVulnerabilityWindowsDetailsArray.to_proto(resource.windows_details)
            )
        if Primitive.to_proto(resource.source_update_time):
            res.source_update_time = Primitive.to_proto(resource.source_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteVulnerability(
            cvss_score=resource.cvss_score,
            severity=resource.severity,
            details=resource.details,
            cvss_v3=resource.cvss_v3,
            windows_details=resource.windows_details,
            source_update_time=resource.source_update_time,
        )


class NoteVulnerabilityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteVulnerability.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteVulnerability.from_proto(i) for i in resources]


class NoteVulnerabilityDetails(object):
    def __init__(
        self,
        severity_name: str = None,
        description: str = None,
        package_type: str = None,
        affected_cpe_uri: str = None,
        affected_package: str = None,
        affected_version_start: dict = None,
        affected_version_end: dict = None,
        fixed_cpe_uri: str = None,
        fixed_package: str = None,
        fixed_version: dict = None,
        is_obsolete: bool = None,
        source_update_time: str = None,
    ):
        self.severity_name = severity_name
        self.description = description
        self.package_type = package_type
        self.affected_cpe_uri = affected_cpe_uri
        self.affected_package = affected_package
        self.affected_version_start = affected_version_start
        self.affected_version_end = affected_version_end
        self.fixed_cpe_uri = fixed_cpe_uri
        self.fixed_package = fixed_package
        self.fixed_version = fixed_version
        self.is_obsolete = is_obsolete
        self.source_update_time = source_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteVulnerabilityDetails()
        if Primitive.to_proto(resource.severity_name):
            res.severity_name = Primitive.to_proto(resource.severity_name)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.package_type):
            res.package_type = Primitive.to_proto(resource.package_type)
        if Primitive.to_proto(resource.affected_cpe_uri):
            res.affected_cpe_uri = Primitive.to_proto(resource.affected_cpe_uri)
        if Primitive.to_proto(resource.affected_package):
            res.affected_package = Primitive.to_proto(resource.affected_package)
        if NoteVulnerabilityDetailsAffectedVersionStart.to_proto(
            resource.affected_version_start
        ):
            res.affected_version_start.CopyFrom(
                NoteVulnerabilityDetailsAffectedVersionStart.to_proto(
                    resource.affected_version_start
                )
            )
        else:
            res.ClearField("affected_version_start")
        if NoteVulnerabilityDetailsAffectedVersionEnd.to_proto(
            resource.affected_version_end
        ):
            res.affected_version_end.CopyFrom(
                NoteVulnerabilityDetailsAffectedVersionEnd.to_proto(
                    resource.affected_version_end
                )
            )
        else:
            res.ClearField("affected_version_end")
        if Primitive.to_proto(resource.fixed_cpe_uri):
            res.fixed_cpe_uri = Primitive.to_proto(resource.fixed_cpe_uri)
        if Primitive.to_proto(resource.fixed_package):
            res.fixed_package = Primitive.to_proto(resource.fixed_package)
        if NoteVulnerabilityDetailsFixedVersion.to_proto(resource.fixed_version):
            res.fixed_version.CopyFrom(
                NoteVulnerabilityDetailsFixedVersion.to_proto(resource.fixed_version)
            )
        else:
            res.ClearField("fixed_version")
        if Primitive.to_proto(resource.is_obsolete):
            res.is_obsolete = Primitive.to_proto(resource.is_obsolete)
        if Primitive.to_proto(resource.source_update_time):
            res.source_update_time = Primitive.to_proto(resource.source_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteVulnerabilityDetails(
            severity_name=resource.severity_name,
            description=resource.description,
            package_type=resource.package_type,
            affected_cpe_uri=resource.affected_cpe_uri,
            affected_package=resource.affected_package,
            affected_version_start=resource.affected_version_start,
            affected_version_end=resource.affected_version_end,
            fixed_cpe_uri=resource.fixed_cpe_uri,
            fixed_package=resource.fixed_package,
            fixed_version=resource.fixed_version,
            is_obsolete=resource.is_obsolete,
            source_update_time=resource.source_update_time,
        )


class NoteVulnerabilityDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteVulnerabilityDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteVulnerabilityDetails.from_proto(i) for i in resources]


class NoteVulnerabilityDetailsAffectedVersionStart(object):
    def __init__(
        self,
        epoch: int = None,
        name: str = None,
        revision: str = None,
        kind: str = None,
        full_name: str = None,
    ):
        self.epoch = epoch
        self.name = name
        self.revision = revision
        self.kind = kind
        self.full_name = full_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStart()
        if Primitive.to_proto(resource.epoch):
            res.epoch = Primitive.to_proto(resource.epoch)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.revision):
            res.revision = Primitive.to_proto(resource.revision)
        if NoteVulnerabilityDetailsAffectedVersionStartKindEnum.to_proto(resource.kind):
            res.kind = NoteVulnerabilityDetailsAffectedVersionStartKindEnum.to_proto(
                resource.kind
            )
        if Primitive.to_proto(resource.full_name):
            res.full_name = Primitive.to_proto(resource.full_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteVulnerabilityDetailsAffectedVersionStart(
            epoch=resource.epoch,
            name=resource.name,
            revision=resource.revision,
            kind=resource.kind,
            full_name=resource.full_name,
        )


class NoteVulnerabilityDetailsAffectedVersionStartArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            NoteVulnerabilityDetailsAffectedVersionStart.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            NoteVulnerabilityDetailsAffectedVersionStart.from_proto(i)
            for i in resources
        ]


class NoteVulnerabilityDetailsAffectedVersionEnd(object):
    def __init__(
        self,
        epoch: int = None,
        name: str = None,
        revision: str = None,
        kind: str = None,
        full_name: str = None,
    ):
        self.epoch = epoch
        self.name = name
        self.revision = revision
        self.kind = kind
        self.full_name = full_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEnd()
        if Primitive.to_proto(resource.epoch):
            res.epoch = Primitive.to_proto(resource.epoch)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.revision):
            res.revision = Primitive.to_proto(resource.revision)
        if NoteVulnerabilityDetailsAffectedVersionEndKindEnum.to_proto(resource.kind):
            res.kind = NoteVulnerabilityDetailsAffectedVersionEndKindEnum.to_proto(
                resource.kind
            )
        if Primitive.to_proto(resource.full_name):
            res.full_name = Primitive.to_proto(resource.full_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteVulnerabilityDetailsAffectedVersionEnd(
            epoch=resource.epoch,
            name=resource.name,
            revision=resource.revision,
            kind=resource.kind,
            full_name=resource.full_name,
        )


class NoteVulnerabilityDetailsAffectedVersionEndArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            NoteVulnerabilityDetailsAffectedVersionEnd.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            NoteVulnerabilityDetailsAffectedVersionEnd.from_proto(i) for i in resources
        ]


class NoteVulnerabilityDetailsFixedVersion(object):
    def __init__(
        self,
        epoch: int = None,
        name: str = None,
        revision: str = None,
        kind: str = None,
        full_name: str = None,
    ):
        self.epoch = epoch
        self.name = name
        self.revision = revision
        self.kind = kind
        self.full_name = full_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteVulnerabilityDetailsFixedVersion()
        if Primitive.to_proto(resource.epoch):
            res.epoch = Primitive.to_proto(resource.epoch)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.revision):
            res.revision = Primitive.to_proto(resource.revision)
        if NoteVulnerabilityDetailsFixedVersionKindEnum.to_proto(resource.kind):
            res.kind = NoteVulnerabilityDetailsFixedVersionKindEnum.to_proto(
                resource.kind
            )
        if Primitive.to_proto(resource.full_name):
            res.full_name = Primitive.to_proto(resource.full_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteVulnerabilityDetailsFixedVersion(
            epoch=resource.epoch,
            name=resource.name,
            revision=resource.revision,
            kind=resource.kind,
            full_name=resource.full_name,
        )


class NoteVulnerabilityDetailsFixedVersionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteVulnerabilityDetailsFixedVersion.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteVulnerabilityDetailsFixedVersion.from_proto(i) for i in resources]


class NoteVulnerabilityCvssV3(object):
    def __init__(
        self,
        base_score: float = None,
        exploitability_score: float = None,
        impact_score: float = None,
        attack_vector: str = None,
        attack_complexity: str = None,
        privileges_required: str = None,
        user_interaction: str = None,
        scope: str = None,
        confidentiality_impact: str = None,
        integrity_impact: str = None,
        availability_impact: str = None,
    ):
        self.base_score = base_score
        self.exploitability_score = exploitability_score
        self.impact_score = impact_score
        self.attack_vector = attack_vector
        self.attack_complexity = attack_complexity
        self.privileges_required = privileges_required
        self.user_interaction = user_interaction
        self.scope = scope
        self.confidentiality_impact = confidentiality_impact
        self.integrity_impact = integrity_impact
        self.availability_impact = availability_impact

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteVulnerabilityCvssV3()
        if Primitive.to_proto(resource.base_score):
            res.base_score = Primitive.to_proto(resource.base_score)
        if Primitive.to_proto(resource.exploitability_score):
            res.exploitability_score = Primitive.to_proto(resource.exploitability_score)
        if Primitive.to_proto(resource.impact_score):
            res.impact_score = Primitive.to_proto(resource.impact_score)
        if NoteVulnerabilityCvssV3AttackVectorEnum.to_proto(resource.attack_vector):
            res.attack_vector = NoteVulnerabilityCvssV3AttackVectorEnum.to_proto(
                resource.attack_vector
            )
        if NoteVulnerabilityCvssV3AttackComplexityEnum.to_proto(
            resource.attack_complexity
        ):
            res.attack_complexity = NoteVulnerabilityCvssV3AttackComplexityEnum.to_proto(
                resource.attack_complexity
            )
        if NoteVulnerabilityCvssV3PrivilegesRequiredEnum.to_proto(
            resource.privileges_required
        ):
            res.privileges_required = NoteVulnerabilityCvssV3PrivilegesRequiredEnum.to_proto(
                resource.privileges_required
            )
        if NoteVulnerabilityCvssV3UserInteractionEnum.to_proto(
            resource.user_interaction
        ):
            res.user_interaction = NoteVulnerabilityCvssV3UserInteractionEnum.to_proto(
                resource.user_interaction
            )
        if NoteVulnerabilityCvssV3ScopeEnum.to_proto(resource.scope):
            res.scope = NoteVulnerabilityCvssV3ScopeEnum.to_proto(resource.scope)
        if NoteVulnerabilityCvssV3ConfidentialityImpactEnum.to_proto(
            resource.confidentiality_impact
        ):
            res.confidentiality_impact = NoteVulnerabilityCvssV3ConfidentialityImpactEnum.to_proto(
                resource.confidentiality_impact
            )
        if NoteVulnerabilityCvssV3IntegrityImpactEnum.to_proto(
            resource.integrity_impact
        ):
            res.integrity_impact = NoteVulnerabilityCvssV3IntegrityImpactEnum.to_proto(
                resource.integrity_impact
            )
        if NoteVulnerabilityCvssV3AvailabilityImpactEnum.to_proto(
            resource.availability_impact
        ):
            res.availability_impact = NoteVulnerabilityCvssV3AvailabilityImpactEnum.to_proto(
                resource.availability_impact
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteVulnerabilityCvssV3(
            base_score=resource.base_score,
            exploitability_score=resource.exploitability_score,
            impact_score=resource.impact_score,
            attack_vector=resource.attack_vector,
            attack_complexity=resource.attack_complexity,
            privileges_required=resource.privileges_required,
            user_interaction=resource.user_interaction,
            scope=resource.scope,
            confidentiality_impact=resource.confidentiality_impact,
            integrity_impact=resource.integrity_impact,
            availability_impact=resource.availability_impact,
        )


class NoteVulnerabilityCvssV3Array(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteVulnerabilityCvssV3.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteVulnerabilityCvssV3.from_proto(i) for i in resources]


class NoteVulnerabilityWindowsDetails(object):
    def __init__(
        self,
        cpe_uri: str = None,
        name: str = None,
        description: str = None,
        fixing_kbs: list = None,
    ):
        self.cpe_uri = cpe_uri
        self.name = name
        self.description = description
        self.fixing_kbs = fixing_kbs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteVulnerabilityWindowsDetails()
        if Primitive.to_proto(resource.cpe_uri):
            res.cpe_uri = Primitive.to_proto(resource.cpe_uri)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if NoteVulnerabilityWindowsDetailsFixingKbsArray.to_proto(resource.fixing_kbs):
            res.fixing_kbs.extend(
                NoteVulnerabilityWindowsDetailsFixingKbsArray.to_proto(
                    resource.fixing_kbs
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteVulnerabilityWindowsDetails(
            cpe_uri=resource.cpe_uri,
            name=resource.name,
            description=resource.description,
            fixing_kbs=resource.fixing_kbs,
        )


class NoteVulnerabilityWindowsDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteVulnerabilityWindowsDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteVulnerabilityWindowsDetails.from_proto(i) for i in resources]


class NoteVulnerabilityWindowsDetailsFixingKbs(object):
    def __init__(self, name: str = None, url: str = None):
        self.name = name
        self.url = url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteVulnerabilityWindowsDetailsFixingKbs()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteVulnerabilityWindowsDetailsFixingKbs(
            name=resource.name, url=resource.url,
        )


class NoteVulnerabilityWindowsDetailsFixingKbsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteVulnerabilityWindowsDetailsFixingKbs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            NoteVulnerabilityWindowsDetailsFixingKbs.from_proto(i) for i in resources
        ]


class NoteBuild(object):
    def __init__(self, builder_version: str = None):
        self.builder_version = builder_version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteBuild()
        if Primitive.to_proto(resource.builder_version):
            res.builder_version = Primitive.to_proto(resource.builder_version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteBuild(builder_version=resource.builder_version,)


class NoteBuildArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteBuild.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteBuild.from_proto(i) for i in resources]


class NoteImage(object):
    def __init__(self, resource_url: str = None, fingerprint: dict = None):
        self.resource_url = resource_url
        self.fingerprint = fingerprint

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteImage()
        if Primitive.to_proto(resource.resource_url):
            res.resource_url = Primitive.to_proto(resource.resource_url)
        if NoteImageFingerprint.to_proto(resource.fingerprint):
            res.fingerprint.CopyFrom(
                NoteImageFingerprint.to_proto(resource.fingerprint)
            )
        else:
            res.ClearField("fingerprint")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteImage(
            resource_url=resource.resource_url, fingerprint=resource.fingerprint,
        )


class NoteImageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteImage.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteImage.from_proto(i) for i in resources]


class NoteImageFingerprint(object):
    def __init__(self, v1_name: str = None, v2_blob: list = None, v2_name: str = None):
        self.v1_name = v1_name
        self.v2_blob = v2_blob
        self.v2_name = v2_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteImageFingerprint()
        if Primitive.to_proto(resource.v1_name):
            res.v1_name = Primitive.to_proto(resource.v1_name)
        if Primitive.to_proto(resource.v2_blob):
            res.v2_blob.extend(Primitive.to_proto(resource.v2_blob))
        if Primitive.to_proto(resource.v2_name):
            res.v2_name = Primitive.to_proto(resource.v2_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteImageFingerprint(
            v1_name=resource.v1_name,
            v2_blob=resource.v2_blob,
            v2_name=resource.v2_name,
        )


class NoteImageFingerprintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteImageFingerprint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteImageFingerprint.from_proto(i) for i in resources]


class NotePackage(object):
    def __init__(self, name: str = None, distribution: list = None):
        self.name = name
        self.distribution = distribution

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNotePackage()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if NotePackageDistributionArray.to_proto(resource.distribution):
            res.distribution.extend(
                NotePackageDistributionArray.to_proto(resource.distribution)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NotePackage(name=resource.name, distribution=resource.distribution,)


class NotePackageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NotePackage.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NotePackage.from_proto(i) for i in resources]


class NotePackageDistribution(object):
    def __init__(
        self,
        cpe_uri: str = None,
        architecture: str = None,
        latest_version: dict = None,
        maintainer: str = None,
        url: str = None,
        description: str = None,
    ):
        self.cpe_uri = cpe_uri
        self.architecture = architecture
        self.latest_version = latest_version
        self.maintainer = maintainer
        self.url = url
        self.description = description

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNotePackageDistribution()
        if Primitive.to_proto(resource.cpe_uri):
            res.cpe_uri = Primitive.to_proto(resource.cpe_uri)
        if NotePackageDistributionArchitectureEnum.to_proto(resource.architecture):
            res.architecture = NotePackageDistributionArchitectureEnum.to_proto(
                resource.architecture
            )
        if NotePackageDistributionLatestVersion.to_proto(resource.latest_version):
            res.latest_version.CopyFrom(
                NotePackageDistributionLatestVersion.to_proto(resource.latest_version)
            )
        else:
            res.ClearField("latest_version")
        if Primitive.to_proto(resource.maintainer):
            res.maintainer = Primitive.to_proto(resource.maintainer)
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NotePackageDistribution(
            cpe_uri=resource.cpe_uri,
            architecture=resource.architecture,
            latest_version=resource.latest_version,
            maintainer=resource.maintainer,
            url=resource.url,
            description=resource.description,
        )


class NotePackageDistributionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NotePackageDistribution.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NotePackageDistribution.from_proto(i) for i in resources]


class NotePackageDistributionLatestVersion(object):
    def __init__(
        self,
        epoch: int = None,
        name: str = None,
        revision: str = None,
        kind: str = None,
        full_name: str = None,
    ):
        self.epoch = epoch
        self.name = name
        self.revision = revision
        self.kind = kind
        self.full_name = full_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNotePackageDistributionLatestVersion()
        if Primitive.to_proto(resource.epoch):
            res.epoch = Primitive.to_proto(resource.epoch)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.revision):
            res.revision = Primitive.to_proto(resource.revision)
        if NotePackageDistributionLatestVersionKindEnum.to_proto(resource.kind):
            res.kind = NotePackageDistributionLatestVersionKindEnum.to_proto(
                resource.kind
            )
        if Primitive.to_proto(resource.full_name):
            res.full_name = Primitive.to_proto(resource.full_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NotePackageDistributionLatestVersion(
            epoch=resource.epoch,
            name=resource.name,
            revision=resource.revision,
            kind=resource.kind,
            full_name=resource.full_name,
        )


class NotePackageDistributionLatestVersionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NotePackageDistributionLatestVersion.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NotePackageDistributionLatestVersion.from_proto(i) for i in resources]


class NoteDiscovery(object):
    def __init__(self, analysis_kind: str = None):
        self.analysis_kind = analysis_kind

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteDiscovery()
        if NoteDiscoveryAnalysisKindEnum.to_proto(resource.analysis_kind):
            res.analysis_kind = NoteDiscoveryAnalysisKindEnum.to_proto(
                resource.analysis_kind
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteDiscovery(analysis_kind=resource.analysis_kind,)


class NoteDiscoveryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteDiscovery.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteDiscovery.from_proto(i) for i in resources]


class NoteDeployment(object):
    def __init__(self, resource_uri: list = None):
        self.resource_uri = resource_uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteDeployment()
        if Primitive.to_proto(resource.resource_uri):
            res.resource_uri.extend(Primitive.to_proto(resource.resource_uri))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteDeployment(resource_uri=resource.resource_uri,)


class NoteDeploymentArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteDeployment.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteDeployment.from_proto(i) for i in resources]


class NoteAttestation(object):
    def __init__(self, hint: dict = None):
        self.hint = hint

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteAttestation()
        if NoteAttestationHint.to_proto(resource.hint):
            res.hint.CopyFrom(NoteAttestationHint.to_proto(resource.hint))
        else:
            res.ClearField("hint")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteAttestation(hint=resource.hint,)


class NoteAttestationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteAttestation.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteAttestation.from_proto(i) for i in resources]


class NoteAttestationHint(object):
    def __init__(self, human_readable_name: str = None):
        self.human_readable_name = human_readable_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisNoteAttestationHint()
        if Primitive.to_proto(resource.human_readable_name):
            res.human_readable_name = Primitive.to_proto(resource.human_readable_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteAttestationHint(human_readable_name=resource.human_readable_name,)


class NoteAttestationHintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteAttestationHint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteAttestationHint.from_proto(i) for i in resources]


class NoteVulnerabilitySeverityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilitySeverityEnum.Value(
            "ContaineranalysisNoteVulnerabilitySeverityEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilitySeverityEnum.Name(resource)[
            len("ContaineranalysisNoteVulnerabilitySeverityEnum") :
        ]


class NoteVulnerabilityDetailsAffectedVersionStartKindEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum.Value(
            "ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum.Name(
            resource
        )[
            len(
                "ContaineranalysisNoteVulnerabilityDetailsAffectedVersionStartKindEnum"
            ) :
        ]


class NoteVulnerabilityDetailsAffectedVersionEndKindEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum.Value(
            "ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum.Name(
            resource
        )[
            len("ContaineranalysisNoteVulnerabilityDetailsAffectedVersionEndKindEnum") :
        ]


class NoteVulnerabilityDetailsFixedVersionKindEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum.Value(
            "ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum.Name(
            resource
        )[
            len("ContaineranalysisNoteVulnerabilityDetailsFixedVersionKindEnum") :
        ]


class NoteVulnerabilityCvssV3AttackVectorEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum.Value(
            "ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum.Name(
            resource
        )[len("ContaineranalysisNoteVulnerabilityCvssV3AttackVectorEnum") :]


class NoteVulnerabilityCvssV3AttackComplexityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum.Value(
            "ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum.Name(
            resource
        )[
            len("ContaineranalysisNoteVulnerabilityCvssV3AttackComplexityEnum") :
        ]


class NoteVulnerabilityCvssV3PrivilegesRequiredEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum.Value(
            "ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum.Name(
            resource
        )[
            len("ContaineranalysisNoteVulnerabilityCvssV3PrivilegesRequiredEnum") :
        ]


class NoteVulnerabilityCvssV3UserInteractionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum.Value(
            "ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum.Name(
            resource
        )[
            len("ContaineranalysisNoteVulnerabilityCvssV3UserInteractionEnum") :
        ]


class NoteVulnerabilityCvssV3ScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum.Value(
            "ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum.Name(
            resource
        )[len("ContaineranalysisNoteVulnerabilityCvssV3ScopeEnum") :]


class NoteVulnerabilityCvssV3ConfidentialityImpactEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum.Value(
            "ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum.Name(
            resource
        )[
            len("ContaineranalysisNoteVulnerabilityCvssV3ConfidentialityImpactEnum") :
        ]


class NoteVulnerabilityCvssV3IntegrityImpactEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum.Value(
            "ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum.Name(
            resource
        )[
            len("ContaineranalysisNoteVulnerabilityCvssV3IntegrityImpactEnum") :
        ]


class NoteVulnerabilityCvssV3AvailabilityImpactEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum.Value(
            "ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum.Name(
            resource
        )[
            len("ContaineranalysisNoteVulnerabilityCvssV3AvailabilityImpactEnum") :
        ]


class NotePackageDistributionArchitectureEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNotePackageDistributionArchitectureEnum.Value(
            "ContaineranalysisNotePackageDistributionArchitectureEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNotePackageDistributionArchitectureEnum.Name(
            resource
        )[len("ContaineranalysisNotePackageDistributionArchitectureEnum") :]


class NotePackageDistributionLatestVersionKindEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNotePackageDistributionLatestVersionKindEnum.Value(
            "ContaineranalysisNotePackageDistributionLatestVersionKindEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNotePackageDistributionLatestVersionKindEnum.Name(
            resource
        )[
            len("ContaineranalysisNotePackageDistributionLatestVersionKindEnum") :
        ]


class NoteDiscoveryAnalysisKindEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteDiscoveryAnalysisKindEnum.Value(
            "ContaineranalysisNoteDiscoveryAnalysisKindEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisNoteDiscoveryAnalysisKindEnum.Name(resource)[
            len("ContaineranalysisNoteDiscoveryAnalysisKindEnum") :
        ]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
