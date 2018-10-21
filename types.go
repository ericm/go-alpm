// types.go - libalpm types.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package alpm

// #cgo CFLAGS: -D_FILE_OFFSET_BITS=64
// #include <alpm.h>
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/jguer/go-alpm/alpm_list"
)

type Depend C.alpm_depend_t

func (dep *Depend) Name() string {
	return C.GoString(dep.name)
}

func (dep *Depend) Version() string {
	return C.GoString(dep.version)
}

func (dep *Depend) Description() string {
	return C.GoString(dep.desc)
}

func (dep *Depend) NameHash() uint {
	return uint(dep.name_hash)
}

func (dep *Depend) Mod() DepMod {
	return DepMod(dep.mod)
}

func (dep *Depend) String() string {
	str := C.alpm_dep_compute_string((*C.alpm_depend_t)(dep))
	return C.GoString(str)
}

func (dep *Depend) Free() {
	C.alpm_dep_free((*C.alpm_depend_t)(dep))
}

type File C.alpm_file_t

func (file *File) Name() string {
	return C.GoString(file.name)
}

func (file *File) Size() int64 {
	return int64(file.size)
}

func (file *File) Mode() uint32 {
	return uint32(file.mode)
}

type FileList C.alpm_filelist_t

func (fl *FileList) Count() uint {
	return uint(fl.count)
}

func (fl *FileList) Slice() []File {
	size := int(fl.Count())

	items := reflect.SliceHeader{
		Len:  size,
		Cap:  size,
		Data: uintptr(unsafe.Pointer(fl.files))}

	return *(*[]File)(unsafe.Pointer(&items))
}

type Backup C.alpm_backup_t

type QuestionAny C.alpm_question_any_t

func (question *QuestionAny) SetAnswer(answer bool) {
	if answer {
		question.answer = 1
	} else {
		question.answer = 0
	}
}

type QuestionInstallIgnorepkg C.alpm_question_install_ignorepkg_t

func (question *QuestionAny) Type() QuestionType {
	return QuestionType(question._type)
}

func (question *QuestionAny) Answer() bool {
	return question.answer == 1
}

func (question *QuestionAny) QuestionInstallIgnorepkg() (QuestionInstallIgnorepkg, error) {
	if question.Type() == QuestionTypeInstallIgnorepkg {
		return *(*QuestionInstallIgnorepkg)(unsafe.Pointer(&question)), nil
	}

	return QuestionInstallIgnorepkg{}, fmt.Errorf("Can not convert to QuestionInstallIgnorepkg")
}

func (question *QuestionAny) QuestionSelectProvider() (QuestionSelectProvider, error) {
	if question.Type() == QuestionTypeSelectProvider {
		return *(*QuestionSelectProvider)(unsafe.Pointer(&question)), nil
	}

	return QuestionSelectProvider{}, fmt.Errorf("Can not convert to QuestionInstallIgnorepkg")
}

func (question *QuestionAny) QuestionReplace() (QuestionReplace, error) {
	if question.Type() == QuestionTypeReplacePkg {
		return *(*QuestionReplace)(unsafe.Pointer(&question)), nil
	}

	return QuestionReplace{}, fmt.Errorf("Can not convert to QuestionReplace")
}

func (question QuestionInstallIgnorepkg) SetInstall(install bool) {
	if install {
		question.install = 1
	} else {
		question.install = 0
	}
}

func (question *QuestionInstallIgnorepkg) Type() QuestionType {
	return QuestionType(question._type)
}

func (question *QuestionInstallIgnorepkg) Install() bool {
	return question.install == 1
}

func (question *QuestionInstallIgnorepkg) Pkg(h *Handle) Package {
	return Package{
		question.pkg,
		*h,
	}
}

type QuestionReplace C.alpm_question_replace_t

func (question *QuestionReplace) Type() QuestionType {
	return QuestionType(question._type)
}

func (question *QuestionReplace) SetReplace(replace bool) {
	if replace {
		question.replace = 1
	} else {
		question.replace = 0
	}
}

func (question *QuestionReplace) Replace() bool {
	return question.replace == 1
}

func (question *QuestionReplace) NewPkg(h *Handle) *Package {
	return &Package{
		question.newpkg,
		*h,
	}
}

func (question *QuestionReplace) OldPkg(h *Handle) *Package {
	return &Package{
		question.oldpkg,
		*h,
	}
}

func (question *QuestionReplace) newDB(h *Handle) *DB {
	return &DB{
		question.newdb,
		*h,
	}
}

type QuestionSelectProvider C.alpm_question_select_provider_t

func (question *QuestionSelectProvider) Type() QuestionType {
	return QuestionType(question._type)
}

func (question *QuestionSelectProvider) SetUseIndex(index int) {
	question.use_index = C.int(index)
}

func (question *QuestionSelectProvider) UseIndex() int {
	return int(question.use_index)
}

func (question *QuestionSelectProvider) Providers(h *Handle) PackageList {
	return PackageList{
		(*alpm_list.List)(unsafe.Pointer(question.providers)),
		*h,
	}
}

func (question *QuestionSelectProvider) Dep() *Depend {
	return (*Depend)(question.depend)
}
