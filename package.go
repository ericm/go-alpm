// package.go - libalpm package type and methods.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package alpm

/*
#include <alpm.h>

int pkg_cmp(const void *v1, const void *v2)
{
    alpm_pkg_t *p1 = (alpm_pkg_t *)v1;
    alpm_pkg_t *p2 = (alpm_pkg_t *)v2;
    off_t s1 = alpm_pkg_get_isize(p1);
    off_t s2 = alpm_pkg_get_isize(p2);

    if (s1 > s2)
        return -1;
    else if (s1 < s2)
        return 1;
    else
        return 0;
}
*/
import "C"

import (
	"time"
	"unsafe"

	"github.com/jguer/go-alpm/alpm_list"
)

// Package describes a single package and associated handle.
type Package struct {
	pmpkg  *C.alpm_pkg_t
	handle Handle
}

// SortBySize returns a PackageList sorted by size.
func (l PackageList) SortBySize() PackageList {
	pkgList := (*C.struct___alpm_list_t)(unsafe.Pointer(l.List))

	pkgCache := (*alpm_list.List)(unsafe.Pointer(
		C.alpm_list_msort(pkgList,
			C.alpm_list_count(pkgList),
			C.alpm_list_fn_cmp(C.pkg_cmp))))

	return PackageList{pkgCache, l.handle}
}

func (pkg *Package) FileName() string {
	return C.GoString(C.alpm_pkg_get_filename(pkg.pmpkg))
}

func (pkg *Package) Base() string {
	return C.GoString(C.alpm_pkg_get_base(pkg.pmpkg))
}

func (pkg *Package) Base64Signature() string {
	return C.GoString(C.alpm_pkg_get_base64_sig(pkg.pmpkg))
}

func (pkg *Package) Validation() Validation {
	return Validation(C.alpm_pkg_get_validation(pkg.pmpkg))
}

// Architecture returns the package target Architecture.
func (pkg *Package) Architecture() string {
	return C.GoString(C.alpm_pkg_get_arch(pkg.pmpkg))
}

func (pkg *Package) Deltas() StringList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_deltas(pkg.pmpkg))
	return StringList{(*alpm_list.List)(ptr)}
}

// Backup returns a list of package backups.
func (pkg *Package) Backup() BackupList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_backup(pkg.pmpkg))
	return BackupList{(*alpm_list.List)(ptr)}
}

// BuildDate returns the BuildDate of the package.
func (pkg *Package) BuildDate() time.Time {
	t := C.alpm_pkg_get_builddate(pkg.pmpkg)
	return time.Unix(int64(t), 0)
}

// Conflicts returns the conflicts of the package as a DependList.
func (pkg *Package) Conflicts() DependList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_conflicts(pkg.pmpkg))
	return DependList{(*alpm_list.List)(ptr)}
}

// DB returns the package's origin database.
func (pkg *Package) DB() *DB {
	ptr := C.alpm_pkg_get_db(pkg.pmpkg)
	if ptr == nil {
		return nil
	}
	return &DB{ptr, pkg.handle}
}

// Depends returns the package's dependency list.
func (pkg *Package) Depends() DependList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_depends(pkg.pmpkg))
	return DependList{(*alpm_list.List)(ptr)}
}

// Depends returns the package's optional dependency list.
func (pkg *Package) OptionalDepends() DependList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_optdepends(pkg.pmpkg))
	return DependList{(*alpm_list.List)(ptr)}
}

// Depends returns the package's check dependency list.
func (pkg *Package) CheckDepends() DependList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_checkdepends(pkg.pmpkg))
	return DependList{(*alpm_list.List)(ptr)}
}

// Depends returns the package's make dependency list.
func (pkg *Package) MakeDepends() DependList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_makedepends(pkg.pmpkg))
	return DependList{(*alpm_list.List)(ptr)}
}

// Description returns the package's description.
func (pkg *Package) Description() string {
	return C.GoString(C.alpm_pkg_get_desc(pkg.pmpkg))
}

// Files returns the file list of the package.
func (pkg *Package) Files() *FileList {
	return (*FileList)(C.alpm_pkg_get_files(pkg.pmpkg))
}

