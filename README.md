# imgutils
Golang image related utility functions.

# Installation
```bash
go get -u github.com/SkyMack/imgutils
```

# Usage
Some example code snippets (incomplete):
```go
	import (
		"fmt"
		"image"
		"image/color"
		"image/draw"
		"math"
		"strconv"
		"strings"
		
		"github.com/SkyMack/imgutils"
		"github.com/golang/freetype/truetype"
		"golang.org/x/image/font"
		"golang.org/x/image/math/fixed"
	)

	fontBorderColor, err := imgutils.ParseHexColor(fontBorderColorStr)
	if err != nil {
		return err
	}

	// create a temp image to draw the text onto
	textImg := image.NewNRGBA(image.Rect(0, 0, conf.textImgWidth, conf.textImgHeight))
	// calc Y level to place the font Drawer dot at, given the font size and DPI
	y := int(math.Ceil(conf.fontSize * fontDPI / 72))
	startDot := fixed.Point26_6{
		X: fixed.I(0 + 2 + conf.fontBorderWidth),
		Y: fixed.I(y + ((2 + conf.fontBorderWidth) * 2)),
	}
	textDrawer := &font.Drawer{
		Dst: textImg,
		Src: conf.fontColor,
		Face: truetype.NewFace(parsedFont, &truetype.Options{
			Size:    conf.fontSize,
			DPI:     fontDPI,
			Hinting: font.HintingFull,
		}),
		Dot: startDot,
	}
	text := fmt.Sprintf("#%v", thumb.paddedNumber)
	// draw the sequence number onto the temp image
	textDrawer.DrawString(text)
	// add the main text border/outline
	imgutils.AddBorders(textImg, conf.fontBorderColor, conf.fontBorderWidth, conf.fontBorderAlphaThresh)

	// 
	textRect := imgutils.OccupiedAreaRect(textImg)
```