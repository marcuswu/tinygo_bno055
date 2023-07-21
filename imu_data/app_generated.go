package imu_data

import (
	karmem "karmem.org/golang"
	"unsafe"
)

var _ unsafe.Pointer

var _Null = make([]byte, 176)
var _NullReader = karmem.NewReader(_Null)

type (
	PacketIdentifier uint64
)

const (
	PacketIdentifierVector       = 6066171078066954460
	PacketIdentifierIMUDataEntry = 15341296650816720624
	PacketIdentifierIMUEntry     = 14804925532808219779
	PacketIdentifierIMUData      = 11720459716770339160
)

type Vector struct {
	X float64
	Y float64
	Z float64
}

func NewVector() Vector {
	return Vector{}
}

func (x *Vector) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierVector
}

func (x *Vector) Reset() {
	x.Read((*VectorViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Vector) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Vector) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(32)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__XOffset := offset + 0
	writer.Write8At(__XOffset, *(*uint64)(unsafe.Pointer(&x.X)))
	__YOffset := offset + 8
	writer.Write8At(__YOffset, *(*uint64)(unsafe.Pointer(&x.Y)))
	__ZOffset := offset + 16
	writer.Write8At(__ZOffset, *(*uint64)(unsafe.Pointer(&x.Z)))

	return offset, nil
}

func (x *Vector) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewVectorViewer(reader, 0), reader)
}

func (x *Vector) Read(viewer *VectorViewer, reader *karmem.Reader) {
	x.X = viewer.X()
	x.Y = viewer.Y()
	x.Z = viewer.Z()
}

type IMUDataEntry struct {
	Acceleration    Vector
	AngularVelocity Vector
	Rotation        Vector
	Magnetometer    Vector
	Gravity         Vector
	EntryTime       int64
}

func NewIMUDataEntry() IMUDataEntry {
	return IMUDataEntry{}
}

func (x *IMUDataEntry) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierIMUDataEntry
}

func (x *IMUDataEntry) Reset() {
	x.Read((*IMUDataEntryViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *IMUDataEntry) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *IMUDataEntry) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(176)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(172))
	__AccelerationOffset := offset + 4
	if _, err := x.Acceleration.Write(writer, __AccelerationOffset); err != nil {
		return offset, err
	}
	__AngularVelocityOffset := offset + 36
	if _, err := x.AngularVelocity.Write(writer, __AngularVelocityOffset); err != nil {
		return offset, err
	}
	__RotationOffset := offset + 68
	if _, err := x.Rotation.Write(writer, __RotationOffset); err != nil {
		return offset, err
	}
	__MagnetometerOffset := offset + 100
	if _, err := x.Magnetometer.Write(writer, __MagnetometerOffset); err != nil {
		return offset, err
	}
	__GravityOffset := offset + 132
	if _, err := x.Gravity.Write(writer, __GravityOffset); err != nil {
		return offset, err
	}
	__EntryTimeOffset := offset + 164
	writer.Write8At(__EntryTimeOffset, *(*uint64)(unsafe.Pointer(&x.EntryTime)))

	return offset, nil
}

func (x *IMUDataEntry) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewIMUDataEntryViewer(reader, 0), reader)
}

func (x *IMUDataEntry) Read(viewer *IMUDataEntryViewer, reader *karmem.Reader) {
	x.Acceleration.Read(viewer.Acceleration(), reader)
	x.AngularVelocity.Read(viewer.AngularVelocity(), reader)
	x.Rotation.Read(viewer.Rotation(), reader)
	x.Magnetometer.Read(viewer.Magnetometer(), reader)
	x.Gravity.Read(viewer.Gravity(), reader)
	x.EntryTime = viewer.EntryTime()
}

type IMUEntry struct {
	Entry IMUDataEntry
}

func NewIMUEntry() IMUEntry {
	return IMUEntry{}
}

func (x *IMUEntry) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierIMUEntry
}

func (x *IMUEntry) Reset() {
	x.Read((*IMUEntryViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *IMUEntry) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *IMUEntry) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(8)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__EntrySize := uint(176)
	__EntryOffset, err := writer.Alloc(__EntrySize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__EntryOffset))
	if _, err := x.Entry.Write(writer, __EntryOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *IMUEntry) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewIMUEntryViewer(reader, 0), reader)
}

func (x *IMUEntry) Read(viewer *IMUEntryViewer, reader *karmem.Reader) {
	x.Entry.Read(viewer.Entry(reader), reader)
}

type IMUData struct {
	Entries []IMUEntry
}

func NewIMUData() IMUData {
	return IMUData{}
}

func (x *IMUData) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierIMUData
}

func (x *IMUData) Reset() {
	x.Read((*IMUDataViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *IMUData) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *IMUData) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(24)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(16))
	__EntriesSize := uint(8 * len(x.Entries))
	__EntriesOffset, err := writer.Alloc(__EntriesSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__EntriesOffset))
	writer.Write4At(offset+4+4, uint32(__EntriesSize))
	writer.Write4At(offset+4+4+4, 8)
	for i := range x.Entries {
		if _, err := x.Entries[i].Write(writer, __EntriesOffset); err != nil {
			return offset, err
		}
		__EntriesOffset += 8
	}

	return offset, nil
}

