object Luhn {
    fun isValid(numberIdString: String): Boolean {
        val numberIdStringStripped = numberIdString.replace("\\s".toRegex(), "")

        if (numberIdStringStripped.length <= 1 || numberIdStringStripped.contains("\\D+".toRegex())) {
            return false
        }
        val numberArr = numberIdStringStripped.map { it.toString().toInt() }
                                                            .toMutableList()
        var numTimeTwo:Int
        for (i in numberArr.size - 2 downTo 0 step 2) {

            numTimeTwo = numberArr[i] * 2
            if (numTimeTwo > 9) numTimeTwo -= 9
            numberArr.removeAt(i)
            numberArr.add(i, numTimeTwo)
        }
        return numberArr.sum() % 10 == 0
    }

}
