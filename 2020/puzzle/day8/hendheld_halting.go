package day8

type Instruction struct {
	Operation string
	Argument  int32
}

func FindAccWhenRecursionDetected(instructionSet []Instruction) int {
	accVal, _ := findRecursion(instructionSet)
	return accVal
}

func FindAccAfterInstructionsFixed(instructionSet []Instruction) int {
	isRecursive := true
	accVal := 0

	originalInstructionSet := instructionSet
	updatedInstructionIndex := -1

	for isRecursive {
		accVal, isRecursive = findRecursion(instructionSet)

		if !isRecursive {
			return accVal
		}

		// update instruction set
		instructionSet, updatedInstructionIndex = updateInstructionSet(originalInstructionSet, updatedInstructionIndex)
	}
	return accVal
}

func updateInstructionSet(instructionSet []Instruction, lastUpdatedInstructionIndex int) ([]Instruction, int) {
	for index, instruction := range instructionSet {
		if index > lastUpdatedInstructionIndex {
			switch instruction.Operation {
			case "nop":
				return updateInstruction(instruction, instructionSet, index, "jmp")
			case "jmp":
				return updateInstruction(instruction, instructionSet, index, "nop")
			}
		}
	}
	return instructionSet, lastUpdatedInstructionIndex
}

func updateInstruction(instruction Instruction, instructionSet []Instruction, index int, operation string) ([]Instruction, int) {
	tmp := Instruction{operation, instruction.Argument}
	tmpSet := []Instruction{}
	if index == 0 {
		tmpSet = append(tmpSet, tmp)
	} else {
		tmpSet = append(tmpSet, instructionSet[:index]...)
		tmpSet = append(tmpSet, tmp)
	}

	tmpSet = append(tmpSet, instructionSet[index+1:]...)
	return tmpSet, index
}

func findRecursion(instructionSet []Instruction) (int, bool) {
	accVal := 0
	iHistory := make(map[int]bool)

	for i := 0; i < len(instructionSet); i++ {
		if _, ok := iHistory[i]; ok {
			return accVal, true
		}
		iHistory[i] = true
		switch instructionSet[i].Operation {
		case "nop":
			break
		case "acc":
			accVal += int(instructionSet[i].Argument)
		case "jmp":
			i += int(instructionSet[i].Argument) - 1
		default:
			panic("Instruction not supported!")
		}
	}

	return accVal, false
}
