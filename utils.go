package main

import (
	"log"
	"strconv"
)

func toInt(num string) int64 {
	id, err := strconv.ParseInt(num, 10, 64)
	log.Println("toInt", err)
	return id
}
