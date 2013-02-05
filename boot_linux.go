// +build linux

package shout

import (
	"os"
)

func init() {
	// Check if there is an the external command to write.
	_, err := os.Stat(CMD_WRITE)
	if os.IsNotExist(err) {
		return
	}

	err = exec.Command(CMD_WRITE, "--ping").Run()
	if _, ok := err.(*exec.ExitError); !ok {

		//	if _, ok, _ := shout.Run(CMD_WRITE + " --ping"); ok {
		USE_CMD_WRITE = true
	}
}
