package alpm_list

/*
#cgo LDFLAGS: -lalpm
#include <alpm_list.h>
*/
import "C"

import "unsafe"

type List C.alpm_list_t

func (l *List) Data() uintptr {
	return uintptr(l.data)
}

/* item mutators */
func (l *List) Add(data uintptr) *List {
	return (*List)(C.alpm_list_add((*C.alpm_list_t)(l), unsafe.Pointer(data)))
}

func Append(l **List, data uintptr) *List {
	return (*List)(C.alpm_list_append((**C.alpm_list_t)((unsafe.Pointer(l))), unsafe.Pointer(data)))
}

func AppendStrdup(l **List, data string) *List {
	str := C.CString(data)
	return (*List)(C.alpm_list_append((**C.alpm_list_t)((unsafe.Pointer(l))), unsafe.Pointer(str)))
}

// TODO: AddSorted

func (l *List) Join(l2 *List) *List {
	return (*List)(C.alpm_list_join((*C.alpm_list_t)(l), (*C.alpm_list_t)(l2)))
}

// TODO: MMerge

// TODO: MSort

func (l *List) RemoveItem(item *List) *List {
	return (*List)(C.alpm_list_remove_item((*C.alpm_list_t)(l), (*C.alpm_list_t)(item)))
}

// TODO: Remove

func (l *List) RemoveStr(str string) *List {
	var data *C.char

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	ret := (*List)(C.alpm_list_remove_str((*C.alpm_list_t)(l), cstr, &data))
	C.free(unsafe.Pointer(data))
	return ret
}

func (l *List) RemoveDupes() *List {
	return (*List)(C.alpm_list_remove_dupes((*C.alpm_list_t)(l)))
}

func (l *List) Strdup() *List {
	return (*List)(C.alpm_list_strdup((*C.alpm_list_t)(l)))
}

func (l *List) Copy() *List {
	return (*List)(C.alpm_list_copy((*C.alpm_list_t)(l)))
}

func (l *List) CopyData(size uint) *List {
	return (*List)(C.alpm_list_copy_data((*C.alpm_list_t)(l), C.size_t(size)))
}

func (l *List) Reverse() *List {
	return (*List)(C.alpm_list_reverse((*C.alpm_list_t)(l)))
}

/* item accessors */
func (l *List) Nth(n uint) *List {
	return (*List)(C.alpm_list_nth((*C.alpm_list_t)(l), C.size_t(n)))
}

func (l *List) Next() *List {
	return (*List)(C.alpm_list_next((*C.alpm_list_t)(l)))
}

func (l *List) Previous() *List {
	return (*List)(C.alpm_list_previous((*C.alpm_list_t)(l)))
}

func (l *List) Last() *List {
	return (*List)(C.alpm_list_last((*C.alpm_list_t)(l)))
}

/* Misc */

func (l *List) Count() uint {
	return (uint)(C.alpm_list_count((*C.alpm_list_t)(l)))
}

// TODO: Find

func (l *List) FindPtr(ptr uintptr) uintptr {
	return (uintptr)(C.alpm_list_find_ptr((*C.alpm_list_t)(l), unsafe.Pointer(ptr)))
}

func (l *List) FindStr(str string) bool {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return C.alpm_list_find_ptr((*C.alpm_list_t)(l), unsafe.Pointer(cstr)) != nil
}

// TODO: DiffSorted

// TODO: Diff

func (l *List) Slice() []uintptr {
	slice := make([]uintptr, l.Count(), 0)
	for i := l; i != nil; i = i.Next() {
		slice = append(slice, i.Data())
	}
	return slice
}

// Custom

func (l *List) String() string {
	return C.GoString((*C.char)(unsafe.Pointer(l.Data())))
}

func (l *List) ForEach(f func(uintptr) error) error {
	for i := l; i != nil; i = i.Next() {
		if err := f(i.Data()); err != nil {
			return err
		}
	}

	return nil
}
