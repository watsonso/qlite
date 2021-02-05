// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OptionCode int32

const (
	OptionCode_NO_ACTION OptionCode = 0
	OptionCode_CREATE    OptionCode = 1
	OptionCode_UPDATE    OptionCode = 2
	OptionCode_DELETE    OptionCode = 3
)

var OptionCode_name = map[int32]string{
	0: "NO_ACTION",
	1: "CREATE",
	2: "UPDATE",
	3: "DELETE",
}

var OptionCode_value = map[string]int32{
	"NO_ACTION": 0,
	"CREATE":    1,
	"UPDATE":    2,
	"DELETE":    3,
}

func (x OptionCode) String() string {
	return proto.EnumName(OptionCode_name, int32(x))
}

func (OptionCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (m *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(m, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

type Pong struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StlInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsOrderly            bool     `protobuf:"varint,2,opt,name=isOrderly,proto3" json:"isOrderly,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StlInfo) Reset()         { *m = StlInfo{} }
func (m *StlInfo) String() string { return proto.CompactTextString(m) }
func (*StlInfo) ProtoMessage()    {}
func (*StlInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *StlInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StlInfo.Unmarshal(m, b)
}
func (m *StlInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StlInfo.Marshal(b, m, deterministic)
}
func (m *StlInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StlInfo.Merge(m, src)
}
func (m *StlInfo) XXX_Size() int {
	return xxx_messageInfo_StlInfo.Size(m)
}
func (m *StlInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StlInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StlInfo proto.InternalMessageInfo

func (m *StlInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StlInfo) GetIsOrderly() bool {
	if m != nil {
		return m.IsOrderly
	}
	return false
}

type ApiDescription struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiDescription) Reset()         { *m = ApiDescription{} }
func (m *ApiDescription) String() string { return proto.CompactTextString(m) }
func (*ApiDescription) ProtoMessage()    {}
func (*ApiDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *ApiDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiDescription.Unmarshal(m, b)
}
func (m *ApiDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiDescription.Marshal(b, m, deterministic)
}
func (m *ApiDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiDescription.Merge(m, src)
}
func (m *ApiDescription) XXX_Size() int {
	return xxx_messageInfo_ApiDescription.Size(m)
}
func (m *ApiDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiDescription.DiscardUnknown(m)
}

var xxx_messageInfo_ApiDescription proto.InternalMessageInfo

func (m *ApiDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApiDescription) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type ApiList struct {
	List                 []*ApiDescription `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ApiList) Reset()         { *m = ApiList{} }
func (m *ApiList) String() string { return proto.CompactTextString(m) }
func (*ApiList) ProtoMessage()    {}
func (*ApiList) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *ApiList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiList.Unmarshal(m, b)
}
func (m *ApiList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiList.Marshal(b, m, deterministic)
}
func (m *ApiList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiList.Merge(m, src)
}
func (m *ApiList) XXX_Size() int {
	return xxx_messageInfo_ApiList.Size(m)
}
func (m *ApiList) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiList.DiscardUnknown(m)
}

var xxx_messageInfo_ApiList proto.InternalMessageInfo

func (m *ApiList) GetList() []*ApiDescription {
	if m != nil {
		return m.List
	}
	return nil
}

type ApiInfo struct {
	IsWriting            bool     `protobuf:"varint,1,opt,name=isWriting,proto3" json:"isWriting,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiInfo) Reset()         { *m = ApiInfo{} }
func (m *ApiInfo) String() string { return proto.CompactTextString(m) }
func (*ApiInfo) ProtoMessage()    {}
func (*ApiInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *ApiInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiInfo.Unmarshal(m, b)
}
func (m *ApiInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiInfo.Marshal(b, m, deterministic)
}
func (m *ApiInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiInfo.Merge(m, src)
}
func (m *ApiInfo) XXX_Size() int {
	return xxx_messageInfo_ApiInfo.Size(m)
}
func (m *ApiInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ApiInfo proto.InternalMessageInfo

func (m *ApiInfo) GetIsWriting() bool {
	if m != nil {
		return m.IsWriting
	}
	return false
}

func (m *ApiInfo) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ApiMap struct {
	Api                  map[string]*ApiInfo `protobuf:"bytes,1,rep,name=api,proto3" json:"api,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ApiMap) Reset()         { *m = ApiMap{} }
func (m *ApiMap) String() string { return proto.CompactTextString(m) }
func (*ApiMap) ProtoMessage()    {}
func (*ApiMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *ApiMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiMap.Unmarshal(m, b)
}
func (m *ApiMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiMap.Marshal(b, m, deterministic)
}
func (m *ApiMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiMap.Merge(m, src)
}
func (m *ApiMap) XXX_Size() int {
	return xxx_messageInfo_ApiMap.Size(m)
}
func (m *ApiMap) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiMap.DiscardUnknown(m)
}

