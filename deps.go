package alpm

/*
#include <alpm.h>
*/
import "C"
import (
	"fmt"
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
		return nil,
			fmt.Errorf("unable to satisfy dependency %s in DBlist", depstring)
	}

	return &Package{ptr, l.handle}, nil
}

func DepFromString(str string) *Depend {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return (*Depend)(C.alpm_dep_from_string(cstr))
}

// FindSatisfier finds a package that satisfies depstring from PkgList
func (l PackageList) FindSatisfier(depstring string) (*Package, error) {
	cDepString := C.CString(depstring)
	defer C.free(unsafe.Pointer(cDepString))

	pkgList := (*C.alpm_list_t)(unsafe.Pointer(l.List))

	ptr := C.alpm_find_satisfier(pkgList, cDepString)
	if ptr == nil {
		return nil,
			fmt.Errorf("unable to find dependency %s in PackageList", depstring)
	}

	return &Package{ptr, l.handle}, nil
}

func (dep *Depend) String() string {
	str := C.alpm_dep_compute_string((*C.alpm_depend_t)(dep))
	return C.GoString(str)
}
