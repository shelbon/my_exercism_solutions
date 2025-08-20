import java.time.LocalDate
import java.time.LocalDateTime

class Gigasecond()
{
    lateinit var date:LocalDateTime
    constructor(localDate: LocalDate) : this() {
        date=localDate.atStartOfDay().plusSeconds(Math.pow(10.0,9.0).toLong())
    }
    constructor( localDateTime: LocalDateTime) : this() {
        this.date=localDateTime.plusSeconds(Math.pow(10.0,9.0).toLong())
    }


}