package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
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
	//injection := "echo \"" + text + "\" > " + procInput
	//cmd := exec.Command("/bin/bash", "-c", injection)
	//_, err := cmd.CombinedOutput()
	//if err != nil {
	//		return
	//	}
	f, err := os.OpenFile(procInput, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%s", text)
	return
}

func ree(bashProcs []string) {
	for _, pid := range bashProcs {
		injectString(pid, "reee")
	}
}

func smashMouth(bashProcs []string) {
	allstar := []string{"SomeBODY once told me the world is gonna roll me",
		"I ain't the sharpest tool in the shed",
		"She was looking kind of dumb with her finger and her thumb",
		"In the shape of an \"L\" on her forehead",
		"WEEELLLL the years start coming and they don't stop coming",
		"Fed to the rules and I hit the groud running",
		"Didn't make sense not to live for fun",
		"Your brain gets smart but your head gets dumb",
		"So much to do, so much to see",
		"So what's wrong with taking the back streets?",
		"You'll never know if you don't go",
		"you'll never shine if you don't glow",
		"Hey now, you're an all-star, get your game on, go play",
		"Hey now, you're a rock star, get the show on, get paid",
		"And all that glitters is gold",
		"Only shootin STARS break the Moooolld",
		"It's a cool place and they it gets colder",
		"You're bundled up now, wait til you get older",
		"But the meteor men beg to differ",
		"Judging by the hole in the satellite picture",
		"The ice we skate is gettin' pretty thin",
		"The water's getting warm so you might as well swim",
		"My world's on fire, how about yours?",
		"That's the way I like it and I never get bored",
		"Hey now, you're an all-star, get your game on, go play",
		"Hey now, you're a rock star, get the show on, get paid",
		"And all that glitters is gold",
		"Only shootin STARS break the Moooolld",
		"Hey now, you're an all-star, get your game on, go play",
		"Hey now, you're a rock star, get the show on, get paid",
		"And all that glitters is gold",
		"Only shootin STARS...",
		"Somebody once asked could I spare some change for gas?",
		"I need to get myself away from this place",
		"I said yep what a concept",
		"I could use a littel fuel myself",
		"And we could all use a little change",
		"WEEELLLL the years start coming and they don't stop coming",
		"Fed to the rules and I hit the groud running",
		"Didn't make sense not to live for fun",
		"Your brain gets smart but your head gets dumb",
		"So much to do, so much to see",
		"So what's wrong with taking the back streets?",
		"You'll never know if you don't go (go!)",
		"you'll never shine if you don't glow",
		"Hey now, you're an all-star, get your game on, go play",
		"Hey now, you're a rock star, get the show on, get paid",
		"And all that glitters is gold",
		"Only shootin STARS break the Moooolld",
		"And all that glitters is gold",
		"Only shootin STARS break the Moooolld"}

	for _, lyric := range allstar {
		for _, pid := range bashProcs {
			injectString(pid, lyric)
		}
		time.Sleep(30 * time.Second)
	}

}

func main() {
	arg := os.Args[1]
	bashProcs := scrapeProc()
	if arg == "-default" {
		ree(bashProcs)
	}
	if arg == "-smashmouth" {
		smashMouth(bashProcs)
	}

}
