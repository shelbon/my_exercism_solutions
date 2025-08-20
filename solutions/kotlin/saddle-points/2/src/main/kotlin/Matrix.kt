class Matrix(private val matrix: List<List<Int>>) {

    val saddlePoints: Set<MatrixCoordinate>
        get() {
            val saddlePoints: MutableSet<MatrixCoordinate> = mutableSetOf()
            var isLessThanInColumn = false
            var greaterNumberInRow: Int?
            var columnPos: Int
            for (row1 in matrix.indices) {
                greaterNumberInRow = matrix[row1].max()
                columnPos = matrix[row1].indexOf(greaterNumberInRow)
                for (row2 in matrix.indices) {
                    greaterNumberInRow?.let {
                        isLessThanInColumn = greaterNumberInRow <= matrix[row2][columnPos]
                    }
                    if (!isLessThanInColumn) {
                        break
                    }
                }
                if (isLessThanInColumn) {
                    saddlePoints.add(MatrixCoordinate(row1, columnPos))
                    isLessThanInColumn = false
                }
            }
            return saddlePoints
        }
}
