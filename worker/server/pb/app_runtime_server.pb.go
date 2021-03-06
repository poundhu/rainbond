// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app_runtime_server.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServiceRequest struct {
	ServiceId            string   `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceRequest) Reset()         { *m = ServiceRequest{} }
func (m *ServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceRequest) ProtoMessage()    {}
func (*ServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f94cf1a886c479d6, []int{0}
}

func (m *ServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceRequest.Unmarshal(m, b)
}
func (m *ServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceRequest.Marshal(b, m, deterministic)
}
func (m *ServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceRequest.Merge(m, src)
}
func (m *ServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceRequest.Size(m)
}
func (m *ServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceRequest proto.InternalMessageInfo

func (m *ServiceRequest) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

type ServicesRequest struct {
	ServiceIds           string   `protobuf:"bytes,1,opt,name=service_ids,json=serviceIds,proto3" json:"service_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServicesRequest) Reset()         { *m = ServicesRequest{} }
func (m *ServicesRequest) String() string { return proto.CompactTextString(m) }
func (*ServicesRequest) ProtoMessage()    {}
func (*ServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f94cf1a886c479d6, []int{1}
}

func (m *ServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServicesRequest.Unmarshal(m, b)
}
func (m *ServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServicesRequest.Marshal(b, m, deterministic)
}
func (m *ServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServicesRequest.Merge(m, src)
}
func (m *ServicesRequest) XXX_Size() int {
	return xxx_messageInfo_ServicesRequest.Size(m)
}
func (m *ServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServicesRequest proto.InternalMessageInfo

func (m *ServicesRequest) GetServiceIds() string {
	if m != nil {
		return m.ServiceIds
	}
	return ""
}

type StatusMessage struct {
	Status               map[string]string `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StatusMessage) Reset()         { *m = StatusMessage{} }
func (m *StatusMessage) String() string { return proto.CompactTextString(m) }
func (*StatusMessage) ProtoMessage()    {}
func (*StatusMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f94cf1a886c479d6, []int{2}
}

func (m *StatusMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusMessage.Unmarshal(m, b)
}
func (m *StatusMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusMessage.Marshal(b, m, deterministic)
}
func (m *StatusMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusMessage.Merge(m, src)
}
func (m *StatusMessage) XXX_Size() int {
	return xxx_messageInfo_StatusMessage.Size(m)
}
func (m *StatusMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StatusMessage proto.InternalMessageInfo

func (m *StatusMessage) GetStatus() map[string]string {
	if m != nil {
		return m.Status
	}
	return nil
}

type DiskMessage struct {
	Disks                map[string]float64 `protobuf:"bytes,1,rep,name=disks,proto3" json:"disks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DiskMessage) Reset()         { *m = DiskMessage{} }
func (m *DiskMessage) String() string { return proto.CompactTextString(m) }
func (*DiskMessage) ProtoMessage()    {}
func (*DiskMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f94cf1a886c479d6, []int{3}
}

func (m *DiskMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiskMessage.Unmarshal(m, b)
}
func (m *DiskMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiskMessage.Marshal(b, m, deterministic)
}
func (m *DiskMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiskMessage.Merge(m, src)
}
func (m *DiskMessage) XXX_Size() int {
	return xxx_messageInfo_DiskMessage.Size(m)
}
func (m *DiskMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DiskMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DiskMessage proto.InternalMessageInfo

func (m *DiskMessage) GetDisks() map[string]float64 {
	if m != nil {
		return m.Disks
	}
	return nil
}

type ServiceAppPodList struct {
	Pods                 []*ServiceAppPod `protobuf:"bytes,1,rep,name=pods,proto3" json:"pods,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ServiceAppPodList) Reset()         { *m = ServiceAppPodList{} }
func (m *ServiceAppPodList) String() string { return proto.CompactTextString(m) }
func (*ServiceAppPodList) ProtoMessage()    {}
func (*ServiceAppPodList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f94cf1a886c479d6, []int{4}
}

func (m *ServiceAppPodList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAppPodList.Unmarshal(m, b)
}
func (m *ServiceAppPodList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAppPodList.Marshal(b, m, deterministic)
}
func (m *ServiceAppPodList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAppPodList.Merge(m, src)
}
func (m *ServiceAppPodList) XXX_Size() int {
	return xxx_messageInfo_ServiceAppPodList.Size(m)
}
func (m *ServiceAppPodList) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAppPodList.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAppPodList proto.InternalMessageInfo

func (m *ServiceAppPodList) GetPods() []*ServiceAppPod {
	if m != nil {
		return m.Pods
	}
	return nil
}

type ServiceAppPod struct {
	ServiceId            string                `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	DeployId             string                `protobuf:"bytes,2,opt,name=deploy_id,json=deployId,proto3" json:"deploy_id,omitempty"`
	DeployType           string                `protobuf:"bytes,3,opt,name=deploy_type,json=deployType,proto3" json:"deploy_type,omitempty"`
	PodName              string                `protobuf:"bytes,4,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	PodIp                string                `protobuf:"bytes,5,opt,name=pod_ip,json=podIp,proto3" json:"pod_ip,omitempty"`
	PodStatus            string                `protobuf:"bytes,6,opt,name=pod_status,json=podStatus,proto3" json:"pod_status,omitempty"`
	Containers           map[string]*Container `protobuf:"bytes,7,rep,name=containers,proto3" json:"containers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ServiceAppPod) Reset()         { *m = ServiceAppPod{} }
func (m *ServiceAppPod) String() string { return proto.CompactTextString(m) }
func (*ServiceAppPod) ProtoMessage()    {}
func (*ServiceAppPod) Descriptor() ([]byte, []int) {
	return fileDescriptor_f94cf1a886c479d6, []int{5}
}

func (m *ServiceAppPod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAppPod.Unmarshal(m, b)
}
func (m *ServiceAppPod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAppPod.Marshal(b, m, deterministic)
}
func (m *ServiceAppPod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAppPod.Merge(m, src)
}
func (m *ServiceAppPod) XXX_Size() int {
	return xxx_messageInfo_ServiceAppPod.Size(m)
}
func (m *ServiceAppPod) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAppPod.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAppPod proto.InternalMessageInfo

func (m *ServiceAppPod) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *ServiceAppPod) GetDeployId() string {
	if m != nil {
		return m.DeployId
	}
	return ""
}

func (m *ServiceAppPod) GetDeployType() string {
	if m != nil {
		return m.DeployType
	}
	return ""
}

func (m *ServiceAppPod) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *ServiceAppPod) GetPodIp() string {
	if m != nil {
		return m.PodIp
	}
	return ""
}

func (m *ServiceAppPod) GetPodStatus() string {
	if m != nil {
		return m.PodStatus
	}
	return ""
}

func (m *ServiceAppPod) GetContainers() map[string]*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

type Container struct {
	ContainerName        string   `protobuf:"bytes,1,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	MemoryLimit          int32    `protobuf:"varint,2,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_f94cf1a886c479d6, []int{6}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *Container) GetMemoryLimit() int32 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

type DeployInfo struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Statefuleset         string            `protobuf:"bytes,2,opt,name=statefuleset,proto3" json:"statefuleset,omitempty"`
	Deployment           string            `protobuf:"bytes,3,opt,name=deployment,proto3" json:"deployment,omitempty"`
	Pods                 map[string]string `protobuf:"bytes,4,rep,name=pods,proto3" json:"pods,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Services             map[string]string `protobuf:"bytes,5,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Secrets              map[string]string `protobuf:"bytes,6,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ingresses            map[string]string `protobuf:"bytes,7,rep,name=ingresses,proto3" json:"ingresses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Replicatset          map[string]string `protobuf:"bytes,8,rep,name=replicatset,proto3" json:"replicatset,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Status               string            `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DeployInfo) Reset()         { *m = DeployInfo{} }
func (m *DeployInfo) String() string { return proto.CompactTextString(m) }
func (*DeployInfo) ProtoMessage()    {}
func (*DeployInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f94cf1a886c479d6, []int{7}
}

func (m *DeployInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployInfo.Unmarshal(m, b)
}
func (m *DeployInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployInfo.Marshal(b, m, deterministic)
}
func (m *DeployInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployInfo.Merge(m, src)
}
func (m *DeployInfo) XXX_Size() int {
	return xxx_messageInfo_DeployInfo.Size(m)
}
func (m *DeployInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeployInfo proto.InternalMessageInfo

func (m *DeployInfo) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeployInfo) GetStatefuleset() string {
	if m != nil {
		return m.Statefuleset
	}
	return ""
}

func (m *DeployInfo) GetDeployment() string {
	if m != nil {
		return m.Deployment
	}
	return ""
}

func (m *DeployInfo) GetPods() map[string]string {
	if m != nil {
		return m.Pods
	}
	return nil
}

func (m *DeployInfo) GetServices() map[string]string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *DeployInfo) GetSecrets() map[string]string {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *DeployInfo) GetIngresses() map[string]string {
	if m != nil {
		return m.Ingresses
	}
	return nil
}

func (m *DeployInfo) GetReplicatset() map[string]string {
	if m != nil {
		return m.Replicatset
	}
	return nil
}

func (m *DeployInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceRequest)(nil), "pb.ServiceRequest")
	proto.RegisterType((*ServicesRequest)(nil), "pb.ServicesRequest")
	proto.RegisterType((*StatusMessage)(nil), "pb.StatusMessage")
	proto.RegisterMapType((map[string]string)(nil), "pb.StatusMessage.StatusEntry")
	proto.RegisterType((*DiskMessage)(nil), "pb.DiskMessage")
	proto.RegisterMapType((map[string]float64)(nil), "pb.DiskMessage.DisksEntry")
	proto.RegisterType((*ServiceAppPodList)(nil), "pb.ServiceAppPodList")
	proto.RegisterType((*ServiceAppPod)(nil), "pb.ServiceAppPod")
	proto.RegisterMapType((map[string]*Container)(nil), "pb.ServiceAppPod.ContainersEntry")
	proto.RegisterType((*Container)(nil), "pb.Container")
	proto.RegisterType((*DeployInfo)(nil), "pb.DeployInfo")
	proto.RegisterMapType((map[string]string)(nil), "pb.DeployInfo.IngressesEntry")
	proto.RegisterMapType((map[string]string)(nil), "pb.DeployInfo.PodsEntry")
	proto.RegisterMapType((map[string]string)(nil), "pb.DeployInfo.ReplicatsetEntry")
	proto.RegisterMapType((map[string]string)(nil), "pb.DeployInfo.SecretsEntry")
	proto.RegisterMapType((map[string]string)(nil), "pb.DeployInfo.ServicesEntry")
}

