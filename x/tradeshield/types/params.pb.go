// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/tradeshield/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	MarketOrderEnabled   bool   `protobuf:"varint,1,opt,name=market_order_enabled,json=marketOrderEnabled,proto3" json:"market_order_enabled,omitempty"`
	StakeEnabled         bool   `protobuf:"varint,2,opt,name=stake_enabled,json=stakeEnabled,proto3" json:"stake_enabled,omitempty"`
	ProcessOrdersEnabled bool   `protobuf:"varint,3,opt,name=process_orders_enabled,json=processOrdersEnabled,proto3" json:"process_orders_enabled,omitempty"`
	SwapEnabled          bool   `protobuf:"varint,4,opt,name=swap_enabled,json=swapEnabled,proto3" json:"swap_enabled,omitempty"`
	PerpetualEnabled     bool   `protobuf:"varint,5,opt,name=perpetual_enabled,json=perpetualEnabled,proto3" json:"perpetual_enabled,omitempty"`
	RewardEnabled        bool   `protobuf:"varint,6,opt,name=reward_enabled,json=rewardEnabled,proto3" json:"reward_enabled,omitempty"`
	LeverageEnabled      bool   `protobuf:"varint,7,opt,name=leverage_enabled,json=leverageEnabled,proto3" json:"leverage_enabled,omitempty"`
	LimitProcessOrder    uint64 `protobuf:"varint,8,opt,name=limit_process_order,json=limitProcessOrder,proto3" json:"limit_process_order,omitempty"`
	// For incentive system v2
	RewardPercentage cosmossdk_io_math.LegacyDec `protobuf:"bytes,9,opt,name=reward_percentage,json=rewardPercentage,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"reward_percentage"`
	MarginError      cosmossdk_io_math.LegacyDec `protobuf:"bytes,10,opt,name=margin_error,json=marginError,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"margin_error"`
	MinimumDeposit   cosmossdk_io_math.Int       `protobuf:"bytes,11,opt,name=minimum_deposit,json=minimumDeposit,proto3,customtype=cosmossdk.io/math.Int" json:"minimum_deposit"`
	Tolerance        cosmossdk_io_math.LegacyDec `protobuf:"bytes,12,opt,name=tolerance,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"tolerance"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_01646d3b2ec43fec, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMarketOrderEnabled() bool {
	if m != nil {
		return m.MarketOrderEnabled
	}
	return false
}

func (m *Params) GetStakeEnabled() bool {
	if m != nil {
		return m.StakeEnabled
	}
	return false
}

func (m *Params) GetProcessOrdersEnabled() bool {
	if m != nil {
		return m.ProcessOrdersEnabled
	}
	return false
}

func (m *Params) GetSwapEnabled() bool {
	if m != nil {
		return m.SwapEnabled
	}
	return false
}

func (m *Params) GetPerpetualEnabled() bool {
	if m != nil {
		return m.PerpetualEnabled
	}
	return false
}

func (m *Params) GetRewardEnabled() bool {
	if m != nil {
		return m.RewardEnabled
	}
	return false
}

func (m *Params) GetLeverageEnabled() bool {
	if m != nil {
		return m.LeverageEnabled
	}
	return false
}

func (m *Params) GetLimitProcessOrder() uint64 {
	if m != nil {
		return m.LimitProcessOrder
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "elys.tradeshield.Params")
}

func init() { proto.RegisterFile("elys/tradeshield/params.proto", fileDescriptor_01646d3b2ec43fec) }

var fileDescriptor_01646d3b2ec43fec = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0x80, 0x1b, 0xd8, 0xca, 0xea, 0x76, 0x5b, 0x1b, 0x0a, 0x0a, 0x43, 0x64, 0x05, 0x84, 0x54,
	0x34, 0x2d, 0x19, 0x82, 0x27, 0x98, 0xba, 0xc3, 0x10, 0xd2, 0xaa, 0x6a, 0x27, 0x0e, 0x44, 0x6e,
	0xf2, 0x2b, 0xb5, 0x1a, 0xc7, 0xd6, 0x6f, 0x97, 0xd2, 0xb7, 0xe0, 0x61, 0x78, 0x04, 0x0e, 0x3b,
	0x4e, 0x9c, 0x10, 0x87, 0x09, 0xb5, 0x2f, 0x82, 0x62, 0x37, 0x59, 0x27, 0x6e, 0xbb, 0xc5, 0xff,
	0xf7, 0xe5, 0x93, 0x2d, 0xcb, 0xe4, 0x05, 0x64, 0x0b, 0x15, 0x6a, 0xa4, 0x09, 0xa8, 0x09, 0x83,
	0x2c, 0x09, 0x25, 0x45, 0xca, 0x55, 0x20, 0x51, 0x68, 0xe1, 0xb6, 0x0b, 0x1c, 0x6c, 0xe0, 0x83,
	0x6e, 0x2a, 0x52, 0x61, 0x60, 0x58, 0x7c, 0x59, 0xef, 0xe0, 0x59, 0x2c, 0x14, 0x17, 0x2a, 0xb2,
	0xc0, 0x2e, 0x2c, 0x7a, 0xf5, 0x73, 0x9b, 0xd4, 0x87, 0xa6, 0xe9, 0x9e, 0x90, 0x2e, 0xa7, 0x38,
	0x05, 0x1d, 0x09, 0x4c, 0x00, 0x23, 0xc8, 0xe9, 0x38, 0x83, 0xc4, 0x73, 0x7a, 0x4e, 0x7f, 0x67,
	0xe4, 0x5a, 0x76, 0x51, 0xa0, 0x33, 0x4b, 0xdc, 0xd7, 0x64, 0x57, 0x69, 0x3a, 0x85, 0x4a, 0x7d,
	0x60, 0xd4, 0x96, 0x19, 0x96, 0xd2, 0x07, 0xf2, 0x54, 0xa2, 0x88, 0x41, 0x29, 0xdb, 0x55, 0x95,
	0xfd, 0xd0, 0xd8, 0xdd, 0x35, 0x35, 0x65, 0x55, 0xfe, 0xf5, 0x92, 0xb4, 0xd4, 0x9c, 0xca, 0xca,
	0xdd, 0x32, 0x6e, 0xb3, 0x98, 0x95, 0xca, 0x11, 0xe9, 0x48, 0x40, 0x09, 0x7a, 0x46, 0xb3, 0xca,
	0xdb, 0x36, 0x5e, 0xbb, 0x02, 0xa5, 0xfc, 0x86, 0xec, 0x21, 0xcc, 0x29, 0x26, 0x95, 0x59, 0x37,
	0xe6, 0xae, 0x9d, 0x96, 0xda, 0x5b, 0xd2, 0xce, 0xe0, 0x2b, 0x20, 0x4d, 0x6f, 0x0f, 0xf5, 0xc8,
	0x88, 0xfb, 0xe5, 0xbc, 0x54, 0x03, 0xf2, 0x38, 0x63, 0x9c, 0xe9, 0xe8, 0xce, 0xe9, 0xbc, 0x9d,
	0x9e, 0xd3, 0xdf, 0x1a, 0x75, 0x0c, 0x1a, 0x6e, 0x9c, 0xcc, 0xfd, 0x42, 0x3a, 0xeb, 0x1d, 0x48,
	0xc0, 0x18, 0x72, 0x4d, 0x53, 0xf0, 0x1a, 0x3d, 0xa7, 0xdf, 0x38, 0x7d, 0x77, 0x75, 0x73, 0x58,
	0xfb, 0x73, 0x73, 0xf8, 0xdc, 0x5e, 0x8d, 0x4a, 0xa6, 0x01, 0x13, 0x21, 0xa7, 0x7a, 0x12, 0x7c,
	0x82, 0x94, 0xc6, 0x8b, 0x01, 0xc4, 0xbf, 0x7e, 0x1c, 0x93, 0xf5, 0xcd, 0x0d, 0x20, 0x1e, 0xb5,
	0x6d, 0x6b, 0x58, 0xa5, 0xdc, 0x4b, 0xd2, 0xe2, 0x14, 0x53, 0x96, 0x47, 0x80, 0x28, 0xd0, 0x23,
	0xf7, 0x4d, 0x37, 0x6d, 0xe6, 0xac, 0xa8, 0xb8, 0x97, 0x64, 0x9f, 0xb3, 0x9c, 0xf1, 0x19, 0x8f,
	0x12, 0x90, 0x42, 0x31, 0xed, 0x35, 0x4d, 0xf8, 0x68, 0x1d, 0x7e, 0xf2, 0x7f, 0xf8, 0x3c, 0xd7,
	0x1b, 0xc9, 0xf3, 0x5c, 0x8f, 0xf6, 0xd6, 0x8d, 0x81, 0x4d, 0xb8, 0x17, 0xa4, 0xa1, 0x45, 0x06,
	0x48, 0xf3, 0x18, 0xbc, 0xd6, 0x7d, 0x37, 0x7a, 0xdb, 0x38, 0xfd, 0x78, 0xb5, 0xf4, 0x9d, 0xeb,
	0xa5, 0xef, 0xfc, 0x5d, 0xfa, 0xce, 0xf7, 0x95, 0x5f, 0xbb, 0x5e, 0xf9, 0xb5, 0xdf, 0x2b, 0xbf,
	0xf6, 0xf9, 0x24, 0x65, 0x7a, 0x32, 0x1b, 0x07, 0xb1, 0xe0, 0x61, 0xf1, 0x5c, 0x8e, 0x73, 0xd0,
	0x73, 0x81, 0x53, 0xb3, 0x08, 0xbf, 0xdd, 0x79, 0x5c, 0x7a, 0x21, 0x41, 0x8d, 0xeb, 0xe6, 0x65,
	0xbc, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x18, 0x92, 0x85, 0x7d, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Tolerance.Size()
		i -= size
		if _, err := m.Tolerance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.MinimumDeposit.Size()
		i -= size
		if _, err := m.MinimumDeposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.MarginError.Size()
		i -= size
		if _, err := m.MarginError.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.RewardPercentage.Size()
		i -= size
		if _, err := m.RewardPercentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.LimitProcessOrder != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.LimitProcessOrder))
		i--
		dAtA[i] = 0x40
	}
	if m.LeverageEnabled {
		i--
		if m.LeverageEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.RewardEnabled {
		i--
		if m.RewardEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.PerpetualEnabled {
		i--
		if m.PerpetualEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.SwapEnabled {
		i--
		if m.SwapEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.ProcessOrdersEnabled {
		i--
		if m.ProcessOrdersEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.StakeEnabled {
		i--
		if m.StakeEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.MarketOrderEnabled {
		i--
		if m.MarketOrderEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MarketOrderEnabled {
		n += 2
	}
	if m.StakeEnabled {
		n += 2
	}
	if m.ProcessOrdersEnabled {
		n += 2
	}
	if m.SwapEnabled {
		n += 2
	}
	if m.PerpetualEnabled {
		n += 2
	}
	if m.RewardEnabled {
		n += 2
	}
	if m.LeverageEnabled {
		n += 2
	}
	if m.LimitProcessOrder != 0 {
		n += 1 + sovParams(uint64(m.LimitProcessOrder))
	}
	l = m.RewardPercentage.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MarginError.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinimumDeposit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Tolerance.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketOrderEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MarketOrderEnabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StakeEnabled = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessOrdersEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ProcessOrdersEnabled = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SwapEnabled = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerpetualEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PerpetualEnabled = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RewardEnabled = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeverageEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LeverageEnabled = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitProcessOrder", wireType)
			}
			m.LimitProcessOrder = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LimitProcessOrder |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardPercentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarginError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MarginError.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinimumDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tolerance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tolerance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
