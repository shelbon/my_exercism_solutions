object Flattener {
    fun flatten(source: Collection<Any?>): List<Any> {
        val result = mutableListOf<Any>()
        source.forEach {
            if (it !is Collection<Any?>) {
                it?.let {
                    result.add(it)
                }
            } else {
                flatten(it).toCollection(result)
            }

        }
        return result
    }
}
