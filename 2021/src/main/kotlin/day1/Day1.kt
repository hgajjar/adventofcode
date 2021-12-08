package day1

import InputReader

class Day1(private var inputReader: InputReader) {
    private val input1 = "day1-I.input"

    fun run() {
        println("Day1 - Part I: ${this.part1()}")
    }

    private fun part1(): Int {
        val input = inputReader.read(input1)

        return input
            .filterIndexed { i, e -> i != 0 && e.toInt() > input[i-1].toInt() }
            .count()
    }
}