func init() { proto.RegisterFile("app_runtime_server.proto", fileDescriptor_f94cf1a886c479d6) }

var fileDescriptor_f94cf1a886c479d6 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x4f, 0x4f, 0xdb, 0x4c,
	0x10, 0xc6, 0x93, 0x40, 0x42, 0x3c, 0x4e, 0x02, 0xec, 0xfb, 0x52, 0xb9, 0x01, 0x0a, 0xb8, 0x42,
	0xe2, 0x50, 0xa5, 0x55, 0x5a, 0xd4, 0x00, 0x55, 0xa5, 0xa8, 0x54, 0x28, 0x12, 0xad, 0x90, 0x69,
	0xcf, 0x91, 0x89, 0x07, 0x64, 0x11, 0xdb, 0x5b, 0xef, 0x06, 0xc9, 0xc7, 0x7e, 0xc1, 0xde, 0xfa,
	0x45, 0xfa, 0x09, 0xaa, 0xfd, 0xe7, 0x38, 0x09, 0x12, 0x4a, 0x6f, 0xf6, 0x33, 0xf3, 0xdb, 0xf1,
	0x3e, 0xb3, 0x3b, 0x06, 0xc7, 0xa7, 0x74, 0x98, 0x4e, 0x62, 0x1e, 0x46, 0x38, 0x64, 0x98, 0x3e,
	0x60, 0xda, 0xa1, 0x69, 0xc2, 0x13, 0x52, 0xa1, 0x37, 0xee, 0x6b, 0x68, 0x5d, 0x63, 0xfa, 0x10,
	0x8e, 0xd0, 0xc3, 0x1f, 0x13, 0x64, 0x9c, 0xec, 0x02, 0x30, 0xa5, 0x0c, 0xc3, 0xc0, 0x29, 0xef,
	0x97, 0x8f, 0x2c, 0xcf, 0xd2, 0xca, 0x20, 0x70, 0xbb, 0xb0, 0xae, 0x01, 0x66, 0x88, 0x3d, 0xb0,
	0xa7, 0x04, 0xd3, 0x08, 0xe4, 0x08, 0x73, 0x7f, 0x96, 0xa1, 0x79, 0xcd, 0x7d, 0x3e, 0x61, 0x5f,
	0x90, 0x31, 0xff, 0x0e, 0xc9, 0x31, 0xd4, 0x98, 0x14, 0x9c, 0xf2, 0xfe, 0xca, 0x91, 0xdd, 0xdd,
	0xed, 0xd0, 0x9b, 0xce, 0x4c, 0x8a, 0x7e, 0xfb, 0x1c, 0xf3, 0x34, 0xf3, 0x74, 0x72, 0xfb, 0x04,
	0xec, 0x82, 0x4c, 0x36, 0x60, 0xe5, 0x1e, 0x33, 0x5d, 0x50, 0x3c, 0x92, 0xff, 0xa1, 0xfa, 0xe0,
	0x8f, 0x27, 0xe8, 0x54, 0xa4, 0xa6, 0x5e, 0x4e, 0x2b, 0xbd, 0xb2, 0x9b, 0x81, 0x7d, 0x1e, 0xb2,
	0x7b, 0xf3, 0x01, 0x6f, 0xa0, 0x1a, 0x84, 0xec, 0xde, 0xd4, 0x6f, 0x8b, 0xfa, 0x85, 0xb8, 0x7c,
	0xd6, 0xc5, 0x55, 0x62, 0xbb, 0x07, 0x30, 0x15, 0x9f, 0x2a, 0x5d, 0x2e, 0x96, 0x3e, 0x85, 0x4d,
	0x6d, 0x59, 0x9f, 0xd2, 0xab, 0x24, 0xb8, 0x0c, 0x19, 0x27, 0x87, 0xb0, 0x4a, 0x93, 0xc0, 0xd4,
	0xdf, 0x94, 0xfb, 0x2f, 0x26, 0x79, 0x32, 0xec, 0xfe, 0xaa, 0x40, 0x73, 0x46, 0x7f, 0xa2, 0x3f,
	0x64, 0x1b, 0xac, 0x00, 0xe9, 0x38, 0xc9, 0x44, 0x54, 0xb9, 0x50, 0x57, 0xc2, 0x20, 0x10, 0x9d,
	0xd2, 0x41, 0x9e, 0x51, 0x74, 0x56, 0x54, 0xa7, 0x94, 0xf4, 0x2d, 0xa3, 0x48, 0x9e, 0x43, 0x9d,
	0x26, 0xc1, 0x30, 0xf6, 0x23, 0x74, 0x56, 0x65, 0x74, 0x8d, 0x26, 0xc1, 0x57, 0x3f, 0x42, 0xb2,
	0x05, 0x35, 0x11, 0x0a, 0xa9, 0x53, 0x55, 0xde, 0xd2, 0x24, 0x18, 0x50, 0xf1, 0x39, 0x42, 0xd6,
	0xdd, 0xac, 0xa9, 0xcf, 0xa1, 0x49, 0xa0, 0xfa, 0x44, 0xfa, 0x00, 0xa3, 0x24, 0xe6, 0x7e, 0x18,
	0x63, 0xca, 0x9c, 0x35, 0xb9, 0xd9, 0x83, 0x85, 0xcd, 0x76, 0x3e, 0xe5, 0x39, 0xca, 0xf3, 0x02,
	0xd4, 0xbe, 0x84, 0xf5, 0xb9, 0xf0, 0x23, 0xee, 0xbf, 0x2c, 0xba, 0x6f, 0x77, 0x9b, 0xa2, 0x44,
	0x4e, 0x15, 0x9b, 0xf1, 0x1d, 0xac, 0x5c, 0x27, 0x87, 0xd0, 0xca, 0x0b, 0xa9, 0x4d, 0xab, 0x25,
	0x9b, 0xb9, 0x2a, 0xb7, 0x7e, 0x00, 0x8d, 0x08, 0xa3, 0x24, 0xcd, 0x86, 0xe3, 0x30, 0x0a, 0xb9,
	0xac, 0x51, 0xf5, 0x6c, 0xa5, 0x5d, 0x0a, 0xc9, 0xfd, 0x5d, 0x05, 0x38, 0x57, 0x36, 0xc7, 0xb7,
	0x09, 0xd9, 0x01, 0x4b, 0x2c, 0xc7, 0xa8, 0x3f, 0x32, 0x6b, 0x4e, 0x05, 0xe2, 0x42, 0x43, 0xf8,
	0x85, 0xb7, 0x93, 0x31, 0x32, 0xe4, 0xba, 0x4d, 0x33, 0x1a, 0x79, 0x01, 0xba, 0x2f, 0x11, 0xc6,
	0x7c, 0xb6, 0x53, 0x42, 0x21, 0xaf, 0xf4, 0xf9, 0x59, 0x95, 0x96, 0x3a, 0xf2, 0xfc, 0xe6, 0xf5,
	0x3b, 0x57, 0x49, 0xa0, 0x9d, 0x94, 0x59, 0xa4, 0x07, 0x75, 0x7d, 0x44, 0x98, 0x53, 0x95, 0xc4,
	0xce, 0x1c, 0x61, 0x2e, 0xb5, 0xa2, 0xf2, 0x6c, 0x72, 0x0c, 0x6b, 0x0c, 0x47, 0x29, 0x72, 0xd1,
	0x5c, 0x01, 0x6e, 0x2f, 0x80, 0x32, 0xaa, 0x38, 0x93, 0x4b, 0xce, 0xc0, 0x0a, 0xe3, 0xbb, 0x14,
	0x19, 0x43, 0xd3, 0xf6, 0xdd, 0x39, 0x70, 0x60, 0xe2, 0x0a, 0x9d, 0xe6, 0x93, 0x3e, 0xd8, 0x29,
	0xd2, 0x71, 0x38, 0xf2, 0xb9, 0xb0, 0xa7, 0x2e, 0xf1, 0xbd, 0x39, 0xdc, 0x9b, 0x66, 0xa8, 0x05,
	0x8a, 0x0c, 0x79, 0x96, 0x0f, 0x18, 0x4b, 0x5a, 0x67, 0x26, 0xc8, 0x7b, 0xb0, 0x72, 0x6f, 0x96,
	0x99, 0x1f, 0xed, 0xb3, 0xfc, 0x1e, 0xfe, 0x03, 0x7c, 0x0a, 0x8d, 0xa2, 0x4d, 0x4b, 0xb1, 0x1f,
	0xa0, 0x35, 0xeb, 0xd4, 0x52, 0xf4, 0x47, 0xd8, 0x98, 0x37, 0x6a, 0x19, 0xbe, 0xfb, 0xa7, 0x0c,
	0xad, 0x3e, 0xa5, 0x9e, 0xfa, 0x7f, 0x5c, 0x67, 0xf1, 0x88, 0xf4, 0xa0, 0x71, 0x81, 0xbc, 0x4f,
	0xa9, 0xbe, 0xe2, 0xff, 0x15, 0xae, 0xb3, 0xf9, 0x27, 0xb4, 0x37, 0x17, 0x06, 0xba, 0x5b, 0x22,
	0xef, 0x00, 0x14, 0x29, 0x06, 0xe9, 0xe3, 0xdc, 0xfa, 0xdc, 0x20, 0x76, 0x4b, 0xe4, 0xc4, 0x50,
	0xa2, 0x71, 0x84, 0x14, 0x28, 0x03, 0x6d, 0x2d, 0x0c, 0x14, 0x31, 0x62, 0xdd, 0x12, 0x39, 0x86,
	0xe6, 0x05, 0xf2, 0xc2, 0xbd, 0x7c, 0x8c, 0x6e, 0xcd, 0x1e, 0x2c, 0xb7, 0x74, 0x53, 0x93, 0xff,
	0xc7, 0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x1e, 0xbb, 0x9c, 0x3b, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppRuntimeSyncClient is the client API for AppRuntimeSync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppRuntimeSyncClient interface {
	GetAppStatus(ctx context.Context, in *ServicesRequest, opts ...grpc.CallOption) (*StatusMessage, error)
	GetAppDisk(ctx context.Context, in *ServicesRequest, opts ...grpc.CallOption) (*DiskMessage, error)
	GetAppPods(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*ServiceAppPodList, error)
	GetDeployInfo(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*DeployInfo, error)
}

