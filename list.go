package alpm

// #include <alpm.h>
import "C"

import (
	"unsafe"

	"github.com/jguer/go-alpm/alpm_list"
)

type StringList struct{ *alpm_list.List }
type BackupList struct{ *alpm_list.List }
type DependList struct{ *alpm_list.List }
type PackageList struct {
	*alpm_list.List
	handle Handle
}
type DBList struct {
	*alpm_list.List
	handle Handle
}

func makeList(l *C.alpm_list_t) *alpm_list.List {
	return (*alpm_list.List)(unsafe.Pointer(l))
}

func makeStringList(l *C.alpm_list_t) StringList {
	return StringList{(*alpm_list.List)((unsafe.Pointer(l)))}
}

func (l StringList) ForEach(f func(str string) error) error {
	return l.List.ForEach(func(ptr uintptr) error {
		return f(C.GoString((*C.char)(unsafe.Pointer(ptr))))
	})
}

func (l StringList) Slice() []string {
	strs := make([]string, 0, l.Count())
	for i := l.List; i != nil; i = i.Next() {
		strs = append(strs, i.String())
	}
	return strs
}

func ToStringList(strs []string) StringList {
	var list *alpm_list.List
	for _, str := range strs {
		alpm_list.AppendStrdup(&list, str)
	}
	return StringList{list}
}

func makeBackupList(l *C.alpm_list_t) BackupList {
	return BackupList{(*alpm_list.List)((unsafe.Pointer(l)))}
}

func (l BackupList) ForEach(f func(*Backup) error) error {
	return l.List.ForEach(func(p uintptr) error {
		b := (*Backup)(unsafe.Pointer(p))
		return f(b)
	})
}

func (l BackupList) Slice() (slice []Backup) {
	l.ForEach(func(f *Backup) error {
		slice = append(slice, *f)
		return nil
	})
	return
}

func makeDependList(l *C.alpm_list_t) DependList {
	return DependList{(*alpm_list.List)((unsafe.Pointer(l)))}
}

// ForEach executes an action on each package of the DependList.
func (l DependList) ForEach(f func(*Depend) error) error {
	return l.List.ForEach(func(p uintptr) error {
		dep := (*Depend)(unsafe.Pointer(p))
		return f(dep)
	})
}

// Slice converts the DependList to a Depend Slice.
func (l DependList) Slice() []*Depend {
	slice := []*Depend{}
	l.ForEach(func(dep *Depend) error {
		slice = append(slice, dep)
		return nil
	})
	return slice
}

func makePackageList(l *C.alpm_list_t, h Handle) PackageList {
	return PackageList{(*alpm_list.List)((unsafe.Pointer(l))), h}
}

// ForEach executes an action on each package of the PackageList.
func (l PackageList) ForEach(f func(*Package) error) error {
	return l.List.ForEach(func(ptr uintptr) error {
		p := Package{(*C.alpm_pkg_t)(unsafe.Pointer(ptr)), l.handle}
		return f(&p)
	})
}

// Slice converts the PackageList to a Package Slice.
func (l PackageList) Slice() []Package {
	slice := []Package{}
	l.ForEach(func(p *Package) error {
		slice = append(slice, *p)
		return nil
	})
	return slice
}

func makeDBList(l *C.alpm_list_t, h Handle) DBList {
	return DBList{(*alpm_list.List)((unsafe.Pointer(l))), h}
}

func (l DBList) ForEach(f func(db *DB) error) error {
	return l.List.ForEach(func(ptr uintptr) error {
		db := DB{(*C.alpm_db_t)(unsafe.Pointer(ptr)), l.handle}
		return f(&db)
	})
}

func (l DBList) Slice() []DB {
	slice := []DB{}
	l.ForEach(func(db *DB) error {
		slice = append(slice, *db)
		return nil
	})
	return slice
}
