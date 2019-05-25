package adb

import (
	"bytes"
	"fmt"
	"image/draw"
	"image/png"
	"os/exec"
)

func (a *ADB) screencapCommand() (draw.Image, error) {
	cmd := execScreencap()

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(bytes.NewReader(out))
	if err != nil {
		return nil, err
	}

	dimg, ok := img.(draw.Image)

	if !ok {
		return nil, fmt.Errorf("Not a drawable image type")
	}

	return dimg, nil
}

func execScreencap() *exec.Cmd {
	return exec.Command("adb", "exec-out", "screencap", "-p")
}
