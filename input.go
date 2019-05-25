package adb

import (
	"math/rand"
	"os/exec"
	"strconv"
)

const (
	evAbs = 0x0003
	evKey = 0x0001
	evSyn = 0x0000
)

const (
	absmTPositionX  = 0x0035
	absmTPositionY  = 0x0036
	absmTPressure   = 0x003a
	absmTTouchMajor = 0x0030
	absmTTouchMinor = 0x0031
	absmTTrackingID = 0x0039
	btnTouch        = 0x014a
	synReport       = 0x0000
)

const (
	down          = 0x00000001
	majorTouchMax = 8
	maxed         = 0xffffffff
	minorTouchMax = 6
	pressureMax   = 25
	up            = 0x00000000
	zero          = 0x00000000
)

var trackingID int

func (a *ADB) touchCommand(x int, y int) error {
	var cmdBuffer [12]*exec.Cmd

	cmdBuffer[0] = execSendEvent(evKey, btnTouch, down)

	cmdBuffer[1] = execSendEvent(evAbs, absmTTrackingID, trackingID)
	trackingID++

	cmdBuffer[2] = execSendEvent(evAbs, absmTPositionX, x)

	cmdBuffer[3] = execSendEvent(evAbs, absmTPositionY, y)

	major, minor := generateMajorMinorPair()
	cmdBuffer[4] = execSendEvent(evAbs, absmTTouchMajor, major)

	cmdBuffer[5] = execSendEvent(evAbs, absmTTouchMinor, minor)

	cmdBuffer[6] = execSendEvent(evAbs, absmTPressure, rand.Intn(pressureMax)+1)

	cmdBuffer[7] = execSendEvent(evSyn, synReport, zero)

	cmdBuffer[8] = execSendEvent(evAbs, absmTPressure, zero)

	cmdBuffer[9] = execSendEvent(evAbs, absmTTrackingID, maxed)

	cmdBuffer[10] = execSendEvent(evKey, btnTouch, up)

	cmdBuffer[11] = execSendEvent(evSyn, synReport, zero)

	for _, cmd := range cmdBuffer {
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func generateMajorMinorPair() (int, int) {
	minor := rand.Intn(minorTouchMax) + 1
	major := rand.Intn(majorTouchMax+1-minor) + minor
	return minor, major
}

func execSendEvent(event int, key int, value int) *exec.Cmd {
	return exec.Command("adb", "shell", "sendevent", "/dev/input/event2", strconv.Itoa(event), strconv.Itoa(key), strconv.Itoa(value))
}
