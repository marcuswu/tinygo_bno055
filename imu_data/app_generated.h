
#include "stdint.h"
#include "stdlib.h"
#include "string.h"
#include "stdbool.h"
#include "karmem.h"

uint8_t _Null[176];



typedef struct {
    uint8_t _data[32];
} VectorViewer;

uint32_t VectorViewerSize(VectorViewer * x) {
    return 32;
}

VectorViewer * NewVectorViewer(KarmemReader * reader, uint32_t offset) {
    if (KarmemReaderIsValidOffset(reader, offset, 32) == false) {
        return (VectorViewer *) &_Null;
    }
    VectorViewer * v = (VectorViewer *) &reader->pointer[offset];
    return v;
}

double VectorViewer_X(VectorViewer * x) {
    double r;
    memcpy(&r, &x->_data[0], 8);
    return r;
}

double VectorViewer_Y(VectorViewer * x) {
    double r;
    memcpy(&r, &x->_data[8], 8);
    return r;
}

double VectorViewer_Z(VectorViewer * x) {
    double r;
    memcpy(&r, &x->_data[16], 8);
    return r;
}

typedef struct {
    uint8_t _data[176];
} IMUDataEntryViewer;

uint32_t IMUDataEntryViewerSize(IMUDataEntryViewer * x) {
    uint32_t r;
    memcpy(&r, x, 4);
    return r;
}

IMUDataEntryViewer * NewIMUDataEntryViewer(KarmemReader * reader, uint32_t offset) {
    if (KarmemReaderIsValidOffset(reader, offset, 8) == false) {
        return (IMUDataEntryViewer *) &_Null;
    }
    IMUDataEntryViewer * v = (IMUDataEntryViewer *) &reader->pointer[offset];
    if (KarmemReaderIsValidOffset(reader, offset, IMUDataEntryViewerSize(v)) == false) {
        return (IMUDataEntryViewer *) &_Null;
    }
    return v;
}

VectorViewer * IMUDataEntryViewer_Acceleration(IMUDataEntryViewer * x) {
    if ((4 + 32) > IMUDataEntryViewerSize(x)) {
        return (VectorViewer *) &_Null;
    }
        return (VectorViewer *) &x->_data[4];
}

VectorViewer * IMUDataEntryViewer_AngularVelocity(IMUDataEntryViewer * x) {
    if ((36 + 32) > IMUDataEntryViewerSize(x)) {
        return (VectorViewer *) &_Null;
    }
        return (VectorViewer *) &x->_data[36];
}

VectorViewer * IMUDataEntryViewer_Rotation(IMUDataEntryViewer * x) {
    if ((68 + 32) > IMUDataEntryViewerSize(x)) {
        return (VectorViewer *) &_Null;
    }
        return (VectorViewer *) &x->_data[68];
}

VectorViewer * IMUDataEntryViewer_Magnetometer(IMUDataEntryViewer * x) {
    if ((100 + 32) > IMUDataEntryViewerSize(x)) {
        return (VectorViewer *) &_Null;
    }
        return (VectorViewer *) &x->_data[100];
}

VectorViewer * IMUDataEntryViewer_Gravity(IMUDataEntryViewer * x) {
    if ((132 + 32) > IMUDataEntryViewerSize(x)) {
        return (VectorViewer *) &_Null;
    }
        return (VectorViewer *) &x->_data[132];
}

int64_t IMUDataEntryViewer_EntryTime(IMUDataEntryViewer * x) {
    if ((164 + 8) > IMUDataEntryViewerSize(x)) {
        return 0;
    }
    int64_t r;
    memcpy(&r, &x->_data[164], 8);
    return r;
}

typedef struct {
    uint8_t _data[8];
} IMUEntryViewer;

uint32_t IMUEntryViewerSize(IMUEntryViewer * x) {
    return 8;
}

IMUEntryViewer * NewIMUEntryViewer(KarmemReader * reader, uint32_t offset) {
    if (KarmemReaderIsValidOffset(reader, offset, 8) == false) {
        return (IMUEntryViewer *) &_Null;
    }
    IMUEntryViewer * v = (IMUEntryViewer *) &reader->pointer[offset];
    return v;
}

IMUDataEntryViewer * IMUEntryViewer_Entry(IMUEntryViewer * x, KarmemReader * reader) {
    uint32_t offset;
    memcpy(&offset, &x->_data[0], 4);
    return NewIMUDataEntryViewer(reader, offset);
}

