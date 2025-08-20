object Hamming
{

         fun compute(dnaStrand1:String, dnaStrand2: String):Int
         {
             var hammingDistance=0
             var  pairOfNucleotides=dnaStrand1.zip(dnaStrand2)
             if(pairOfNucleotides.count()==0)
             {
                 return hammingDistance
             }
             require(dnaStrand1.length==dnaStrand2.length){ "left and right strands must be of equal length." }
             for( nucleotides in pairOfNucleotides)
             {
                if(nucleotides.first != nucleotides.second)
                    hammingDistance++
             }
             return hammingDistance
         }
}