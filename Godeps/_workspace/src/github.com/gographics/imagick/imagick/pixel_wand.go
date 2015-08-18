// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

import (
	"unsafe"
)

type PixelWand struct {
	pw *C.PixelWand
}

// Returns a new pixel wand
func NewPixelWand() *PixelWand {
	return &PixelWand{C.NewPixelWand()}
}

// Clears resources associated with the wand
func (pw *PixelWand) Clear() {
	C.ClearPixelWand(pw.pw)
}

// Makes an exact copy of the wand
func (pw *PixelWand) Clone() *PixelWand {
	return &PixelWand{C.ClonePixelWand(pw.pw)}
}

// Deallocates resources associated with a pixel wand
func (pw *PixelWand) Destroy() {
	if pw.pw == nil {
		return
	}
	pw.pw = C.DestroyPixelWand(pw.pw)
	C.free(unsafe.Pointer(pw.pw))
	pw.pw = nil
}

// Returns true if the distance between two colors is less than the specified distance
func (pw *PixelWand) IsSimilar(pixelWand *PixelWand, fuzz float64) bool {
	return 1 == C.int(C.IsPixelWandSimilar(pw.pw, pixelWand.pw, C.double(fuzz)))
}

// Returns true if the wand is verified as a pixel wand
func (pw *PixelWand) IsVerified() bool {
	if pw.pw != nil {
		return 1 == C.int(C.IsPixelWand(pw.pw))
	}
	return false
}

// Returns the normalized alpha color of the pixel wand
func (pw *PixelWand) GetAlpha() float64 {
	return float64(C.PixelGetAlpha(pw.pw))
}

// Returns the alpha value of the pixel wand
func (pw *PixelWand) GetAlphaQuantum() Quantum {
	return Quantum(C.PixelGetAlphaQuantum(pw.pw))
}

// Returns the normalized black color of the pixel wand
func (pw *PixelWand) GetBlack() float64 {
	return float64(C.PixelGetBlack(pw.pw))
}

// Returns the black color of the pixel wand
func (pw *PixelWand) GetBlackQuantum() Quantum {
	return Quantum(C.PixelGetBlackQuantum(pw.pw))
}

// Returns the normalized blue color of the pixel wand
func (pw *PixelWand) GetBlue() float64 {
	return float64(C.PixelGetBlue(pw.pw))
}

// Returns the blue color of the pixel wand
func (pw *PixelWand) GetBlueQuantum() Quantum {
	return Quantum(C.PixelGetBlueQuantum(pw.pw))
}

// Returns the color of the pixel wand as a string
func (pw *PixelWand) GetColorAsString() string {
	return C.GoString(C.PixelGetColorAsString(pw.pw))
}

// Returns the normalized color of the pixel wand as string
func (pw *PixelWand) GetColorAsNormalizedString() string {
	return C.GoString(C.PixelGetColorAsNormalizedString(pw.pw))
}

// Returns the color count associated with this color
func (pw *PixelWand) GetColorCount() uint {
	return uint(C.PixelGetColorCount(pw.pw))
}

// Returns the normalized cyan color of the pixel wand
func (pw *PixelWand) GetCyan() float64 {
	return float64(C.PixelGetCyan(pw.pw))
}

// Returns the cyan color of the pixel wand
func (pw *PixelWand) GetCyanQuantum() Quantum {
	return Quantum(C.PixelGetCyanQuantum(pw.pw))
}

// Returns the normalized fuzz value of the pixel wand
func (pw *PixelWand) GetFuzz() float64 {
	return float64(C.PixelGetFuzz(pw.pw))
}

// Returns the normalized green color of the pixel wand
func (pw *PixelWand) GetGreen() float64 {
	return float64(C.PixelGetGreen(pw.pw))
}

// Returns the green color of the pixel wand
func (pw *PixelWand) GetGreenQuantum() Quantum {
	return Quantum(C.PixelGetGreenQuantum(pw.pw))
}

