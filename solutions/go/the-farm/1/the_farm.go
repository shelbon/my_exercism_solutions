package thefarm
import ("errors"
        "fmt")

 type SillyNephewError struct {
     nbCows int
 } 
 var(
     ErrDivisionByZero=errors.New("division by zero")
     ErrNegativeFodder=errors.New("negative fodder")
     )

func (err *SillyNephewError)Error() string{
    return fmt.Sprintf("silly nephew, there cannot be %d cows",err.nbCows)
}
// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
      amountFooder,err:=weightFodder.FodderAmount()
    switch {
        case cows==0 :return 0.0,ErrDivisionByZero
    
        case cows<0: return 0.0,&SillyNephewError{
                nbCows:cows,
        }
        
        case err==ErrScaleMalfunction && amountFooder>=0:return  (amountFooder*2)/float64(cows),nil
        
        case err==ErrScaleMalfunction && amountFooder<0 :
         return 0.0,ErrNegativeFodder
        
        case err!=nil : return 0.0,err
        
        case amountFooder<0 : return 0.0,ErrNegativeFodder
        
        default: return amountFooder/float64(cows),nil
    }
    }

