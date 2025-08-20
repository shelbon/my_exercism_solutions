object Hamming
{

         fun compute(dnaStrand1:String, dnaStrand2: String):Int
         {
             require(dnaStrand1.length==dnaStrand2.length){ "left and right strands must be of equal length." }
             return dnaStrand1.zip(dnaStrand2).filter { it.first !=it.second }.count()

         }
}