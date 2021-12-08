package day13

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readCurrentTimeAndBusIds(fileName string) (time, busCollection) {
	contents, _ := ioutil.ReadFile(fileName)
	lines := strings.Split(string(contents), "\n")
	mins, _ := strconv.Atoi(string(lines[0]))
	currentTime := time{mins}
	busColl := busCollection{}

	for _, id := range strings.Split(string(lines[1]), ",") {
		if id == "x" {
			id = "0"
		}
		intId, _ := strconv.Atoi(id)
		busColl = append(busColl, bus{intId})
	}

	return currentTime, busColl
}