// Groups returns the groups the package belongs to.
func (pkg *Package) Groups() StringList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_groups(pkg.pmpkg))
	return StringList{(*alpm_list.List)(ptr)}
}

// ISize returns the package installed size.
func (pkg *Package) ISize() int64 {
	t := C.alpm_pkg_get_isize(pkg.pmpkg)
	return int64(t)
}

// InstallDate returns the package install date.
func (pkg *Package) InstallDate() time.Time {
	t := C.alpm_pkg_get_installdate(pkg.pmpkg)
	return time.Unix(int64(t), 0)
}

// Licenses returns the package license list.
func (pkg *Package) Licenses() StringList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_licenses(pkg.pmpkg))
	return StringList{(*alpm_list.List)(ptr)}
}

// SHA256Sum returns package SHA256Sum.
func (pkg *Package) SHA256Sum() string {
	return C.GoString(C.alpm_pkg_get_sha256sum(pkg.pmpkg))
}

// MD5Sum returns package MD5Sum.
func (pkg *Package) MD5Sum() string {
	return C.GoString(C.alpm_pkg_get_md5sum(pkg.pmpkg))
}

// Name returns package name.
func (pkg *Package) Name() string {
	return C.GoString(C.alpm_pkg_get_name(pkg.pmpkg))
}

// Packager returns package packager name.
func (pkg *Package) Packager() string {
	return C.GoString(C.alpm_pkg_get_packager(pkg.pmpkg))
}

// Provides returns DependList of packages provides by package.
func (pkg *Package) Provides() DependList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_provides(pkg.pmpkg))
	return DependList{(*alpm_list.List)(ptr)}
}

// Reason returns package install reason.
func (pkg *Package) Reason() PkgReason {
	reason := C.alpm_pkg_get_reason(pkg.pmpkg)
	return PkgReason(reason)
}

// Origin returns package origin.
func (pkg *Package) Origin() PkgFrom {
	origin := C.alpm_pkg_get_origin(pkg.pmpkg)
	return PkgFrom(origin)
}

// Replaces returns a DependList with the packages this package replaces.
func (pkg *Package) Replaces() DependList {
	ptr := unsafe.Pointer(C.alpm_pkg_get_replaces(pkg.pmpkg))
	return DependList{(*alpm_list.List)(ptr)}
}

// Size returns the packed package size.
func (pkg *Package) Size() int64 {
	t := C.alpm_pkg_get_size(pkg.pmpkg)
	return int64(t)
}

// URL returns the upstream URL of the package.
func (pkg *Package) URL() string {
	return C.GoString(C.alpm_pkg_get_url(pkg.pmpkg))
}

// Version returns the package version.
func (pkg *Package) Version() string {
	return C.GoString(C.alpm_pkg_get_version(pkg.pmpkg))
}

// ComputeRequiredBy returns the names of reverse dependencies of a package
func (pkg *Package) ComputeRequiredBy() []string {
	result := C.alpm_pkg_compute_requiredby(pkg.pmpkg)
	requiredby := make([]string, 0)
	for i := (*alpm_list.List)(unsafe.Pointer(result)); i != nil; i = i.Next() {
		defer C.free(unsafe.Pointer(i))
		data := unsafe.Pointer(i.Data())
		if data != nil {
			defer C.free(data)
			name := C.GoString((*C.char)(data))
			requiredby = append(requiredby, name)
		}
	}
	return requiredby
}

// ComputeOptionalFor returns the names of packages that optionally require the given package
func (pkg *Package) ComputeOptionalFor() []string {
	result := C.alpm_pkg_compute_optionalfor(pkg.pmpkg)
	optionalfor := make([]string, 0)
	for i := (*alpm_list.List)(unsafe.Pointer(result)); i != nil; i = i.Next() {
		defer C.free(unsafe.Pointer(i))
		data := unsafe.Pointer(i.Data())
		if data != nil {
			defer C.free(data)
			name := C.GoString((*C.char)(data))
			optionalfor = append(optionalfor, name)
		}
	}
	return optionalfor
}

func (pkg *Package) ShouldIgnore() bool {
	result := C.alpm_pkg_should_ignore(pkg.handle.ptr, pkg.pmpkg)
	return result == 1
}
