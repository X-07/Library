// Package common provides simple gtk helpers.
package gotk3

//
//------------------------------------------------------------------[ WINDOW ]--

// Special hack to prevent threads related crashs.
// http://stackoverflow.com/questions/18647475/threading-problems-with-gtk
// http://stackoverflow.com/questions/13351297/what-is-the-downside-of-xinitthreads

// #include <X11/Xlib.h>
// #cgo pkg-config: x11
import "C"

// GRRTHREADS is a dirty hack to prevent threads related crashs.
func GRRTHREADS() {
	C.XInitThreads()
}
