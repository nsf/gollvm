package llvm

/*
#include <llvm-c/Core.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// TODO: Add comments
// TODO: Promote *InContext functions to Context methods?
// TODO: Use Go's reflection in order to simplify bindings?
// TODO: Add IsNil for all ref types
// TODO: Add type safety?

type (
	Context struct {
		c C.LLVMContextRef
	}
	Module struct {
		c C.LLVMModuleRef
	}
	Type struct {
		c C.LLVMTypeRef
	}
	TypeHandle struct {
		c C.LLVMTypeHandleRef
	}
	Value struct {
		c C.LLVMValueRef
	}
	BasicBlock struct {
		c C.LLVMBasicBlockRef
	}
	Builder struct {
		c C.LLVMBuilderRef
	}
	ModuleProvider struct {
		c C.LLVMModuleProviderRef
	}
	MemoryBuffer struct {
		c C.LLVMMemoryBufferRef
	}
	PassManager struct {
		c C.LLVMPassManagerRef
	}
	Use struct {
		c C.LLVMUseRef
	}
	Attribute C.LLVMAttribute
	Opcode C.LLVMOpcode
	TypeKind C.LLVMTypeKind
	Linkage C.LLVMLinkage
	Visibility C.LLVMVisibility
	CallConv C.LLVMCallConv
	IntPredicate C.LLVMIntPredicate
	FloatPredicate C.LLVMRealPredicate
)

// helpers
func llvmTypeRefPtr(t *Type) *C.LLVMTypeRef    { return (*C.LLVMTypeRef)(unsafe.Pointer(t)) }
func llvmValueRefPtr(t *Value) *C.LLVMValueRef { return (*C.LLVMValueRef)(unsafe.Pointer(t)) }
func llvmBasicBlockRefPtr(t *BasicBlock) *C.LLVMBasicBlockRef {
	return (*C.LLVMBasicBlockRef)(unsafe.Pointer(t))
}
func boolToLLVMBool(b bool) C.LLVMBool {
	if b {
		return C.LLVMBool(1)
	}
	return C.LLVMBool(0)
}

//-------------------------------------------------------------------------
// llvm.Attribute
//-------------------------------------------------------------------------
const (
	NoneAttribute            = 0
	ZExtAttribute            = C.LLVMZExtAttribute
	SExtAttribute            = C.LLVMSExtAttribute
	NoReturnAttribute        = C.LLVMNoReturnAttribute
	InRegAttribute           = C.LLVMInRegAttribute
	StructRetAttribute       = C.LLVMStructRetAttribute
	NoUnwindAttribute        = C.LLVMNoUnwindAttribute
	NoAliasAttribute         = C.LLVMNoAliasAttribute
	ByValAttribute           = C.LLVMByValAttribute
	NestAttribute            = C.LLVMNestAttribute
	ReadNoneAttribute        = C.LLVMReadNoneAttribute
	ReadOnlyAttribute        = C.LLVMReadOnlyAttribute
	NoInlineAttribute        = C.LLVMNoInlineAttribute
	AlwaysInlineAttribute    = C.LLVMAlwaysInlineAttribute
	OptimizeForSizeAttribute = C.LLVMOptimizeForSizeAttribute
	StackProtectAttribute    = C.LLVMStackProtectAttribute
	StackProtectReqAttribute = C.LLVMStackProtectReqAttribute
	Alignment                = C.LLVMAlignment
	NoCaptureAttribute       = C.LLVMNoCaptureAttribute
	NoRedZoneAttribute       = C.LLVMNoRedZoneAttribute
	NoImplicitFloatAttribute = C.LLVMNoImplicitFloatAttribute
	NakedAttribute           = C.LLVMNakedAttribute
	InlineHintAttribute      = C.LLVMInlineHintAttribute
	StackAlignment           = C.LLVMStackAlignment
)

//-------------------------------------------------------------------------
// llvm.Opcode
//-------------------------------------------------------------------------
const (
	Ret         = C.LLVMRet
	Br          = C.LLVMBr
	Switch      = C.LLVMSwitch
	IndirectBr  = C.LLVMIndirectBr
	Invoke      = C.LLVMInvoke
	Unwind      = C.LLVMUnwind
	Unreachable = C.LLVMUnreachable

	// Standard Binary Operators
	Add  = C.LLVMAdd
	FAdd = C.LLVMFAdd
	Sub  = C.LLVMSub
	FSub = C.LLVMFSub
	Mul  = C.LLVMMul
	FMul = C.LLVMFMul
	UDiv = C.LLVMUDiv
	SDiv = C.LLVMSDiv
	FDiv = C.LLVMFDiv
	URem = C.LLVMURem
	SRem = C.LLVMSRem
	FRem = C.LLVMFRem

	// Logical Operators
	Shl  = C.LLVMShl
	LShr = C.LLVMLShr
	AShr = C.LLVMAShr
	And  = C.LLVMAnd
	Or   = C.LLVMOr
	Xor  = C.LLVMXor

	// Memory Operators
	Alloca        = C.LLVMAlloca
	Load          = C.LLVMLoad
	Store         = C.LLVMStore
	GetElementPtr = C.LLVMGetElementPtr

	// Cast Operators
	Trunc    = C.LLVMTrunc
	ZExt     = C.LLVMZExt
	SExt     = C.LLVMSExt
	FPToUI   = C.LLVMFPToUI
	FPToSI   = C.LLVMFPToSI
	UIToFP   = C.LLVMUIToFP
	SIToFP   = C.LLVMSIToFP
	FPTrunc  = C.LLVMFPTrunc
	FPExt    = C.LLVMFPExt
	PtrToInt = C.LLVMPtrToInt
	IntToPtr = C.LLVMIntToPtr
	BitCast  = C.LLVMBitCast

	// Other Operators
	ICmp   = C.LLVMICmp
	FCmp   = C.LLVMFCmp
	PHI    = C.LLVMPHI
	Call   = C.LLVMCall
	Select = C.LLVMSelect
	// UserOp1
	// UserOp2
	VAArg          = C.LLVMVAArg
	ExtractElement = C.LLVMExtractElement
	InsertElement  = C.LLVMInsertElement
	ShuffleVector  = C.LLVMShuffleVector
	ExtractValue   = C.LLVMExtractValue
	InsertValue    = C.LLVMInsertValue
)

//-------------------------------------------------------------------------
// llvm.TypeKind
//-------------------------------------------------------------------------

const (
	VoidTypeKind      = C.LLVMVoidTypeKind
	FloatTypeKind     = C.LLVMFloatTypeKind
	DoubleTypeKind    = C.LLVMDoubleTypeKind
	X86_FP80TypeKind  = C.LLVMX86_FP80TypeKind
	FP128TypeKind     = C.LLVMFP128TypeKind
	PPC_FP128TypeKind = C.LLVMPPC_FP128TypeKind
	LabelTypeKind     = C.LLVMLabelTypeKind
	IntegerTypeKind   = C.LLVMIntegerTypeKind
	FunctionTypeKind  = C.LLVMFunctionTypeKind
	StructTypeKind    = C.LLVMStructTypeKind
	ArrayTypeKind     = C.LLVMArrayTypeKind
	PointerTypeKind   = C.LLVMPointerTypeKind
	OpaqueTypeKind    = C.LLVMOpaqueTypeKind
	VectorTypeKind    = C.LLVMVectorTypeKind
	MetadataTypeKind  = C.LLVMMetadataTypeKind
)

//-------------------------------------------------------------------------
// llvm.Linkage
//-------------------------------------------------------------------------

const (
	ExternalLinkage                 = C.LLVMExternalLinkage
	AvailableExternallyLinkage      = C.LLVMAvailableExternallyLinkage
	LinkOnceAnyLinkage              = C.LLVMLinkOnceAnyLinkage
	LinkOnceODRLinkage              = C.LLVMLinkOnceODRLinkage
	WeakAnyLinkage                  = C.LLVMWeakAnyLinkage
	WeakODRLinkage                  = C.LLVMWeakODRLinkage
	AppendingLinkage                = C.LLVMAppendingLinkage
	InternalLinkage                 = C.LLVMInternalLinkage
	PrivateLinkage                  = C.LLVMPrivateLinkage
	DLLImportLinkage                = C.LLVMDLLImportLinkage
	DLLExportLinkage                = C.LLVMDLLExportLinkage
	ExternalWeakLinkage             = C.LLVMExternalWeakLinkage
	GhostLinkage                    = C.LLVMGhostLinkage
	CommonLinkage                   = C.LLVMCommonLinkage
	LinkerPrivateLinkage            = C.LLVMLinkerPrivateLinkage
	LinkerPrivateWeakLinkage        = C.LLVMLinkerPrivateWeakLinkage
	LinkerPrivateWeakDefAutoLinkage = C.LLVMLinkerPrivateWeakDefAutoLinkage
)

//-------------------------------------------------------------------------
// llvm.Visibility
//-------------------------------------------------------------------------
const (
	DefaultVisibility   = C.LLVMDefaultVisibility
	HiddenVisibility    = C.LLVMHiddenVisibility
	ProtectedVisibility = C.LLVMProtectedVisibility
)

//-------------------------------------------------------------------------
// llvm.CallConv
//-------------------------------------------------------------------------

const (
	CCallConv           = C.LLVMCCallConv
	FastCallConv        = C.LLVMFastCallConv
	ColdCallConv        = C.LLVMColdCallConv
	X86StdcallCallConv  = C.LLVMX86StdcallCallConv
	X86FastcallCallConv = C.LLVMX86FastcallCallConv
)

//-------------------------------------------------------------------------
// llvm.IntPredicate
//-------------------------------------------------------------------------

const (
	IntEQ  = C.LLVMIntEQ
	IntNE  = C.LLVMIntNE
	IntUGT = C.LLVMIntUGT
	IntUGE = C.LLVMIntUGE
	IntULT = C.LLVMIntULT
	IntULE = C.LLVMIntULE
	IntSGT = C.LLVMIntSGT
	IntSGE = C.LLVMIntSGE
	IntSLT = C.LLVMIntSLT
	IntSLE = C.LLVMIntSLE
)

//-------------------------------------------------------------------------
// llvm.FloatPredicate
//-------------------------------------------------------------------------

const (
	FloatPredicateFalse = C.LLVMRealPredicateFalse
	FloatOEQ            = C.LLVMRealOEQ
	FloatOGT            = C.LLVMRealOGT
	FloatOGE            = C.LLVMRealOGE
	FloatOLT            = C.LLVMRealOLT
	FloatOLE            = C.LLVMRealOLE
	FloatONE            = C.LLVMRealONE
	FloatORD            = C.LLVMRealORD
	FloatUNO            = C.LLVMRealUNO
	FloatUEQ            = C.LLVMRealUEQ
	FloatUGT            = C.LLVMRealUGT
	FloatUGE            = C.LLVMRealUGE
	FloatULT            = C.LLVMRealULT
	FloatULE            = C.LLVMRealULE
	FloatUNE            = C.LLVMRealUNE
	FloatPredicateTrue  = C.LLVMRealPredicateTrue
)

//-------------------------------------------------------------------------
// llvm.Context
//-------------------------------------------------------------------------

func NewContext() Context    { return Context{C.LLVMContextCreate()} }
func GlobalContext() Context { return Context{C.LLVMGetGlobalContext()} }
func (c Context) Dispose()   { C.LLVMContextDispose(c.c) }

/*
unsigned LLVMGetMDKindIDInContext(LLVMContextRef C, const char* Name,
                                  unsigned SLen);
unsigned LLVMGetMDKindID(const char* Name, unsigned SLen);
*/

