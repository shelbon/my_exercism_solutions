object Flattener {
    fun flatten(source: Collection<Any?>): List<Any> {
        return source.fold(emptyList()) { acc, n ->
            when (n) {
                null -> acc
                !is Collection<*> -> acc.plus(n)
                else -> acc.plus(flatten(n))
            }
        }

    }
}
