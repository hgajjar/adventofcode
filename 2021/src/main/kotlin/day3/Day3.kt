package day3

import InputReader

class Day3(private var inputReader: InputReader) {
    private val input = "day3.input"

    fun run() {
        println("Day3 - Part I: ${this.part1()}")
        println("Day3 - Part II: ${this.part2()}")
    }

    private fun part1(): Int {
        return this.generateGammaRate() * this.generateEpsilonRate()
    }

    private fun part2(): Int {
        val oxygenGeneratorRetingCriteria = fun(ones: Int, zeros: Int, col: Int, report: Matrix): Matrix {
            return if (ones > zeros) {
                report.filter { it[col] == 1 }.toTypedArray()
            } else if (zeros > ones) {
                report.filter { it[col] == 0 }.toTypedArray()
            } else {
                report.filter { it[col] == 1 }.toTypedArray()
            }
        }

        val co2ScrubberRatingCriteria = fun(ones: Int, zeros: Int, col: Int, report: Matrix): Matrix {
            return if (ones > zeros) {
                report.filter { it[col] == 0 }.toTypedArray()
            } else if (zeros > ones) {
                report.filter { it[col] == 1 }.toTypedArray()
            } else {
                report.filter { it[col] == 0 }.toTypedArray()
            }
        }

        val o2 = this.filterByBitCriteria(this.parseInput(), 0, oxygenGeneratorRetingCriteria)
        val co2 = this.filterByBitCriteria(this.parseInput(), 0, co2ScrubberRatingCriteria)

        return o2 * co2
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

    private fun filterByBitCriteria(report: Matrix, currentCol: Int, criteria: (ones: Int, zeros: Int, col: Int, report: Matrix) -> Matrix): Int {
        var filteredReport: Matrix
        for (col in currentCol until  report[0].size) {
            var zeros = 0
            var ones = 0
            for (element in report) {
                if (element[col] == 0) {
                    zeros++
                } else {
                    ones++
                }
            }

            filteredReport = criteria(ones, zeros, col, report)

            return if (filteredReport.size != 1) {
                this.filterByBitCriteria(filteredReport, col+1, criteria)
            } else {
                filteredReport[0].joinToString(separator = "").toInt(2)
            }
        }

        return 0
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