object HandshakeCalculator {
    private const val REVERSE_CODE=4
    fun calculateHandshake(number: Int): List<Signal> {
        val handShakeEvent = Signal.values().filter { signal ->
            isBitExist(signal.ordinal, number)
        }
        if(isBitExist(REVERSE_CODE,number))
        {
            return handShakeEvent.reversed()
        }
        return  handShakeEvent
    }

    private fun isBitExist(position: Int, number: Int): Boolean = number shr position and 1 == 1
}



