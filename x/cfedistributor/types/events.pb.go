// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cfedistributor/events.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type EventDistributionFinished struct {
	SubDistributionFinished []*SubDistributionFinished `protobuf:"bytes,1,rep,name=subDistributionFinished,proto3" json:"subDistributionFinished,omitempty"`
}

func (m *EventDistributionFinished) Reset()         { *m = EventDistributionFinished{} }
func (m *EventDistributionFinished) String() string { return proto.CompactTextString(m) }
func (*EventDistributionFinished) ProtoMessage()    {}
func (*EventDistributionFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ccf2e9cad100e99, []int{0}
}
func (m *EventDistributionFinished) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventDistributionFinished) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventDistributionFinished.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDistributionFinished) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDistributionFinished.Merge(m, src)
}
func (m *EventDistributionFinished) XXX_Size() int {
	return m.Size()
}
func (m *EventDistributionFinished) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDistributionFinished.DiscardUnknown(m)
}

var xxx_messageInfo_EventDistributionFinished proto.InternalMessageInfo

func (m *EventDistributionFinished) GetSubDistributionFinished() []*SubDistributionFinished {
	if m != nil {
		return m.SubDistributionFinished
	}
	return nil
}

type SubDistributionFinished struct {
	Sources              []string                `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	ProcessedDestination []*ProcessedDestination `protobuf:"bytes,2,rep,name=processedDestination,proto3" json:"processedDestination,omitempty"`
}

func (m *SubDistributionFinished) Reset()         { *m = SubDistributionFinished{} }
func (m *SubDistributionFinished) String() string { return proto.CompactTextString(m) }
func (*SubDistributionFinished) ProtoMessage()    {}
func (*SubDistributionFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ccf2e9cad100e99, []int{1}
}
func (m *SubDistributionFinished) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubDistributionFinished) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubDistributionFinished.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubDistributionFinished) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubDistributionFinished.Merge(m, src)
}
func (m *SubDistributionFinished) XXX_Size() int {
	return m.Size()
}
func (m *SubDistributionFinished) XXX_DiscardUnknown() {
	xxx_messageInfo_SubDistributionFinished.DiscardUnknown(m)
}

var xxx_messageInfo_SubDistributionFinished proto.InternalMessageInfo

func (m *SubDistributionFinished) GetSources() []string {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *SubDistributionFinished) GetProcessedDestination() []*ProcessedDestination {
	if m != nil {
		return m.ProcessedDestination
	}
	return nil
}

type ProcessedDestination struct {
	Name       string                                      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CoinSended github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=coin_sended,json=coinSended,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"coin_sended" yaml:"coin_state"`
}

func (m *ProcessedDestination) Reset()         { *m = ProcessedDestination{} }
func (m *ProcessedDestination) String() string { return proto.CompactTextString(m) }
func (*ProcessedDestination) ProtoMessage()    {}
func (*ProcessedDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ccf2e9cad100e99, []int{2}
}
func (m *ProcessedDestination) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessedDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessedDestination.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProcessedDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessedDestination.Merge(m, src)
}
func (m *ProcessedDestination) XXX_Size() int {
	return m.Size()
}
func (m *ProcessedDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessedDestination.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessedDestination proto.InternalMessageInfo

func (m *ProcessedDestination) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProcessedDestination) GetCoinSended() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.CoinSended
	}
	return nil
}

func init() {
	proto.RegisterType((*EventDistributionFinished)(nil), "chain4energy.c4echain.cfedistributor.EventDistributionFinished")
	proto.RegisterType((*SubDistributionFinished)(nil), "chain4energy.c4echain.cfedistributor.SubDistributionFinished")
	proto.RegisterType((*ProcessedDestination)(nil), "chain4energy.c4echain.cfedistributor.ProcessedDestination")
}

func init() { proto.RegisterFile("cfedistributor/events.proto", fileDescriptor_3ccf2e9cad100e99) }

