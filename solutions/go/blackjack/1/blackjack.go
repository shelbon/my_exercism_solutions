package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardValue:=0
    switch {
       case card=="ace": cardValue=11
       case card=="two": cardValue=2
       case card=="three": cardValue=3
       case card=="four": cardValue=4
       case card=="five": cardValue=5
       case card=="six": cardValue=6
       case card=="seven": cardValue=7
       case card=="eight": cardValue=8
       case card=="nine": cardValue=9
       case card=="ten"||card =="jack"||
            card=="queen"||card=="king":cardValue=10
       default: cardValue=0
    }
    return cardValue
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	card1Value:=ParseCard(card1)
    card2Value:=ParseCard(card2)
    return card1Value+card2Value==21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	switch{
         case isBlackjack && dealerScore<10: return "W"
         case !isBlackjack && dealerScore>=10: return "P"
    	 default: return "S"
    }
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	 switch{
         case handScore >=17:return "S"
         case handScore <=11:return "H"
         case handScore >=12 && handScore <=16 && dealerScore<7 :return "S"
         default: return "H"
     }
}
