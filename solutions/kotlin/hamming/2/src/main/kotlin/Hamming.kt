object Hamming
{

         fun compute(dnaStrand1:String, dnaStrand2: String):Int
         {
             var hammingDistance=0
             if(dnaStrand1.isEmpty() && dnaStrand2.isEmpty())
             {
                 return hammingDistance
             }
             require(dnaStrand1.length==dnaStrand2.length){ "left and right strands must be of equal length." }
             for( i in dnaStrand1.indices)
             {
                if(dnaStrand1[i]!=dnaStrand2[i])
                    hammingDistance++
             }
             return hammingDistance
         }
}