var fileDescriptor_3ccf2e9cad100e99 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0xae, 0xd2, 0x40,
	0x14, 0x86, 0x3b, 0x6a, 0x34, 0x0c, 0x2b, 0x1b, 0x12, 0x2a, 0x9a, 0x42, 0x1a, 0x17, 0x24, 0x86,
	0x99, 0xa0, 0x2c, 0x0c, 0x89, 0x1b, 0x44, 0xe3, 0xd2, 0x94, 0x9d, 0x1b, 0xd3, 0x4e, 0x8f, 0x65,
	0xa2, 0x9d, 0x21, 0x3d, 0x53, 0x94, 0x27, 0x70, 0xeb, 0xc2, 0x37, 0x70, 0xe7, 0xda, 0x87, 0x60,
	0xc9, 0xd2, 0x15, 0x1a, 0x78, 0x83, 0xfb, 0x04, 0x37, 0x9d, 0x42, 0xc2, 0xbd, 0x29, 0x09, 0xab,
	0xce, 0xf4, 0xcc, 0xff, 0x9d, 0x3f, 0xe7, 0x3f, 0xf4, 0xb1, 0xf8, 0x04, 0x89, 0x44, 0x93, 0xcb,
	0xb8, 0x30, 0x3a, 0xe7, 0xb0, 0x04, 0x65, 0x90, 0x2d, 0x72, 0x6d, 0xb4, 0xfb, 0x54, 0xcc, 0x23,
	0xa9, 0x46, 0xa0, 0x20, 0x4f, 0x57, 0x4c, 0x8c, 0xc0, 0xde, 0xd9, 0x4d, 0x49, 0xa7, 0x95, 0xea,
	0x54, 0x5b, 0x01, 0x2f, 0x4f, 0x95, 0xb6, 0xe3, 0x0b, 0x8d, 0x99, 0x46, 0x1e, 0x47, 0x08, 0x7c,
	0x39, 0x8c, 0xc1, 0x44, 0x43, 0x2e, 0xb4, 0x54, 0x55, 0x3d, 0xf8, 0x49, 0xe8, 0xa3, 0x37, 0x65,
	0xb3, 0xe9, 0x11, 0x25, 0xb5, 0x7a, 0x2b, 0x95, 0xc4, 0x39, 0x24, 0xee, 0x57, 0xda, 0xc6, 0x22,
	0xae, 0x2b, 0x79, 0xa4, 0x77, 0xb7, 0xdf, 0x7c, 0xfe, 0x8a, 0x5d, 0xe2, 0x8d, 0xcd, 0xea, 0x21,
	0xe1, 0x39, 0x7a, 0xf0, 0x8b, 0xd0, 0xf6, 0x19, 0x91, 0xeb, 0xd1, 0x07, 0xa8, 0x8b, 0x5c, 0x00,
	0x5a, 0x13, 0x8d, 0xf0, 0x78, 0x75, 0x15, 0x6d, 0x2d, 0x72, 0x2d, 0x00, 0x11, 0x92, 0x29, 0xa0,
	0x91, 0x2a, 0x2a, 0x95, 0xde, 0x1d, 0xeb, 0x75, 0x7c, 0x99, 0xd7, 0xf7, 0x35, 0x84, 0xb0, 0x96,
	0x1b, 0xfc, 0x21, 0xb4, 0x55, 0xf7, 0xdc, 0x75, 0xe9, 0x3d, 0x15, 0x65, 0xe0, 0x91, 0x1e, 0xe9,
	0x37, 0x42, 0x7b, 0x76, 0xbf, 0x13, 0xda, 0x2c, 0x07, 0xff, 0x11, 0x41, 0x25, 0x90, 0x1c, 0x4c,
	0x3d, 0x61, 0x55, 0x40, 0xac, 0x0c, 0x88, 0x1d, 0x02, 0x62, 0x53, 0x10, 0xaf, 0xb5, 0x54, 0x93,
	0x77, 0xeb, 0x6d, 0xd7, 0xb9, 0xda, 0x76, 0x1f, 0xae, 0xa2, 0xec, 0xcb, 0x38, 0xa8, 0xe4, 0x26,
	0x32, 0x10, 0xfc, 0xfe, 0xd7, 0x7d, 0x96, 0x4a, 0x33, 0x2f, 0x62, 0x26, 0x74, 0xc6, 0x0f, 0x29,
	0x57, 0x9f, 0x01, 0x26, 0x9f, 0xb9, 0x59, 0x2d, 0x00, 0x8f, 0x20, 0x0c, 0x69, 0xa9, 0x9d, 0xd9,
	0xce, 0x93, 0x70, 0xbd, 0xf3, 0xc9, 0x66, 0xe7, 0x93, 0xff, 0x3b, 0x9f, 0xfc, 0xd8, 0xfb, 0xce,
	0x66, 0xef, 0x3b, 0x7f, 0xf7, 0xbe, 0xf3, 0xe1, 0xe5, 0x29, 0xf2, 0x64, 0x58, 0x5c, 0x8c, 0x60,
	0x60, 0x7f, 0xf0, 0x6f, 0xfc, 0xd6, 0xaa, 0xda, 0x46, 0xf1, 0x7d, 0xbb, 0x4e, 0x2f, 0xae, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x5a, 0xf6, 0xbf, 0x6c, 0xc9, 0x02, 0x00, 0x00,
}

func (m *EventDistributionFinished) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDistributionFinished) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventDistributionFinished) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubDistributionFinished) > 0 {
		for iNdEx := len(m.SubDistributionFinished) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubDistributionFinished[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SubDistributionFinished) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubDistributionFinished) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubDistributionFinished) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProcessedDestination) > 0 {
		for iNdEx := len(m.ProcessedDestination) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProcessedDestination[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Sources) > 0 {
		for iNdEx := len(m.Sources) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Sources[iNdEx])
			copy(dAtA[i:], m.Sources[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.Sources[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ProcessedDestination) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessedDestination) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProcessedDestination) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CoinSended) > 0 {
		for iNdEx := len(m.CoinSended) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoinSended[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventDistributionFinished) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SubDistributionFinished) > 0 {
		for _, e := range m.SubDistributionFinished {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *SubDistributionFinished) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Sources) > 0 {
		for _, s := range m.Sources {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	if len(m.ProcessedDestination) > 0 {
		for _, e := range m.ProcessedDestination {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *ProcessedDestination) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.CoinSended) > 0 {
		for _, e := range m.CoinSended {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventDistributionFinished) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventDistributionFinished: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventDistributionFinished: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubDistributionFinished", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubDistributionFinished = append(m.SubDistributionFinished, &SubDistributionFinished{})
			if err := m.SubDistributionFinished[len(m.SubDistributionFinished)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *SubDistributionFinished) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: SubDistributionFinished: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubDistributionFinished: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sources", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sources = append(m.Sources, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessedDestination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProcessedDestination = append(m.ProcessedDestination, &ProcessedDestination{})
			if err := m.ProcessedDestination[len(m.ProcessedDestination)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *ProcessedDestination) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: ProcessedDestination: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessedDestination: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinSended", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoinSended = append(m.CoinSended, types.DecCoin{})
			if err := m.CoinSended[len(m.CoinSended)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)