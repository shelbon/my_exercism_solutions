class Triangle(private val side1: Any, private val side2: Any, private val side3: Any) {

    init {
        isValid(side1, side2, side3)
    }

    val isScalene: Boolean
        get() {
            return nbUniqueSide == 3
        }
    val isIsosceles: Boolean
        get() {
            print(nbUniqueSide)
            return nbUniqueSide <= 2
        }
    val isEquilateral: Boolean
        get() {
            return nbUniqueSide == 1
        }

    private fun isValid(side1: Any, side2: Any, side3: Any) {
        if (side1 is Int && side2 is Int && side3 is Int) {
            require(side3 > side1 - side2 && side3 < side1 + side2) { "it's not a triangle " }
        } else if (side1 is Double && side2 is Double && side3 is Double) {
            require(side3 > side1 - side2 && side3 < side1 + side2) { "it's not a triangle" }
        }
    }

    private val nbUniqueSide: Int=setOf(side1, side2, side3).count()

}