// Returns the normalized HSL color of the pixel wand
func (pw *PixelWand) GetHSL() (hue, saturation, brightness float64) {
	var cdhue, cdsaturation, cdbrightness C.double
	C.PixelGetHSL(pw.pw, &cdhue, &cdsaturation, &cdbrightness)
	hue, saturation, brightness = float64(cdhue), float64(cdsaturation), float64(cdbrightness)
	return
}

// Returns the colormap index from the pixel wand
func (pw *PixelWand) GetIndex() IndexPacket {
	cip := C.PixelGetIndex(pw.pw)
	return IndexPacket(cip)
}

// Returns the normalized magenta color of the pixel wand
func (pw *PixelWand) GetMagenta() float64 {
	return float64(C.PixelGetMagenta(pw.pw))
}

// Returns the magenta color of the pixel wand
func (pw *PixelWand) GetMagentaQuantum() Quantum {
	return Quantum(C.PixelGetMagentaQuantum(pw.pw))
}

// Gets the magick color of the pixel wand
func (pw *PixelWand) GetMagickColor() *MagickPixelPacket {
	var mpp C.MagickPixelPacket
	C.PixelGetMagickColor(pw.pw, &mpp)
	return newMagickPixelPacketFromCAPI(&mpp)
}

// Returns the normalized opacity color of the pixel wand
func (pw *PixelWand) GetOpacity() float64 {
	return float64(C.PixelGetOpacity(pw.pw))
}

// Returns the opacity color of the pixel wand
func (pw *PixelWand) GetOpacityQuantum() Quantum {
	return Quantum(C.PixelGetOpacityQuantum(pw.pw))
}

// Gets the color of the pixel wand as a PixelPacket
func (pw *PixelWand) GetQuantumColor() *PixelPacket {
	var pp C.PixelPacket
	C.PixelGetQuantumColor(pw.pw, &pp)
	return newPixelPacketFromCAPI(&pp)
}

// Returns the normalized red color of the pixel wand
func (pw *PixelWand) GetRed() float64 {
	return float64(C.PixelGetRed(pw.pw))
}

// Returns the red color of the pixel wand
func (pw *PixelWand) GetRedQuantum() Quantum {
	return Quantum(C.PixelGetRedQuantum(pw.pw))
}

// Returns the normalized yellow color of the pixel wand
func (pw *PixelWand) GetYellow() float64 {
	return float64(C.PixelGetYellow(pw.pw))
}

// Returns the yellow color of the pixel wand
func (pw *PixelWand) GetYellowQuantum() Quantum {
	return Quantum(C.PixelGetYellowQuantum(pw.pw))
}

// Sets the normalized alpha color of the pixel wand.
// 1.0 is fully opaque and 0.0 is fully transparent.
func (pw *PixelWand) SetAlpha(alpha float64) {
	C.PixelSetAlpha(pw.pw, C.double(alpha))
}

// Sets the alpha color of the pixel wand
func (pw *PixelWand) SetAlphaQuantum(opacity Quantum) {
	C.PixelSetAlphaQuantum(pw.pw, C.Quantum(opacity))
}

// Sets the normalized black color of the pixel wand
func (pw *PixelWand) SetBlack(black float64) {
	C.PixelSetBlack(pw.pw, C.double(black))
}

// Sets the black color of the pixel wand
func (pw *PixelWand) SetBlackQuantum(black Quantum) {
	C.PixelSetBlackQuantum(pw.pw, C.Quantum(black))
}

// Sets the normalized blue color of the pixel wand
func (pw *PixelWand) SetBlue(blue float64) {
	C.PixelSetBlue(pw.pw, C.double(blue))
}

// Sets the blue color of the pixel wand
func (pw *PixelWand) SetBlueQuantum(blue Quantum) {
	C.PixelSetBlueQuantum(pw.pw, C.Quantum(blue))
}

