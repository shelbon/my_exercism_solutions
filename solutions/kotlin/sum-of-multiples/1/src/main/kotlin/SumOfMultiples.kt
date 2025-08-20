object SumOfMultiples {
    fun sum(factors: Set<Int>, limit: Int): Int {
        val multiples=mutableSetOf<Int>()
        val factorsWithoutZero= factors.takeWhile { it>0 }
        factorsWithoutZero.forEach{
            factor->
               multiples.addAll(getMultipleOfUntil(factor,limit))
        }
        return multiples.sum()
    }
     private fun getMultipleOfUntil(targetNumber: Int, upperLimit: Int):List<Int>{
        val multipleOfTargetNumber=mutableListOf<Int>()
         val numberOfMultiple=upperLimit/targetNumber
         for(number in 1..numberOfMultiple){
             val multiple=targetNumber * number
             if (multiple < upperLimit) {
                 multipleOfTargetNumber.add(multiple)
             }
        }
        return multipleOfTargetNumber
     }
}
