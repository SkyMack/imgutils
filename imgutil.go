package imgutil

import (
	"encoding/hex"
	"errors"
	"image"
	"image/color"
)

var (
	errInvalidFormat = errors.New("invalid RGB hex code")
)

// IsEmptyPixel returns true if the pixel at the given point in img has an alpha value of 0
func IsEmptyPixel(img *image.NRGBA, pixX int, pixY int) bool {
	if img.NRGBAAt(pixX, pixY).A <= 0 {
		return true
	}
	return false
}

// AddBorders draws a colr colored border borderWidth pixels wide around all shapes/lines in img
func AddBorders(img *image.NRGBA, colr color.Color, borderWidth int, aThresh uint8) {
	for i := 1; i <= borderWidth; i++ {
		addBorderPixels(img, colr, aThresh)
	}
}

// ParseHexColor converts a 6 rune string of hexidecimal characters into a color.NRGBA
func ParseHexColor(s string) (colr color.NRGBA, err error) {
	colr.A = 0xff
	if len([]rune(s)) != 6 {
		return colr, errInvalidFormat
	}
	colrBytes, err := hex.DecodeString(s)
	if err != nil {
		return colr, err
	}

	colr.R = colrBytes[0]
	colr.G = colrBytes[1]
	colr.B = colrBytes[2]
	return
}

// setPixelUnderAlphaThreshold changes the pixel at point in img to colr, but only if the point is currently empty
func setPixelUnderAlphaThreshold(img *image.NRGBA, pixX int, pixY int, colr color.Color, aThresh uint8) bool {
	if img.NRGBAAt(pixX, pixY).A <= aThresh {
		img.Set(pixX, pixY, colr)
		return true
	}
	return false
}

// addBorderPixels toggles pixels with an alpha value less than or equal to alphaThresh to colr above, below, before,
// and after any pixels in img with an alpha value more than alphaThresh
func addBorderPixels(img *image.NRGBA, colr color.Color, alphaThresh uint8) {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	for y := 0; y <= imgHeight; y++ {
		for x := 0; x <= imgWidth; x++ {
			if img.NRGBAAt(x, y).A > alphaThresh {
				setPixelUnderAlphaThreshold(img, x-1, y, colr, alphaThresh)
			}
		}
		for x := imgWidth; x >= 0; x-- {
			if img.NRGBAAt(x, y).A > alphaThresh {
				setPixelUnderAlphaThreshold(img, x+1, y, colr, alphaThresh)
			}
		}
	}

	for x := 0; x <= imgWidth; x++ {
		for y := 0; y <= imgHeight; y++ {
			if img.NRGBAAt(x, y).A > alphaThresh {
				setPixelUnderAlphaThreshold(img, x, y-1, colr, alphaThresh)
			}
		}
		for y := imgHeight; y >= 0; y-- {
			if img.NRGBAAt(x, y).A > alphaThresh {
				setPixelUnderAlphaThreshold(img, x, y+1, colr, alphaThresh)
			}
		}
	}
}

// OccupiedAreaRect returns a rectangle bounding the area of img that is occupied by visible pixels
func OccupiedAreaRect(img *image.NRGBA) image.Rectangle {
	var minX, minY, maxX, maxY int
	var valueFound bool

	imgHeight := img.Bounds().Dy()
	imgWidth := img.Bounds().Dx()

	valueFound = false
	for y := 0; y <= imgHeight; y++ {
		for x := 0; x <= imgWidth; x++ {
			if !IsEmptyPixel(img, x, y) {
				minY = y
				valueFound = true
				break
			}
		}
		if valueFound {
			break
		}
	}
	valueFound = false
	for y := imgHeight; y >= 0; y-- {
		for x := imgWidth; x >= 0; x-- {
			if !IsEmptyPixel(img, x, y) {
				maxY = y
				valueFound = true
				break
			}
		}
		if valueFound {
			break
		}
	}

	valueFound = false
	for x := 0; x <= imgWidth; x++ {
		for y := 0; y <= imgHeight; y++ {
			if !IsEmptyPixel(img, x, y) {
				minX = x
				valueFound = true
				break
			}
		}
		if valueFound {
			break
		}
	}
	valueFound = false
	for x := imgWidth; x >= 0; x-- {
		for y := imgHeight; y >= 0; y-- {
			if !IsEmptyPixel(img, x, y) {
				maxX = x
				valueFound = true
				break
			}
		}
		if valueFound {
			break
		}
	}

	return image.Rect(minX-1, minY-1, maxX+1, maxY+1)
}
