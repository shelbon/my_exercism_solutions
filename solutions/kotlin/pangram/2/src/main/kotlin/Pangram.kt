import java.util.*

object Pangram {

    fun isPangram(text:String) :Boolean
    {
        if(text.isEmpty()) return false
         val  matchingChar= mutableListOf<Char>()
         val cleansedText= text.replace("\\s|_|[0-9]|[\\/ \\.\\\"]".toRegex(),"")
                 .lowercase(Locale.getDefault())
         for(char in cleansedText)
         {
             if(!matchingChar.contains(char))  matchingChar.add(char)
         }
        if(matchingChar.count()<26) return false

        return true
    }

}
