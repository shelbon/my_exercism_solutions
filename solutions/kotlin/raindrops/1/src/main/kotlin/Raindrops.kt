import java.lang.StringBuilder

class Raindrops {
    companion object {
        fun convert(number: Int):String {
             var raindrops=StringBuilder()
             var factorsAndOutputs= sortedMapOf(3 to "Pling",5 to "Plang",7 to "Plong")
             for(entry in factorsAndOutputs)
             {
                 if(asFactor(entry.key,number)) raindrops.append(entry.value)
             }
            return raindrops.toString().ifEmpty { number.toString()}
        }
        private fun asFactor(factor:Int, number:Int):Boolean
        {
            if(number%factor==0)
            {
                return true
            }
            return  false
        }

    }

}
