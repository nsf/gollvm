include $(GOROOT)/src/Make.inc

TARG=llvm
CGOFILES=core.go executionengine.go target.go

CGO_CFLAGS=`llvm-config --cflags`
CGO_LDFLAGS=-lstdc++ `llvm-config --ldflags --libs all`

include $(GOROOT)/src/Make.pkg
