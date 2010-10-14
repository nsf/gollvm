package llvm

/*
#include <llvm-c/Transforms/Scalar.h>
#include <stdlib.h>
*/
import "C"

func (pm PassManager) AddAggressiveDCEPass() { C.LLVMAddAggressiveDCEPass(pm.C) }
func (pm PassManager) AddCFGSimplificationPass() { C.LLVMAddCFGSimplificationPass(pm.C) }
func (pm PassManager) AddDeadStoreEliminationPass() { C.LLVMAddDeadStoreEliminationPass(pm.C) }
func (pm PassManager) AddGVNPass() { C.LLVMAddGVNPass(pm.C) }
func (pm PassManager) AddIndVarSimplifyPass() { C.LLVMAddIndVarSimplifyPass(pm.C) }
func (pm PassManager) AddInstructionCombiningPass() { C.LLVMAddInstructionCombiningPass(pm.C) }
func (pm PassManager) AddJumpThreadingPass() { C.LLVMAddJumpThreadingPass(pm.C) }
func (pm PassManager) AddLICMPass() { C.LLVMAddLICMPass(pm.C) }
func (pm PassManager) AddLoopDeletionPass() { C.LLVMAddLoopDeletionPass(pm.C) }
func (pm PassManager) AddLoopIndexSplitPass() { C.LLVMAddLoopIndexSplitPass(pm.C) }
func (pm PassManager) AddLoopRotatePass() { C.LLVMAddLoopRotatePass(pm.C) }
func (pm PassManager) AddLoopUnrollPass() { C.LLVMAddLoopUnrollPass(pm.C) }
func (pm PassManager) AddLoopUnswitchPass() { C.LLVMAddLoopUnswitchPass(pm.C) }
func (pm PassManager) AddMemCpyOptPass() { C.LLVMAddMemCpyOptPass(pm.C) }
func (pm PassManager) AddPromoteMemoryToRegisterPass() { C.LLVMAddPromoteMemoryToRegisterPass(pm.C) }
func (pm PassManager) AddReassociatePass() { C.LLVMAddReassociatePass(pm.C) }
func (pm PassManager) AddSCCPPass() { C.LLVMAddSCCPPass(pm.C) }
func (pm PassManager) AddScalarReplAggregatesPass() { C.LLVMAddScalarReplAggregatesPass(pm.C) }
func (pm PassManager) AddScalarReplAggregatesPassWithThreshold(threshold int) { C.LLVMAddScalarReplAggregatesPassWithThreshold(pm.C, C.int(threshold)) }
func (pm PassManager) AddSimplifyLibCallsPass() { C.LLVMAddSimplifyLibCallsPass(pm.C) }
func (pm PassManager) AddTailCallEliminationPass() { C.LLVMAddTailCallEliminationPass(pm.C) }
func (pm PassManager) AddConstantPropagationPass() { C.LLVMAddConstantPropagationPass(pm.C) }
func (pm PassManager) AddDemoteMemoryToRegisterPass() { C.LLVMAddDemoteMemoryToRegisterPass(pm.C) }
func (pm PassManager) AddVerifierPass() { C.LLVMAddVerifierPass(pm.C) }
/*

// See llvm::createAggressiveDCEPass function.
void LLVMAddAggressiveDCEPass(LLVMPassManagerRef PM);

// See llvm::createCFGSimplificationPass function.
void LLVMAddCFGSimplificationPass(LLVMPassManagerRef PM);

// See llvm::createDeadStoreEliminationPass function.
void LLVMAddDeadStoreEliminationPass(LLVMPassManagerRef PM);

// See llvm::createGVNPass function.
void LLVMAddGVNPass(LLVMPassManagerRef PM);

// See llvm::createIndVarSimplifyPass function.
void LLVMAddIndVarSimplifyPass(LLVMPassManagerRef PM);

// See llvm::createInstructionCombiningPass function.
void LLVMAddInstructionCombiningPass(LLVMPassManagerRef PM);

// See llvm::createJumpThreadingPass function.
void LLVMAddJumpThreadingPass(LLVMPassManagerRef PM);

// See llvm::createLICMPass function.
void LLVMAddLICMPass(LLVMPassManagerRef PM);

// See llvm::createLoopDeletionPass function.
void LLVMAddLoopDeletionPass(LLVMPassManagerRef PM);

// See llvm::createLoopIndexSplitPass function.
void LLVMAddLoopIndexSplitPass(LLVMPassManagerRef PM);

// See llvm::createLoopRotatePass function.
void LLVMAddLoopRotatePass(LLVMPassManagerRef PM);

// See llvm::createLoopUnrollPass function.
void LLVMAddLoopUnrollPass(LLVMPassManagerRef PM);

// See llvm::createLoopUnswitchPass function.
void LLVMAddLoopUnswitchPass(LLVMPassManagerRef PM);

// See llvm::createMemCpyOptPass function.
void LLVMAddMemCpyOptPass(LLVMPassManagerRef PM);

// See llvm::createPromoteMemoryToRegisterPass function.
void LLVMAddPromoteMemoryToRegisterPass(LLVMPassManagerRef PM);

// See llvm::createReassociatePass function.
void LLVMAddReassociatePass(LLVMPassManagerRef PM);

// See llvm::createSCCPPass function.
void LLVMAddSCCPPass(LLVMPassManagerRef PM);

// See llvm::createScalarReplAggregatesPass function.
void LLVMAddScalarReplAggregatesPass(LLVMPassManagerRef PM);

// See llvm::createScalarReplAggregatesPass function.
void LLVMAddScalarReplAggregatesPassWithThreshold(LLVMPassManagerRef PM,
                                                  int Threshold);

// See llvm::createSimplifyLibCallsPass function.
void LLVMAddSimplifyLibCallsPass(LLVMPassManagerRef PM);

// See llvm::createTailCallEliminationPass function.
void LLVMAddTailCallEliminationPass(LLVMPassManagerRef PM);

// See llvm::createConstantPropagationPass function.
void LLVMAddConstantPropagationPass(LLVMPassManagerRef PM);

// See llvm::demotePromoteMemoryToRegisterPass function.
void LLVMAddDemoteMemoryToRegisterPass(LLVMPassManagerRef PM);

// See llvm::createVerifierPass function.
void LLVMAddVerifierPass(LLVMPassManagerRef PM);

*/
