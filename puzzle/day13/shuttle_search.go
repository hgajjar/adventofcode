package day13

type time struct {
	mins int
}

type bus struct {
	ID int
}

type busCollection []bus

func findProductOfNextBusAndItsWaitTime(currentTime time, busColl busCollection) int {
	earliestBus := bus{}
	minWaitTime := time{}
	for _, bus := range busColl {
		if bus.ID == 0 {
			continue
		}
		waitTime := bus.ID - (currentTime.mins % bus.ID)
		if waitTime < minWaitTime.mins || minWaitTime.mins == 0 {
			minWaitTime.mins = waitTime
			earliestBus = bus
		}
	}

	return earliestBus.ID * minWaitTime.mins
}

func findEarliestTimeWhereBusesDepartEveryMinute(busColl busCollection) int {
	time := time{}
	runningProduct := 1
	for i, bus := range busColl {
		if bus.ID == 0 {
			continue
		}
		for (time.mins+i)%bus.ID != 0 {
			time.mins += runningProduct
		}
		runningProduct *= bus.ID
	}

	return time.mins
}
