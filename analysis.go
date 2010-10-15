package llvm

/*
#include <llvm-c/Analysis.h>
#include <stdlib.h>
*/
import "C"
import "os"

type VerifierFailureAction C.LLVMVerifierFailureAction

const (
	// verifier will print to stderr and abort()
	AbortProcessAction = C.LLVMAbortProcessAction
	// verifier will print to stderr and return 1
	PrintMessageAction = C.LLVMPrintMessageAction
	// verifier will just return 1
	ReturnStatusAction = C.LLVMReturnStatusAction
)

// Verifies that a module is valid, taking the specified action if not.
// Optionally returns a human-readable description of any invalid constructs.
func VerifyModule(m Module, a VerifierFailureAction) os.Error {
	var cmsg *C.char
	broken := C.LLVMVerifyModule(m.C, C.LLVMVerifierFailureAction(a), &cmsg)

	// C++'s verifyModule means isModuleBroken, so it returns false if
	// there are no errors
	if broken == 0 {
		return nil
	}
	err := os.NewError(C.GoString(cmsg))
	C.LLVMDisposeMessage(cmsg)
	return err
}

var verifyFunctionError = os.NewError("Function is broken")

// Verifies that a single function is valid, taking the specified action.
// Useful for debugging.
func VerifyFunction(f Value, a VerifierFailureAction) os.Error {
	broken := C.LLVMVerifyFunction(f.C, C.LLVMVerifierFailureAction(a))

	// C++'s verifyFunction means isFunctionBroken, so it returns false if
	// there are no errors
	if broken == 0 {
		return nil
	}
	return verifyFunctionError
}

// Open up a ghostview window that displays the CFG of the current function.
// Useful for debugging.
func ViewFunctionCFG(f Value)     { C.LLVMViewFunctionCFG(f.C) }
func ViewFunctionCFGOnly(f Value) { C.LLVMViewFunctionCFGOnly(f.C) }
