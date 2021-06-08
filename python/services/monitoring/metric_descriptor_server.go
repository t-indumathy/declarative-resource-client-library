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
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	monitoringpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/monitoring_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
)

// Server implements the gRPC interface for MetricDescriptor.
type MetricDescriptorServer struct{}

// ProtoToMetricDescriptorDescriptorLabelsValueTypeEnum converts a MetricDescriptorDescriptorLabelsValueTypeEnum enum from its proto representation.
func ProtoToMonitoringMetricDescriptorDescriptorLabelsValueTypeEnum(e monitoringpb.MonitoringMetricDescriptorDescriptorLabelsValueTypeEnum) *monitoring.MetricDescriptorDescriptorLabelsValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringMetricDescriptorDescriptorLabelsValueTypeEnum_name[int32(e)]; ok {
		e := monitoring.MetricDescriptorDescriptorLabelsValueTypeEnum(n[len("MonitoringMetricDescriptorDescriptorLabelsValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorMetricKindEnum converts a MetricDescriptorMetricKindEnum enum from its proto representation.
func ProtoToMonitoringMetricDescriptorMetricKindEnum(e monitoringpb.MonitoringMetricDescriptorMetricKindEnum) *monitoring.MetricDescriptorMetricKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringMetricDescriptorMetricKindEnum_name[int32(e)]; ok {
		e := monitoring.MetricDescriptorMetricKindEnum(n[len("MonitoringMetricDescriptorMetricKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorValueTypeEnum converts a MetricDescriptorValueTypeEnum enum from its proto representation.
func ProtoToMonitoringMetricDescriptorValueTypeEnum(e monitoringpb.MonitoringMetricDescriptorValueTypeEnum) *monitoring.MetricDescriptorValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringMetricDescriptorValueTypeEnum_name[int32(e)]; ok {
		e := monitoring.MetricDescriptorValueTypeEnum(n[len("MonitoringMetricDescriptorValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorMetadataLaunchStageEnum converts a MetricDescriptorMetadataLaunchStageEnum enum from its proto representation.
func ProtoToMonitoringMetricDescriptorMetadataLaunchStageEnum(e monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum) *monitoring.MetricDescriptorMetadataLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum_name[int32(e)]; ok {
		e := monitoring.MetricDescriptorMetadataLaunchStageEnum(n[len("MonitoringMetricDescriptorMetadataLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorLaunchStageEnum converts a MetricDescriptorLaunchStageEnum enum from its proto representation.
func ProtoToMonitoringMetricDescriptorLaunchStageEnum(e monitoringpb.MonitoringMetricDescriptorLaunchStageEnum) *monitoring.MetricDescriptorLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringMetricDescriptorLaunchStageEnum_name[int32(e)]; ok {
		e := monitoring.MetricDescriptorLaunchStageEnum(n[len("MonitoringMetricDescriptorLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorDescriptorLabels converts a MetricDescriptorDescriptorLabels resource from its proto representation.
func ProtoToMonitoringMetricDescriptorDescriptorLabels(p *monitoringpb.MonitoringMetricDescriptorDescriptorLabels) *monitoring.MetricDescriptorDescriptorLabels {
	if p == nil {
		return nil
	}
	obj := &monitoring.MetricDescriptorDescriptorLabels{
		Key:         dcl.StringOrNil(p.Key),
		ValueType:   ProtoToMonitoringMetricDescriptorDescriptorLabelsValueTypeEnum(p.GetValueType()),
		Description: dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToMetricDescriptorMetadata converts a MetricDescriptorMetadata resource from its proto representation.
func ProtoToMonitoringMetricDescriptorMetadata(p *monitoringpb.MonitoringMetricDescriptorMetadata) *monitoring.MetricDescriptorMetadata {
	if p == nil {
		return nil
	}
	obj := &monitoring.MetricDescriptorMetadata{
		LaunchStage:  ProtoToMonitoringMetricDescriptorMetadataLaunchStageEnum(p.GetLaunchStage()),
		SamplePeriod: dcl.StringOrNil(p.SamplePeriod),
		IngestDelay:  dcl.StringOrNil(p.IngestDelay),
	}
	return obj
}

// ProtoToMetricDescriptor converts a MetricDescriptor resource from its proto representation.
func ProtoToMetricDescriptor(p *monitoringpb.MonitoringMetricDescriptor) *monitoring.MetricDescriptor {
	obj := &monitoring.MetricDescriptor{
		SelfLink:    dcl.StringOrNil(p.SelfLink),
		Type:        dcl.StringOrNil(p.Type),
		MetricKind:  ProtoToMonitoringMetricDescriptorMetricKindEnum(p.GetMetricKind()),
		ValueType:   ProtoToMonitoringMetricDescriptorValueTypeEnum(p.GetValueType()),
		Unit:        dcl.StringOrNil(p.Unit),
		Description: dcl.StringOrNil(p.Description),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		Metadata:    ProtoToMonitoringMetricDescriptorMetadata(p.GetMetadata()),
		LaunchStage: ProtoToMonitoringMetricDescriptorLaunchStageEnum(p.GetLaunchStage()),
		Project:     dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetDescriptorLabels() {
		obj.DescriptorLabels = append(obj.DescriptorLabels, *ProtoToMonitoringMetricDescriptorDescriptorLabels(r))
	}
	for _, r := range p.GetMonitoredResourceTypes() {
		obj.MonitoredResourceTypes = append(obj.MonitoredResourceTypes, r)
	}
	return obj
}

// MetricDescriptorDescriptorLabelsValueTypeEnumToProto converts a MetricDescriptorDescriptorLabelsValueTypeEnum enum to its proto representation.
func MonitoringMetricDescriptorDescriptorLabelsValueTypeEnumToProto(e *monitoring.MetricDescriptorDescriptorLabelsValueTypeEnum) monitoringpb.MonitoringMetricDescriptorDescriptorLabelsValueTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringMetricDescriptorDescriptorLabelsValueTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringMetricDescriptorDescriptorLabelsValueTypeEnum_value["MetricDescriptorDescriptorLabelsValueTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringMetricDescriptorDescriptorLabelsValueTypeEnum(v)
	}
	return monitoringpb.MonitoringMetricDescriptorDescriptorLabelsValueTypeEnum(0)
}

// MetricDescriptorMetricKindEnumToProto converts a MetricDescriptorMetricKindEnum enum to its proto representation.
func MonitoringMetricDescriptorMetricKindEnumToProto(e *monitoring.MetricDescriptorMetricKindEnum) monitoringpb.MonitoringMetricDescriptorMetricKindEnum {
	if e == nil {
		return monitoringpb.MonitoringMetricDescriptorMetricKindEnum(0)
	}
	if v, ok := monitoringpb.MonitoringMetricDescriptorMetricKindEnum_value["MetricDescriptorMetricKindEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringMetricDescriptorMetricKindEnum(v)
	}
	return monitoringpb.MonitoringMetricDescriptorMetricKindEnum(0)
}

// MetricDescriptorValueTypeEnumToProto converts a MetricDescriptorValueTypeEnum enum to its proto representation.
func MonitoringMetricDescriptorValueTypeEnumToProto(e *monitoring.MetricDescriptorValueTypeEnum) monitoringpb.MonitoringMetricDescriptorValueTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringMetricDescriptorValueTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringMetricDescriptorValueTypeEnum_value["MetricDescriptorValueTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringMetricDescriptorValueTypeEnum(v)
	}
	return monitoringpb.MonitoringMetricDescriptorValueTypeEnum(0)
}

// MetricDescriptorMetadataLaunchStageEnumToProto converts a MetricDescriptorMetadataLaunchStageEnum enum to its proto representation.
func MonitoringMetricDescriptorMetadataLaunchStageEnumToProto(e *monitoring.MetricDescriptorMetadataLaunchStageEnum) monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum {
	if e == nil {
		return monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum(0)
	}
	if v, ok := monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum_value["MetricDescriptorMetadataLaunchStageEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum(v)
	}
	return monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum(0)
}

// MetricDescriptorLaunchStageEnumToProto converts a MetricDescriptorLaunchStageEnum enum to its proto representation.
func MonitoringMetricDescriptorLaunchStageEnumToProto(e *monitoring.MetricDescriptorLaunchStageEnum) monitoringpb.MonitoringMetricDescriptorLaunchStageEnum {
	if e == nil {
		return monitoringpb.MonitoringMetricDescriptorLaunchStageEnum(0)
	}
	if v, ok := monitoringpb.MonitoringMetricDescriptorLaunchStageEnum_value["MetricDescriptorLaunchStageEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringMetricDescriptorLaunchStageEnum(v)
	}
	return monitoringpb.MonitoringMetricDescriptorLaunchStageEnum(0)
}

// MetricDescriptorDescriptorLabelsToProto converts a MetricDescriptorDescriptorLabels resource to its proto representation.
func MonitoringMetricDescriptorDescriptorLabelsToProto(o *monitoring.MetricDescriptorDescriptorLabels) *monitoringpb.MonitoringMetricDescriptorDescriptorLabels {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringMetricDescriptorDescriptorLabels{
		Key:         dcl.ValueOrEmptyString(o.Key),
		ValueType:   MonitoringMetricDescriptorDescriptorLabelsValueTypeEnumToProto(o.ValueType),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// MetricDescriptorMetadataToProto converts a MetricDescriptorMetadata resource to its proto representation.
func MonitoringMetricDescriptorMetadataToProto(o *monitoring.MetricDescriptorMetadata) *monitoringpb.MonitoringMetricDescriptorMetadata {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringMetricDescriptorMetadata{
		LaunchStage:  MonitoringMetricDescriptorMetadataLaunchStageEnumToProto(o.LaunchStage),
		SamplePeriod: dcl.ValueOrEmptyString(o.SamplePeriod),
		IngestDelay:  dcl.ValueOrEmptyString(o.IngestDelay),
	}
	return p
}

// MetricDescriptorToProto converts a MetricDescriptor resource to its proto representation.
func MetricDescriptorToProto(resource *monitoring.MetricDescriptor) *monitoringpb.MonitoringMetricDescriptor {
	p := &monitoringpb.MonitoringMetricDescriptor{
		SelfLink:    dcl.ValueOrEmptyString(resource.SelfLink),
		Type:        dcl.ValueOrEmptyString(resource.Type),
		MetricKind:  MonitoringMetricDescriptorMetricKindEnumToProto(resource.MetricKind),
		ValueType:   MonitoringMetricDescriptorValueTypeEnumToProto(resource.ValueType),
		Unit:        dcl.ValueOrEmptyString(resource.Unit),
		Description: dcl.ValueOrEmptyString(resource.Description),
		DisplayName: dcl.ValueOrEmptyString(resource.DisplayName),
		Metadata:    MonitoringMetricDescriptorMetadataToProto(resource.Metadata),
		LaunchStage: MonitoringMetricDescriptorLaunchStageEnumToProto(resource.LaunchStage),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.DescriptorLabels {
		p.DescriptorLabels = append(p.DescriptorLabels, MonitoringMetricDescriptorDescriptorLabelsToProto(&r))
	}
	for _, r := range resource.MonitoredResourceTypes {
		p.MonitoredResourceTypes = append(p.MonitoredResourceTypes, r)
	}

	return p
}

// ApplyMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Apply() method.
func (s *MetricDescriptorServer) applyMetricDescriptor(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringMetricDescriptorRequest) (*monitoringpb.MonitoringMetricDescriptor, error) {
	p := ProtoToMetricDescriptor(request.GetResource())
	res, err := c.ApplyMetricDescriptor(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetricDescriptorToProto(res)
	return r, nil
}

// ApplyMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Apply() method.
func (s *MetricDescriptorServer) ApplyMonitoringMetricDescriptor(ctx context.Context, request *monitoringpb.ApplyMonitoringMetricDescriptorRequest) (*monitoringpb.MonitoringMetricDescriptor, error) {
	cl, err := createConfigMetricDescriptor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyMetricDescriptor(ctx, cl, request)
}

// DeleteMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Delete() method.
func (s *MetricDescriptorServer) DeleteMonitoringMetricDescriptor(ctx context.Context, request *monitoringpb.DeleteMonitoringMetricDescriptorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMetricDescriptor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMetricDescriptor(ctx, ProtoToMetricDescriptor(request.GetResource()))

}

// ListMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptorList() method.
func (s *MetricDescriptorServer) ListMonitoringMetricDescriptor(ctx context.Context, request *monitoringpb.ListMonitoringMetricDescriptorRequest) (*monitoringpb.ListMonitoringMetricDescriptorResponse, error) {
	cl, err := createConfigMetricDescriptor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMetricDescriptor(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*monitoringpb.MonitoringMetricDescriptor
	for _, r := range resources.Items {
		rp := MetricDescriptorToProto(r)
		protos = append(protos, rp)
	}
	return &monitoringpb.ListMonitoringMetricDescriptorResponse{Items: protos}, nil
}

func createConfigMetricDescriptor(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
