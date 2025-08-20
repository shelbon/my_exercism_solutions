object Flattener {
    fun flatten(source: Collection<Any?>): List<Any> {
        return "[-]*\\d+".toRegex().findAll(source.toString())
                .map { it.value.toInt() }
                .toList()
    }
}
