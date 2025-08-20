
enum class Classification {
    DEFICIENT, PERFECT, ABUNDANT
}

fun checkNumber(number: Int) {
    if (number <= 0) {
        throw RuntimeException()
    }
}

fun classify(naturalNumber: Int): Classification {
    checkNumber(naturalNumber)
    val properDivisor = mutableListOf<Int>()
    for (number in 1 until naturalNumber) {
        if (naturalNumber.rem(number) == 0) {
            properDivisor.add(number)
        }
    }

    if (properDivisor.isEmpty()) return Classification.DEFICIENT

    val sum = properDivisor.reduce { acc: Int, i: Int -> acc + i }
    return when {
        sum == naturalNumber -> Classification.PERFECT
        sum > naturalNumber -> Classification.ABUNDANT
        else -> Classification.DEFICIENT
    }

}
