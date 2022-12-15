package imgutil

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	blueSqrHeight   = 3
	blueSqrWidth    = 3
	purpleSqrAlpha  = 120
	purpleSqrHeight = 2
	purpleSqrWidth  = 2
	redSqrHeight    = 3
	redSqrWidth     = 3
	testImgHeight   = 13
	testImgWidth    = 13
)

var (
	blueSqrOffset   = image.Point{X: 7, Y: 7}
	blueUniform     = &image.Uniform{C: color.NRGBA{R: 0, G: 0, B: 255, A: 255}}
	cyanNRGBA       = color.NRGBA{R: 0, G: 100, B: 100, A: 255}
	cyanHex         = "006464"
	debugMode       = false
	purpleSqrOffset = image.Point{X: 2, Y: 9}
	purpleUniform   = &image.Uniform{C: color.NRGBA{R: 200, G: 0, B: 200, A: purpleSqrAlpha}}
	redSqrOffset    = image.Point{X: 4, Y: 4}
	redUniform      = &image.Uniform{C: color.NRGBA{R: 255, G: 0, B: 0, A: 255}}
	testOutputDir   = filepath.Join("testdata", "output")

	blueSqrBottomRightPixel = blueSqrOffset.Add(image.Point{X: blueSqrWidth - 1, Y: blueSqrWidth - 1})
	bottomRightMostPixel    = image.Point{X: testImgWidth - 1, Y: testImgHeight - 1}
	purpleSqrTopLeftPixel   = purpleSqrOffset
	redSqrTopLeftPixel      = redSqrOffset
	topLeftMostPixel        = image.Point{X: 0, Y: 0}
	topRightMostPixel       = image.Point{X: testImgWidth - 1, Y: 0}
)

func init() {
	if val, ok := os.LookupEnv("DEBUG"); ok && val == "1" {
		debugMode = true

		img := generateTestImage()
		if err := saveImage(img, "test_image"); err != nil {
			panic(err)
		}
	}
}

func generateTestImage() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, testImgWidth, testImgHeight))
	redRect := image.Rect(0, 0, redSqrWidth, redSqrHeight)
	blueRect := image.Rect(0, 0, blueSqrWidth, blueSqrHeight)
	purpleRect := image.Rect(0, 0, purpleSqrWidth, purpleSqrHeight)
	draw.Draw(img, redRect.Add(redSqrOffset), redUniform, image.Point{}, draw.Src)
	draw.Draw(img, blueRect.Add(blueSqrOffset), blueUniform, image.Point{}, draw.Src)
	draw.Draw(img, purpleRect.Add(purpleSqrOffset), purpleUniform, image.Point{}, draw.Src)
	return img
}

func check1PixRedSqrBorder(t *testing.T, img *image.NRGBA, shouldBeEmpty bool) {
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, redSqrTopLeftPixel.X-1, redSqrTopLeftPixel.Y))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, redSqrTopLeftPixel.X, redSqrTopLeftPixel.Y-1))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, redSqrTopLeftPixel.X+redSqrWidth, redSqrTopLeftPixel.Y))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, redSqrTopLeftPixel.X, redSqrTopLeftPixel.Y+redSqrHeight))
}

func check1PixBlueSqrBorder(t *testing.T, img *image.NRGBA, shouldBeEmpty bool) {
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, blueSqrBottomRightPixel.X+1, blueSqrBottomRightPixel.Y))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, blueSqrBottomRightPixel.X, blueSqrBottomRightPixel.Y+1))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, blueSqrBottomRightPixel.X-blueSqrWidth, blueSqrBottomRightPixel.Y))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, blueSqrBottomRightPixel.X, blueSqrBottomRightPixel.Y-blueSqrHeight))
}