typedef struct {
    uint8_t _data[24];
} IMUDataViewer;

uint32_t IMUDataViewerSize(IMUDataViewer * x) {
    uint32_t r;
    memcpy(&r, x, 4);
    return r;
}

IMUDataViewer * NewIMUDataViewer(KarmemReader * reader, uint32_t offset) {
    if (KarmemReaderIsValidOffset(reader, offset, 8) == false) {
        return (IMUDataViewer *) &_Null;
    }
    IMUDataViewer * v = (IMUDataViewer *) &reader->pointer[offset];
    if (KarmemReaderIsValidOffset(reader, offset, IMUDataViewerSize(v)) == false) {
        return (IMUDataViewer *) &_Null;
    }
    return v;
}
uint32_t IMUDataViewer_EntriesLength(IMUDataViewer * x, KarmemReader * reader) {
    if ((4 + 12) > IMUDataViewerSize(x)) {
        return 0;
    }
    uint32_t offset;
    memcpy(&offset, &x->_data[4], 4);
    uint32_t size;
    memcpy(&size, &x->_data[4 + 4], 4);
    if (KarmemReaderIsValidOffset(reader, offset, size) == false) {
        return 0;
    }
    uint32_t length = size / 8;
    return length;
}

IMUEntryViewer * IMUDataViewer_Entries(IMUDataViewer * x, KarmemReader * reader) {
    if ((4 + 12) > IMUDataViewerSize(x)) {
        return (IMUEntryViewer *) &_Null;
    }
    uint32_t offset;
    memcpy(&offset, &x->_data[4], 4);
    uint32_t size;
    memcpy(&size, &x->_data[4 + 4], 4);
    if (KarmemReaderIsValidOffset(reader, offset, size) == false) {
        return (IMUEntryViewer *) &_Null;
    }
    uint32_t length = size / 8;
    return (IMUEntryViewer *) &reader->pointer[offset];
}
typedef uint64_t EnumPacketIdentifier;

const EnumPacketIdentifier EnumPacketIdentifierVector = 6066171078066954460UL;
const EnumPacketIdentifier EnumPacketIdentifierIMUDataEntry = 15341296650816720624UL;
const EnumPacketIdentifier EnumPacketIdentifierIMUEntry = 14804925532808219779UL;
const EnumPacketIdentifier EnumPacketIdentifierIMUData = 11720459716770339160UL;

    
typedef struct {
    double X;
    double Y;
    double Z;
} Vector;

EnumPacketIdentifier VectorPacketIdentifier(Vector * x) {
    return EnumPacketIdentifierVector;
}

uint32_t VectorWrite(Vector * x, KarmemWriter * writer, uint32_t start) {
    uint32_t offset = start;
    uint32_t size = 32;
    if (offset == 0) {
        offset = KarmemWriterAlloc(writer, size);
        if (offset == 0xFFFFFFFF) {
            return 0;
        }
    }
    uint32_t __XOffset = offset + 0;
    KarmemWriterWriteAt(writer, __XOffset, (void *) &x->X, 8);
    uint32_t __YOffset = offset + 8;
    KarmemWriterWriteAt(writer, __YOffset, (void *) &x->Y, 8);
    uint32_t __ZOffset = offset + 16;
    KarmemWriterWriteAt(writer, __ZOffset, (void *) &x->Z, 8);

    return offset;
}

uint32_t VectorWriteAsRoot(Vector * x, KarmemWriter * writer) {
    return VectorWrite(x, writer, 0);
}

void VectorRead(Vector * x, VectorViewer * viewer, KarmemReader * reader) {
    x->X = VectorViewer_X(viewer);
    x->Y = VectorViewer_Y(viewer);
    x->Z = VectorViewer_Z(viewer);
}

void VectorReadAsRoot(Vector * x, KarmemReader * reader) {
    return VectorRead(x, NewVectorViewer(reader, 0), reader);
}

void VectorReset(Vector * x) {
    KarmemReader reader = KarmemNewReader(&_Null[0], 176);
    VectorRead(x, (VectorViewer *) &_Null, &reader);
}

Vector NewVector() {
    Vector r;
    memset(&r, 0, sizeof(Vector));
    return r;
}
typedef struct {
    Vector Acceleration;
    Vector AngularVelocity;
    Vector Rotation;
    Vector Magnetometer;
    Vector Gravity;
    int64_t EntryTime;
} IMUDataEntry;

