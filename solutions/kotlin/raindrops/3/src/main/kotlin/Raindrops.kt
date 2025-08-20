object Raindrops {
    fun convert(number: Int): String {

        var factorsAndOutputs = sortedMapOf(3 to "Pling", 5 to "Plang", 7 to "Plong")
        var raindrops = factorsAndOutputs.filter { asFactor(it.key, number) }.values.joinToString("")

        return raindrops.ifEmpty { number.toString() }
    }

    private fun asFactor(factor: Int, number: Int): Boolean = number % factor == 0

}

