package alpm

import (
	"fmt"

	"github.com/jguer/go-alpm/alpm_list"
)

func ExampleServer() {
	h, _ := Initialize(root, dbpath)
	defer h.Release()
	h.RegisterSyncDB("core", 0)

	dbs, _ := h.SyncDBs()

	dbs.ForEach(func(db *DB) error {
		var servers *alpm_list.List
		alpm_list.AppendStrdup(&servers, "a")
		alpm_list.AppendStrdup(&servers, "b")
		alpm_list.AppendStrdup(&servers, "c")

		db.SetServers(StringList{servers})
		s := db.Servers().Slice()
		fmt.Println(s)
		return nil
	})
	// Output: [a b c]

}
