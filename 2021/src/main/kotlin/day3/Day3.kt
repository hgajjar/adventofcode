package day3

import InputReader

class Day3(private var inputReader: InputReader) {
    private val input = "day3.input"

    fun run() {
        println("Day3 - Part I: ${this.part1()}")
    }

    private fun part1(): Int {
        return this.generateGammaRate() * this.generateEpsilonRate()
    }

    private fun generateGammaRate(): Int {
        val condition = fun (v: Vector): Int { val zeros = v.count { i -> i == 0 }; val ones = v.count { e -> e == 1 }; return if (zeros > ones) 0 else 1 }
        return this.generateRate(condition)
    }

    private fun generateEpsilonRate(): Int {
        val condition = fun (v: Vector): Int { val zeros = v.count { i -> i == 0 }; val ones = v.count { e -> e == 1 }; return if (zeros > ones) 1 else 0 }
        return this.generateRate(condition)
    }

    private fun generateRate(condition: (v: Vector) -> Int): Int {
        return this
            .parseInput()
            .transpose()
            .map(condition)
            .joinToString(separator = "")
            .toInt(2)
    }

    private fun parseInput(): Matrix {
         return inputReader
            .read(input)
            .map { it.chunked(1).map{ i -> i.toInt() }.toIntArray() }
             .toTypedArray()
    }
}

typealias Vector = IntArray
typealias Matrix = Array<Vector>

fun Matrix.transpose(): Matrix {
    val rows = this.size
    val cols = this[0].size
    val trans = Matrix(cols) { Vector(rows) }
    for (i in 0 until cols) {
        for (j in 0 until rows) trans[i][j] = this[j][i]
    }
    return trans
}