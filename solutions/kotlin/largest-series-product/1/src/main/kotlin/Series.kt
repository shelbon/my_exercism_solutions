class Series(private val s: String) {
    init {
        require(s.matches(Regex("^[0-9]*")))
    }

    fun getLargestProduct(numberOfDigit: Int): Int {
        require(numberOfDigit in 0..s.length)
        if (s.isEmpty() || numberOfDigit == 0) {
            return 1
        }
        return s.map { it.toString().toInt() }
                .windowed(numberOfDigit) {
                    it.reduce { acc, i -> acc * i }
                }.max()!!
    }

}
