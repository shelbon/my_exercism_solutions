import kotlin.math.pow

class Squares(private val maxNaturalNumber: Int) {


    fun squareOfSum(): Int {
        var sum = 0

        (1..maxNaturalNumber).forEach { sum += it }

        return sum.toDouble().pow(2).toInt()
    }

    fun sumOfSquares(): Int {
        var sum = 0

        (1..maxNaturalNumber).forEach { sum += it.toDouble().pow(2).toInt() }

        return sum
    }

    fun difference(): Int {
        return squareOfSum() - sumOfSquares()
    }

}