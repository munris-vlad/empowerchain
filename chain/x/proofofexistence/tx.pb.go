// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: empowerchain/proofofexistence/tx.proto

package proofofexistence

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgCreateProof is the message used for creating a new proof of existence
type MsgCreateProof struct {
	// creator is the account address that created the proof
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// hash is the SHA-256 hash in HEX format
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *MsgCreateProof) Reset()         { *m = MsgCreateProof{} }
func (m *MsgCreateProof) String() string { return proto.CompactTextString(m) }
func (*MsgCreateProof) ProtoMessage()    {}
func (*MsgCreateProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_b45d71116a87d100, []int{0}
}
func (m *MsgCreateProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateProof.Merge(m, src)
}
func (m *MsgCreateProof) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateProof) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateProof.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateProof proto.InternalMessageInfo

func (m *MsgCreateProof) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateProof) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// MsgCreateProofResponse is response from creating a new proof of existence
type MsgCreateProofResponse struct {
}

func (m *MsgCreateProofResponse) Reset()         { *m = MsgCreateProofResponse{} }
func (m *MsgCreateProofResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateProofResponse) ProtoMessage()    {}
func (*MsgCreateProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b45d71116a87d100, []int{1}
}
func (m *MsgCreateProofResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateProofResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateProofResponse.Merge(m, src)
}
func (m *MsgCreateProofResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateProofResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateProof)(nil), "empowerchain.proofofexistence.MsgCreateProof")
	proto.RegisterType((*MsgCreateProofResponse)(nil), "empowerchain.proofofexistence.MsgCreateProofResponse")
}

func init() {
	proto.RegisterFile("empowerchain/proofofexistence/tx.proto", fileDescriptor_b45d71116a87d100)
}

var fileDescriptor_b45d71116a87d100 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0xcd, 0x2d, 0xc8,
	0x2f, 0x4f, 0x2d, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x28, 0xca, 0xcf, 0x4f, 0xcb, 0x4f,
	0x4b, 0xad, 0xc8, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x92, 0x45, 0x56, 0xa7, 0x87, 0xae, 0x4e, 0x4a, 0x3c, 0x39, 0xbf, 0x38, 0x37,
	0xbf, 0x58, 0x3f, 0xb7, 0x38, 0x5d, 0xbf, 0xcc, 0x10, 0x44, 0x41, 0xf4, 0x49, 0x49, 0x42, 0x24,
	0xe2, 0xc1, 0x3c, 0x7d, 0x08, 0x07, 0x22, 0xa5, 0x94, 0xc6, 0xc5, 0xe7, 0x5b, 0x9c, 0xee, 0x5c,
	0x94, 0x9a, 0x58, 0x92, 0x1a, 0x00, 0x32, 0x50, 0xc8, 0x88, 0x8b, 0x3d, 0x19, 0xc4, 0xcd, 0x2f,
	0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x92, 0xb8, 0xb4, 0x45, 0x57, 0x04, 0xaa, 0xc9, 0x31,
	0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x38, 0xb8, 0xa4, 0x28, 0x33, 0x2f, 0x3d, 0x08, 0xa6, 0x50, 0x48,
	0x88, 0x8b, 0x25, 0x23, 0xb1, 0x38, 0x43, 0x82, 0x09, 0xa4, 0x21, 0x08, 0xcc, 0xb6, 0xe2, 0x69,
	0x7a, 0xbe, 0x41, 0x0b, 0xa6, 0x42, 0x49, 0x82, 0x4b, 0x0c, 0xd5, 0x9e, 0xa0, 0xd4, 0xe2, 0x82,
	0xfc, 0xbc, 0xe2, 0x54, 0xa3, 0x2a, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x62, 0x2e, 0x6e, 0x64,
	0x57, 0xe8, 0xea, 0xe1, 0xf5, 0xab, 0x1e, 0xaa, 0x61, 0x52, 0xa6, 0x24, 0x29, 0x87, 0xd9, 0xed,
	0x14, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xe6, 0xe9, 0x99, 0x25,
	0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x28, 0xb1, 0x83, 0xc2, 0xa9, 0xc0, 0x88, 0xac,
	0x24, 0x36, 0x70, 0xb8, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x32, 0x66, 0x89, 0xf6, 0xd4,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateProof creates a new proof of existence
	CreateProof(ctx context.Context, in *MsgCreateProof, opts ...grpc.CallOption) (*MsgCreateProofResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateProof(ctx context.Context, in *MsgCreateProof, opts ...grpc.CallOption) (*MsgCreateProofResponse, error) {
	out := new(MsgCreateProofResponse)
	err := c.cc.Invoke(ctx, "/empowerchain.proofofexistence.Msg/CreateProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateProof creates a new proof of existence
	CreateProof(context.Context, *MsgCreateProof) (*MsgCreateProofResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateProof(ctx context.Context, req *MsgCreateProof) (*MsgCreateProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProof not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateProof)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/empowerchain.proofofexistence.Msg/CreateProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateProof(ctx, req.(*MsgCreateProof))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "empowerchain.proofofexistence.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProof",
			Handler:    _Msg_CreateProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "empowerchain/proofofexistence/tx.proto",
}

func (m *MsgCreateProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateProof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateProof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateProofResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateProofResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateProofResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateProofResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateProofResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateProofResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateProofResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