EnumPacketIdentifier IMUDataEntryPacketIdentifier(IMUDataEntry * x) {
    return EnumPacketIdentifierIMUDataEntry;
}

uint32_t IMUDataEntryWrite(IMUDataEntry * x, KarmemWriter * writer, uint32_t start) {
    uint32_t offset = start;
    uint32_t size = 176;
    if (offset == 0) {
        offset = KarmemWriterAlloc(writer, size);
        if (offset == 0xFFFFFFFF) {
            return 0;
        }
    }
    uint32_t sizeData = 172;
    KarmemWriterWriteAt(writer, offset, (void *)&sizeData, 4);
    uint32_t __AccelerationOffset = offset + 4;
    if (VectorWrite(&x->Acceleration, writer, __AccelerationOffset) == 0) {
        return 0;
    }
    uint32_t __AngularVelocityOffset = offset + 36;
    if (VectorWrite(&x->AngularVelocity, writer, __AngularVelocityOffset) == 0) {
        return 0;
    }
    uint32_t __RotationOffset = offset + 68;
    if (VectorWrite(&x->Rotation, writer, __RotationOffset) == 0) {
        return 0;
    }
    uint32_t __MagnetometerOffset = offset + 100;
    if (VectorWrite(&x->Magnetometer, writer, __MagnetometerOffset) == 0) {
        return 0;
    }
    uint32_t __GravityOffset = offset + 132;
    if (VectorWrite(&x->Gravity, writer, __GravityOffset) == 0) {
        return 0;
    }
    uint32_t __EntryTimeOffset = offset + 164;
    KarmemWriterWriteAt(writer, __EntryTimeOffset, (void *) &x->EntryTime, 8);

    return offset;
}

uint32_t IMUDataEntryWriteAsRoot(IMUDataEntry * x, KarmemWriter * writer) {
    return IMUDataEntryWrite(x, writer, 0);
}

void IMUDataEntryRead(IMUDataEntry * x, IMUDataEntryViewer * viewer, KarmemReader * reader) {
    VectorRead(&x->Acceleration, IMUDataEntryViewer_Acceleration(viewer), reader);
    VectorRead(&x->AngularVelocity, IMUDataEntryViewer_AngularVelocity(viewer), reader);
    VectorRead(&x->Rotation, IMUDataEntryViewer_Rotation(viewer), reader);
    VectorRead(&x->Magnetometer, IMUDataEntryViewer_Magnetometer(viewer), reader);
    VectorRead(&x->Gravity, IMUDataEntryViewer_Gravity(viewer), reader);
    x->EntryTime = IMUDataEntryViewer_EntryTime(viewer);
}

void IMUDataEntryReadAsRoot(IMUDataEntry * x, KarmemReader * reader) {
    return IMUDataEntryRead(x, NewIMUDataEntryViewer(reader, 0), reader);
}

void IMUDataEntryReset(IMUDataEntry * x) {
    KarmemReader reader = KarmemNewReader(&_Null[0], 176);
    IMUDataEntryRead(x, (IMUDataEntryViewer *) &_Null, &reader);
}

IMUDataEntry NewIMUDataEntry() {
    IMUDataEntry r;
    memset(&r, 0, sizeof(IMUDataEntry));
    return r;
}
typedef struct {
    IMUDataEntry Entry;
} IMUEntry;

EnumPacketIdentifier IMUEntryPacketIdentifier(IMUEntry * x) {
    return EnumPacketIdentifierIMUEntry;
}

uint32_t IMUEntryWrite(IMUEntry * x, KarmemWriter * writer, uint32_t start) {
    uint32_t offset = start;
    uint32_t size = 8;
    if (offset == 0) {
        offset = KarmemWriterAlloc(writer, size);
        if (offset == 0xFFFFFFFF) {
            return 0;
        }
    }
    uint32_t __EntrySize = 176;
    uint32_t __EntryOffset = KarmemWriterAlloc(writer, __EntrySize);

    KarmemWriterWriteAt(writer, offset+0, (void *) &__EntryOffset, 4);
    if (IMUDataEntryWrite(&x->Entry, writer, __EntryOffset) == 0) {
        return 0;
    }

    return offset;
}

uint32_t IMUEntryWriteAsRoot(IMUEntry * x, KarmemWriter * writer) {
    return IMUEntryWrite(x, writer, 0);
}

void IMUEntryRead(IMUEntry * x, IMUEntryViewer * viewer, KarmemReader * reader) {
    IMUDataEntryRead(&x->Entry, IMUEntryViewer_Entry(viewer, reader), reader);
}