type appRuntimeSyncClient struct {
	cc *grpc.ClientConn
}

func NewAppRuntimeSyncClient(cc *grpc.ClientConn) AppRuntimeSyncClient {
	return &appRuntimeSyncClient{cc}
}

func (c *appRuntimeSyncClient) GetAppStatus(ctx context.Context, in *ServicesRequest, opts ...grpc.CallOption) (*StatusMessage, error) {
	out := new(StatusMessage)
	err := c.cc.Invoke(ctx, "/pb.AppRuntimeSync/GetAppStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRuntimeSyncClient) GetAppDisk(ctx context.Context, in *ServicesRequest, opts ...grpc.CallOption) (*DiskMessage, error) {
	out := new(DiskMessage)
	err := c.cc.Invoke(ctx, "/pb.AppRuntimeSync/GetAppDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRuntimeSyncClient) GetAppPods(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*ServiceAppPodList, error) {
	out := new(ServiceAppPodList)
	err := c.cc.Invoke(ctx, "/pb.AppRuntimeSync/GetAppPods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRuntimeSyncClient) GetDeployInfo(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*DeployInfo, error) {
	out := new(DeployInfo)
	err := c.cc.Invoke(ctx, "/pb.AppRuntimeSync/GetDeployInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppRuntimeSyncServer is the server API for AppRuntimeSync service.
type AppRuntimeSyncServer interface {
	GetAppStatus(context.Context, *ServicesRequest) (*StatusMessage, error)
	GetAppDisk(context.Context, *ServicesRequest) (*DiskMessage, error)
	GetAppPods(context.Context, *ServiceRequest) (*ServiceAppPodList, error)
	GetDeployInfo(context.Context, *ServiceRequest) (*DeployInfo, error)
}

func RegisterAppRuntimeSyncServer(s *grpc.Server, srv AppRuntimeSyncServer) {
	s.RegisterService(&_AppRuntimeSync_serviceDesc, srv)
}

func _AppRuntimeSync_GetAppStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRuntimeSyncServer).GetAppStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AppRuntimeSync/GetAppStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRuntimeSyncServer).GetAppStatus(ctx, req.(*ServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRuntimeSync_GetAppDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRuntimeSyncServer).GetAppDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AppRuntimeSync/GetAppDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRuntimeSyncServer).GetAppDisk(ctx, req.(*ServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRuntimeSync_GetAppPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRuntimeSyncServer).GetAppPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AppRuntimeSync/GetAppPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRuntimeSyncServer).GetAppPods(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRuntimeSync_GetDeployInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRuntimeSyncServer).GetDeployInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AppRuntimeSync/GetDeployInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRuntimeSyncServer).GetDeployInfo(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppRuntimeSync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AppRuntimeSync",
	HandlerType: (*AppRuntimeSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppStatus",
			Handler:    _AppRuntimeSync_GetAppStatus_Handler,
		},
		{
			MethodName: "GetAppDisk",
			Handler:    _AppRuntimeSync_GetAppDisk_Handler,
		},
		{
			MethodName: "GetAppPods",
			Handler:    _AppRuntimeSync_GetAppPods_Handler,
		},
		{
			MethodName: "GetDeployInfo",
			Handler:    _AppRuntimeSync_GetDeployInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app_runtime_server.proto",
}
