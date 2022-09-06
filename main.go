package main

import (
	"go-playground/plays"
)

func main() {
	entrances := []string{
		"result1",
		"result2",
		"result3",
		"result4",
		"result5",
	}
	result, err := plays.WaitTimeout(entrances)
	if err != nil {
		println(err.Error())
		return
	}

	for _, str := range result {
		println(str)
	}
}
