object Pangram {

    fun isPangram(text:String) :Boolean
    {
        if(text.isEmpty()) return false
        var matchingChar= mutableListOf<Char>()
         var cleansedText=text.replace("\\s|_|[0-9]|[\\/ \\. \\\"]".toRegex(),"").toLowerCase()
        for(char in cleansedText)
        {
            if(!matchingChar.contains(char))  matchingChar.add(char)
        }
        if(matchingChar.count()<26) return false

        return true
    }

}
