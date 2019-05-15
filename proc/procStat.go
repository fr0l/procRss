package proc

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Process struct {
	PID  int
	PPID int
	RSS  int
}

func Stats() []Process {
	lines := processLines()
	return parsePsStats(lines)
}

func CalculateRss(stats []Process, pid int) (int, int) {
	var rssPid = -1
	var rssDescendants = 0

	for _, stat := range stats {
		if stat.PID == pid {
			rssPid = stat.RSS
		} else if stat.PPID == pid {
			childRss, descendantsRss := CalculateRss(stats, stat.PID)
			rssDescendants += childRss + descendantsRss
		}
	}

	return rssPid, rssDescendants
}

func processLines() []string {
	out, err := exec.
		Command("ps", "ax", "-o", "pid,ppid,rss").
		Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(out), "\n")
	return lines[1:]
}

func parsePsStats(lines []string) []Process {
	var ps = make([]Process, len(lines))

	for i, line := range lines {
		if line == "" {
			continue
		}

		st := strings.Fields(line)
		pid := toInt(st[0])
		ppid := toInt(st[1])
		rss := toInt(st[2])
		ps[i] = Process{pid, ppid, rss}
	}

	return ps
}

func toInt(s string) int {
	i, e := strconv.Atoi(s)

	if e != nil {
		log.Fatal(e)
	}

	return i
}
