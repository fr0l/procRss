package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"procRss/proc"
	"strconv"
)

func main() {
	pids, err := parsePID(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	stats := proc.Stats()

	for _, pid := range pids {
		showStats(stats, pid)
	}
}

func showStats(stats []proc.Process, pid int) {
	rssPid, rssDescendants := proc.CalculateRss(stats, pid)

	if rssPid == -1 {
		log.Fatalf("Unknown process by pid %d", pid)
	}

	pidRamMB := rssPid / 1024
	descendantsRamMB := rssDescendants / 1024
	totalRamMB := pidRamMB + descendantsRamMB

	fmt.Printf("PID: %7d, process RSS %7d MB, descendants RSS %7d MB, total RSS %7d MB\n", pid, pidRamMB, descendantsRamMB, totalRamMB)
}

func parsePID(args []string) ([]int, error) {
	if len(args) == 1 {
		return []int{}, errors.New("pass PID or PIDs as arguments\n")
	}
	fmt.Printf("%s\n", args)

	var pids = make([]int, len(args)-1)
	for i, p := range args[1:] {
		pid, err := strconv.Atoi(p)
		if err != nil {
			log.Fatal(err)
		}
		pids[i] = pid
	}
	return pids, nil
}
