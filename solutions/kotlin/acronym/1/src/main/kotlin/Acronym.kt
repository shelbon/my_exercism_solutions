object Acronym {
    fun generate(phrase: String): String {
        val tabPhrase=phrase.filter{ it!='-'}
                                       .split(" ")
                                       .filter { it.isNotEmpty() }
       return  tabPhrase.map{ it[0].toString().toUpperCase()}
                        .reduce { acc, s ->acc+s  }
    }

}
