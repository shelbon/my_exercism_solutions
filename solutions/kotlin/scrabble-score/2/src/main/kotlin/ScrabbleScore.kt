object ScrabbleScore {

    fun scoreWord(word: String): Int {

        val letterValue = mapOf("AEIOULNRST" to 1, "DG" to 2, "BCMP" to 3,
                                               "FHVWY" to 4, "K" to 5, "JX" to 8, "QZ" to 10)

        return word.groupingBy {it}.eachCount().entries.sumBy {it.value*
                letterValue.filterKeys { it2->it2.contains(it.key.toUpperCase()) }.values.first()}
    }

}
