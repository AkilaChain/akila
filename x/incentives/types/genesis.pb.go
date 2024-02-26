// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akila/incentives/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the module's genesis state.
type GenesisState struct {
	// params are the incentives module parameters
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// incentives is a slice of active incentives
	Incentives []Incentive `protobuf:"bytes,2,rep,name=incentives,proto3" json:"incentives"`
	// gas_meters is a slice of active Gasmeters
	GasMeters []GasMeter `protobuf:"bytes,3,rep,name=gas_meters,json=gasMeters,proto3" json:"gas_meters"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c48f663a6f69e41, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetIncentives() []Incentive {
	if m != nil {
		return m.Incentives
	}
	return nil
}

func (m *GenesisState) GetGasMeters() []GasMeter {
	if m != nil {
		return m.GasMeters
	}
	return nil
}

// Params defines the incentives module params
type Params struct {
	// enable_incentives is the parameter to enable incentives
	EnableIncentives bool `protobuf:"varint,1,opt,name=enable_incentives,json=enableIncentives,proto3" json:"enable_incentives,omitempty"`
	// allocation_limit is the maximum percentage an incentive can allocate per denomination
	AllocationLimit github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=allocation_limit,json=allocationLimit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"allocation_limit"`
	// incentives_epoch_identifier for the epochs module hooks
	IncentivesEpochIdentifier string `protobuf:"bytes,3,opt,name=incentives_epoch_identifier,json=incentivesEpochIdentifier,proto3" json:"incentives_epoch_identifier,omitempty"`
	// reward_scaler is the scaling factor for capping rewards
	RewardScaler github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=reward_scaler,json=rewardScaler,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_scaler"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c48f663a6f69e41, []int{1}
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

func (m *Params) GetEnableIncentives() bool {
	if m != nil {
		return m.EnableIncentives
	}
	return false
}

func (m *Params) GetIncentivesEpochIdentifier() string {
	if m != nil {
		return m.IncentivesEpochIdentifier
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "akila.incentives.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "akila.incentives.v1.Params")
}

func init() { proto.RegisterFile("akila/incentives/v1/genesis.proto", fileDescriptor_7c48f663a6f69e41) }

var fileDescriptor_7c48f663a6f69e41 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x8a, 0xda, 0x40,
	0x18, 0xc7, 0x13, 0x15, 0xa9, 0xa3, 0xa5, 0x36, 0xed, 0x21, 0xd5, 0x36, 0x5a, 0x29, 0x45, 0x28,
	0x9d, 0xa0, 0x05, 0xa1, 0x97, 0x1e, 0x82, 0x45, 0x84, 0x16, 0x96, 0x78, 0xda, 0xbd, 0x84, 0x31,
	0xce, 0xc6, 0xc1, 0x24, 0x13, 0x32, 0xb3, 0xd9, 0xdd, 0xb7, 0xd8, 0x97, 0xd8, 0x77, 0xf1, 0x28,
	0xec, 0x65, 0xd9, 0x83, 0x2c, 0xfa, 0x22, 0xcb, 0x4c, 0xb2, 0x9b, 0x1c, 0x72, 0xda, 0x53, 0x66,
	0x92, 0xdf, 0xff, 0xf7, 0xff, 0x02, 0x1f, 0xf8, 0x8a, 0x36, 0xc4, 0x47, 0x26, 0x09, 0x5d, 0x1c,
	0x72, 0x92, 0x60, 0x66, 0x26, 0x23, 0xd3, 0xc3, 0x21, 0x66, 0x84, 0xc1, 0x28, 0xa6, 0x9c, 0x6a,
	0x1f, 0x24, 0x02, 0x73, 0x04, 0x26, 0xa3, 0xce, 0xb7, 0xb2, 0x5c, 0x01, 0x91, 0xd1, 0xce, 0x47,
	0x8f, 0x7a, 0x54, 0x1e, 0x4d, 0x71, 0x4a, 0xdf, 0x0e, 0xee, 0x54, 0xd0, 0x9a, 0xa5, 0x15, 0x0b,
	0x8e, 0x38, 0xd6, 0x7e, 0x83, 0x7a, 0x84, 0x62, 0x14, 0x30, 0x5d, 0xed, 0xab, 0xc3, 0xe6, 0xb8,
	0x0b, 0x4b, 0x2a, 0xe1, 0x89, 0x44, 0xac, 0xda, 0x76, 0xdf, 0x53, 0xec, 0x2c, 0xa0, 0x4d, 0x01,
	0xc8, 0x29, 0xbd, 0xd2, 0xaf, 0x0e, 0x9b, 0x63, 0xa3, 0x34, 0x3e, 0x7f, 0xbe, 0x65, 0x86, 0x42,
	0x4e, 0xb3, 0x00, 0xf0, 0x10, 0x73, 0x02, 0xcc, 0x71, 0xcc, 0xf4, 0xaa, 0xb4, 0x7c, 0x29, 0xb5,
	0xcc, 0x10, 0xfb, 0x2f, 0xa8, 0x4c, 0xd2, 0xf0, 0xb2, 0x3b, 0x1b, 0xdc, 0x56, 0x40, 0x3d, 0x1d,
	0x51, 0xfb, 0x01, 0xde, 0xe3, 0x10, 0x2d, 0x7d, 0xec, 0x14, 0x66, 0x13, 0xbf, 0xf6, 0xc6, 0x6e,
	0xa7, 0x1f, 0xe6, 0x79, 0xf7, 0x29, 0x68, 0x23, 0xdf, 0xa7, 0x2e, 0xe2, 0x84, 0x86, 0x8e, 0x4f,
	0x02, 0xc2, 0xf5, 0x4a, 0x5f, 0x1d, 0x36, 0x2c, 0x28, 0x2a, 0x1e, 0xf6, 0xbd, 0xef, 0x1e, 0xe1,
	0xeb, 0x8b, 0x25, 0x74, 0x69, 0x60, 0xba, 0x94, 0x05, 0x94, 0x65, 0x8f, 0x9f, 0x6c, 0xb5, 0x31,
	0xf9, 0x75, 0x84, 0x19, 0x9c, 0x62, 0xd7, 0x7e, 0x97, 0x7b, 0xfe, 0x09, 0x8d, 0xf6, 0x07, 0x74,
	0xf3, 0x01, 0x1c, 0x1c, 0x51, 0x77, 0xed, 0x90, 0x95, 0xb8, 0x9f, 0x13, 0x1c, 0xeb, 0x55, 0xd1,
	0x62, 0x7f, 0xca, 0x91, 0xbf, 0x82, 0x98, 0xbf, 0x00, 0xda, 0x02, 0xbc, 0x8d, 0xf1, 0x25, 0x8a,
	0x57, 0x0e, 0x73, 0x91, 0x8f, 0x63, 0xbd, 0xf6, 0xaa, 0xb9, 0x5a, 0xa9, 0x64, 0x21, 0x1d, 0xd6,
	0x64, 0x7b, 0x30, 0xd4, 0xdd, 0xc1, 0x50, 0x1f, 0x0f, 0x86, 0x7a, 0x73, 0x34, 0x94, 0xdd, 0xd1,
	0x50, 0xee, 0x8f, 0x86, 0x72, 0xf6, 0x39, 0xdd, 0xa9, 0x64, 0x34, 0x31, 0xaf, 0x8a, 0x9b, 0x25,
	0x4d, 0xcb, 0xba, 0x5c, 0x9e, 0x5f, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xb8, 0xd5, 0x24,
	0xb2, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GasMeters) > 0 {
		for iNdEx := len(m.GasMeters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasMeters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Incentives) > 0 {
		for iNdEx := len(m.Incentives) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Incentives[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
		size := m.RewardScaler.Size()
		i -= size
		if _, err := m.RewardScaler.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.IncentivesEpochIdentifier) > 0 {
		i -= len(m.IncentivesEpochIdentifier)
		copy(dAtA[i:], m.IncentivesEpochIdentifier)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IncentivesEpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.AllocationLimit.Size()
		i -= size
		if _, err := m.AllocationLimit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.EnableIncentives {
		i--
		if m.EnableIncentives {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Incentives) > 0 {
		for _, e := range m.Incentives {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GasMeters) > 0 {
		for _, e := range m.GasMeters {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableIncentives {
		n += 2
	}
	l = m.AllocationLimit.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.IncentivesEpochIdentifier)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.RewardScaler.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Incentives", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Incentives = append(m.Incentives, Incentive{})
			if err := m.Incentives[len(m.Incentives)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasMeters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasMeters = append(m.GasMeters, GasMeter{})
			if err := m.GasMeters[len(m.GasMeters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
				return fmt.Errorf("proto: wrong wireType = %d for field EnableIncentives", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.EnableIncentives = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationLimit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AllocationLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentivesEpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IncentivesEpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardScaler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardScaler.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)