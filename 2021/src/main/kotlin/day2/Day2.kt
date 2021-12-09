package day2

import InputReader

const val FORWARD = "forward"
const val UP = "up"
const val DOWN = "down"

class Day2(private var inputReader: InputReader) {
    private val input = "day2.input"

    fun run() {
        println("Day2 - Part I: ${this.part1()}")
        println("Day2 - Part II: ${this.part2()}")
    }

    private fun part1(): Int {
        var horizontalPos = 0
        var depth = 0

        this
            .parseInput()
            .forEach {
                when (it.direction) {
                    FORWARD -> horizontalPos += it.units
                    DOWN -> depth += it.units
                    UP -> depth -= it.units
                }
            }

        return horizontalPos * depth
    }

    private fun part2(): Int {
        var horizontalPos = 0
        var depth = 0
        var aim = 0

        this
            .parseInput()
            .forEach {
                when(it.direction) {
                    FORWARD -> {
                        horizontalPos += it.units
                        depth += aim * it.units
                    }
                    DOWN -> aim += it.units
                    UP -> aim -= it.units
                }
            }

        return horizontalPos * depth
    }

    private fun parseInput(): List<Command> {
        return inputReader
            .read(input)
            .map { val parts = it.split(" "); Command(parts[0], parts[1].toInt()) }
    }
}

data class Command(val direction: String, val units: Int)