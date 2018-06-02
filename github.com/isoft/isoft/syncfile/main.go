package main

import "isoft/syncfile/sync"

func main() {
	//ok := flag.Bool("static_sync", false, "static_sync_all")
	//fmt.Print(*ok)

	SyncFile := sync.ReadSyncFile("D:\\zhourui\\program\\go\\goland_workspace\\src\\isoft\\syncfile\\sync\\static.xml")
	sync.StartAllSyncFile(SyncFile, "")

}
