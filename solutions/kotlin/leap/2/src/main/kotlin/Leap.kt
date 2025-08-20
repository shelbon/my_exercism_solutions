data class Year(val year: Int) {

    // TODO: Implement proper constructor

    val isLeap: Boolean =  year.mod(4) == 0 && ( year.mod(100) != 0 ||  year.mod( 400) == 0)
}