//-------------------------------------------------------------------------
// llvm.Module
//-------------------------------------------------------------------------

// Create and destroy modules.
// See llvm::Module::Module.
func NewModule(name string) (m Module) {
	cname := C.CString(name)
	m.c = C.LLVMModuleCreateWithName(cname)
	C.free(unsafe.Pointer(cname))
	return
}

func NewModuleInContext(name string, c Context) (m Module) {
	cname := C.CString(name)
	m.c = C.LLVMModuleCreateWithNameInContext(cname, c.c)
	C.free(unsafe.Pointer(cname))
	return
}

// See llvm::Module::~Module
func (m Module) Dispose() { C.LLVMDisposeModule(m.c) }

// Data layout. See Module::getDataLayout.
func (m Module) DataLayout() string {
	clayout := C.LLVMGetDataLayout(m.c)
	return C.GoString(clayout)
}

func (m Module) SetDataLayout(layout string) {
	clayout := C.CString(layout)
	C.LLVMSetDataLayout(m.c, clayout)
	C.free(unsafe.Pointer(clayout))
}

// Target triple. See Module::getTargetTriple.
func (m Module) Target() string {
	ctarget := C.LLVMGetTarget(m.c)
	return C.GoString(ctarget)
}
func (m Module) SetTarget(target string) {
	ctarget := C.CString(target)
	C.LLVMSetTarget(m.c, ctarget)
	C.free(unsafe.Pointer(ctarget))
}

