package main

import (
	"regexp"
	"strconv"
)

func calcProcessStats(processes []map[string]string) []map[string]string {

	reUser := regexp.MustCompile("(system user)|(repl)|(root)")
	reCommand := regexp.MustCompile("Sleep")

	var (
		max    int
		result []map[string]string
	)

	for _, process := range processes {
		if !reUser.MatchString(process["User"]) && !reCommand.MatchString(process["Command"]) {
			time, _ := strconv.Atoi(process["Time"])
			if time > max {
				max = time
			}
		}
	}

	stats := make(map[string]string)
	stats["processlist_count"] = strconv.Itoa(len(processes))
	stats["query_max_time"] = strconv.Itoa(max)
	result = append(result, stats)

	return result
}
