// Same copyright and license as the rest of the files in this project
// TODO: glib build tags for since GLib 2.46

package glib

// // #include <gio/gio.h>
// // #include <glib.h>
// // #include <glib-object.h>
// // #include "glib.go.h"
// // #include "glib_since_2_44.go.h"
// import "C"

/*
 * GListStore
 */

// Sort is a wrapper around g_list_store_sort().
// func (v *ListStore) Sort(compareFunc CompareDataFunc, userData ...interface{}) {
// 		// TODO: figure out a way to determine when we can clean up
// 		compareDataFuncRegistry.Lock()
// 		id := compareDataFuncRegistry.next
// 		compareDataFuncRegistry.next++
// 		compareDataFuncRegistry.m[id] = compareDataFuncData{fn: compareFunc, userData: userData}
// 		compareDataFuncRegistry.Unlock()
	
// 		C._g_list_store_sort(v.native(), C.gpointer(uintptr(id)))
// }
