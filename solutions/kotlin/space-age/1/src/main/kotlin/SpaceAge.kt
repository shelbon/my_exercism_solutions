import java.math.BigDecimal
import java.math.RoundingMode

class SpaceAge(private val age: Long) {
    enum class  OrbitalPeriodOfPlanet(val earthYear:Double)
    {
        EARTH(1.0),
        MERCURY(0.2408467),
        VENUS(0.61519726),
        MARS(1.8808158),
        JUPITER(11.862615),
        SATURN(29.447498),
        URANUS(84.016846),
        NEPTUNE(164.79132)


    }
    fun onEarth(): Double = rounding(age/yearToSecond(OrbitalPeriodOfPlanet.EARTH.earthYear))
    fun onMercury(): Double =rounding(age/yearToSecond(OrbitalPeriodOfPlanet.MERCURY.earthYear))

    fun onVenus(): Double {
        return rounding(age/yearToSecond(OrbitalPeriodOfPlanet.VENUS.earthYear))
    }

    fun onMars(): Double {
        return rounding(age/yearToSecond(OrbitalPeriodOfPlanet.MARS.earthYear))
    }

    fun onJupiter(): Double {
        return rounding(age/yearToSecond(OrbitalPeriodOfPlanet.JUPITER.earthYear))
    }

    fun onSaturn(): Double {
        return rounding(age/yearToSecond( OrbitalPeriodOfPlanet.SATURN.earthYear))
    }

    fun onUranus(): Double {
        return rounding(age/yearToSecond(OrbitalPeriodOfPlanet.URANUS.earthYear))
    }

    fun onNeptune(): Double {
        return rounding(age/yearToSecond(OrbitalPeriodOfPlanet.NEPTUNE.earthYear))
    }
    private fun rounding(number:Double) :Double
    {
      return   BigDecimal(number).setScale(2, RoundingMode.HALF_EVEN).toDouble()
    }

    private fun yearToSecond(year:Double):Double
    {
        return year.times(31557600)
    }

}