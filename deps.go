package alpm

/*
#include <alpm.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

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
