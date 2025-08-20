class Hamming
{
    companion object{
         fun compute(dnaStrand1:String, dnaStrand2: String):Int
         {
             var hammingDistance=0
             if(dnaStrand1.isEmpty() && dnaStrand2.isEmpty())
             {
                 return hammingDistance
             }
             handleException(dnaStrand1, dnaStrand2)
             for( i in dnaStrand1.indices)
             {
                if(dnaStrand1[i]!=dnaStrand2[i])
                    hammingDistance++
             }
             return hammingDistance
         }
        private  fun handleException(dnaStrand1:String, dnaStrand2: String)
        {
            if(dnaStrand1.length!=dnaStrand2.length || dnaStrand2.length != dnaStrand1.length)
            {
                throw IllegalArgumentException("left and right strands must be of equal length.")
            }
        }
    }
}