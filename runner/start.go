package runner

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	startChannel chan string
	stopChannel  chan bool
	mainLog      logFunc
	watcherLog   logFunc
	runnerLog    logFunc
	buildLog     logFunc
	appLog       logFunc
)

// apparently this shouldn't even be necessary, but I don't know enough about how channels work.
func flushEvents() {
	for {
		select {
		case eventName := <-startChannel:
			eventName += ""
		default:
			return
		}
	}
}

func start() {
	loopIndex := 0
	buildDelay := buildDelay()

	started := false

	go func() {
		for {
			loopIndex++

			// Flush. I have no idea how channels work, I'm just adopting what was here
			eventName := <-startChannel
			eventName += ""

			time.Sleep(buildDelay * time.Millisecond)

			flushEvents()

			errorMessage, ok := build()
			if !ok {
				mainLog("Build failed:\n\033[0m%s\033[0m", errorMessage)
			} else {
				if started {
					stopChannel <- true
				}
				run()
			}

			started = true
		}
	}()
}

func init() {
	startChannel = make(chan string, 1000)
	stopChannel = make(chan bool)
}

func initLogFuncs() {
	mainLog = newLogFunc("main")
	watcherLog = newLogFunc("watcher")
	runnerLog = newLogFunc("runner")
	buildLog = newLogFunc("build")
	appLog = newLogFunc("app")
}

func setEnvVars() {
	os.Setenv("DEV_RUNNER", "1")
	wd, err := os.Getwd()
	if err == nil {
		os.Setenv("RUNNER_WD", wd)
	}

	for k, v := range settings {
		key := strings.ToUpper(fmt.Sprintf("%s%s", envSettingsPrefix, k))
		os.Setenv(key, v)
	}
}

// Watches for file changes in the root directory.
// After each file system event it builds and (re)starts the application.
func Start() {
	initLogFuncs()
	initLimit()
	initSettings()
	setEnvVars()
	watch()
	start()
	startChannel <- "/"

	<-make(chan int)
}
