package llvm

/*
#include <llvm-c/Target.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "os"

type (
	TargetData struct {
		C C.LLVMTargetDataRef
	}
	StructLayout struct {
		C C.LLVMStructLayoutRef
	}
	ByteOrdering C.enum_LLVMByteOrdering
)

const (
	BigEndian    = C.LLVMBigEndian
	LittleEndian = C.LLVMLittleEndian
)

// Declare all of the target-initialization functions that are available.

//#define LLVM_TARGET(TargetName) \
//  func Initialize##TargetName##TargetInfo() { C.LLVMInitialize##TargetName##TargetInfo() }
//#include "llvm/Config/Targets.def"
//#undef LLVM_TARGET  // Explicit undef to make SWIG happier
func InitializeMBlazeTargetInfo()     { C.LLVMInitializeMBlazeTargetInfo() }
func InitializeCppBackendTargetInfo() { C.LLVMInitializeCppBackendTargetInfo() }
func InitializeCBackendTargetInfo()   { C.LLVMInitializeCBackendTargetInfo() }
func InitializeBlackfinTargetInfo()   { C.LLVMInitializeBlackfinTargetInfo() }
func InitializeSystemZTargetInfo()    { C.LLVMInitializeSystemZTargetInfo() }
func InitializeMSP430TargetInfo()     { C.LLVMInitializeMSP430TargetInfo() }
func InitializeXCoreTargetInfo()      { C.LLVMInitializeXCoreTargetInfo() }
func InitializePIC16TargetInfo()      { C.LLVMInitializePIC16TargetInfo() }
func InitializeCellSPUTargetInfo()    { C.LLVMInitializeCellSPUTargetInfo() }
func InitializeMipsTargetInfo()       { C.LLVMInitializeMipsTargetInfo() }
func InitializeARMTargetInfo()        { C.LLVMInitializeARMTargetInfo() }
func InitializeAlphaTargetInfo()      { C.LLVMInitializeAlphaTargetInfo() }
func InitializePowerPCTargetInfo()    { C.LLVMInitializePowerPCTargetInfo() }
func InitializeSparcTargetInfo()      { C.LLVMInitializeSparcTargetInfo() }
func InitializeX86TargetInfo()        { C.LLVMInitializeX86TargetInfo() }

//#define LLVM_TARGET(TargetName) \
//  func Initialize##TargetName##Target() { C.LLVMInitialize##TargetName##Target() }
//#include "llvm/Config/Targets.def"
//#undef LLVM_TARGET  // Explicit undef to make SWIG happier
func InitializeMBlazeTarget()     { C.LLVMInitializeMBlazeTarget() }
func InitializeCppBackendTarget() { C.LLVMInitializeCppBackendTarget() }
func InitializeCBackendTarget()   { C.LLVMInitializeCBackendTarget() }
func InitializeBlackfinTarget()   { C.LLVMInitializeBlackfinTarget() }
func InitializeSystemZTarget()    { C.LLVMInitializeSystemZTarget() }
func InitializeMSP430Target()     { C.LLVMInitializeMSP430Target() }
func InitializeXCoreTarget()      { C.LLVMInitializeXCoreTarget() }
func InitializePIC16Target()      { C.LLVMInitializePIC16Target() }
func InitializeCellSPUTarget()    { C.LLVMInitializeCellSPUTarget() }
func InitializeMipsTarget()       { C.LLVMInitializeMipsTarget() }
func InitializeARMTarget()        { C.LLVMInitializeARMTarget() }
func InitializeAlphaTarget()      { C.LLVMInitializeAlphaTarget() }
func InitializePowerPCTarget()    { C.LLVMInitializePowerPCTarget() }
func InitializeSparcTarget()      { C.LLVMInitializeSparcTarget() }
func InitializeX86Target()        { C.LLVMInitializeX86Target() }

// LLVMInitializeAllTargetInfos - The main program should call this function if
// it wants access to all available targets that LLVM is configured to
// support.
func InitializeAllTargetInfos() { C.LLVMInitializeAllTargetInfos() }

// LLVMInitializeAllTargets - The main program should call this function if it
// wants to link in all available targets that LLVM is configured to
// support.
func InitializeAllTargets() { C.LLVMInitializeAllTargets() }

// LLVMInitializeNativeTarget - The main program should call this function to
// initialize the native target corresponding to the host. This is useful
// for JIT applications to ensure that the target gets linked in correctly.
var initializeNativeTargetError = os.NewError("Failed to initialize native target")

func InitializeNativeTarget() os.Error {
	fail := C.LLVMInitializeNativeTarget()
	if fail == 0 {
		return nil
	}
	return initializeNativeTargetError
}

//-------------------------------------------------------------------------
// llvm.TargetData
//-------------------------------------------------------------------------

// Creates target data from a target layout string.
// See the constructor llvm::TargetData::TargetData.
func NewTargetData(rep string) (td TargetData) {
	crep := C.CString(rep)
	td.C = C.LLVMCreateTargetData(crep)
	C.free(unsafe.Pointer(crep))
	return
}

// Adds target data information to a pass manager. This does not take ownership
// of the target data.
// See the method llvm::PassManagerBase::add.
func (pm PassManager) Add(td TargetData) {
	C.LLVMAddTargetData(td.C, pm.C)
}

// Converts target data to a target layout string. The string must be disposed
// with LLVMDisposeMessage.
// See the constructor llvm::TargetData::TargetData.
func (td TargetData) String() (s string) {
	cmsg := C.LLVMCopyStringRepOfTargetData(td.C)
	s = C.GoString(cmsg)
	C.LLVMDisposeMessage(cmsg)
	return
}

// Returns the byte order of a target, either LLVMBigEndian or
// LLVMLittleEndian.
// See the method llvm::TargetData::isLittleEndian.
func (td TargetData) ByteOrder() ByteOrdering { return ByteOrdering(C.LLVMByteOrder(td.C)) }

// Returns the pointer size in bytes for a target.
// See the method llvm::TargetData::getPointerSize.
func (td TargetData) PointerSize() int { return int(C.LLVMPointerSize(td.C)) }

// Returns the integer type that is the same size as a pointer on a target.
// See the method llvm::TargetData::getIntPtrType.
func (td TargetData) IntPtrType() (t Type) { t.C = C.LLVMIntPtrType(td.C); return }

// Computes the size of a type in bytes for a target.
// See the method llvm::TargetData::getTypeSizeInBits.
func (td TargetData) TypeSizeInBits(t Type) uint64 {
	return uint64(C.LLVMSizeOfTypeInBits(td.C, t.C))
}

// Computes the storage size of a type in bytes for a target.
// See the method llvm::TargetData::getTypeStoreSize.
func (td TargetData) TypeStoreSize(t Type) uint64 {
	return uint64(C.LLVMStoreSizeOfType(td.C, t.C))
}

// Computes the ABI size of a type in bytes for a target.
// See the method llvm::TargetData::getTypeAllocSize.
func (td TargetData) TypeAllocSize(t Type) uint64 {
	return uint64(C.LLVMABISizeOfType(td.C, t.C))
}

// Computes the ABI alignment of a type in bytes for a target.
// See the method llvm::TargetData::getABITypeAlignment.
func (td TargetData) ABITypeAlignment(t Type) int {
	return int(C.LLVMABIAlignmentOfType(td.C, t.C))
}

// Computes the call frame alignment of a type in bytes for a target.
// See the method llvm::TargetData::getCallFrameTypeAlignment.
func (td TargetData) CallFrameTypeAlignment(t Type) int {
	return int(C.LLVMCallFrameAlignmentOfType(td.C, t.C))
}

// Computes the preferred alignment of a type in bytes for a target.
// See the method llvm::TargetData::getPrefTypeAlignment.
func (td TargetData) PrefTypeAlignment(t Type) int {
	return int(C.LLVMPreferredAlignmentOfType(td.C, t.C))
}

// Computes the preferred alignment of a global variable in bytes for a target.
// See the method llvm::TargetData::getPreferredAlignment.
func (td TargetData) PreferredAlignment(g Value) int {
	return int(C.LLVMPreferredAlignmentOfGlobal(td.C, g.C))
}

// Computes the structure element that contains the byte offset for a target.
// See the method llvm::StructLayout::getElementContainingOffset.
func (td TargetData) ElementContainingOffset(t Type, offset uint64) int {
	return int(C.LLVMElementAtOffset(td.C, t.C, C.ulonglong(offset)))
}

// Computes the byte offset of the indexed struct element for a target.
// See the method llvm::StructLayout::getElementOffset.
func (td TargetData) ElementOffset(t Type, element int) uint64 {
	return uint64(C.LLVMOffsetOfElement(td.C, t.C, C.unsigned(element)))
}

// Struct layouts are speculatively cached. If a TargetDataRef is alive when
// types are being refined and removed, this method must be called whenever a
// struct type is removed to avoid a dangling pointer in this cache.
// See the method llvm::TargetData::InvalidateStructLayoutInfo.
func (td TargetData) InvalidateStructLayoutInfo(t Type) {
	C.LLVMInvalidateStructLayout(td.C, t.C)
}

// Deallocates a TargetData.
// See the destructor llvm::TargetData::~TargetData.
func (td TargetData) Dispose() { C.LLVMDisposeTargetData(td.C) }
