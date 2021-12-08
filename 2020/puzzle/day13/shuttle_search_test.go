package day13

import "testing"

var (
	testCurrentTime, testBusColl = readCurrentTimeAndBusIds("../../data/test/day13.txt")
	currentTime, busColl         = readCurrentTimeAndBusIds("../../data/day13.txt")
)

func TestFindProductOfNextBusAndItsWaitTime(t *testing.T) {
	testOp := findProductOfNextBusAndItsWaitTime(testCurrentTime, testBusColl)

	if testOp != 295 {
		t.Errorf("Expected product: %d, found %d", 295, testOp)
	}

	op := findProductOfNextBusAndItsWaitTime(currentTime, busColl)

	if op != 136 {
		t.Errorf("Expected product: %d, found %d", 136, op)
	}
}

func TestFindEarliestTimeWhereBusesDepartEveryMinute(t *testing.T) {
	testOp := findEarliestTimeWhereBusesDepartEveryMinute(testBusColl)

	if testOp != 1068781 {
		t.Errorf("Expected earliest timestamp where buses depart sequencially is: %d, found %d", 1068781, testOp)
	}

	op := findEarliestTimeWhereBusesDepartEveryMinute(busColl)

	if op != 305068317272992 {
		t.Errorf("Expected earliest timestamp where buses depart sequencially is: %d, found %d", 305068317272992, op)
	}
}
