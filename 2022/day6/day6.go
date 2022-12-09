package day6

import (
	_ "embed"
	"fmt"
	"github.com/samber/lo"
)

var (
	//go:embed input/sample.txt
	sampleInput string

	//go:embed input/actual.txt
	actualInput string
)

const (
	packetLength int = 14
)

func Execute() {
	fmt.Println(findStartOfPacketMarker(sampleInput))
	fmt.Println(findStartOfPacketMarker(actualInput))
}

func findStartOfPacketMarker(input string) int {
	for i := 0; i < (len(input) - packetLength + 1); i++ {
		if len(lo.Uniq[byte]([]byte(input[i:i+packetLength]))) == packetLength {
			return i + packetLength
		}
	}

	return 0
}