func check1PixPurpleSqrBorder(t *testing.T, img *image.NRGBA, shouldBeEmpty bool) {
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, purpleSqrTopLeftPixel.X-1, purpleSqrTopLeftPixel.Y))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, purpleSqrTopLeftPixel.X, purpleSqrTopLeftPixel.Y-1))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, purpleSqrTopLeftPixel.X+purpleSqrWidth, purpleSqrTopLeftPixel.Y))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, purpleSqrTopLeftPixel.X, purpleSqrTopLeftPixel.Y+purpleSqrHeight))
}

func check2PixRedSqrBorder(t *testing.T, img *image.NRGBA, shouldBeEmpty bool) {
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, redSqrTopLeftPixel.X-2, redSqrTopLeftPixel.Y))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, redSqrTopLeftPixel.X, redSqrTopLeftPixel.Y-2))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, redSqrTopLeftPixel.X+redSqrWidth+1, redSqrTopLeftPixel.Y))
}

func check2PixBlueSqrBorder(t *testing.T, img *image.NRGBA, shouldBeEmpty bool) {
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, blueSqrBottomRightPixel.X+2, blueSqrBottomRightPixel.Y))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, blueSqrBottomRightPixel.X, blueSqrBottomRightPixel.Y+2))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, blueSqrBottomRightPixel.X-blueSqrWidth-1, blueSqrBottomRightPixel.Y))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, blueSqrBottomRightPixel.X, blueSqrBottomRightPixel.Y-blueSqrHeight-1))
}

func check2PixPurpleSqrBorder(t *testing.T, img *image.NRGBA, shouldBeEmpty bool) {
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, purpleSqrTopLeftPixel.X, purpleSqrTopLeftPixel.Y-2))
	assert.Equal(t, shouldBeEmpty, IsEmptyPixel(img, purpleSqrTopLeftPixel.X+purpleSqrWidth+1, purpleSqrTopLeftPixel.Y))
}

func checkKnowFilledNotEmpty(t *testing.T, img *image.NRGBA) {
	assert.Equal(t, false, IsEmptyPixel(img, redSqrTopLeftPixel.X, redSqrTopLeftPixel.Y))
	assert.Equal(t, false, IsEmptyPixel(img, blueSqrBottomRightPixel.X, blueSqrBottomRightPixel.Y))
}

func checkKnownEmptyAreEmpty(t *testing.T, img *image.NRGBA) {
	assert.Equal(t, true, IsEmptyPixel(img, topLeftMostPixel.X, topLeftMostPixel.Y))
	assert.Equal(t, true, IsEmptyPixel(img, topRightMostPixel.X, topRightMostPixel.Y))
	assert.Equal(t, true, IsEmptyPixel(img, bottomRightMostPixel.X, bottomRightMostPixel.Y))
}

func checkStartingState(t *testing.T, img *image.NRGBA) {
	// Ensure known-filled pixels are not empty and known empty pixels are empty
	checkKnowFilledNotEmpty(t, img)
	checkKnownEmptyAreEmpty(t, img)

	// Ensure pixels 1 space out are currently empty
	check1PixRedSqrBorder(t, img, true)
	check1PixBlueSqrBorder(t, img, true)
	check1PixPurpleSqrBorder(t, img, true)

	// Ensure pixels 2 spaces out are also empty
	check2PixRedSqrBorder(t, img, true)
}

func saveImage(img *image.NRGBA, name string) error {
	if debugMode {
		fh, err := os.Create(filepath.Join(testOutputDir, fmt.Sprintf("%s.png", strings.Replace(name, "/", "_", -1))))
		if err != nil {
			return err
		}
		defer fh.Close()
		if err := png.Encode(fh, img); err != nil {
			return err
		}
	}
	return nil
}

func TestIsEmptyPixel(t *testing.T) {
	img := generateTestImage()
	t.Run("Empty Pixel Returns True", func(t *testing.T) {
		checkKnownEmptyAreEmpty(t, img)
	})

	t.Run("Active Pixel Returns False", func(t *testing.T) {
		checkKnowFilledNotEmpty(t, img)
	})
}

