object ScrabbleScore {

    fun scoreWord(word: String): Int {

        val letterValue = mapOf("AEIOULNRST" to 1, "DG" to 2, "BCMP" to 3,
                                               "FHVWY" to 4, "K" to 5, "JX" to 8, "QZ" to 10)
        var score = 0
        var charCount = word.groupingBy { it }.eachCount()
        for ((char, number) in charCount) {
            score += number * letterValue.filterKeys { it.contains(char.toUpperCase()) }.values.first()
        }
        return score
    }

}
