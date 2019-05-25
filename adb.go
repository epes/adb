package adb

import (
	"fmt"
	"image/draw"
)

// ADB defines the device and provides controls over the
// android device bridge
type ADB struct {
	screenWidth  int
	screenHeight int
	rotated      bool
}

// NewADB returns an instance of ADB to control the device
func NewADB(screenWidth int, screenHeight int, rotated bool) *ADB {
	return &ADB{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		rotated:      rotated,
	}
}

// Touch sends a click to the specified X and Y coordinates
func (a *ADB) Touch(x int, y int) error {
	if a.rotated {
		err := a.touchCommand(a.screenWidth-y, x)
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else {
		err := a.touchCommand(x, y)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

// Screencap returns a screencap of the device over ADB
func (a *ADB) Screencap() (draw.Image, error) {
	return a.screencapCommand()
}
