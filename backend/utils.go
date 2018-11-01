package main

import "log"

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