void IMUEntryReadAsRoot(IMUEntry * x, KarmemReader * reader) {
    return IMUEntryRead(x, NewIMUEntryViewer(reader, 0), reader);
}

void IMUEntryReset(IMUEntry * x) {
    KarmemReader reader = KarmemNewReader(&_Null[0], 176);
    IMUEntryRead(x, (IMUEntryViewer *) &_Null, &reader);
}

IMUEntry NewIMUEntry() {
    IMUEntry r;
    memset(&r, 0, sizeof(IMUEntry));
    return r;
}
typedef struct {
    IMUEntry * Entries;
    uint32_t _Entries_len;
    uint32_t _Entries_cap;
} IMUData;

EnumPacketIdentifier IMUDataPacketIdentifier(IMUData * x) {
    return EnumPacketIdentifierIMUData;
}

uint32_t IMUDataWrite(IMUData * x, KarmemWriter * writer, uint32_t start) {
    uint32_t offset = start;
    uint32_t size = 24;
    if (offset == 0) {
        offset = KarmemWriterAlloc(writer, size);
        if (offset == 0xFFFFFFFF) {
            return 0;
        }
    }
    uint32_t sizeData = 16;
    KarmemWriterWriteAt(writer, offset, (void *)&sizeData, 4);
    uint32_t __EntriesSize = 8 * x->_Entries_len;
    uint32_t __EntriesOffset = KarmemWriterAlloc(writer, __EntriesSize);

    KarmemWriterWriteAt(writer, offset+4, (void *) &__EntriesOffset, 4);
    KarmemWriterWriteAt(writer, offset+4+4, (void *) &__EntriesSize, 4);
    uint32_t __EntriesSizeEach = 8;
    KarmemWriterWriteAt(writer, offset+4+4+4, (void *) &__EntriesSizeEach, 4);
    uint32_t __EntriesIndex = 0;
    uint32_t __EntriesEnd = __EntriesOffset + __EntriesSize;
    while (__EntriesOffset < __EntriesEnd) {
        if (IMUEntryWrite(&x->Entries[__EntriesIndex], writer, __EntriesOffset) == 0) {
            return 0;
        }
        __EntriesOffset = __EntriesOffset + 8;
        __EntriesIndex = __EntriesIndex + 1;
    }

    return offset;
}

uint32_t IMUDataWriteAsRoot(IMUData * x, KarmemWriter * writer) {
    return IMUDataWrite(x, writer, 0);
}

void IMUDataRead(IMUData * x, IMUDataViewer * viewer, KarmemReader * reader) {
    IMUEntryViewer * __EntriesSlice = IMUDataViewer_Entries(viewer, reader);
    uint32_t __EntriesLen = IMUDataViewer_EntriesLength(viewer, reader);
    if (__EntriesLen > x->_Entries_cap) {
        uint32_t __EntriesCapacityTarget = __EntriesLen;
        x->Entries = (IMUEntry *) realloc(x->Entries, __EntriesCapacityTarget * sizeof(IMUEntry));
        uint32_t __EntriesNewIndex = x->_Entries_cap;
        while (__EntriesNewIndex < __EntriesCapacityTarget) {
            x->Entries[__EntriesNewIndex] = NewIMUEntry();
            __EntriesNewIndex = __EntriesNewIndex + 1;
        }
        x->_Entries_cap = __EntriesCapacityTarget;
    }
    if (__EntriesLen > x->_Entries_len) {
        x->_Entries_len = __EntriesLen;
    }
    uint32_t __EntriesIndex = 0;
    while (__EntriesIndex < __EntriesLen) {
        IMUEntryRead(&x->Entries[__EntriesIndex], &__EntriesSlice[__EntriesIndex], reader);
         __EntriesIndex = __EntriesIndex + 1;
    }
    x->_Entries_len = __EntriesLen;
}

void IMUDataReadAsRoot(IMUData * x, KarmemReader * reader) {
    return IMUDataRead(x, NewIMUDataViewer(reader, 0), reader);
}

void IMUDataReset(IMUData * x) {
    KarmemReader reader = KarmemNewReader(&_Null[0], 176);
    IMUDataRead(x, (IMUDataViewer *) &_Null, &reader);
}

IMUData NewIMUData() {
    IMUData r;
    memset(&r, 0, sizeof(IMUData));
    return r;
}