// See Module::addTypeName.
func (m Module) AddTypeName(name string, ty Type) bool {
	cname := C.CString(name)
	result := C.LLVMAddTypeName(m.c, cname, ty.c) != 0
	C.free(unsafe.Pointer(cname))
	return result
}
func (m Module) DeleteTypeName(name string) {
	cname := C.CString(name)
	C.LLVMDeleteTypeName(m.c, cname)
	C.free(unsafe.Pointer(cname))
}
func (m Module) GetTypeByName(name string) (t Type) {
	cname := C.CString(name)
	t.c = C.LLVMGetTypeByName(m.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}

// See Module::dump.
func (m Module) Dump() {
	C.LLVMDumpModule(m.c)
}

// See Module::setModuleInlineAsm.
func (m Module) SetInlineAsm(asm string) {
	casm := C.CString(asm)
	C.LLVMSetModuleInlineAsm(m.c, casm)
	C.free(unsafe.Pointer(casm))
}

//-------------------------------------------------------------------------
// llvm.Type
//-------------------------------------------------------------------------

// LLVM types conform to the following hierarchy:
// 
//   types:
//     integer type
//     real type
//     function type
//     sequence types:
//       array type
//       pointer type
//       vector type
//     void type
//     label type
//     opaque type

// See llvm::LLVMTypeKind::getTypeID.
func (t Type) TypeKind() TypeKind { return TypeKind(C.LLVMGetTypeKind(t.c)) }

// See llvm::LLVMType::getContext. 
func (t Type) Context() (c Context) {
	c.c = C.LLVMGetTypeContext(t.c)
	return
}

// Operations on integer types
func Int1TypeInContext(c Context) (t Type)  { t.c = C.LLVMInt1TypeInContext(c.c); return }
func Int8TypeInContext(c Context) (t Type)  { t.c = C.LLVMInt8TypeInContext(c.c); return }
func Int16TypeInContext(c Context) (t Type) { t.c = C.LLVMInt16TypeInContext(c.c); return }
func Int32TypeInContext(c Context) (t Type) { t.c = C.LLVMInt32TypeInContext(c.c); return }
func Int64TypeInContext(c Context) (t Type) { t.c = C.LLVMInt64TypeInContext(c.c); return }

func IntTypeInContext(c Context) (t Type, numbits int) {
	t.c = C.LLVMIntTypeInContext(c.c, C.unsigned(numbits))
	return
}

func Int1Type() (t Type)  { t.c = C.LLVMInt1Type(); return }
func Int8Type() (t Type)  { t.c = C.LLVMInt8Type(); return }
func Int16Type() (t Type) { t.c = C.LLVMInt16Type(); return }
func Int32Type() (t Type) { t.c = C.LLVMInt32Type(); return }
func Int64Type() (t Type) { t.c = C.LLVMInt64Type(); return }

func IntType(numbits int) (t Type) {
	t.c = C.LLVMIntType(C.unsigned(numbits))
	return
}

func (t Type) IntTypeWidth() int {
	return int(C.LLVMGetIntTypeWidth(t.c))
}

// Operations on real types
func FloatTypeInContext(c Context) (t Type)    { t.c = C.LLVMFloatTypeInContext(c.c); return }
func DoubleTypeInContext(c Context) (t Type)   { t.c = C.LLVMDoubleTypeInContext(c.c); return }
func X86FP80TypeInContext(c Context) (t Type)  { t.c = C.LLVMX86FP80TypeInContext(c.c); return }
func FP128TypeInContext(c Context) (t Type)    { t.c = C.LLVMFP128TypeInContext(c.c); return }
func PPCFP128TypeInContext(c Context) (t Type) { t.c = C.LLVMPPCFP128TypeInContext(c.c); return }

func FloatType() (t Type)    { t.c = C.LLVMFloatType(); return }
func DoubleType() (t Type)   { t.c = C.LLVMDoubleType(); return }
func X86FP80Type() (t Type)  { t.c = C.LLVMX86FP80Type(); return }
func FP128Type() (t Type)    { t.c = C.LLVMFP128Type(); return }
func PPCFP128Type() (t Type) { t.c = C.LLVMPPCFP128Type(); return }

// Operations on function types
func FunctionType(returnType Type, paramTypes []Type, isVarArg bool) (t Type) {
	t.c = C.LLVMFunctionType(returnType.c,
		llvmTypeRefPtr(&paramTypes[0]),
		C.unsigned(len(paramTypes)),
		boolToLLVMBool(isVarArg))
	return
}

func (t Type) IsFunctionVarArg() bool { return C.LLVMIsFunctionVarArg(t.c) != 0 }
func (t Type) ReturnType() (rt Type)  { rt.c = C.LLVMGetReturnType(t.c); return }
func (t Type) ParamTypesCount() int   { return int(C.LLVMCountParamTypes(t.c)) }
func (t Type) ParamTypes() []Type {
	out := make([]Type, t.ParamTypesCount())
	C.LLVMGetParamTypes(t.c, llvmTypeRefPtr(&out[0]))
	return out
}

// Operations on struct types
func StructTypeInContext(c Context, elementTypes []Type, packed bool) (t Type) {
	t.c = C.LLVMStructTypeInContext(c.c,
		llvmTypeRefPtr(&elementTypes[0]),
		C.unsigned(len(elementTypes)),
		boolToLLVMBool(packed))
	return
}

func StructType(elementTypes []Type, packed bool) (t Type) {
	t.c = C.LLVMStructType(llvmTypeRefPtr(&elementTypes[0]),
		C.unsigned(len(elementTypes)),
		boolToLLVMBool(packed))
	return
}

func (t Type) IsStructPacked() bool         { return C.LLVMIsPackedStruct(t.c) != 0 }
func (t Type) StructElementTypesCount() int { return int(C.LLVMCountStructElementTypes(t.c)) }
func (t Type) StructElementTypes() []Type {
	out := make([]Type, t.StructElementTypesCount())
	C.LLVMGetStructElementTypes(t.c, llvmTypeRefPtr(&out[0]))
	return out
}

// Operations on array, pointer, and vector types (sequence types)
func ArrayType(elementType Type, elementCount int) (t Type) {
	t.c = C.LLVMArrayType(elementType.c, C.unsigned(elementCount))
	return
}
func PointerType(elementType Type, addressSpace int) (t Type) {
	t.c = C.LLVMPointerType(elementType.c, C.unsigned(addressSpace))
	return
}
func VectorType(elementType Type, elementCount int) (t Type) {
	t.c = C.LLVMVectorType(elementType.c, C.unsigned(elementCount))
	return
}

func (t Type) ElementType() (rt Type)   { rt.c = C.LLVMGetElementType(t.c); return }
func (t Type) ArrayLength() int         { return int(C.LLVMGetArrayLength(t.c)) }
func (t Type) PointerAddressSpace() int { return int(C.LLVMGetPointerAddressSpace(t.c)) }
func (t Type) VectorSize() int          { return int(C.LLVMGetVectorSize(t.c)) }

// Operations on other types
func VoidTypeInContext(c Context) (t Type)   { t.c = C.LLVMVoidTypeInContext(c.c); return }
func LabelTypeInContext(c Context) (t Type)  { t.c = C.LLVMLabelTypeInContext(c.c); return }
func OpaqueTypeInContext(c Context) (t Type) { t.c = C.LLVMOpaqueTypeInContext(c.c); return }

func VoidType() (t Type)   { t.c = C.LLVMVoidType(); return }
func LabelType() (t Type)  { t.c = C.LLVMLabelType(); return }
func OpaqueType() (t Type) { t.c = C.LLVMOpaqueType(); return }

// Operations on type handles
func (t Type) TypeHandle() (th TypeHandle) {
	th.c = C.LLVMCreateTypeHandle(t.c)
	return
}
func (t Type) Refine(concrete Type) { C.LLVMRefineType(t.c, concrete.c) }
func (th TypeHandle) Get() (t Type) { t.c = C.LLVMResolveTypeHandle(th.c); return }
func (th TypeHandle) Dispose()      { C.LLVMDisposeTypeHandle(th.c) }

//-------------------------------------------------------------------------
// llvm.Value
//-------------------------------------------------------------------------
// The bulk of LLVM's object model consists of values, which comprise a very
// rich type hierarchy.

// TODO
//#define LLVM_FOR_EACH_VALUE_SUBCLASS(macro) \
//  macro(Argument)                           \
//  macro(BasicBlock)                         \
//  macro(InlineAsm)                          \
//  macro(User)                               \
//    macro(Constant)                         \
//      macro(ConstantAggregateZero)          \
//      macro(ConstantArray)                  \
//      macro(ConstantExpr)                   \
//      macro(ConstantFP)                     \
//      macro(ConstantInt)                    \
//      macro(ConstantPointerNull)            \
//      macro(ConstantStruct)                 \
//      macro(ConstantVector)                 \
//      macro(GlobalValue)                    \
//        macro(Function)                     \
//        macro(GlobalAlias)                  \
//        macro(GlobalVariable)               \
//      macro(UndefValue)                     \
//    macro(Instruction)                      \
//      macro(BinaryOperator)                 \
//      macro(CallInst)                       \
//        macro(IntrinsicInst)                \
//          macro(DbgInfoIntrinsic)           \
//            macro(DbgDeclareInst)           \
//          macro(EHSelectorInst)             \
//          macro(MemIntrinsic)               \
//            macro(MemCpyInst)               \
//            macro(MemMoveInst)              \
//            macro(MemSetInst)               \
//      macro(CmpInst)                        \
//      macro(FCmpInst)                       \
//      macro(ICmpInst)                       \
//      macro(ExtractElementInst)             \
//      macro(GetElementPtrInst)              \
//      macro(InsertElementInst)              \
//      macro(InsertValueInst)                \
//      macro(PHINode)                        \
//      macro(SelectInst)                     \
//      macro(ShuffleVectorInst)              \
//      macro(StoreInst)                      \
//      macro(TerminatorInst)                 \
//        macro(BranchInst)                   \
//        macro(InvokeInst)                   \
//        macro(ReturnInst)                   \
//        macro(SwitchInst)                   \
//        macro(UnreachableInst)              \
//        macro(UnwindInst)                   \
//    macro(UnaryInstruction)                 \
//      macro(AllocaInst)                     \
//      macro(CastInst)                       \
//        macro(BitCastInst)                  \
//        macro(FPExtInst)                    \
//        macro(FPToSIInst)                   \
//        macro(FPToUIInst)                   \
//        macro(FPTruncInst)                  \
//        macro(IntToPtrInst)                 \
//        macro(PtrToIntInst)                 \
//        macro(SExtInst)                     \
//        macro(SIToFPInst)                   \
//        macro(TruncInst)                    \
//        macro(UIToFPInst)                   \
//        macro(ZExtInst)                     \
//      macro(ExtractValueInst)               \
//      macro(LoadInst)                       \
//      macro(VAArgInst)

// Operations on all values
func (v Value) Type() (t Type) { t.c = C.LLVMTypeOf(v.c); return }
func (v Value) Name() string   { return C.GoString(C.LLVMGetValueName(v.c)) }
func (v Value) SetName(name string) {
	cname := C.CString(name)
	C.LLVMSetValueName(v.c, cname)
	C.free(unsafe.Pointer(cname))
}
func (v Value) Dump()                       { C.LLVMDumpValue(v.c) }
func (v Value) ReplaceAllUsesWith(nv Value) { C.LLVMReplaceAllUsesWith(v.c, nv.c) }
func (v Value) HasMetadata() bool           { return C.LLVMHasMetadata(v.c) != 0 }
func (v Value) Metadata(kind int) (rv Value) {
	rv.c = C.LLVMGetMetadata(v.c, C.unsigned(kind))
	return
}
func (v Value) SetMetadata(kind int, node Value) {
	C.LLVMSetMetadata(v.c, C.unsigned(kind), node.c)
}

// TODO
// Conversion functions. Return the input value if it is an instance of the
// specified class, otherwise NULL. See llvm::dyn_cast_or_null<>.
//#define LLVM_DECLARE_VALUE_CAST(name) \
//  LLVMValueRef LLVMIsA##name(LLVMValueRef Val);
//LLVM_FOR_EACH_VALUE_SUBCLASS(LLVM_DECLARE_VALUE_CAST)

// Operations on Uses
func (v Value) FirstUse() (u Use)  { u.c = C.LLVMGetFirstUse(v.c); return }
func (u Use) NextUse() (ru Use)    { ru.c = C.LLVMGetNextUse(u.c); return }
func (u Use) User() (v Value)      { v.c = C.LLVMGetUser(u.c); return }
func (u Use) UsedValue() (v Value) { v.c = C.LLVMGetUsedValue(u.c); return }

// Operations on Users
func (v Value) Operand(i int) (rv Value)   { rv.c = C.LLVMGetOperand(v.c, C.unsigned(i)); return }
func (v Value) SetOperand(i int, op Value) { C.LLVMSetOperand(v.c, C.unsigned(i), op.c) }
func (v Value) OperandsCount() int         { return int(C.LLVMGetNumOperands(v.c)) }

// Operations on constants of any type
func ConstNull(t Type) (v Value)        { v.c = C.LLVMConstNull(t.c); return }
func ConstAllOnes(t Type) (v Value)     { v.c = C.LLVMConstAllOnes(t.c); return }
func Undef(t Type) (v Value)            { v.c = C.LLVMGetUndef(t.c); return }
func (v Value) IsConstant() bool        { return C.LLVMIsConstant(v.c) != 0 }
func (v Value) IsNull() bool            { return C.LLVMIsNull(v.c) != 0 }
func (v Value) IsUndef() bool           { return C.LLVMIsUndef(v.c) != 0 }
func ConstPointerNull(t Type) (v Value) { v.c = C.LLVMConstPointerNull(t.c); return }

// Operations on metadata
func MDStringInContext(c Context, str string) (v Value) {
	cstr := C.CString(str)
	v.c = C.LLVMMDStringInContext(c.c, cstr, C.unsigned(len(str)))
	C.free(unsafe.Pointer(cstr))
	return
}
func MDString(str string) (v Value) {
	cstr := C.CString(str)
	v.c = C.LLVMMDString(cstr, C.unsigned(len(str)))
	C.free(unsafe.Pointer(cstr))
	return
}
func MDNodeInContext(c Context, vals []Value) (v Value) {
	v.c = C.LLVMMDNodeInContext(c.c, llvmValueRefPtr(&vals[0]), C.unsigned(len(vals)))
	return
}
func MDNode(vals []Value) (v Value) {
	v.c = C.LLVMMDNode(llvmValueRefPtr(&vals[0]), C.unsigned(len(vals)))
	return
}

// Operations on scalar constants
func ConstInt(t Type, n uint64, signExtend bool) (v Value) {
	v.c = C.LLVMConstInt(t.c,
		C.ulonglong(n),
		boolToLLVMBool(signExtend))
	return
}
func ConstIntFromString(t Type, str string, radix int) (v Value) {
	cstr := C.CString(str)
	v.c = C.LLVMConstIntOfString(t.c, cstr, C.uint8_t(radix))
	C.free(unsafe.Pointer(cstr))
	return
}
func ConstFloat(t Type, n float64) (v Value) {
	v.c = C.LLVMConstReal(t.c, C.double(n))
	return
}
func ConstFloatFromString(t Type, str string) (v Value) {
	cstr := C.CString(str)
	v.c = C.LLVMConstRealOfString(t.c, cstr)
	C.free(unsafe.Pointer(cstr))
	return
}

func (v Value) ZExtValue() uint64 { return uint64(C.LLVMConstIntGetZExtValue(v.c)) }
func (v Value) SExtValue() int64  { return int64(C.LLVMConstIntGetSExtValue(v.c)) }

// Operations on composite constants
func ConstStringInContext(c Context, str string, addnull bool) (v Value) {
	cstr := C.CString(str)
	v.c = C.LLVMConstStringInContext(c.c, cstr,
		C.unsigned(len(str)), boolToLLVMBool(!addnull))
	C.free(unsafe.Pointer(cstr))
	return
}
func ConstStructInContext(c Context, constVals []Value, packed bool) (v Value) {
	v.c = C.LLVMConstStructInContext(c.c,
		llvmValueRefPtr(&constVals[0]),
		C.unsigned(len(constVals)),
		boolToLLVMBool(packed))
	return
}
func ConstString(str string, addnull bool) (v Value) {
	cstr := C.CString(str)
	v.c = C.LLVMConstString(cstr,
		C.unsigned(len(str)), boolToLLVMBool(!addnull))
	C.free(unsafe.Pointer(cstr))
	return
}
func ConstArray(t Type, constVals []Value) (v Value) {
	v.c = C.LLVMConstArray(t.c,
		llvmValueRefPtr(&constVals[0]),
		C.unsigned(len(constVals)))
	return
}
func ConstStruct(constVals []Value, packed bool) (v Value) {
	v.c = C.LLVMConstStruct(llvmValueRefPtr(&constVals[0]),
		C.unsigned(len(constVals)),
		boolToLLVMBool(packed))
	return
}
func ConstVector(scalarConstVals []Value, packed bool) (v Value) {
	v.c = C.LLVMConstVector(llvmValueRefPtr(&scalarConstVals[0]),
		C.unsigned(len(scalarConstVals)))
	return
}

// Constant expressions 
func (v Value) Opcode() Opcode            { return Opcode(C.LLVMGetConstOpcode(v.c)) }
func AlignOf(t Type) (v Value)                { v.c = C.LLVMAlignOf(t.c); return }
func SizeOf(t Type) (v Value)                 { v.c = C.LLVMSizeOf(t.c); return }
func ConstNeg(v Value) (rv Value)             { rv.c = C.LLVMConstNeg(v.c); return }
func ConstNSWNeg(v Value) (rv Value)          { rv.c = C.LLVMConstNSWNeg(v.c); return }
func ConstNUWNeg(v Value) (rv Value)          { rv.c = C.LLVMConstNUWNeg(v.c); return }
func ConstFNeg(v Value) (rv Value)            { rv.c = C.LLVMConstFNeg(v.c); return }
func ConstNot(v Value) (rv Value)             { rv.c = C.LLVMConstNot(v.c); return }
func ConstAdd(lhs, rhs Value) (v Value)       { v.c = C.LLVMConstAdd(lhs.c, rhs.c); return }
func ConstNSWAdd(lhs, rhs Value) (v Value)    { v.c = C.LLVMConstNSWAdd(lhs.c, rhs.c); return }
func ConstNUWAdd(lhs, rhs Value) (v Value)    { v.c = C.LLVMConstNUWAdd(lhs.c, rhs.c); return }
func ConstFAdd(lhs, rhs Value) (v Value)      { v.c = C.LLVMConstFAdd(lhs.c, rhs.c); return }
func ConstSub(lhs, rhs Value) (v Value)       { v.c = C.LLVMConstSub(lhs.c, rhs.c); return }
func ConstNSWSub(lhs, rhs Value) (v Value)    { v.c = C.LLVMConstNSWSub(lhs.c, rhs.c); return }
func ConstNUWSub(lhs, rhs Value) (v Value)    { v.c = C.LLVMConstNUWSub(lhs.c, rhs.c); return }
func ConstFSub(lhs, rhs Value) (v Value)      { v.c = C.LLVMConstFSub(lhs.c, rhs.c); return }
func ConstMul(lhs, rhs Value) (v Value)       { v.c = C.LLVMConstMul(lhs.c, rhs.c); return }
func ConstNSWMul(lhs, rhs Value) (v Value)    { v.c = C.LLVMConstNSWMul(lhs.c, rhs.c); return }
func ConstNUWMul(lhs, rhs Value) (v Value)    { v.c = C.LLVMConstNUWMul(lhs.c, rhs.c); return }
func ConstFMul(lhs, rhs Value) (v Value)      { v.c = C.LLVMConstFMul(lhs.c, rhs.c); return }
func ConstUDiv(lhs, rhs Value) (v Value)      { v.c = C.LLVMConstUDiv(lhs.c, rhs.c); return }
func ConstSDiv(lhs, rhs Value) (v Value)      { v.c = C.LLVMConstSDiv(lhs.c, rhs.c); return }
func ConstExactSDiv(lhs, rhs Value) (v Value) { v.c = C.LLVMConstExactSDiv(lhs.c, rhs.c); return }
func ConstFDiv(lhs, rhs Value) (v Value)      { v.c = C.LLVMConstFDiv(lhs.c, rhs.c); return }
func ConstURem(lhs, rhs Value) (v Value)      { v.c = C.LLVMConstURem(lhs.c, rhs.c); return }
func ConstSRem(lhs, rhs Value) (v Value)      { v.c = C.LLVMConstSRem(lhs.c, rhs.c); return }
func ConstFRem(lhs, rhs Value) (v Value)      { v.c = C.LLVMConstFRem(lhs.c, rhs.c); return }
func ConstAnd(lhs, rhs Value) (v Value)       { v.c = C.LLVMConstAnd(lhs.c, rhs.c); return }
func ConstOr(lhs, rhs Value) (v Value)        { v.c = C.LLVMConstOr(lhs.c, rhs.c); return }
func ConstXor(lhs, rhs Value) (v Value)       { v.c = C.LLVMConstXor(lhs.c, rhs.c); return }

func ConstICmp(pred IntPredicate, lhs, rhs Value) (v Value) {
	v.c = C.LLVMConstICmp(C.LLVMIntPredicate(pred), lhs.c, rhs.c)
	return
}
func ConstFCmp(pred FloatPredicate, lhs, rhs Value) (v Value) {
	v.c = C.LLVMConstFCmp(C.LLVMRealPredicate(pred), lhs.c, rhs.c)
	return
}

func ConstShl(lhs, rhs Value) (v Value)  { v.c = C.LLVMConstShl(lhs.c, rhs.c); return }
func ConstLShr(lhs, rhs Value) (v Value) { v.c = C.LLVMConstLShr(lhs.c, rhs.c); return }
func ConstAShr(lhs, rhs Value) (v Value) { v.c = C.LLVMConstAShr(lhs.c, rhs.c); return }

func ConstGEP(v Value, indices []Value) (rv Value) {
	rv.c = C.LLVMConstGEP(v.c,
		llvmValueRefPtr(&indices[0]), C.unsigned(len(indices)))
	return
}
func ConstInBoundsGEP(v Value, indices []Value) (rv Value) {
	rv.c = C.LLVMConstInBoundsGEP(v.c,
		llvmValueRefPtr(&indices[0]), C.unsigned(len(indices)))
	return
}
func ConstTrunc(v Value, t Type) (rv Value)         { rv.c = C.LLVMConstTrunc(v.c, t.c); return }
func ConstSExt(v Value, t Type) (rv Value)          { rv.c = C.LLVMConstSExt(v.c, t.c); return }
func ConstZExt(v Value, t Type) (rv Value)          { rv.c = C.LLVMConstZExt(v.c, t.c); return }
func ConstFPTrunc(v Value, t Type) (rv Value)       { rv.c = C.LLVMConstFPTrunc(v.c, t.c); return }
func ConstFPExt(v Value, t Type) (rv Value)         { rv.c = C.LLVMConstFPExt(v.c, t.c); return }
func ConstUIToFP(v Value, t Type) (rv Value)        { rv.c = C.LLVMConstUIToFP(v.c, t.c); return }
func ConstSIToFP(v Value, t Type) (rv Value)        { rv.c = C.LLVMConstSIToFP(v.c, t.c); return }
func ConstFPToUI(v Value, t Type) (rv Value)        { rv.c = C.LLVMConstFPToUI(v.c, t.c); return }
func ConstFPToSI(v Value, t Type) (rv Value)        { rv.c = C.LLVMConstFPToSI(v.c, t.c); return }
func ConstPtrToInt(v Value, t Type) (rv Value)      { rv.c = C.LLVMConstPtrToInt(v.c, t.c); return }
func ConstIntToPtr(v Value, t Type) (rv Value)      { rv.c = C.LLVMConstIntToPtr(v.c, t.c); return }
func ConstBitCast(v Value, t Type) (rv Value)       { rv.c = C.LLVMConstBitCast(v.c, t.c); return }
func ConstZExtOrBitCast(v Value, t Type) (rv Value) { rv.c = C.LLVMConstZExtOrBitCast(v.c, t.c); return }
func ConstSExtOrBitCast(v Value, t Type) (rv Value) { rv.c = C.LLVMConstSExtOrBitCast(v.c, t.c); return }
func ConstTruncOrBitCast(v Value, t Type) (rv Value) {
	rv.c = C.LLVMConstTruncOrBitCast(v.c, t.c)
	return
}
func ConstPointerCast(v Value, t Type) (rv Value) { rv.c = C.LLVMConstPointerCast(v.c, t.c); return }
func ConstIntCast(v Value, t Type, signed bool) (rv Value) {
	rv.c = C.LLVMConstIntCast(v.c, t.c, boolToLLVMBool(signed))
	return
}
func ConstFPCast(v Value, t Type) (rv Value) { rv.c = C.LLVMConstFPCast(v.c, t.c); return }
func ConstSelect(cond, iftrue, iffalse Value) (rv Value) {
	rv.c = C.LLVMConstSelect(cond.c, iftrue.c, iffalse.c)
	return
}
func ConstExtractElement(vec, i Value) (rv Value) {
	rv.c = C.LLVMConstExtractElement(vec.c, i.c)
	return
}
func ConstInsertElement(vec, elem, i Value) (rv Value) {
	rv.c = C.LLVMConstInsertElement(vec.c, elem.c, i.c)
	return
}
func ConstShuffleVector(veca, vecb, mask Value) (rv Value) {
	rv.c = C.LLVMConstShuffleVector(veca.c, vecb.c, mask.c)
	return
}
//TODO
//LLVMValueRef LLVMConstExtractValue(LLVMValueRef AggConstant, unsigned *IdxList,
//                                   unsigned NumIdx);
//LLVMValueRef LLVMConstInsertValue(LLVMValueRef AggConstant,
//                                  LLVMValueRef ElementValueConstant,
//                                  unsigned *IdxList, unsigned NumIdx);
func BlockAddress(f Value, bb BasicBlock) (v Value) {
	v.c = C.LLVMBlockAddress(f.c, bb.c)
	return
}

// Operations on global variables, functions, and aliases (globals)
func (v Value) GlobalParent() (m Module) { m.c = C.LLVMGetGlobalParent(v.c); return }
func (v Value) IsDeclaration() bool      { return C.LLVMIsDeclaration(v.c) != 0 }
func (v Value) Linkage() Linkage     { return Linkage(C.LLVMGetLinkage(v.c)) }
func (v Value) SetLinkage(l Linkage)     { C.LLVMSetLinkage(v.c, C.LLVMLinkage(l)) }
func (v Value) Section() string          { return C.GoString(C.LLVMGetSection(v.c)) }
func (v Value) SetSection(str string) {
	cstr := C.CString(str)
	C.LLVMSetSection(v.c, cstr)
	C.free(unsafe.Pointer(cstr))
}
func (v Value) Visibility() Visibility { return Visibility(C.LLVMGetVisibility(v.c)) }
func (v Value) SetVisibility(vi Visibility) { C.LLVMSetVisibility(v.c, C.LLVMVisibility(vi)) }
func (v Value) Alignment() int              { return int(C.LLVMGetAlignment(v.c)) }
func (v Value) SetAlignment(a int)          { C.LLVMSetAlignment(v.c, C.unsigned(a)) }

// Operations on global variables
func AddGlobal(m Module, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMAddGlobal(m.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func AddGlobalInAddressSpace(m Module, t Type, name string, addressSpace int) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMAddGlobalInAddressSpace(m.c, t.c, cname, C.unsigned(addressSpace))
	C.free(unsafe.Pointer(cname))
	return
}
func (m Module) NamedGlobal(name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMGetNamedGlobal(m.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (m Module) FirstGlobal() (v Value)   { v.c = C.LLVMGetFirstGlobal(m.c); return }
func (m Module) LastGlobal() (v Value)    { v.c = C.LLVMGetLastGlobal(m.c); return }
func NextGlobal(v Value) (rv Value)       { rv.c = C.LLVMGetNextGlobal(v.c); return }
func PrevGlobal(v Value) (rv Value)       { rv.c = C.LLVMGetPreviousGlobal(v.c); return }
func (v Value) EraseFromParentAsGlobal()  { C.LLVMDeleteGlobal(v.c) }
func (v Value) Initializer() (rv Value)   { rv.c = C.LLVMGetInitializer(v.c); return }
func (v Value) SetInitializer(cv Value)   { C.LLVMSetInitializer(v.c, cv.c) }
func (v Value) IsThreadLocal() bool       { return C.LLVMIsThreadLocal(v.c) != 0 }
func (v Value) SetThreadLocal(tl bool)    { C.LLVMSetThreadLocal(v.c, boolToLLVMBool(tl)) }
func (v Value) IsGlobalConstant() bool    { return C.LLVMIsGlobalConstant(v.c) != 0 }
func (v Value) SetGlobalConstant(gc bool) { C.LLVMSetGlobalConstant(v.c, boolToLLVMBool(gc)) }

// Operations on aliases
func AddAlias(m Module, t Type, aliasee Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMAddAlias(m.c, t.c, aliasee.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}

// Operations on functions
func AddFunction(m Module, name string, ft Type) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMAddFunction(m.c, cname, ft.c)
	C.free(unsafe.Pointer(cname))
	return
}

func (m Module) NamedFunction(name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMGetNamedFunction(m.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}

func (m Module) FirstFunction() (v Value)  { v.c = C.LLVMGetFirstFunction(m.c); return }
func (m Module) LastFunction() (v Value)   { v.c = C.LLVMGetLastFunction(m.c); return }
func NextFunction(v Value) (rv Value)      { rv.c = C.LLVMGetNextFunction(v.c); return }
func PrevFunction(v Value) (rv Value)      { rv.c = C.LLVMGetPreviousFunction(v.c); return }
func (v Value) EraseFromParentAsFunction() { C.LLVMDeleteFunction(v.c) }
func (v Value) IntrinsicID() int           { return int(C.LLVMGetIntrinsicID(v.c)) }
func (v Value) FunctionCallConv() CallConv {
	return CallConv(C.LLVMCallConv(C.LLVMGetFunctionCallConv(v.c)))
}
func (v Value) SetFunctionCallConv(cc CallConv) { C.LLVMSetFunctionCallConv(v.c, C.unsigned(cc)) }
func (v Value) GC() string                      { return C.GoString(C.LLVMGetGC(v.c)) }
func (v Value) SetGC(name string) {
	cname := C.CString(name)
	C.LLVMSetGC(v.c, cname)
	C.free(unsafe.Pointer(cname))
}
func (v Value) AddFunctionAttr(a Attribute) { C.LLVMAddFunctionAttr(v.c, C.LLVMAttribute(a)) }
//TODO
//LLVMAttribute LLVMGetFunctionAttr(LLVMValueRef Fn);
//void LLVMRemoveFunctionAttr(LLVMValueRef Fn, LLVMAttribute PA);

// Operations on parameters
func (v Value) ParamsCount() int { return int(C.LLVMCountParams(v.c)) }
func (v Value) Params() []Value {
	out := make([]Value, v.ParamsCount())
	C.LLVMGetParams(v.c, llvmValueRefPtr(&out[0]))
	return out
}
func (v Value) Param(i int) (rv Value)  { rv.c = C.LLVMGetParam(v.c, C.unsigned(i)); return }
func (v Value) ParamParent() (rv Value) { rv.c = C.LLVMGetParamParent(v.c); return }
func (v Value) FirstParam() (rv Value)  { rv.c = C.LLVMGetFirstParam(v.c); return }
func (v Value) LastParam() (rv Value)   { rv.c = C.LLVMGetLastParam(v.c); return }
func NextParam(v Value) (rv Value)      { rv.c = C.LLVMGetNextParam(v.c); return }
func PrevParam(v Value) (rv Value)      { rv.c = C.LLVMGetPreviousParam(v.c); return }

//TODO
//void LLVMAddAttribute(LLVMValueRef Arg, LLVMAttribute PA);
//void LLVMRemoveAttribute(LLVMValueRef Arg, LLVMAttribute PA);
//LLVMAttribute LLVMGetAttribute(LLVMValueRef Arg);

func (v Value) SetParamAlignment(align int) { C.LLVMSetParamAlignment(v.c, C.unsigned(align)) }

// Operations on basic blocks
func (bb BasicBlock) AsValue() (v Value)      { v.c = C.LLVMBasicBlockAsValue(bb.c); return }
func (v Value) IsBasicBlock() bool            { return C.LLVMValueIsBasicBlock(v.c) != 0 }
func (v Value) AsBasicBlock() (bb BasicBlock) { bb.c = C.LLVMValueAsBasicBlock(v.c); return }
func (bb BasicBlock) Parent() (v Value)       { v.c = C.LLVMGetBasicBlockParent(bb.c); return }
func (v Value) BasicBlocksCount() int         { return int(C.LLVMCountBasicBlocks(v.c)) }
func (v Value) BasicBlocks() []BasicBlock {
	out := make([]BasicBlock, v.BasicBlocksCount())
	C.LLVMGetBasicBlocks(v.c, llvmBasicBlockRefPtr(&out[0]))
	return out
}
func (v Value) FirstBasicBlock() (bb BasicBlock)    { bb.c = C.LLVMGetFirstBasicBlock(v.c); return }
func (v Value) LastBasicBlock() (bb BasicBlock)     { bb.c = C.LLVMGetLastBasicBlock(v.c); return }
func NextBasicBlock(bb BasicBlock) (rbb BasicBlock) { rbb.c = C.LLVMGetNextBasicBlock(bb.c); return }
func PrevBasicBlock(bb BasicBlock) (rbb BasicBlock) { rbb.c = C.LLVMGetPreviousBasicBlock(bb.c); return }
func (v Value) EntryBasicBlock() (bb BasicBlock)    { bb.c = C.LLVMGetEntryBasicBlock(v.c); return }
func AddBasicBlockInContext(c Context, f Value, name string) (bb BasicBlock) {
	cname := C.CString(name)
	bb.c = C.LLVMAppendBasicBlockInContext(c.c, f.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func InsertBasicBlockInContext(c Context, ref BasicBlock, name string) (bb BasicBlock) {
	cname := C.CString(name)
	bb.c = C.LLVMInsertBasicBlockInContext(c.c, ref.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func AddBasicBlock(f Value, name string) (bb BasicBlock) {
	cname := C.CString(name)
	bb.c = C.LLVMAppendBasicBlock(f.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func InsertBasicBlock(ref BasicBlock, name string) (bb BasicBlock) {
	cname := C.CString(name)
	bb.c = C.LLVMInsertBasicBlock(ref.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (bb BasicBlock) EraseFromParent() { C.LLVMDeleteBasicBlock(bb.c) }

//TODO
//void LLVMMoveBasicBlockBefore(LLVMBasicBlockRef BB, LLVMBasicBlockRef MovePos);
//void LLVMMoveBasicBlockAfter(LLVMBasicBlockRef BB, LLVMBasicBlockRef MovePos);

// Operations on instructions
func (v Value) InstructionParent() (bb BasicBlock) { bb.c = C.LLVMGetInstructionParent(v.c); return }
func (bb BasicBlock) FirstInstruction() (v Value)  { v.c = C.LLVMGetFirstInstruction(bb.c); return }
func (bb BasicBlock) LastInstruction() (v Value)   { v.c = C.LLVMGetLastInstruction(bb.c); return }
func NextInstruction(v Value) (rv Value)           { rv.c = C.LLVMGetNextInstruction(v.c); return }
func PrevInstruction(v Value) (rv Value)           { rv.c = C.LLVMGetPreviousInstruction(v.c); return }

// Operations on call sites
func (v Value) SetInstructionCallConv(cc CallConv) {
	C.LLVMSetInstructionCallConv(v.c, C.unsigned(cc))
}
func (v Value) InstructionCallConv() CallConv {
	return CallConv(C.LLVMCallConv(C.LLVMGetInstructionCallConv(v.c)))
}

//TODO
//void LLVMAddInstrAttribute(LLVMValueRef Instr, unsigned index, LLVMAttribute);
//void LLVMRemoveInstrAttribute(LLVMValueRef Instr, unsigned index, 
//                              LLVMAttribute);

func (v Value) SetInstrParamAlignment(i int, align int) {
	C.LLVMSetInstrParamAlignment(v.c, C.unsigned(i), C.unsigned(align))
}

// Operations on call instructions (only)
func (v Value) IsTailCall() bool    { return C.LLVMIsTailCall(v.c) != 0 }
func (v Value) SetTailCall(is bool) { C.LLVMSetTailCall(v.c, boolToLLVMBool(is)) }

// Operations on phi nodes
func (v Value) AddIncoming(vals []Value, blocks []BasicBlock) {
	C.LLVMAddIncoming(v.c, llvmValueRefPtr(&vals[0]),
		llvmBasicBlockRefPtr(&blocks[0]), C.unsigned(len(vals)))
}
func (v Value) IncomingCount() int { return int(C.LLVMCountIncoming(v.c)) }
func (v Value) IncomingValue(i int) (rv Value) {
	rv.c = C.LLVMGetIncomingValue(v.c, C.unsigned(i))
	return
}
func (v Value) IncomingBlock(i int) (bb BasicBlock) {
	bb.c = C.LLVMGetIncomingBlock(v.c, C.unsigned(i))
	return
}

//-------------------------------------------------------------------------
// llvm.Builder
//-------------------------------------------------------------------------

// An instruction builder represents a point within a basic block, and is the
// exclusive means of building instructions using the C interface.

func NewBuilderInContext(c Context) (b Builder) { b.c = C.LLVMCreateBuilderInContext(c.c); return }
func NewBuilder() (b Builder)                   { b.c = C.LLVMCreateBuilder(); return }
func (b Builder) SetInsertPoint(block BasicBlock, instr Value) {
	C.LLVMPositionBuilder(b.c, block.c, instr.c)
}
func (b Builder) SetInsertPointBefore(instr Value)     { C.LLVMPositionBuilderBefore(b.c, instr.c) }
func (b Builder) SetInsertPointAtEnd(block BasicBlock) { C.LLVMPositionBuilderBefore(b.c, block.c) }
func (b Builder) GetInsertBlock() (bb BasicBlock)      { bb.c = C.LLVMGetInsertBlock(b.c); return }
func (b Builder) ClearInsertionPoint()                 { C.LLVMClearInsertionPosition(b.c) }
func (b Builder) Insert(instr Value)                   { C.LLVMInsertIntoBuilder(b.c, instr.c) }
func (b Builder) InsertWithName(instr Value, name string) {
	cname := C.CString(name)
	C.LLVMInsertIntoBuilderWithName(b.c, instr.c, cname)
	C.free(unsafe.Pointer(cname))
}
func (b Builder) Dispose() { C.LLVMDisposeBuilder(b.c) }

// Metadata
func (b Builder) SetCurrentDebugLocation(v Value) { C.LLVMSetCurrentDebugLocation(b.c, v.c) }
func (b Builder) CurrentDebugLocation() (v Value) { v.c = C.LLVMGetCurrentDebugLocation(b.c); return }
func (b Builder) SetInstDebugLocation(v Value)    { C.LLVMSetCurrentDebugLocation(b.c, v.c) }

// Terminators
func (b Builder) CreateRetVoid() (rv Value)    { rv.c = C.LLVMBuildRetVoid(b.c); return }
func (b Builder) CreateRet(v Value) (rv Value) { rv.c = C.LLVMBuildRet(b.c, v.c); return }
func (b Builder) CreateAggregateRet(vs []Value) (rv Value) {
	rv.c = C.LLVMBuildAggregateRet(b.c, llvmValueRefPtr(&vs[0]), C.unsigned(len(vs)))
	return
}
func (b Builder) CreateBr(bb BasicBlock) (rv Value) { rv.c = C.LLVMBuildBr(b.c, bb.c); return }
func (b Builder) CreateCondBr(ifv Value, thenb, elseb BasicBlock) (rv Value) {
	rv.c = C.LLVMBuildCondBr(b.c, ifv.c, thenb.c, elseb.c)
	return
}
func (b Builder) CreateSwitch(v Value, elseb BasicBlock, numCases int) (rv Value) {
	rv.c = C.LLVMBuildSwitch(b.c, v.c, elseb.c, C.unsigned(numCases))
	return
}
func (b Builder) CreateIndirectBr(addr Value, numDests int) (rv Value) {
	rv.c = C.LLVMBuildIndirectBr(b.c, addr.c, C.unsigned(numDests))
	return
}
func (b Builder) CreateInvoke(fn Value, args []Value, then, catch BasicBlock, name string) (rv Value) {
	cname := C.CString(name)
	rv.c = C.LLVMBuildInvoke(b.c, fn.c, llvmValueRefPtr(&args[0]),
		C.unsigned(len(args)), then.c, catch.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateUnwind() (rv Value)      { rv.c = C.LLVMBuildUnwind(b.c); return }
func (b Builder) CreateUnreachable() (rv Value) { rv.c = C.LLVMBuildUnreachable(b.c); return }

// Add a case to the switch instruction
func (v Value) AddCase(on Value, dest BasicBlock) { C.LLVMAddCase(v.c, on.c, dest.c) }

// Add a destination to the indirectbr instruction
func (v Value) AddDest(dest BasicBlock) { C.LLVMAddDestination(v.c, dest.c) }

// Arithmetic
func (b Builder) CreateAdd(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildAdd(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateNSWAdd(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildNSWAdd(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateNUWAdd(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildNUWAdd(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFAdd(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildFAdd(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateSub(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildSub(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateNSWSub(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildNSWSub(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateNUWSub(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildNUWSub(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFSub(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildFSub(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateMul(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildMul(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateNSWMul(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildNSWMul(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateNUWMul(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildNUWMul(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFMul(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildFMul(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateUDiv(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildUDiv(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateSDiv(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildSDiv(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateExactSDiv(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildExactSDiv(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFDiv(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildFDiv(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateURem(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildURem(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateSRem(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildSRem(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFRem(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildFRem(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateShl(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildShl(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateLShr(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildLShr(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateAShr(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildAShr(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateAnd(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildAnd(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateOr(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildOr(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateXor(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildXor(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateBinOp(op Opcode, lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildBinOp(b.c, C.LLVMOpcode(op), lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateNeg(v Value, name string) (rv Value) {
	cname := C.CString(name)
	rv.c = C.LLVMBuildNeg(b.c, v.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateNSWNeg(v Value, name string) (rv Value) {
	cname := C.CString(name)
	rv.c = C.LLVMBuildNSWNeg(b.c, v.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateNUWNeg(v Value, name string) (rv Value) {
	cname := C.CString(name)
	rv.c = C.LLVMBuildNUWNeg(b.c, v.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFNeg(v Value, name string) (rv Value) {
	cname := C.CString(name)
	rv.c = C.LLVMBuildFNeg(b.c, v.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateNot(v Value, name string) (rv Value) {
	cname := C.CString(name)
	rv.c = C.LLVMBuildNot(b.c, v.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}

// Memory

func (b Builder) CreateMalloc(t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildMalloc(b.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateArrayMalloc(t Type, val Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildArrayMalloc(b.c, t.c, val.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateAlloca(t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildAlloca(b.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateArrayAlloca(t Type, val Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildArrayAlloca(b.c, t.c, val.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFree(p Value) (v Value) {
	v.c = C.LLVMBuildFree(b.c, p.c)
	return
}
func (b Builder) CreateLoad(p Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildLoad(b.c, p.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateStore(val Value, p Value) (v Value) {
	v.c = C.LLVMBuildStore(b.c, val.c, p.c)
	return
}
func (b Builder) CreateGEP(p Value, indices []Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildGEP(b.c, p.c,
		llvmValueRefPtr(&indices[0]), C.unsigned(len(indices)), cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateInBoundsGEP(p Value, indices []Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildInBoundsGEP(b.c, p.c,
		llvmValueRefPtr(&indices[0]), C.unsigned(len(indices)), cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateStructGEP(p Value, i int, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildStructGEP(b.c, p.c, C.unsigned(i), cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateGlobalString(str, name string) (v Value) {
	cstr := C.CString(str)
	cname := C.CString(name)
	v.c = C.LLVMBuildGlobalString(b.c, cstr, cname)
	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(cstr))
	return
}
func (b Builder) CreateGlobalStringPtr(str, name string) (v Value) {
	cstr := C.CString(str)
	cname := C.CString(name)
	v.c = C.LLVMBuildGlobalStringPtr(b.c, cstr, cname)
	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(cstr))
	return
}

// Casts
func (b Builder) CreateTrunc(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildTrunc(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateZExt(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildZExt(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateSExt(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildSExt(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFPToUI(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildFPToUI(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFPToSI(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildFPToSI(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateUIToFP(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildUIToFP(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateSIToFP(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildSIToFP(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFPTrunc(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildFPTrunc(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFPExt(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildFPExt(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreatePtrToInt(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildPtrToInt(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateIntToPtr(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildIntToPtr(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateBitCast(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildBitCast(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateZExtOrBitCast(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildZExtOrBitCast(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateSExtOrBitCast(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildSExtOrBitCast(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateTruncOrBitCast(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildTruncOrBitCast(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateCast(val Value, op Opcode, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildCast(b.c, C.LLVMOpcode(op), val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
} //
func (b Builder) CreatePointerCast(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildPointerCast(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateIntCast(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildIntCast(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFPCast(val Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildFPCast(b.c, val.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}

// Comparisons
func (b Builder) CreateICmp(pred IntPredicate, lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildICmp(b.c, C.LLVMIntPredicate(pred), lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateFCmp(pred FloatPredicate, lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildFCmp(b.c, C.LLVMRealPredicate(pred), lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}

// Miscellaneous instructions
func (b Builder) CreatePHI(t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildPhi(b.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateCall(fn Value, args []Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildCall(b.c, fn.c,
		llvmValueRefPtr(&args[0]), C.unsigned(len(args)), cname)
	C.free(unsafe.Pointer(cname))
	return
}

func (b Builder) CreateSelect(ifv, thenv, elsev Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildSelect(b.c, ifv.c, thenv.c, elsev.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}

func (b Builder) CreateVAArg(list Value, t Type, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildVAArg(b.c, list.c, t.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateExtractElement(vec, i Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildExtractElement(b.c, vec.c, i.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateInsertElement(vec, elt, i Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildInsertElement(b.c, vec.c, elt.c, i.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateShuffleVector(v1, v2, mask Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildShuffleVector(b.c, v1.c, v2.c, mask.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateExtractValue(agg Value, i int, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildExtractValue(b.c, agg.c, C.unsigned(i), cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateInsertValue(agg, elt Value, i int, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildInsertValue(b.c, agg.c, elt.c, C.unsigned(i), cname)
	C.free(unsafe.Pointer(cname))
	return
}

func (b Builder) CreateIsNull(val Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildIsNull(b.c, val.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreateIsNotNull(val Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildIsNotNull(b.c, val.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}
func (b Builder) CreatePtrDiff(lhs, rhs Value, name string) (v Value) {
	cname := C.CString(name)
	v.c = C.LLVMBuildPtrDiff(b.c, lhs.c, rhs.c, cname)
	C.free(unsafe.Pointer(cname))
	return
}

//-------------------------------------------------------------------------
// llvm.ModuleProvider
//-------------------------------------------------------------------------

// Changes the type of M so it can be passed to FunctionPassManagers and the
// JIT. They take ModuleProviders for historical reasons.
func NewModuleProviderForModule(m Module) (mp ModuleProvider) {
	mp.c = C.LLVMCreateModuleProviderForExistingModule(m.c)
	return
}


// Destroys the module M.
func (mp ModuleProvider) Dispose() { C.LLVMDisposeModuleProvider(mp.c) }

//-------------------------------------------------------------------------
// llvm.MemoryBuffer
//-------------------------------------------------------------------------

//TODO
//LLVMBool LLVMCreateMemoryBufferWithContentsOfFile(const char *Path,
//                                                  LLVMMemoryBufferRef *OutMemBuf,
//                                                  char **OutMessage);
//LLVMBool LLVMCreateMemoryBufferWithSTDIN(LLVMMemoryBufferRef *OutMemBuf,
//                                         char **OutMessage);
//void LLVMDisposeMemoryBuffer(LLVMMemoryBufferRef MemBuf);

//-------------------------------------------------------------------------
// llvm.PassManager
//-------------------------------------------------------------------------

// Constructs a new whole-module pass pipeline. This type of pipeline is
// suitable for link-time optimization and whole-module transformations.
// See llvm::PassManager::PassManager.
func NewPassManager() (pm PassManager) { pm.c = C.LLVMCreatePassManager(); return }

// Constructs a new function-by-function pass pipeline over the module
// provider. It does not take ownership of the module provider. This type of
// pipeline is suitable for code generation and JIT compilation tasks.
// See llvm::FunctionPassManager::FunctionPassManager.
func NewFunctionPassManagerForModule(m Module) (pm PassManager) {
	pm.c = C.LLVMCreateFunctionPassManagerForModule(m.c)
	return
}

// Deprecated: Use LLVMCreateFunctionPassManagerForModule instead.
//LLVMPassManagerRef LLVMCreateFunctionPassManager(LLVMModuleProviderRef MP);
//XXX: don't port this

// Initializes, executes on the provided module, and finalizes all of the
// passes scheduled in the pass manager. Returns 1 if any of the passes
// modified the module, 0 otherwise. See llvm::PassManager::run(Module&).
func (pm PassManager) Run(m Module) bool { return C.LLVMRunPassManager(pm.c, m.c) != 0 }

// Initializes all of the function passes scheduled in the function pass
// manager. Returns 1 if any of the passes modified the module, 0 otherwise.
// See llvm::FunctionPassManager::doInitialization.
func (pm PassManager) InitializeFunc() bool { return C.LLVMInitializeFunctionPassManager(pm.c) != 0 }

// Executes all of the function passes scheduled in the function pass manager
// on the provided function. Returns 1 if any of the passes modified the
// function, false otherwise.
// See llvm::FunctionPassManager::run(Function&).
func (pm PassManager) RunFunc(f Value) bool { return C.LLVMRunFunctionPassManager(pm.c, f.c) != 0 }

// Finalizes all of the function passes scheduled in in the function pass
// manager. Returns 1 if any of the passes modified the module, 0 otherwise.
// See llvm::FunctionPassManager::doFinalization.
func (pm PassManager) FinalizeFunc() bool { return C.LLVMFinalizeFunctionPassManager(pm.c) != 0 }

// Frees the memory of a pass pipeline. For function pipelines, does not free
// the module provider.
// See llvm::PassManagerBase::~PassManagerBase.
func (pm PassManager) Dispose() { C.LLVMDisposePassManager(pm.c) }
