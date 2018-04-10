package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	hierr "github.com/reconquest/hierr-go"
)

var (
	statName  string
	statValue string
	result    []map[string]string
)

func getRow(dsn, query string) (*sql.Rows, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, hierr.Errorf(err, "can't open %s.", dsn)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func getGlobalStats(query, dsn string) ([]map[string]string, error) {

	stats := make(map[string]string)

	rows, err := getRow(dsn, query)
	if err != nil {
		return nil, hierr.Errorf(err, "can't do query %s.", query)
	}

	for rows.Next() {
		err = rows.Scan(&statName, &statValue)
		if err != nil {
			return nil, hierr.Errorf(err, "can't get value from row.")
		}
		stats[statName] = statValue
	}

	result = append(result, stats)

	if len(result) == 0 {
		return nil, fmt.Errorf("return null value in query %s", query)
	}

	return result, nil
}

func getStats(query, dsn string) ([]map[string]string, error) {

	rows, err := getRow(dsn, query)
	if err != nil {
		return nil, hierr.Errorf(err, "can't do query %s.", query)
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, hierr.Errorf(err, "can't get columns.")
	}
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, hierr.Errorf(err, "can't get value.")
		}

		stats := make(map[string]string)
		for i, value := range values {
			statName = columns[i]
			statValue = string(value)
			stats[statName] = statValue
		}
		result = append(result, stats)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("return null value in query %s", query)
	}

	return result, nil
}
