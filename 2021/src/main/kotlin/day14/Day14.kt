package day14

import InputReader
import kotlin.io.path.bufferedReader
import kotlin.io.path.writeText

class Day14(private var inputReader: InputReader) {
    private val input = "day14.input"

    fun run() {
        println("Day14 - Part I: ${this.calculate(10)}")
        println("Day14 - Part II: ${this.calculate(40)}")
    }

    private fun calculate(iterations: Int): Int {
        val (initialTemplate, insertionRules) = parseInput()
        var template = initialTemplate;

        val tmpFile = kotlin.io.path.createTempFile()
        println(tmpFile.toAbsolutePath())

        repeat(iterations) { index ->
            val begin = System.currentTimeMillis()

            tmpFile.writeText(template)

            var partitionSize = 20
            if (index > 20) {
                partitionSize = index * 4
            }

            println("Start Memory: " + (Runtime.getRuntime().totalMemory() - Runtime.getRuntime().freeMemory())/1024/1024 + "mb")


            var chunk = charArrayOf()
            println(tmpFile.bufferedReader().read(chunk, 0, 2))

            template = template.windowedSequence(partitionSize, partitionSize - 1, true).mapIndexed(fun (index: Int, template: String): String {
                val t = this.polymerize(template, insertionRules)
                if (index == 0) {
                    return t
                }
                return t.drop(1)
            }).joinToString(separator = "")



//            template = this.polymerize(template, insertionRules)

//            println(template)

            val end = System.currentTimeMillis()

            println("Step: $index: ; Memory: " + (Runtime.getRuntime().totalMemory() - Runtime.getRuntime().freeMemory())/1024/1024 + "mb; Elapsed time: ${(end-begin) / 1000}sec")
        }

        val allChars = template.toCharArray().distinct()

        val maxChar = allChars.maxWithOrNull(compareBy { e -> template.count { it == e } })
        val minChar = allChars.minWithOrNull(compareBy { e -> template.count { it == e } })

        return template.count { it == maxChar } - template.count { it == minChar }
    }

    private fun polymerize(template: String, insertionRules: List<InsertionRule>): String {
        val hasMatch: (InsertionRule,String) -> Boolean = { rule: InsertionRule, pair: String -> pair[0] == rule.left && pair[1] == rule.right  }

        return template
            .windowedSequence(2, 1)
            .mapIndexed(fun (index: Int, pair: String): String {
                val matchingRule = insertionRules.find { hasMatch(it, pair) }
                if (index == 0) {
                    return listOf(pair[0], matchingRule?.insertion, pair[1]).joinToString(separator = "")
                }
                return listOf(matchingRule?.insertion, pair[1]).joinToString(separator = "")
            })
            .joinToString(separator = "")
    }

    private fun parseInput(): Pair<Template, List<InsertionRule>> {
        val data = inputReader.read(input)
        val template = data[0]
        val insertionRules = data.drop(2)
            .map { val tokens = it.split(" -> "); InsertionRule(tokens[0][0], tokens[0][1], tokens[1][0]) }

        return Pair(template, insertionRules)
    }
}

typealias Template = String

data class InsertionRule(val left: Char, val right: Char, val insertion: Char)