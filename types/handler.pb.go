// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: handler.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Handler is a handler specification.
type Handler struct {
	// Name is the unique identifier for a handler.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type is the handler type, i.e. pipe.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Mutator is the handler event data mutator.
	Mutator string `protobuf:"bytes,3,opt,name=mutator,proto3" json:"mutator,omitempty"`
	// Command is the command to be executed for a pipe handler.
	Command string `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
	// Timeout is the handler timeout in seconds.
	Timeout uint32 `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout"`
	// Socket contains configuration for a TCP or UDP handler.
	Socket *HandlerSocket `protobuf:"bytes,6,opt,name=socket" json:"socket,omitempty"`
	// Handlers is a list of handlers for a handler set.
	Handlers []string `protobuf:"bytes,7,rep,name=handlers" json:"handlers"`
	// Filters is a list of filters name to evaluate before executing this handler
	Filters []string `protobuf:"bytes,8,rep,name=filters" json:"filters"`
	// EnvVars is a list of environment variables to use with command execution
	EnvVars []string `protobuf:"bytes,9,rep,name=env_vars,json=envVars" json:"env_vars"`
	// RuntimeAssets are a list of assets required to execute a handler.
	RuntimeAssets []string `protobuf:"bytes,12,rep,name=runtime_assets,json=runtimeAssets" json:"runtime_assets"`
	// Namespace to which the handler belongs
	Namespace            string   `protobuf:"bytes,13,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Handler) Reset()         { *m = Handler{} }
func (m *Handler) String() string { return proto.CompactTextString(m) }
func (*Handler) ProtoMessage()    {}
func (*Handler) Descriptor() ([]byte, []int) {
	return fileDescriptor_handler_062bdf20c53515aa, []int{0}
}
func (m *Handler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Handler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Handler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Handler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handler.Merge(dst, src)
}
func (m *Handler) XXX_Size() int {
	return m.Size()
}
func (m *Handler) XXX_DiscardUnknown() {
	xxx_messageInfo_Handler.DiscardUnknown(m)
}

var xxx_messageInfo_Handler proto.InternalMessageInfo

func (m *Handler) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Handler) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Handler) GetMutator() string {
	if m != nil {
		return m.Mutator
	}
	return ""
}

func (m *Handler) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Handler) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Handler) GetSocket() *HandlerSocket {
	if m != nil {
		return m.Socket
	}
	return nil
}

func (m *Handler) GetHandlers() []string {
	if m != nil {
		return m.Handlers
	}
	return nil
}

func (m *Handler) GetFilters() []string {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *Handler) GetEnvVars() []string {
	if m != nil {
		return m.EnvVars
	}
	return nil
}

func (m *Handler) GetRuntimeAssets() []string {
	if m != nil {
		return m.RuntimeAssets
	}
	return nil
}

func (m *Handler) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// HandlerSocket contains configuration for a TCP or UDP handler.
type HandlerSocket struct {
	// Host is the socket peer address.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Port is the socket peer port.
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandlerSocket) Reset()         { *m = HandlerSocket{} }
func (m *HandlerSocket) String() string { return proto.CompactTextString(m) }
func (*HandlerSocket) ProtoMessage()    {}
func (*HandlerSocket) Descriptor() ([]byte, []int) {
	return fileDescriptor_handler_062bdf20c53515aa, []int{1}
}
func (m *HandlerSocket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandlerSocket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HandlerSocket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HandlerSocket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandlerSocket.Merge(dst, src)
}
func (m *HandlerSocket) XXX_Size() int {
	return m.Size()
}
func (m *HandlerSocket) XXX_DiscardUnknown() {
	xxx_messageInfo_HandlerSocket.DiscardUnknown(m)
}

var xxx_messageInfo_HandlerSocket proto.InternalMessageInfo

func (m *HandlerSocket) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HandlerSocket) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*Handler)(nil), "sensu.types.Handler")
	proto.RegisterType((*HandlerSocket)(nil), "sensu.types.HandlerSocket")
}
func (this *Handler) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Handler)
	if !ok {
		that2, ok := that.(Handler)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Mutator != that1.Mutator {
		return false
	}
	if this.Command != that1.Command {
		return false
	}
	if this.Timeout != that1.Timeout {
		return false
	}
	if !this.Socket.Equal(that1.Socket) {
		return false
	}
	if len(this.Handlers) != len(that1.Handlers) {
		return false
	}
	for i := range this.Handlers {
		if this.Handlers[i] != that1.Handlers[i] {
			return false
		}
	}
	if len(this.Filters) != len(that1.Filters) {
		return false
	}
	for i := range this.Filters {
		if this.Filters[i] != that1.Filters[i] {
			return false
		}
	}
	if len(this.EnvVars) != len(that1.EnvVars) {
		return false
	}
	for i := range this.EnvVars {
		if this.EnvVars[i] != that1.EnvVars[i] {
			return false
		}
	}
	if len(this.RuntimeAssets) != len(that1.RuntimeAssets) {
		return false
	}
	for i := range this.RuntimeAssets {
		if this.RuntimeAssets[i] != that1.RuntimeAssets[i] {
			return false
		}
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HandlerSocket) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HandlerSocket)
	if !ok {
		that2, ok := that.(HandlerSocket)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Host != that1.Host {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *Handler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Handler) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Mutator) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Mutator)))
		i += copy(dAtA[i:], m.Mutator)
	}
	if len(m.Command) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Command)))
		i += copy(dAtA[i:], m.Command)
	}
	if m.Timeout != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintHandler(dAtA, i, uint64(m.Timeout))
	}
	if m.Socket != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintHandler(dAtA, i, uint64(m.Socket.Size()))
		n1, err := m.Socket.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Handlers) > 0 {
		for _, s := range m.Handlers {
			dAtA[i] = 0x3a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Filters) > 0 {
		for _, s := range m.Filters {
			dAtA[i] = 0x42
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.EnvVars) > 0 {
		for _, s := range m.EnvVars {
			dAtA[i] = 0x4a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.RuntimeAssets) > 0 {
		for _, s := range m.RuntimeAssets {
			dAtA[i] = 0x62
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Namespace) > 0 {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *HandlerSocket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandlerSocket) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Host) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Host)))
		i += copy(dAtA[i:], m.Host)
	}
	if m.Port != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintHandler(dAtA, i, uint64(m.Port))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintHandler(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedHandler(r randyHandler, easy bool) *Handler {
	this := &Handler{}
	this.Name = string(randStringHandler(r))
	this.Type = string(randStringHandler(r))
	this.Mutator = string(randStringHandler(r))
	this.Command = string(randStringHandler(r))
	this.Timeout = uint32(r.Uint32())
	if r.Intn(10) != 0 {
		this.Socket = NewPopulatedHandlerSocket(r, easy)
	}
	v1 := r.Intn(10)
	this.Handlers = make([]string, v1)
	for i := 0; i < v1; i++ {
		this.Handlers[i] = string(randStringHandler(r))
	}
	v2 := r.Intn(10)
	this.Filters = make([]string, v2)
	for i := 0; i < v2; i++ {
		this.Filters[i] = string(randStringHandler(r))
	}
	v3 := r.Intn(10)
	this.EnvVars = make([]string, v3)
	for i := 0; i < v3; i++ {
		this.EnvVars[i] = string(randStringHandler(r))
	}
	v4 := r.Intn(10)
	this.RuntimeAssets = make([]string, v4)
	for i := 0; i < v4; i++ {
		this.RuntimeAssets[i] = string(randStringHandler(r))
	}
	this.Namespace = string(randStringHandler(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedHandler(r, 14)
	}
	return this
}

func NewPopulatedHandlerSocket(r randyHandler, easy bool) *HandlerSocket {
	this := &HandlerSocket{}
	this.Host = string(randStringHandler(r))
	this.Port = uint32(r.Uint32())
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedHandler(r, 3)
	}
	return this
}

type randyHandler interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneHandler(r randyHandler) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringHandler(r randyHandler) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneHandler(r)
	}
	return string(tmps)
}
func randUnrecognizedHandler(r randyHandler, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldHandler(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldHandler(dAtA []byte, r randyHandler, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateHandler(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Handler) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	l = len(m.Mutator)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	l = len(m.Command)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovHandler(uint64(m.Timeout))
	}
	if m.Socket != nil {
		l = m.Socket.Size()
		n += 1 + l + sovHandler(uint64(l))
	}
	if len(m.Handlers) > 0 {
		for _, s := range m.Handlers {
			l = len(s)
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if len(m.Filters) > 0 {
		for _, s := range m.Filters {
			l = len(s)
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if len(m.EnvVars) > 0 {
		for _, s := range m.EnvVars {
			l = len(s)
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if len(m.RuntimeAssets) > 0 {
		for _, s := range m.RuntimeAssets {
			l = len(s)
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HandlerSocket) Size() (n int) {
	var l int
	_ = l
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovHandler(uint64(m.Port))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHandler(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHandler(x uint64) (n int) {
	return sovHandler(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Handler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Handler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Handler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mutator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mutator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Command = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Socket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Socket == nil {
				m.Socket = &HandlerSocket{}
			}
			if err := m.Socket.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handlers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handlers = append(m.Handlers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filters = append(m.Filters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnvVars", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnvVars = append(m.EnvVars, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeAssets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RuntimeAssets = append(m.RuntimeAssets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHandler
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HandlerSocket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HandlerSocket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandlerSocket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHandler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHandler
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHandler(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHandler
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHandler
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHandler(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHandler = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHandler   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("handler.proto", fileDescriptor_handler_062bdf20c53515aa) }

var fileDescriptor_handler_062bdf20c53515aa = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0xcd, 0xce, 0x93, 0x40,
	0x14, 0x75, 0x2c, 0x2d, 0x30, 0x2d, 0x2e, 0x66, 0x35, 0x69, 0x1a, 0x20, 0x35, 0x46, 0x36, 0xd2,
	0x44, 0x37, 0xba, 0x2c, 0x2b, 0xd7, 0x63, 0xe2, 0xc2, 0x4d, 0x33, 0xa5, 0xd3, 0x9f, 0x58, 0x18,
	0xc2, 0x0c, 0x4d, 0x7c, 0x13, 0x7d, 0x03, 0x1f, 0xc1, 0x47, 0xe8, 0xd2, 0x27, 0x20, 0x8a, 0x3b,
	0x9e, 0xc0, 0xa5, 0x99, 0x0b, 0xb4, 0x7e, 0xdf, 0x86, 0x9c, 0x73, 0xee, 0xe1, 0xe6, 0xde, 0x73,
	0x07, 0x7b, 0x47, 0x9e, 0xef, 0xce, 0xa2, 0x8c, 0x8b, 0x52, 0x6a, 0x49, 0xa6, 0x4a, 0xe4, 0xaa,
	0x8a, 0xf5, 0x97, 0x42, 0xa8, 0xf9, 0xab, 0xc3, 0x49, 0x1f, 0xab, 0x6d, 0x9c, 0xca, 0x6c, 0x75,
	0x90, 0x07, 0xb9, 0x02, 0xcf, 0xb6, 0xda, 0x03, 0x03, 0x02, 0xa8, 0xfb, 0x77, 0xf9, 0x6d, 0x84,
	0xed, 0xf7, 0x5d, 0x37, 0x42, 0xb0, 0x95, 0xf3, 0x4c, 0x50, 0x14, 0xa2, 0xc8, 0x65, 0x80, 0x8d,
	0x66, 0xfa, 0xd2, 0xa7, 0x9d, 0x66, 0x30, 0xa1, 0xd8, 0xce, 0x2a, 0xcd, 0xb5, 0x2c, 0xe9, 0x08,
	0xe4, 0x81, 0x9a, 0x4a, 0x2a, 0xb3, 0x8c, 0xe7, 0x3b, 0x6a, 0x75, 0x95, 0x9e, 0x92, 0x17, 0xd8,
	0xd6, 0xa7, 0x4c, 0xc8, 0x4a, 0xd3, 0x71, 0x88, 0x22, 0x2f, 0x99, 0xb6, 0x75, 0x30, 0x48, 0x6c,
	0x00, 0xe4, 0x2d, 0x9e, 0x28, 0x99, 0x7e, 0x16, 0x9a, 0x4e, 0x42, 0x14, 0x4d, 0x5f, 0xcf, 0xe3,
	0xff, 0x76, 0x8b, 0xfb, 0x41, 0x3f, 0x80, 0x23, 0xb1, 0xae, 0x75, 0x80, 0x58, 0xef, 0x27, 0x11,
	0x76, 0xfa, 0x54, 0x14, 0xb5, 0xc3, 0x51, 0xe4, 0x26, 0xb3, 0xb6, 0x0e, 0x6e, 0x1a, 0xbb, 0x21,
	0x33, 0xca, 0xfe, 0x74, 0xd6, 0xc6, 0xe8, 0x80, 0x11, 0x46, 0xe9, 0x25, 0x36, 0x00, 0xf2, 0x12,
	0x3b, 0x22, 0xbf, 0x6c, 0x2e, 0xbc, 0x54, 0xd4, 0xbd, 0x37, 0x1c, 0x34, 0x66, 0x8b, 0xfc, 0xf2,
	0x91, 0x97, 0x8a, 0xbc, 0xc3, 0xcf, 0xca, 0x2a, 0x37, 0x1b, 0x6c, 0xb8, 0x52, 0x42, 0x2b, 0x3a,
	0x03, 0x3b, 0x69, 0xeb, 0xe0, 0x51, 0x85, 0x79, 0x3d, 0x5f, 0x03, 0x25, 0x0b, 0xec, 0x9a, 0x94,
	0x55, 0xc1, 0x53, 0x41, 0x3d, 0x48, 0xec, 0x2e, 0x2c, 0xd7, 0xd8, 0x7b, 0xb0, 0xb1, 0x39, 0xc6,
	0x51, 0x2a, 0x3d, 0x1c, 0xc8, 0x60, 0xb2, 0xc0, 0x56, 0x21, 0x4b, 0x0d, 0x07, 0xf2, 0x12, 0xa7,
	0xad, 0x03, 0xe0, 0x0c, 0xbe, 0xc9, 0xf3, 0xbf, 0xbf, 0x7d, 0xf4, 0xbd, 0xf1, 0xd1, 0x8f, 0xc6,
	0x47, 0xd7, 0xc6, 0x47, 0x3f, 0x1b, 0x1f, 0xfd, 0x6a, 0x7c, 0xf4, 0xf5, 0x8f, 0xff, 0xe4, 0xd3,
	0x18, 0x62, 0xdd, 0x4e, 0xe0, 0x29, 0xbc, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0x42, 0x8c, 0x7d,
	0xf7, 0x57, 0x02, 0x00, 0x00,
}
