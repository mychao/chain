package gorocksdb

// #cgo CXXFLAGS: -std=c++11
// #cgo CPPFLAGS: -I${SRCDIR}/../../facebook/rocksdb/include
// #cgo CFLAGS: -I${SRCDIR}/../../facebook/rocksdb/include
// #cgo LDFLAGS: -I${SRCDIR}/../../facebook/rocksdb
// #cgo LDFLAGS: ${SRCDIR}/../../facebook/rocksdb/librocksdb.a
// #cgo LDFLAGS: -lstdc++
// #cgo LDFLAGS: -lm
// #cgo LDFLAGS: -lz
// #cgo LDFLAGS: -lbz2
// #cgo LDFLAGS: -lsnappy
// #cgo darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
// #cgo !darwin LDFLAGS: -Wl,-unresolved-symbols=ignore-all -lrt
import "C"