// Sets the color of the pixel wand with a string (e.g. "blue", "#0000ff", "rgb(0,0,255)", "cmyk(100,100,100,10)", etc.)
func (pw *PixelWand) SetColor(color string) bool {
	cscolor := C.CString(color)
	defer C.free(unsafe.Pointer(cscolor))
	return 1 == int(C.PixelSetColor(pw.pw, cscolor))
}

// Sets the color count of the pixel wand
func (pw *PixelWand) SetColorCount(count uint) {
	C.PixelSetColorCount(pw.pw, C.size_t(count))
}

// Sets the color of the pixel wand from another one
func (pw *PixelWand) SetColorFromWand(pixelWand *PixelWand) {
	C.PixelSetColorFromWand(pw.pw, pixelWand.pw)
}

// Sets the normalized cyan color of the pixel wand
func (pw *PixelWand) SetCyan(cyan float64) {
	C.PixelSetCyan(pw.pw, C.double(cyan))
}

// Sets the cyan color of the pixel wand
func (pw *PixelWand) SetCyanQuantum(cyan Quantum) {
	C.PixelSetCyanQuantum(pw.pw, C.Quantum(cyan))
}

// Sets the fuzz value of the pixel wand
func (pw *PixelWand) SetFuzz(fuzz float64) {
	C.PixelSetFuzz(pw.pw, C.double(fuzz))
}

// Sets the normalized green color of the pixel wand
func (pw *PixelWand) SetGreen(green float64) {
	C.PixelSetGreen(pw.pw, C.double(green))
}

// Sets the green color of the pixel wand
func (pw *PixelWand) SetGreenQuantum(green Quantum) {
	C.PixelSetGreenQuantum(pw.pw, C.Quantum(green))
}

// Sets the normalized HSL color of the pixel wand
func (pw *PixelWand) SetHSL(hue, saturation, brightness float64) {
	C.PixelSetHSL(pw.pw, C.double(hue), C.double(saturation), C.double(brightness))
}

// Sets the colormap index of the pixel wand
func (pw *PixelWand) SetIndex(index *IndexPacket) {
	C.PixelSetIndex(pw.pw, C.IndexPacket(*index))
}

// Sets the normalized magenta color of the pixel wand
func (pw *PixelWand) SetMagenta(magenta float64) {
	C.PixelSetMagenta(pw.pw, C.double(magenta))
}

// Sets the magenta color of the pixel wand
func (pw *PixelWand) SetMagentaQuantum(magenta Quantum) {
	C.PixelSetMagentaQuantum(pw.pw, C.Quantum(magenta))
}

// Sets the color of the pixel wand
func (pw *PixelWand) SetMagickColor(color *MagickPixelPacket) {
	C.PixelSetMagickColor(pw.pw, color.mpp)
}

// Sets the normalized opacity color of the pixel wand
func (pw *PixelWand) SetOpacity(opacity float64) {
	C.PixelSetOpacity(pw.pw, C.double(opacity))
}

// Sets the opacity color of the pixel wand
func (pw *PixelWand) SetOpacityQuantum(opacity Quantum) {
	C.PixelSetOpacityQuantum(pw.pw, C.Quantum(opacity))
}

// Sets the color of the pixel wand
func (pw *PixelWand) SetQuantumColor(color *PixelPacket) {
	C.PixelSetQuantumColor(pw.pw, color.pp)
}

// Sets the normalized red color of the pixel wand
func (pw *PixelWand) SetRed(red float64) {
	C.PixelSetRed(pw.pw, C.double(red))
}

// Sets the red color of the pixel wand
func (pw *PixelWand) SetRedQuantum(red Quantum) {
	C.PixelSetRedQuantum(pw.pw, C.Quantum(red))
}

// Sets the normalized yellow color of the pixel wand
func (pw *PixelWand) SetYellow(yellow float64) {
	C.PixelSetYellow(pw.pw, C.double(yellow))
}

// Sets the yellow color of the pixel wand
func (pw *PixelWand) SetYellowQuantum(yellow Quantum) {
	C.PixelSetYellowQuantum(pw.pw, C.Quantum(yellow))
}
