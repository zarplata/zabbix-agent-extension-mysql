package main

import (
	"encoding/json"
	"fmt"

	hierr "github.com/reconquest/hierr-go"
)

func discovery(dsn string) error {
	discoveryData := make(map[string][]map[string]string)
	var discoveredItems []map[string]string

	statsGalera, err := getGlobalStats(queryGalera, dsn)
	if err != nil {
		return hierr.Errorf(err, "can't get galera stats.")
	}
	if len(statsGalera) == 1 {
		discoveredItem := make(map[string]string)
		discoveredItem["{#TYPE}"] = "galera"
		discoveredItems = append(discoveredItems, discoveredItem)
	}

	statsSlave, err := getStats(querySlave, dsn)
	if err != nil {
		return hierr.Errorf(err, "can't get slave stats.")
	}
	if len(statsSlave) > 1 {
		discoveredItem := make(map[string]string)
		discoveredItem["{#TYPE}"] = "slave"
		discoveredItems = append(discoveredItems, discoveredItem)
	}

	discoveryData["data"] = discoveredItems

	out, err := json.Marshal(discoveryData)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", out)
	return nil
}