var xxx_messageInfo_ApiMap proto.InternalMessageInfo

func (m *ApiMap) GetApi() map[string]*ApiInfo {
	if m != nil {
		return m.Api
	}
	return nil
}

type PendingMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Opt                  string   `protobuf:"bytes,2,opt,name=opt,proto3" json:"opt,omitempty"`
	Msg                  []byte   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Time                 int64    `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PendingMessage) Reset()         { *m = PendingMessage{} }
func (m *PendingMessage) String() string { return proto.CompactTextString(m) }
func (*PendingMessage) ProtoMessage()    {}
func (*PendingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *PendingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PendingMessage.Unmarshal(m, b)
}
func (m *PendingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PendingMessage.Marshal(b, m, deterministic)
}
func (m *PendingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingMessage.Merge(m, src)
}
func (m *PendingMessage) XXX_Size() int {
	return xxx_messageInfo_PendingMessage.Size(m)
}
func (m *PendingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PendingMessage proto.InternalMessageInfo

func (m *PendingMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PendingMessage) GetOpt() string {
	if m != nil {
		return m.Opt
	}
	return ""
}

func (m *PendingMessage) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *PendingMessage) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type Result struct {
	Code                 OptionCode `protobuf:"varint,1,opt,name=code,proto3,enum=OptionCode" json:"code,omitempty"`
	NewId                string     `protobuf:"bytes,2,opt,name=newId,proto3" json:"newId,omitempty"`
	Msg                  []byte     `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCode() OptionCode {
	if m != nil {
		return m.Code
	}
	return OptionCode_NO_ACTION
}

func (m *Result) GetNewId() string {
	if m != nil {
		return m.NewId
	}
	return ""
}

func (m *Result) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type Request struct {
	Id                   []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Opt                  string   `protobuf:"bytes,2,opt,name=opt,proto3" json:"opt,omitempty"`
	Msg                  []byte   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Request) GetOpt() string {
	if m != nil {
		return m.Opt
	}
	return ""
}

