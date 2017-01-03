// Code generated by protoc-gen-go.
// source: certificate.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	certificate.proto

It has these top-level messages:
	CertificateListRequest
	CertificateListResponse
	CertificateDescribeRequest
	CertificateDescribeResponse
	Certificate
	CertificateLoadRequest
	CertificateDeleteRequest
	CertificateDeployRequest
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CertificateListRequest struct {
}

func (m *CertificateListRequest) Reset()                    { *m = CertificateListRequest{} }
func (m *CertificateListRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateListRequest) ProtoMessage()               {}
func (*CertificateListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CertificateListResponse struct {
	Status       *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificates []*Certificate          `protobuf:"bytes,2,rep,name=certificates" json:"certificates,omitempty"`
}

func (m *CertificateListResponse) Reset()                    { *m = CertificateListResponse{} }
func (m *CertificateListResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateListResponse) ProtoMessage()               {}
func (*CertificateListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CertificateListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CertificateListResponse) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type CertificateDescribeRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateDescribeRequest) Reset()                    { *m = CertificateDescribeRequest{} }
func (m *CertificateDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDescribeRequest) ProtoMessage()               {}
func (*CertificateDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CertificateDescribeRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type CertificateDescribeResponse struct {
	Status      *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificate *Certificate            `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
}

func (m *CertificateDescribeResponse) Reset()                    { *m = CertificateDescribeResponse{} }
func (m *CertificateDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateDescribeResponse) ProtoMessage()               {}
func (*CertificateDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CertificateDescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CertificateDescribeResponse) GetCertificate() *Certificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type Certificate struct {
	Phid       string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CommonName string `protobuf:"bytes,3,opt,name=common_name,json=commonName" json:"common_name,omitempty"`
	IssuedBy   string `protobuf:"bytes,4,opt,name=issued_by,json=issuedBy" json:"issued_by,omitempty"`
	ValidFrom  int64  `protobuf:"varint,5,opt,name=valid_from,json=validFrom" json:"valid_from,omitempty"`
	ExpireDate int64  `protobuf:"varint,6,opt,name=expire_date,json=expireDate" json:"expire_date,omitempty"`
	// those feilds will not included into list response.
	// only describe response will include the underlying
	// feilds.
	Sans         []string `protobuf:"bytes,7,rep,name=sans" json:"sans,omitempty"`
	Cert         string   `protobuf:"bytes,8,opt,name=cert" json:"cert,omitempty"`
	Key          string   `protobuf:"bytes,9,opt,name=key" json:"key,omitempty"`
	Version      int32    `protobuf:"varint,10,opt,name=version" json:"version,omitempty"`
	SerialNumber string   `protobuf:"bytes,11,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
}

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Certificate) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Certificate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Certificate) GetCommonName() string {
	if m != nil {
		return m.CommonName
	}
	return ""
}

func (m *Certificate) GetIssuedBy() string {
	if m != nil {
		return m.IssuedBy
	}
	return ""
}

func (m *Certificate) GetValidFrom() int64 {
	if m != nil {
		return m.ValidFrom
	}
	return 0
}

func (m *Certificate) GetExpireDate() int64 {
	if m != nil {
		return m.ExpireDate
	}
	return 0
}

func (m *Certificate) GetSans() []string {
	if m != nil {
		return m.Sans
	}
	return nil
}

func (m *Certificate) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func (m *Certificate) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Certificate) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Certificate) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

type CertificateLoadRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	CertData string `protobuf:"bytes,2,opt,name=cert_data,json=certData" json:"cert_data,omitempty"`
	KeyData  string `protobuf:"bytes,3,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
}

func (m *CertificateLoadRequest) Reset()                    { *m = CertificateLoadRequest{} }
func (m *CertificateLoadRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateLoadRequest) ProtoMessage()               {}
func (*CertificateLoadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CertificateLoadRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CertificateLoadRequest) GetCertData() string {
	if m != nil {
		return m.CertData
	}
	return ""
}

func (m *CertificateLoadRequest) GetKeyData() string {
	if m != nil {
		return m.KeyData
	}
	return ""
}

type CertificateDeleteRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateDeleteRequest) Reset()                    { *m = CertificateDeleteRequest{} }
func (m *CertificateDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDeleteRequest) ProtoMessage()               {}
func (*CertificateDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CertificateDeleteRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type CertificateDeployRequest struct {
	Uid         string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	SecretName  string `protobuf:"bytes,2,opt,name=secret_name,json=secretName" json:"secret_name,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	Namespace   string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *CertificateDeployRequest) Reset()                    { *m = CertificateDeployRequest{} }
