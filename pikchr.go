package pikchr

// #include "pikchr.h"
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

// Render pikchr source as an SVG.
func Render(src string, in ...Option) (*RenderResult, bool) {
	opts := &RenderOptions{}
	for _, set := range in {
		set(opts)
	}

	rflags := C.PIKCHR_PLAINTEXT_ERRORS
	if opts.DarkMode {
		rflags |= C.PIKCHR_DARK_MODE
	}
	if opts.HTMLTextError {
		rflags &^= C.PIKCHR_PLAINTEXT_ERRORS
	}

	var cname *C.char
	if opts.SVGClassName != "" {
		cname = C.CString(opts.SVGClassName)
	}

	w, h := C.int(0), C.int(0)

	csrc := C.CString(src)
	cres := C.pikchr(
		csrc,
		cname,
		C.uint(rflags),
		&w,
		&h,
	)
	C.free(unsafe.Pointer(csrc))
	C.free(unsafe.Pointer(cname))
	defer C.free(unsafe.Pointer(cres))

	res := &RenderResult{
		Data:   C.GoString(cres),
		Width:  int(w),
		Height: int(h),
	}

	return res, int(w) > 0
}

type RenderResult struct {
	// svg or text of error
	Data string

	Width, Height int
}

type RenderOptions struct {
	HTMLTextError bool
	DarkMode      bool
	SVGClassName  string
}

type Option func(*RenderOptions)

func SVGClass(name string) Option {
	return func(opts *RenderOptions) {
		opts.SVGClassName = name
	}
}

func Dark() Option {
	return func(opts *RenderOptions) {
		opts.DarkMode = true
	}
}

func HTMLError() Option {
	return func(opts *RenderOptions) {
		opts.HTMLTextError = true
	}
}
