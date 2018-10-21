package alpm

/*
#include <alpm.h>
*/
import "C"
import (
	"unsafe"
)

func (dep *Depend) Free() {
	C.alpm_dep_free((*C.alpm_depend_t)(dep))
}

// FindSatisfier searches a DBList for a package that satisfies depstring
// Example "glibc>=2.12"
func (l DBList) FindSatisfier(depstring string) (*Package, error) {
	cDepString := C.CString(depstring)
	defer C.free(unsafe.Pointer(cDepString))

	pkgList := (*C.alpm_list_t)(unsafe.Pointer(l.List))
	pkgHandle := (*C.alpm_handle_t)(unsafe.Pointer(l.handle.ptr))

	ptr := C.alpm_find_dbs_satisfier(pkgHandle, pkgList, cDepString)
	if ptr == nil {
		return nil, l.handle.LastError()
	}

	return &Package{ptr, l.handle}, nil
}

func DepFromString(str string) *Depend {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return (*Depend)(C.alpm_dep_from_string(cstr))
}

// FindSatisfier finds a package that satisfies depstring from PkgList
func (l PackageList) FindSatisfier(depstring string) *Package {
	cDepString := C.CString(depstring)
	defer C.free(unsafe.Pointer(cDepString))

	pkgList := (*C.alpm_list_t)(unsafe.Pointer(l.List))

	ptr := C.alpm_find_satisfier(pkgList, cDepString)
	if ptr == nil {
		return nil
	}

	return &Package{ptr, l.handle}
}

func (dep *Depend) String() string {
	str := C.alpm_dep_compute_string((*C.alpm_depend_t)(dep))
	defer C.free(unsafe.Pointer(str))
	return C.GoString(str)
}