func (m *Request) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type Response struct {
	Msg                  []byte   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterEnum("OptionCode", OptionCode_name, OptionCode_value)
	proto.RegisterType((*Null)(nil), "null")
	proto.RegisterType((*Pong)(nil), "pong")
	proto.RegisterType((*StlInfo)(nil), "stlInfo")
	proto.RegisterType((*ApiDescription)(nil), "apiDescription")
	proto.RegisterType((*ApiList)(nil), "apiList")
	proto.RegisterType((*ApiInfo)(nil), "apiInfo")
	proto.RegisterType((*ApiMap)(nil), "apiMap")
	proto.RegisterMapType((map[string]*ApiInfo)(nil), "apiMap.ApiEntry")
	proto.RegisterType((*PendingMessage)(nil), "pendingMessage")
	proto.RegisterType((*Result)(nil), "result")
	proto.RegisterType((*Request)(nil), "request")
	proto.RegisterType((*Response)(nil), "response")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x14, 0x24, 0x4d, 0x9a, 0x8f, 0xb7, 0xd0, 0xad, 0xac, 0x45, 0xaa, 0xca, 0x6a, 0xa9, 0xcc, 0x25,
	0xe2, 0xe0, 0x43, 0xb9, 0xac, 0x40, 0x08, 0xaa, 0x36, 0x87, 0x4a, 0xbb, 0x5b, 0x88, 0x8a, 0xe0,
	0x86, 0xb2, 0x8d, 0x37, 0xb2, 0x36, 0xb5, 0x4d, 0xec, 0x80, 0xfa, 0xf7, 0xf8, 0x65, 0xc8, 0x8e,
	0xb3, 0x1f, 0x82, 0x03, 0xb7, 0xe9, 0x1b, 0xf5, 0xbd, 0xf1, 0xcc, 0x04, 0x92, 0x42, 0x32, 0x22,
	0x1b, 0xa1, 0x05, 0x0e, 0x21, 0xe0, 0x6d, 0x5d, 0xe3, 0x09, 0x04, 0x52, 0xf0, 0x0a, 0x8d, 0xc1,
	0xdf, 0xab, 0x6a, 0xe2, 0xcd, 0xbc, 0x34, 0xc9, 0x0d, 0xc4, 0xef, 0x20, 0x52, 0xba, 0x5e, 0xf3,
	0x1b, 0x81, 0x10, 0x04, 0xbc, 0xd8, 0x53, 0xc7, 0x5a, 0x8c, 0x4e, 0x21, 0x61, 0x6a, 0xd3, 0x94,
	0xb4, 0xa9, 0x0f, 0x93, 0xc1, 0xcc, 0x4b, 0xe3, 0xfc, 0x7e, 0x80, 0xcf, 0x61, 0x54, 0x48, 0xb6,
	0xa2, 0x6a, 0xd7, 0x30, 0xa9, 0x99, 0xe0, 0xff, 0xdc, 0x81, 0x20, 0x60, 0xfc, 0x46, 0xd8, 0xbf,
	0x27, 0xb9, 0xc5, 0x98, 0x40, 0x54, 0x48, 0x76, 0xc1, 0x94, 0x46, 0xaf, 0x20, 0xa8, 0x99, 0xd2,
	0x13, 0x6f, 0xe6, 0xa7, 0x47, 0xf3, 0x63, 0xf2, 0x78, 0x63, 0x6e, 0x49, 0x23, 0xb3, 0x90, 0xcc,
	0xca, 0xb4, 0x92, 0xbe, 0x36, 0x4c, 0x33, 0xde, 0xbd, 0xc4, 0x4a, 0x72, 0x03, 0x73, 0x4c, 0x1f,
	0x24, 0xb5, 0xc7, 0x86, 0xb9, 0xc5, 0x98, 0x43, 0x58, 0x48, 0x76, 0x59, 0x48, 0x84, 0xc1, 0x2f,
	0x24, 0x73, 0xa7, 0xc6, 0xa4, 0x9b, 0x92, 0x85, 0x64, 0x19, 0xd7, 0xcd, 0x21, 0x37, 0xe4, 0xf4,
	0x23, 0xc4, 0xfd, 0xc0, 0xf8, 0x75, 0x4b, 0x0f, 0xbd, 0x5f, 0xb7, 0xf4, 0x80, 0xce, 0x60, 0xf8,
	0xb3, 0xa8, 0xdb, 0xee, 0xc0, 0xd1, 0x3c, 0x26, 0x4e, 0x56, 0xde, 0x8d, 0xdf, 0x0e, 0xce, 0x3d,
	0xfc, 0x0d, 0x46, 0x92, 0xf2, 0x92, 0xf1, 0xea, 0x92, 0x2a, 0x55, 0x54, 0x14, 0x8d, 0x60, 0xc0,
	0x4a, 0xb7, 0x66, 0xc0, 0x4a, 0xb3, 0x57, 0x48, 0xed, 0x1c, 0x31, 0xb0, 0x4f, 0xc6, 0x9f, 0x79,
	0xe9, 0x53, 0x9b, 0x8c, 0x7d, 0x09, 0xdb, 0xd3, 0x49, 0x30, 0xf3, 0x52, 0x3f, 0xb7, 0x18, 0x7f,
	0x86, 0xb0, 0xa1, 0xaa, 0xad, 0x35, 0x7a, 0x09, 0xc1, 0x4e, 0x94, 0x9d, 0xd1, 0xa3, 0xf9, 0x11,
	0x11, 0xd6, 0xad, 0xa5, 0x28, 0x69, 0x6e, 0x09, 0x74, 0x02, 0x43, 0x4e, 0x7f, 0xad, 0x4b, 0x77,
	0xa4, 0xfb, 0xf1, 0xf7, 0x19, 0xfc, 0x1e, 0xa2, 0x86, 0xfe, 0x68, 0xa9, 0xd2, 0x77, 0x2a, 0xfd,
	0xff, 0x57, 0x89, 0x4f, 0x21, 0x6e, 0xa8, 0x92, 0x82, 0x2b, 0xfa, 0xb0, 0x5d, 0x1d, 0xfb, 0xfa,
	0x03, 0xc0, 0xbd, 0x30, 0xf4, 0x0c, 0x92, 0xab, 0xcd, 0xf7, 0xc5, 0x72, 0xbb, 0xde, 0x5c, 0x8d,
	0x9f, 0x20, 0x80, 0x70, 0x99, 0x67, 0x8b, 0x6d, 0x36, 0xf6, 0x0c, 0xfe, 0xf2, 0x69, 0x65, 0xf0,
	0xc0, 0xe0, 0x55, 0x76, 0x91, 0x6d, 0xb3, 0xb1, 0x3f, 0xff, 0xed, 0x81, 0xaf, 0x74, 0x8d, 0x4e,
	0x20, 0x90, 0x26, 0xde, 0x21, 0x31, 0x7d, 0x9e, 0x0e, 0x89, 0xad, 0xf3, 0x14, 0xa2, 0x8a, 0x6a,
	0xdb, 0x0a, 0x47, 0xc4, 0xa4, 0x6f, 0xf3, 0x0b, 0x48, 0x2a, 0xaa, 0x17, 0x5d, 0xee, 0x8e, 0x8d,
	0x5c, 0xe2, 0x28, 0x85, 0xe7, 0x1d, 0xf9, 0xa0, 0x69, 0xb6, 0x8c, 0x77, 0x6b, 0xfa, 0x76, 0x62,
	0x08, 0x55, 0x7b, 0xbd, 0x67, 0x1a, 0x1d, 0x93, 0xc7, 0xa1, 0x4e, 0x23, 0xe2, 0xb2, 0x38, 0x83,
	0x68, 0x27, 0xf6, 0xb2, 0xd5, 0x14, 0xc5, 0xc4, 0x99, 0x39, 0x4d, 0x48, 0xef, 0xcb, 0x75, 0x68,
	0x3f, 0xc6, 0x37, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x45, 0xfa, 0x45, 0x99, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StlClient is the client API for Stl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StlClient interface {
	Ping(ctx context.Context, in *Null, opts ...grpc.CallOption) (*Pong, error)
	GetInfo(ctx context.Context, in *Null, opts ...grpc.CallOption) (*StlInfo, error)
	GetApiMap(ctx context.Context, in *Null, opts ...grpc.CallOption) (*ApiMap, error)
	GetApiDescriptionList(ctx context.Context, in *Null, opts ...grpc.CallOption) (*ApiList, error)
	Submit(ctx context.Context, in *PendingMessage, opts ...grpc.CallOption) (*Result, error)
	Compute(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type stlClient struct {
	cc *grpc.ClientConn
}

func NewStlClient(cc *grpc.ClientConn) StlClient {
	return &stlClient{cc}
}

func (c *stlClient) Ping(ctx context.Context, in *Null, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/stl/ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stlClient) GetInfo(ctx context.Context, in *Null, opts ...grpc.CallOption) (*StlInfo, error) {
	out := new(StlInfo)
	err := c.cc.Invoke(ctx, "/stl/getInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stlClient) GetApiMap(ctx context.Context, in *Null, opts ...grpc.CallOption) (*ApiMap, error) {
	out := new(ApiMap)
	err := c.cc.Invoke(ctx, "/stl/getApiMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stlClient) GetApiDescriptionList(ctx context.Context, in *Null, opts ...grpc.CallOption) (*ApiList, error) {
	out := new(ApiList)
	err := c.cc.Invoke(ctx, "/stl/getApiDescriptionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stlClient) Submit(ctx context.Context, in *PendingMessage, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/stl/submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stlClient) Compute(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/stl/compute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StlServer is the server API for Stl service.
type StlServer interface {
	Ping(context.Context, *Null) (*Pong, error)
	GetInfo(context.Context, *Null) (*StlInfo, error)
	GetApiMap(context.Context, *Null) (*ApiMap, error)
	GetApiDescriptionList(context.Context, *Null) (*ApiList, error)
	Submit(context.Context, *PendingMessage) (*Result, error)
	Compute(context.Context, *Request) (*Response, error)
}

// UnimplementedStlServer can be embedded to have forward compatible implementations.
type UnimplementedStlServer struct {
}

func (*UnimplementedStlServer) Ping(ctx context.Context, req *Null) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedStlServer) GetInfo(ctx context.Context, req *Null) (*StlInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (*UnimplementedStlServer) GetApiMap(ctx context.Context, req *Null) (*ApiMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiMap not implemented")
}
func (*UnimplementedStlServer) GetApiDescriptionList(ctx context.Context, req *Null) (*ApiList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiDescriptionList not implemented")
}
func (*UnimplementedStlServer) Submit(ctx context.Context, req *PendingMessage) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (*UnimplementedStlServer) Compute(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compute not implemented")
}

func RegisterStlServer(s *grpc.Server, srv StlServer) {
	s.RegisterService(&_Stl_serviceDesc, srv)
}

func _Stl_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StlServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stl/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StlServer).Ping(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stl_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StlServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stl/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StlServer).GetInfo(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stl_GetApiMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StlServer).GetApiMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stl/GetApiMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StlServer).GetApiMap(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stl_GetApiDescriptionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StlServer).GetApiDescriptionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stl/GetApiDescriptionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StlServer).GetApiDescriptionList(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stl_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PendingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StlServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stl/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StlServer).Submit(ctx, req.(*PendingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stl_Compute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StlServer).Compute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stl/Compute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StlServer).Compute(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stl",
	HandlerType: (*StlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ping",
			Handler:    _Stl_Ping_Handler,
		},
		{
			MethodName: "getInfo",
			Handler:    _Stl_GetInfo_Handler,
		},
		{
			MethodName: "getApiMap",
			Handler:    _Stl_GetApiMap_Handler,
		},
		{
			MethodName: "getApiDescriptionList",
			Handler:    _Stl_GetApiDescriptionList_Handler,
		},
		{
			MethodName: "submit",
			Handler:    _Stl_Submit_Handler,
		},
		{
			MethodName: "compute",
			Handler:    _Stl_Compute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
