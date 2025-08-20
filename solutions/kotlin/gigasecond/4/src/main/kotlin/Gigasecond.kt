import java.time.LocalDate
import java.time.LocalDateTime

class Gigasecond(birthTime: LocalDateTime) {

    val date: LocalDateTime = birthTime.plusSeconds(1E9.toLong())


    constructor(birthTime: LocalDate) : this(birthTime.atStartOfDay())


}