class Series(private val input: String) {
    init {
        require(input.matches(Regex("^[0-9]*")))
    }


    fun getLargestProduct(numberOfDigit: Int): Long {
        require(numberOfDigit in 0..input.length)
        if (input.isEmpty() || numberOfDigit == 0) {
            return 1
        }
        return input.map { it.toString().toLong() }
            .windowed(numberOfDigit) {
                it.reduce { acc, i -> acc * i }
            }.maxOrNull()?:-1;
    }
}
