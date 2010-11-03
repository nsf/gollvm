package llvm

/*
#include <llvm-c/BitWriter.h>
#include <stdlib.h>
*/
import "C"
import "os"
import "unsafe"

var writeBitcodeToFileErr = os.NewError("Failed to write bitcode to file")

func WriteBitcodeToFile(m Module, filename string) os.Error {
	cfilename := C.CString(filename)
	result := C.LLVMWriteBitcodeToFile(m.C, cfilename)
	C.free(unsafe.Pointer(cfilename))
	if result == 0 {
		return nil
	}
	return writeBitcodeToFileErr
}

// TODO(nsf): Figure out way how to make it work with io.Writer
