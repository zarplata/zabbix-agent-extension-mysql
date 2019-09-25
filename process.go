package main

import (
	"regexp"
	"strconv"
)

func calcProcessStats(processes []map[string]string) []map[string]string {

	reUser := regexp.MustCompile("(system user)|(repl)")
	reCommand := regexp.MustCompile("Sleep")

	var (
		max    int
		count  int = 0
		result []map[string]string
	)

	for _, process := range processes {
		if reUser.MatchString(process["User"]) {
			continue
		}
		if !reCommand.MatchString(process["Command"]) {
			count++
			time, _ := strconv.Atoi(process["Time"])
			if time > max {
				max = time
			}
		}
	}

	stats := make(map[string]string)
	stats["processlist_count"] = strconv.Itoa(count)
	stats["query_max_time"] = strconv.Itoa(max)
	result = append(result, stats)

	return result
}
