package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
		injectString(pid, "REEEEEEEEEEEEEEEEEEEEEEEEEE")
	}
}

func fb(bashProcs []string) {
	for _, pid := range bashProcs {
		injectString(pid, "Move slow and repair things.")
	}
}

func bees(bashProcs []string) {
	for _, pid := range bashProcs {
		injectString(pid, "bzzzzzzzzz I'm a bee stuck in your terminal bzzzz bzzzz")
	}
}

func custom(bashProcs []string, text string) {
	for _, pid := range bashProcs {
		injectString(pid, text)
	}
}

func smashMouth(bashProcs []string) {
	sleeptime := 30
	if len(os.Args) == 3 {
		sleeptime, _ = strconv.Atoi(os.Args[2])
	}
	allstar := []string{"SomeBODY once told me the world is gonna roll me\n",
		"I ain't the sharpest tool in the shed\n",
		"She was looking kind of dumb with her finger and her thumb\n",
		"In the shape of an \"L\" on her forehead\n",
		"WEEELLLL the years start coming and they don't stop coming\n",
		"Fed to the rules and I hit the groud running\n",
		"Didn't make sense not to live for fun\n",
		"Your brain gets smart but your head gets dumb\n",
		"So much to do, so much to see\n",
		"So what's wrong with taking the back streets?\n",
		"You'll never know if you don't go\n",
		"you'll never shine if you don't glow\n",
		"Hey now, you're an all-star, get your game on, go play\n",
		"Hey now, you're a rock star, get the show on, get paid\n",
		"And all that glitters is gold\n",
		"Only shootin STARS break the Moooolld\n",
		"It's a cool place and they it gets colder\n",
		"You're bundled up now, wait til you get older\n",
		"But the meteor men beg to differ\n",
		"Judging by the hole in the satellite picture\n",
		"The ice we skate is gettin' pretty thin\n",
		"The water's getting warm so you might as well swim\n",
		"My world's on fire, how about yours?\n",
		"That's the way I like it and I never get bored\n",
		"Hey now, you're an all-star, get your game on, go play\n",
		"Hey now, you're a rock star, get the show on, get paid\n",
		"And all that glitters is gold\n",
		"Only shootin STARS break the Moooolld\n",
		"Hey now, you're an all-star, get your game on, go play\n",
		"Hey now, you're a rock star, get the show on, get paid\n",
		"And all that glitters is gold\n",
		"Only shootin STARS...\n",
		"Somebody once asked could I spare some change for gas?\n",
		"I need to get myself away from this place\n",
		"I said yep what a concept\n",
		"I could use a littel fuel myself\n",
		"And we could all use a little change\n",
		"WEEELLLL the years start coming and they don't stop coming\n",
		"Fed to the rules and I hit the groud running\n",
		"Didn't make sense not to live for fun\n",
		"Your brain gets smart but your head gets dumb\n",
		"So much to do, so much to see\n",
		"So what's wrong with taking the back streets?\n",
		"You'll never know if you don't go (go!)\n",
		"you'll never shine if you don't glow\n",
		"Hey now, you're an all-star, get your game on, go play\n",
		"Hey now, you're a rock star, get the show on, get paid\n",
		"And all that glitters is gold\n",
		"Only shootin STARS break the Moooolld\n",
		"And all that glitters is gold\n",
		"Only shootin STARS break the Moooolld\n"}

	for _, lyric := range allstar {
		for _, pid := range bashProcs {
			injectString(pid, lyric)
		}
		time.Sleep(time.Duration(sleeptime) * time.Second)
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
	if arg == "-fb" {
		fb(bashProcs)
	}
	if arg == "-bees" {
		bees(bashProcs)
	}
	if arg == "-custom" {
		st := os.Args[2]
		custom(bashProcs, st)

	}

}