func (m *CertificateDeployRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDeployRequest) ProtoMessage()               {}
func (*CertificateDeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CertificateDeployRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CertificateDeployRequest) GetSecretName() string {
	if m != nil {
		return m.SecretName
	}
	return ""
}

func (m *CertificateDeployRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *CertificateDeployRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func init() {
	proto.RegisterType((*CertificateListRequest)(nil), "appscode.certificate.v1beta1.CertificateListRequest")
	proto.RegisterType((*CertificateListResponse)(nil), "appscode.certificate.v1beta1.CertificateListResponse")
	proto.RegisterType((*CertificateDescribeRequest)(nil), "appscode.certificate.v1beta1.CertificateDescribeRequest")
	proto.RegisterType((*CertificateDescribeResponse)(nil), "appscode.certificate.v1beta1.CertificateDescribeResponse")
	proto.RegisterType((*Certificate)(nil), "appscode.certificate.v1beta1.Certificate")
	proto.RegisterType((*CertificateLoadRequest)(nil), "appscode.certificate.v1beta1.CertificateLoadRequest")
	proto.RegisterType((*CertificateDeleteRequest)(nil), "appscode.certificate.v1beta1.CertificateDeleteRequest")
	proto.RegisterType((*CertificateDeployRequest)(nil), "appscode.certificate.v1beta1.CertificateDeployRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Certificates service

type CertificatesClient interface {
	List(ctx context.Context, in *CertificateListRequest, opts ...grpc.CallOption) (*CertificateListResponse, error)
	Describe(ctx context.Context, in *CertificateDescribeRequest, opts ...grpc.CallOption) (*CertificateDescribeResponse, error)
	Load(ctx context.Context, in *CertificateLoadRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *CertificateDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Deploy(ctx context.Context, in *CertificateDeployRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type certificatesClient struct {
	cc *grpc.ClientConn
}

func NewCertificatesClient(cc *grpc.ClientConn) CertificatesClient {
	return &certificatesClient{cc}
}

func (c *certificatesClient) List(ctx context.Context, in *CertificateListRequest, opts ...grpc.CallOption) (*CertificateListResponse, error) {
	out := new(CertificateListResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Describe(ctx context.Context, in *CertificateDescribeRequest, opts ...grpc.CallOption) (*CertificateDescribeResponse, error) {
	out := new(CertificateDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Load(ctx context.Context, in *CertificateLoadRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/Load", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Delete(ctx context.Context, in *CertificateDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Deploy(ctx context.Context, in *CertificateDeployRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Certificates service

type CertificatesServer interface {
	List(context.Context, *CertificateListRequest) (*CertificateListResponse, error)
	Describe(context.Context, *CertificateDescribeRequest) (*CertificateDescribeResponse, error)
	Load(context.Context, *CertificateLoadRequest) (*appscode_dtypes.VoidResponse, error)
	Delete(context.Context, *CertificateDeleteRequest) (*appscode_dtypes.VoidResponse, error)
	Deploy(context.Context, *CertificateDeployRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterCertificatesServer(s *grpc.Server, srv CertificatesServer) {
	s.RegisterService(&_Certificates_serviceDesc, srv)
}

func _Certificates_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.certificate.v1beta1.Certificates/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).List(ctx, req.(*CertificateListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.certificate.v1beta1.Certificates/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Describe(ctx, req.(*CertificateDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateLoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.certificate.v1beta1.Certificates/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Load(ctx, req.(*CertificateLoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.certificate.v1beta1.Certificates/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Delete(ctx, req.(*CertificateDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateDeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.certificate.v1beta1.Certificates/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Deploy(ctx, req.(*CertificateDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Certificates_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.certificate.v1beta1.Certificates",
	HandlerType: (*CertificatesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Certificates_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Certificates_Describe_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _Certificates_Load_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Certificates_Delete_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _Certificates_Deploy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "certificate.proto",
}

func init() { proto.RegisterFile("certificate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 754 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xbb, 0x6e, 0x1b, 0x39,
	0x14, 0x05, 0x25, 0x5b, 0x8f, 0x2b, 0x2d, 0xe0, 0x65, 0xb1, 0x9e, 0x1d, 0xdb, 0x6b, 0xed, 0xec,
	0x16, 0x5a, 0x63, 0xa1, 0x81, 0xbd, 0xeb, 0x3c, 0x0c, 0xa4, 0xb1, 0x95, 0x00, 0x79, 0x19, 0x86,
	0x82, 0xb8, 0x48, 0x23, 0x50, 0x33, 0xb4, 0x43, 0x48, 0x1a, 0x4e, 0x86, 0x94, 0x91, 0x41, 0xe0,
	0xc6, 0x1f, 0x90, 0x26, 0x55, 0x90, 0x26, 0x45, 0xca, 0x54, 0x09, 0x92, 0x2f, 0xc8, 0x1f, 0xe4,
	0x17, 0xf2, 0x21, 0x01, 0xc9, 0x91, 0x44, 0xf9, 0x2d, 0x35, 0x02, 0x75, 0x0e, 0xef, 0xf0, 0x9c,
	0xfb, 0x20, 0xe1, 0xd7, 0x80, 0x26, 0x92, 0x1d, 0xb0, 0x80, 0x48, 0xda, 0x88, 0x13, 0x2e, 0x39,
	0x5e, 0x26, 0x71, 0x2c, 0x02, 0x1e, 0xd2, 0x86, 0xcd, 0x1d, 0xad, 0x77, 0xa8, 0x24, 0xeb, 0xee,
	0xf2, 0x21, 0xe7, 0x87, 0x3d, 0xea, 0x93, 0x98, 0xf9, 0x24, 0x8a, 0xb8, 0x24, 0x92, 0xf1, 0x48,
	0x98, 0x58, 0xf7, 0x8f, 0x61, 0xec, 0x05, 0xfc, 0xea, 0x04, 0x1f, 0xca, 0x34, 0xa6, 0xc2, 0xd7,
	0xbf, 0x66, 0x83, 0xe7, 0xc0, 0x6f, 0x3b, 0xe3, 0x53, 0x1f, 0x31, 0x21, 0x5b, 0xf4, 0xc5, 0x80,
	0x0a, 0xe9, 0xbd, 0x45, 0xb0, 0x78, 0x86, 0x12, 0x31, 0x8f, 0x04, 0xc5, 0x3e, 0x14, 0x84, 0x24,
	0x72, 0x20, 0x1c, 0x54, 0x43, 0xf5, 0xca, 0xc6, 0x62, 0x63, 0xe4, 0xc1, 0x9c, 0xd1, 0x78, 0xa2,
	0xe9, 0x56, 0xb6, 0x0d, 0x3f, 0x86, 0xaa, 0x65, 0x4e, 0x38, 0xb9, 0x5a, 0xbe, 0x5e, 0xd9, 0xf8,
	0xa7, 0x71, 0x99, 0xf5, 0x86, 0x75, 0x7a, 0x6b, 0x22, 0xdc, 0x6b, 0x80, 0x6b, 0x91, 0x4d, 0x2a,
	0x82, 0x84, 0x75, 0x68, 0xa6, 0x1c, 0x2f, 0x40, 0x7e, 0xc0, 0x42, 0x2d, 0xad, 0xdc, 0x52, 0x4b,
	0xef, 0x1d, 0x82, 0xa5, 0x73, 0x03, 0x66, 0xf5, 0xf3, 0x10, 0x2a, 0x96, 0x20, 0x27, 0xa7, 0xa3,
	0xa6, 0xb0, 0x63, 0x47, 0x7b, 0x1f, 0x73, 0x50, 0xb1, 0x48, 0x8c, 0x61, 0x2e, 0x7e, 0x3e, 0x32,
	0xa0, 0xd7, 0x0a, 0x8b, 0x48, 0xdf, 0x9c, 0x54, 0x6e, 0xe9, 0x35, 0x5e, 0x85, 0x4a, 0xc0, 0xfb,
	0x7d, 0x1e, 0xb5, 0x35, 0x95, 0xd7, 0x14, 0x18, 0x68, 0x57, 0x6d, 0x58, 0x82, 0x32, 0x13, 0x62,
	0x40, 0xc3, 0x76, 0x27, 0x75, 0xe6, 0x34, 0x5d, 0x32, 0xc0, 0x76, 0x8a, 0x57, 0x00, 0x8e, 0x48,
	0x8f, 0x85, 0xed, 0x83, 0x84, 0xf7, 0x9d, 0xf9, 0x1a, 0xaa, 0xe7, 0x5b, 0x65, 0x8d, 0xdc, 0x4b,
	0x78, 0x5f, 0x7d, 0x9c, 0xbe, 0x8c, 0x59, 0x42, 0xdb, 0xa1, 0x72, 0x58, 0xd0, 0x3c, 0x18, 0xa8,
	0x99, 0xa9, 0x14, 0x24, 0x12, 0x4e, 0xb1, 0x96, 0x57, 0x8a, 0xd4, 0x5a, 0x61, 0xca, 0x98, 0x53,
	0x32, 0x2a, 0xd5, 0x5a, 0x55, 0xa3, 0x4b, 0x53, 0xa7, 0x6c, 0xaa, 0xd1, 0xa5, 0x29, 0x76, 0xa0,
	0x78, 0x44, 0x13, 0xc1, 0x78, 0xe4, 0x40, 0x0d, 0xd5, 0xe7, 0x5b, 0xc3, 0xbf, 0xf8, 0x2f, 0xf8,
	0x45, 0xd0, 0x84, 0x91, 0x5e, 0x3b, 0x1a, 0xf4, 0x3b, 0x34, 0x71, 0x2a, 0x3a, 0xaa, 0x6a, 0xc0,
	0x5d, 0x8d, 0x79, 0xe1, 0x64, 0xcb, 0x72, 0x12, 0x0e, 0x0b, 0x3f, 0x4c, 0x12, 0xb2, 0x92, 0xb4,
	0x04, 0x65, 0x25, 0x43, 0xb9, 0x20, 0x59, 0xf6, 0x4a, 0x0a, 0x68, 0x12, 0x49, 0xf0, 0xef, 0x50,
	0xea, 0xd2, 0xd4, 0x70, 0x26, 0x7d, 0xc5, 0x2e, 0x4d, 0x15, 0xe5, 0xfd, 0x0b, 0xce, 0x44, 0xc7,
	0xf4, 0xa8, 0xbc, 0xa4, 0xc1, 0x5e, 0xa3, 0x53, 0xdb, 0xe3, 0x1e, 0x4f, 0x2f, 0xdc, 0xae, 0x92,
	0x2b, 0x68, 0x90, 0x50, 0xd9, 0xb6, 0x8a, 0x0a, 0x06, 0xd2, 0x95, 0xfb, 0x13, 0xaa, 0x41, 0x6f,
	0x20, 0x24, 0x4d, 0xec, 0xda, 0x56, 0x32, 0x4c, 0x6f, 0x59, 0x86, 0xb2, 0xa2, 0x44, 0x4c, 0x02,
	0x9a, 0x15, 0x77, 0x0c, 0x6c, 0xbc, 0x2f, 0x42, 0xd5, 0x12, 0x24, 0xf0, 0x27, 0x04, 0x73, 0x6a,
	0x86, 0xf1, 0xff, 0xd7, 0xee, 0x52, 0xeb, 0x36, 0x70, 0x37, 0xa7, 0x8c, 0x32, 0x83, 0xe5, 0xdd,
	0x39, 0xf9, 0xe2, 0xe4, 0x4a, 0xe8, 0xe4, 0xfb, 0x8f, 0x37, 0xb9, 0x75, 0xec, 0xfb, 0xed, 0x89,
	0xfb, 0xc8, 0xfa, 0x92, 0x9f, 0x7d, 0xc9, 0xc6, 0x04, 0xfe, 0x86, 0xa0, 0x34, 0x1c, 0x56, 0x7c,
	0xeb, 0xda, 0x12, 0x4e, 0x5d, 0x08, 0xee, 0xed, 0x19, 0x22, 0x33, 0x03, 0x3b, 0x96, 0x81, 0x9b,
	0x78, 0x73, 0x4a, 0x03, 0xfe, 0xab, 0x01, 0x0b, 0x8f, 0xf1, 0x67, 0x95, 0x7b, 0x4e, 0xc2, 0x69,
	0x72, 0x3f, 0x6e, 0x6b, 0x77, 0xe5, 0xcc, 0x6d, 0xb4, 0xcf, 0x59, 0x38, 0x92, 0xb8, 0x6f, 0x49,
	0x7c, 0xe0, 0xde, 0x9d, 0x5a, 0xa2, 0x6a, 0x9a, 0x63, 0x9f, 0x04, 0xfa, 0xdd, 0xf0, 0x79, 0x47,
	0x12, 0x16, 0x6d, 0xa1, 0x35, 0xfc, 0x01, 0x41, 0xc1, 0xf4, 0x3d, 0xbe, 0x31, 0x45, 0x02, 0xad,
	0x41, 0xb9, 0x4a, 0xf9, 0x44, 0x72, 0xd7, 0x66, 0x4c, 0xee, 0x57, 0x2d, 0x53, 0xcd, 0xdb, 0x54,
	0x32, 0xad, 0x01, 0xbd, 0x4a, 0xe6, 0x53, 0x4b, 0xe6, 0x7d, 0xb7, 0x39, 0x93, 0xcc, 0x51, 0x7e,
	0x43, 0x7d, 0xf2, 0x16, 0x5a, 0xdb, 0xde, 0x81, 0xbf, 0x03, 0xde, 0x1f, 0x1f, 0x4d, 0x62, 0x76,
	0x9e, 0xec, 0xed, 0x05, 0x4b, 0xf7, 0x9e, 0x7a, 0xb4, 0xf7, 0xd0, 0xb3, 0x62, 0x46, 0x76, 0x0a,
	0xfa, 0x19, 0xff, 0xef, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0x3b, 0x39, 0xcf, 0x58, 0x08,
	0x00, 0x00,
}