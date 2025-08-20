class Triangle() {
    private lateinit var side1: Any
    private lateinit var side2: Any
    private lateinit var side3: Any

    constructor(side1: Int, side2: Int, side3: Int) : this() {
        isValid(side1,side2,side3)
        this.side1 = side1
        this.side2 = side2
        this.side3 = side3

    }

    constructor(side1: Double, side2: Double, side3: Double) : this() {
        isValid(side1, side2, side3)
        this.side1 = side1
        this.side2 = side2
        this.side3 = side3

    }

    val isScalene: Boolean
        get() {
            if (side1 != side2 && side2 != side3) {
                return true
            }
            return false
        }
    val isIsosceles: Boolean
        get() {
            var nbOfEqualSide = 0
            if (side1 == side2) {
                nbOfEqualSide++
            }
            if (side1 == side3) {
                nbOfEqualSide++
            }
            if (side2 == side3) {
                nbOfEqualSide++
            }
            if (nbOfEqualSide >= 1) {
                return true
            }
            return false
        }
    val isEquilateral: Boolean
        get() {
            if (side1 == side2 && side2 == side3) {
                return true
            }
            return false
        }

    private fun isValid(side1:Any,side2:Any,side3:Any) {
        if (side1 is Int && side2 is Int && side3 is Int ) {

            require( side1 >0 && side2>0 && side3>0 && side1 + side2 >= side3
                    && side2 + side3 >= side1
                    && side1 + side3 >= side2) { "it's not a triangle " }

        } else if (side1 is Double && side2 is Double && side3 is Double ) {

            require( side1 >0 && side2>0 && side3>0 &&  side1 + side2 >= side3
                    && side2 + side3 >= side1
                    && side1 + side3 >= side2) { "it's not a triangle" }
        }
    }

}