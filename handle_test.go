package alpm

import "fmt"
import "github.com/jguer/go-alpm/alpm_list"

// Tests package attribute getters.
func ExampleHandke() {
	root := "/"
	dbpath := "/var/lib/pacman"

	h, _ := Initialize(root, dbpath)
	defer h.Release()

	var list *alpm_list.List
	alpm_list.AppendStrdup(&list, "a/")
	alpm_list.AppendStrdup(&list, "b/")
	alpm_list.AppendStrdup(&list, "c/")

	h.SetArch("foo")
	fmt.Println(h.Root())
	fmt.Println(h.DBPath())
	fmt.Println(h.Arch())

	h.SetCacheDirs(StringList{list})
	h.AddCacheDir("d/")
	h.RemoveCacheDir("b/")
	cd, err := h.CacheDirs()
	fmt.Println(cd.Slice(), err)

	fmt.Println(h.UseSyslog())
	h.SetUseSyslog(true)
	fmt.Println(h.UseSyslog())
	// Output:
	// / <nil>
	// /var/lib/pacman/ <nil>
	// foo <nil>
	// [a/ c/ d/] <nil>
	// false <nil>
	// true <nil>
}