func (x *IMUData) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewIMUDataViewer(reader, 0), reader)
}

func (x *IMUData) Read(viewer *IMUDataViewer, reader *karmem.Reader) {
	__EntriesSlice := viewer.Entries(reader)
	__EntriesLen := len(__EntriesSlice)
	if __EntriesLen > cap(x.Entries) {
		x.Entries = append(x.Entries, make([]IMUEntry, __EntriesLen-len(x.Entries))...)
	}
	if __EntriesLen > len(x.Entries) {
		x.Entries = x.Entries[:__EntriesLen]
	}
	for i := 0; i < __EntriesLen; i++ {
		x.Entries[i].Read(&__EntriesSlice[i], reader)
	}
	x.Entries = x.Entries[:__EntriesLen]
}

type VectorViewer struct {
	_data [32]byte
}

func NewVectorViewer(reader *karmem.Reader, offset uint32) (v *VectorViewer) {
	if !reader.IsValidOffset(offset, 32) {
		return (*VectorViewer)(unsafe.Pointer(&_Null))
	}
	v = (*VectorViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *VectorViewer) size() uint32 {
	return 32
}
func (x *VectorViewer) X() (v float64) {
	return *(*float64)(unsafe.Add(unsafe.Pointer(&x._data), 0))
}
func (x *VectorViewer) Y() (v float64) {
	return *(*float64)(unsafe.Add(unsafe.Pointer(&x._data), 8))
}
func (x *VectorViewer) Z() (v float64) {
	return *(*float64)(unsafe.Add(unsafe.Pointer(&x._data), 16))
}

type IMUDataEntryViewer struct {
	_data [176]byte
}

func NewIMUDataEntryViewer(reader *karmem.Reader, offset uint32) (v *IMUDataEntryViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return (*IMUDataEntryViewer)(unsafe.Pointer(&_Null))
	}
	v = (*IMUDataEntryViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*IMUDataEntryViewer)(unsafe.Pointer(&_Null))
	}
	return v
}

func (x *IMUDataEntryViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *IMUDataEntryViewer) Acceleration() (v *VectorViewer) {
	if 4+32 > x.size() {
		return (*VectorViewer)(unsafe.Pointer(&_Null))
	}
	return (*VectorViewer)(unsafe.Add(unsafe.Pointer(&x._data), 4))
}
func (x *IMUDataEntryViewer) AngularVelocity() (v *VectorViewer) {
	if 36+32 > x.size() {
		return (*VectorViewer)(unsafe.Pointer(&_Null))
	}
	return (*VectorViewer)(unsafe.Add(unsafe.Pointer(&x._data), 36))
}
func (x *IMUDataEntryViewer) Rotation() (v *VectorViewer) {
	if 68+32 > x.size() {
		return (*VectorViewer)(unsafe.Pointer(&_Null))
	}
	return (*VectorViewer)(unsafe.Add(unsafe.Pointer(&x._data), 68))
}
func (x *IMUDataEntryViewer) Magnetometer() (v *VectorViewer) {
	if 100+32 > x.size() {
		return (*VectorViewer)(unsafe.Pointer(&_Null))
	}
	return (*VectorViewer)(unsafe.Add(unsafe.Pointer(&x._data), 100))
}
func (x *IMUDataEntryViewer) Gravity() (v *VectorViewer) {
	if 132+32 > x.size() {
		return (*VectorViewer)(unsafe.Pointer(&_Null))
	}
	return (*VectorViewer)(unsafe.Add(unsafe.Pointer(&x._data), 132))
}
func (x *IMUDataEntryViewer) EntryTime() (v int64) {
	if 164+8 > x.size() {
		return v
	}
	return *(*int64)(unsafe.Add(unsafe.Pointer(&x._data), 164))
}

type IMUEntryViewer struct {
	_data [8]byte
}

func NewIMUEntryViewer(reader *karmem.Reader, offset uint32) (v *IMUEntryViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return (*IMUEntryViewer)(unsafe.Pointer(&_Null))
	}
	v = (*IMUEntryViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *IMUEntryViewer) size() uint32 {
	return 8
}
func (x *IMUEntryViewer) Entry(reader *karmem.Reader) (v *IMUDataEntryViewer) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0))
	return NewIMUDataEntryViewer(reader, offset)
}

type IMUDataViewer struct {
	_data [24]byte
}

func NewIMUDataViewer(reader *karmem.Reader, offset uint32) (v *IMUDataViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return (*IMUDataViewer)(unsafe.Pointer(&_Null))
	}
	v = (*IMUDataViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*IMUDataViewer)(unsafe.Pointer(&_Null))
	}
	return v
}

func (x *IMUDataViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *IMUDataViewer) Entries(reader *karmem.Reader) (v []IMUEntryViewer) {
	if 4+12 > x.size() {
		return []IMUEntryViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return []IMUEntryViewer{}
	}
	length := uintptr(size / 8)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]IMUEntryViewer)(unsafe.Pointer(&slice))
}