func TestAddBorders(t *testing.T) {
	t.Run("Draw1PixBorder", func(t *testing.T) {
		img := generateTestImage()
		borderWidth := 1

		checkStartingState(t, img)

		AddBorders(img, cyanNRGBA, borderWidth, 0)
		err := saveImage(img, t.Name())
		assert.NoError(t, err)

		// Ensure known-filled pixels are still not empty and known empty pixels are still empty
		checkKnowFilledNotEmpty(t, img)
		checkKnownEmptyAreEmpty(t, img)

		// Ensure expected border pixels were filled in
		check1PixRedSqrBorder(t, img, false)
		check1PixBlueSqrBorder(t, img, false)
		check1PixPurpleSqrBorder(t, img, false)

		// Ensure pixels 2 spaces out are still empty
		check2PixRedSqrBorder(t, img, true)
		check2PixBlueSqrBorder(t, img, true)
		check2PixPurpleSqrBorder(t, img, true)
	})

	t.Run("Draw2PixBorder", func(t *testing.T) {
		img := generateTestImage()
		borderWidth := 2

		checkStartingState(t, img)

		AddBorders(img, cyanNRGBA, borderWidth, 0)
		err := saveImage(img, t.Name())
		assert.NoError(t, err)

		// Ensure known-filled pixels are still not empty and known empty pixels are still empty
		checkKnowFilledNotEmpty(t, img)
		checkKnownEmptyAreEmpty(t, img)

		// Ensure pixels 2 spaces out from the rectangles are now filled
		check2PixRedSqrBorder(t, img, false)
		check2PixBlueSqrBorder(t, img, false)
		check2PixPurpleSqrBorder(t, img, false)
	})

	t.Run("Draw1PixBorderWithAlphaThreshold", func(t *testing.T) {
		img := generateTestImage()
		borderWidth := 1

		checkStartingState(t, img)

		AddBorders(img, cyanNRGBA, borderWidth, purpleSqrAlpha+5)
		err := saveImage(img, t.Name())
		assert.NoError(t, err)

		// Ensure known-filled pixels are still not empty and known empty pixels are still empty
		checkKnowFilledNotEmpty(t, img)
		checkKnownEmptyAreEmpty(t, img)

		// Ensure expected border pixels were filled in
		check1PixRedSqrBorder(t, img, false)
		check1PixBlueSqrBorder(t, img, false)
		check1PixPurpleSqrBorder(t, img, true)

		// Ensure pixels 2 spaces out are still empty
		check2PixRedSqrBorder(t, img, true)
	})
}

func TestParseHexColor(t *testing.T) {
	t.Run("ValidHexCode", func(t *testing.T) {
		testColor, err := ParseHexColor(cyanHex)
		assert.NoError(t, err)
		assert.Equal(t, cyanNRGBA.R, testColor.R, "red value is correct")
		assert.Equal(t, cyanNRGBA.G, testColor.G, "green value is correct")
		assert.Equal(t, cyanNRGBA.B, testColor.B, "blue value is correct")
		assert.Equal(t, cyanNRGBA.A, testColor.A, "alpha value is correct")
	})

	t.Run("StringToShort", func(t *testing.T) {
		_, err := ParseHexColor("123")
		assert.Error(t, err)
	})

	t.Run("StringToLong", func(t *testing.T) {
		_, err := ParseHexColor("123456789")
		assert.Error(t, err)
	})

	t.Run("StringContainsInvalidCharacter", func(t *testing.T) {
		_, err := ParseHexColor("X23456")
		assert.Error(t, err)
	})
}

func TestOccupiedAreaRect(t *testing.T) {
	t.Run("BorderlessTestImage", func(t *testing.T) {
		img := generateTestImage()
		rect := OccupiedAreaRect(img)
		assert.Equal(t, 1, rect.Min.X)
		assert.Equal(t, 3, rect.Min.Y)
		assert.Equal(t, 10, rect.Max.X)
		assert.Equal(t, 11, rect.Max.Y)
	})
}
