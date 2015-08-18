// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type DrawingWandException struct {
	kind        ExceptionType
	description string
}

func (dwe *DrawingWandException) Error() string {
	return fmt.Sprintf("%s: %s", dwe.kind.String(), dwe.description)
}

// Clears any exceptions associated with the wand
func (dw *DrawingWand) clearException() bool {
	return 1 == C.int(C.DrawClearException(dw.dw))
}

// Returns the kind, reason and description of any error that occurs when using other methods in this API
func (dw *DrawingWand) GetLastError() error {
	var et C.ExceptionType
	csdescription := C.DrawGetException(dw.dw, &et)
	defer C.free(unsafe.Pointer(csdescription))
	if ExceptionType(et) != EXCEPTION_UNDEFINED {
		dw.clearException()
		return &DrawingWandException{ExceptionType(C.int(et)), C.GoString(csdescription)}
	}
	return nil
}
