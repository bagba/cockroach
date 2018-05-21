// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: roachpb/io-formats.proto

package roachpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type IOFileFormat_FileFormat int32

const (
	IOFileFormat_Unknown IOFileFormat_FileFormat = 0
	IOFileFormat_CSV     IOFileFormat_FileFormat = 1
)

var IOFileFormat_FileFormat_name = map[int32]string{
	0: "Unknown",
	1: "CSV",
}
var IOFileFormat_FileFormat_value = map[string]int32{
	"Unknown": 0,
	"CSV":     1,
}

func (x IOFileFormat_FileFormat) Enum() *IOFileFormat_FileFormat {
	p := new(IOFileFormat_FileFormat)
	*p = x
	return p
}
func (x IOFileFormat_FileFormat) String() string {
	return proto.EnumName(IOFileFormat_FileFormat_name, int32(x))
}
func (x *IOFileFormat_FileFormat) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IOFileFormat_FileFormat_value, data, "IOFileFormat_FileFormat")
	if err != nil {
		return err
	}
	*x = IOFileFormat_FileFormat(value)
	return nil
}
func (IOFileFormat_FileFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorIoFormats, []int{0, 0}
}

type IOFileFormat struct {
	Format IOFileFormat_FileFormat `protobuf:"varint,1,opt,name=format,enum=cockroach.roachpb.IOFileFormat_FileFormat" json:"format"`
	Csv    CSVOptions              `protobuf:"bytes,2,opt,name=csv" json:"csv"`
}

func (m *IOFileFormat) Reset()                    { *m = IOFileFormat{} }
func (m *IOFileFormat) String() string            { return proto.CompactTextString(m) }
func (*IOFileFormat) ProtoMessage()               {}
func (*IOFileFormat) Descriptor() ([]byte, []int) { return fileDescriptorIoFormats, []int{0} }

// CSVOptions describe the format of csv data (delimiter, comment, etc).
type CSVOptions struct {
	// comma is an delimiter used by the CSV file; defaults to a comma.
	Comma int32 `protobuf:"varint,1,opt,name=comma" json:"comma"`
	// comment is an comment rune; zero value means comments not enabled.
	Comment int32 `protobuf:"varint,2,opt,name=comment" json:"comment"`
	// null_encoding, if not nil, is the string which identifies a NULL. Can be the empty string.
	NullEncoding *string `protobuf:"bytes,3,opt,name=null_encoding,json=nullEncoding" json:"null_encoding,omitempty"`
	// skip the first N lines of the input (e.g. to ignore column headers) when reading.
	Skip uint32 `protobuf:"varint,4,opt,name=skip" json:"skip"`
}

func (m *CSVOptions) Reset()                    { *m = CSVOptions{} }
func (m *CSVOptions) String() string            { return proto.CompactTextString(m) }
func (*CSVOptions) ProtoMessage()               {}
func (*CSVOptions) Descriptor() ([]byte, []int) { return fileDescriptorIoFormats, []int{1} }

func init() {
	proto.RegisterType((*IOFileFormat)(nil), "cockroach.roachpb.IOFileFormat")
	proto.RegisterType((*CSVOptions)(nil), "cockroach.roachpb.CSVOptions")
	proto.RegisterEnum("cockroach.roachpb.IOFileFormat_FileFormat", IOFileFormat_FileFormat_name, IOFileFormat_FileFormat_value)
}
func (m *IOFileFormat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IOFileFormat) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Format))
	dAtA[i] = 0x12
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Csv.Size()))
	n1, err := m.Csv.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *CSVOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CSVOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Comma))
	dAtA[i] = 0x10
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Comment))
	if m.NullEncoding != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintIoFormats(dAtA, i, uint64(len(*m.NullEncoding)))
		i += copy(dAtA[i:], *m.NullEncoding)
	}
	dAtA[i] = 0x20
	i++
	i = encodeVarintIoFormats(dAtA, i, uint64(m.Skip))
	return i, nil
}

func encodeVarintIoFormats(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *IOFileFormat) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovIoFormats(uint64(m.Format))
	l = m.Csv.Size()
	n += 1 + l + sovIoFormats(uint64(l))
	return n
}

func (m *CSVOptions) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovIoFormats(uint64(m.Comma))
	n += 1 + sovIoFormats(uint64(m.Comment))
	if m.NullEncoding != nil {
		l = len(*m.NullEncoding)
		n += 1 + l + sovIoFormats(uint64(l))
	}
	n += 1 + sovIoFormats(uint64(m.Skip))
	return n
}

