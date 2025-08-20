class Matrix(private val matrix: List<List<Int>>) {


    val saddlePoints: Set<MatrixCoordinate>

    init {
        saddlePoints = getSaddlePointsCoordinate()
    }

    private fun getSaddlePointsCoordinate(): Set<MatrixCoordinate> {
        val saddlePoints = mutableSetOf<MatrixCoordinate>()
        if (matrix[0].isNotEmpty()) {
            matrix.mapIndexed { index, list ->
                val greaterNumberInRow = list.max()!!
                val columnPos = list.indexOf(greaterNumberInRow)
                val isLessEqualInColumn = matrix.mapIndexed { _, list2 ->
                    greaterNumberInRow <= list2[columnPos]
                }.reduce { acc, b -> b && acc }
                if (isLessEqualInColumn) {
                    saddlePoints.add(MatrixCoordinate(index, columnPos))
                }
                saddlePoints
            }
        }
        return saddlePoints
    }
}
