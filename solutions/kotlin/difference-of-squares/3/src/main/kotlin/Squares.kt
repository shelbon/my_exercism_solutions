import kotlin.math.pow

class Squares(private val maxNaturalNumber: Int) {


    fun squareOfSum(): Int {
        return  (1..maxNaturalNumber).sum()*(1..maxNaturalNumber).sum()
    }

    fun sumOfSquares(): Int {
        return (1..maxNaturalNumber).reduce { acc, i -> acc + i*i }
    }

    fun difference(): Int {
        return squareOfSum() - sumOfSquares()
    }

}