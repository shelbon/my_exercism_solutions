package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
     
	 switch{
         case balance< 0 : return 321.3/100
     	 case balance >=0 && balance <1000: return 50.0/100
         case balance >=1000 && balance<5000: return 162.1/100
         default: return 247.5/100
     }
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interestRate:=InterestRate(balance)
    return balance * float64(interestRate) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance+Interest(balance);
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    years:=0;
    for balance < targetBalance {
        balance = AnnualBalanceUpdate(balance)
        years++
    }
    return years;
}