func sovIoFormats(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIoFormats(x uint64) (n int) {
	return sovIoFormats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IOFileFormat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIoFormats
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
			return fmt.Errorf("proto: IOFileFormat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IOFileFormat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			m.Format = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Format |= (IOFileFormat_FileFormat(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Csv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
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
				return ErrInvalidLengthIoFormats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Csv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIoFormats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIoFormats
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
func (m *CSVOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIoFormats
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
			return fmt.Errorf("proto: CSVOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CSVOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comma", wireType)
			}
			m.Comma = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Comma |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comment", wireType)
			}
			m.Comment = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Comment |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NullEncoding", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
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
				return ErrInvalidLengthIoFormats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.NullEncoding = &s
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Skip", wireType)
			}
			m.Skip = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoFormats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Skip |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIoFormats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIoFormats
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
func skipIoFormats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIoFormats
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
					return 0, ErrIntOverflowIoFormats
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
					return 0, ErrIntOverflowIoFormats
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
				return 0, ErrInvalidLengthIoFormats
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIoFormats
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
				next, err := skipIoFormats(dAtA[start:])
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
	ErrInvalidLengthIoFormats = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIoFormats   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("roachpb/io-formats.proto", fileDescriptorIoFormats) }

var fileDescriptorIoFormats = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0xc6, 0xf3, 0x6c, 0x6b, 0xf0, 0xb5, 0x95, 0x3a, 0xb8, 0x18, 0x0a, 0x8e, 0x31, 0xab, 0x28,
	0x98, 0x42, 0xc1, 0x0b, 0xb4, 0x58, 0x74, 0x55, 0x68, 0xb1, 0x0b, 0x37, 0x52, 0xc7, 0x5a, 0x43,
	0x9b, 0x79, 0x21, 0x89, 0x7a, 0x0a, 0xc1, 0xbb, 0x78, 0x89, 0x2c, 0x5d, 0xba, 0x12, 0x8d, 0x17,
	0x91, 0x4c, 0xa6, 0xb4, 0xa0, 0xbb, 0xc7, 0xf7, 0xe7, 0xc7, 0xf7, 0x90, 0xc7, 0x34, 0x95, 0x0f,
	0xd1, 0x6d, 0x27, 0xa0, 0xd3, 0x7b, 0x8a, 0xc3, 0x69, 0x9a, 0xf8, 0x51, 0x4c, 0x29, 0xb1, 0x3d,
	0x49, 0x72, 0xa1, 0x5d, 0xdf, 0x64, 0xda, 0xfb, 0x73, 0x9a, 0x93, 0x76, 0x3b, 0xc5, 0x55, 0x06,
	0xdd, 0x37, 0xc0, 0xc6, 0xe5, 0x70, 0x10, 0x2c, 0x67, 0x03, 0x0d, 0x60, 0x17, 0xb8, 0x5d, 0xa2,
	0x38, 0x38, 0xe0, 0xed, 0x76, 0x4f, 0xfc, 0x3f, 0x28, 0x7f, 0xb3, 0xe0, 0xaf, 0xcf, 0x5e, 0x35,
	0xfb, 0x3c, 0xb4, 0x46, 0xa6, 0xcf, 0xce, 0xb0, 0x22, 0x93, 0x27, 0xbe, 0xe5, 0x80, 0x57, 0xef,
	0x1e, 0xfc, 0x83, 0xe9, 0x8f, 0x27, 0xc3, 0x28, 0x0d, 0x48, 0x25, 0xa6, 0x59, 0xe4, 0x5d, 0x17,
	0x71, 0x63, 0x4e, 0x1d, 0xed, 0x2b, 0xb5, 0x50, 0xf4, 0xac, 0x5a, 0x16, 0xb3, 0xb1, 0xd2, 0x1f,
	0x4f, 0x5a, 0xe0, 0xbe, 0x00, 0xe2, 0xba, 0xcd, 0xda, 0x58, 0x93, 0x14, 0x86, 0x53, 0x3d, 0xb9,
	0x66, 0x60, 0xa5, 0xc4, 0x04, 0xda, 0xc5, 0x31, 0x53, 0xa9, 0x5e, 0xb2, 0x72, 0x57, 0x22, 0x3b,
	0xc6, 0xa6, 0x7a, 0x5c, 0x2e, 0x6f, 0x66, 0x4a, 0xd2, 0x5d, 0xa0, 0xe6, 0xbc, 0xe2, 0x80, 0xb7,
	0xa3, 0x53, 0x30, 0x6a, 0x14, 0xd6, 0xb9, 0x71, 0x18, 0xc7, 0x6a, 0xb2, 0x08, 0x22, 0x5e, 0x75,
	0xc0, 0x6b, 0x1a, 0x8e, 0x56, 0x7a, 0x47, 0xd9, 0xb7, 0xb0, 0xb2, 0x5c, 0xc0, 0x7b, 0x2e, 0xe0,
	0x23, 0x17, 0xf0, 0x95, 0x0b, 0x78, 0xfd, 0x11, 0xd6, 0xb5, 0x6d, 0x9e, 0xfd, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x10, 0x9e, 0xdd, 0x0c, 0xac, 0x01, 0x00, 0x00,
}