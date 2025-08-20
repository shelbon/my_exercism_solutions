import java.time.LocalDate
import java.time.LocalDateTime

class Gigasecond(birthTime: LocalDateTime) {

    val date: LocalDateTime = birthTime.plusSeconds(Math.pow(10.0, 9.0).toLong())


    constructor(birthTime: LocalDate) : this(birthTime.atStartOfDay()) {
        val date = birthTime.atStartOfDay().plusSeconds(Math.pow(10.0, 9.0).toLong())
    }


}