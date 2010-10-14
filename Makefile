include $(GOROOT)/src/Make.inc

TARG=llvm
CGOFILES=llvm.go

CGO_CFLAGS=`llvm-config --cflags`
CGO_LDFLAGS=`llvm-config --ldflags --libs all`

include $(GOROOT)/src/Make.pkg
