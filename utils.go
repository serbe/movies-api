package main

import (
	"log"
	"strconv"
)

func toInt(num string) int {
	id, err := strconv.Atoi(num)
	if err != nil {
		log.Println("toInt", err)
	}
	return id
}

func toInt64(num string) int64 {
	id, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		log.Println("toInt64", err)
	}
	return id
}
