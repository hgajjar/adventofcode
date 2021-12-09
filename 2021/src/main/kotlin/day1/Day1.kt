package day1

import InputReader

class Day1(private var inputReader: InputReader) {
    private val input = "day1.input"

    fun run() {
        println("Day1 - Part I: ${this.part1()}")
        println("Day1 - Part II: ${this.part2()}")
    }

    private fun part1(): Int {
        val measurements = inputReader.read(input)

        return measurements
            .filterIndexed { i, e -> i != 0 && e.toInt() > measurements[i-1].toInt() }
            .count()
    }

    private fun part2(): Int {
        val sumOfMeasurements = inputReader
            .read(input)
            .map { it.toInt() }
            .windowed(3, 1)
            .map { it.sum() }

        return sumOfMeasurements
            .filterIndexed { i, e -> i != 0 && e > sumOfMeasurements[i-1] }
            .count()
    }
}