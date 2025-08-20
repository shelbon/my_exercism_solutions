object Flattener {
    fun flatten(source: Collection<Any?>): List<Any> {
        return source.fold(emptyList()) { acc, n ->
            n?.let {
                if (it !is Collection<*>) {
                    acc.plus(it)
                } else {
                    acc.plus(flatten(it))
                }
            } ?: acc
        }

    }
}
