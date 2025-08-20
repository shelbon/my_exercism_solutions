import kotlin.math.pow

class Squares(private val maxNaturalNumber: Int) {


    fun squareOfSum(): Int {
        return (1..maxNaturalNumber).sum().toDouble().pow(2).toInt()
    }

    fun sumOfSquares(): Int {


        return (1..maxNaturalNumber).reduce { acc, i -> acc + (i.toDouble().pow(2).toInt()) }

    }

    fun difference(): Int {
        return squareOfSum() - sumOfSquares()
    }

}