package runner

import (
	// "io"
	// "io/ioutil"
	// "os"
	"os/exec"
)

func build() (string, bool) {
	buildLog("Fetching dependencies.")

	out, err := exec.Command("go", "get").CombinedOutput()
	buildLog("\033[0m" + string(out))
	if err != nil {
		return err.Error(), false
	}

	buildLog("Building.")

	out, err = exec.Command("go", "build", "-o", buildPath(), root()).CombinedOutput()

	buildLog("\033[0m" + string(out))
	if err != nil {
		return err.Error(), false
	}

	return "", true
}
