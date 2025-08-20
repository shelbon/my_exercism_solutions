package diffsquares
import  "math"
func byExponent2(n int) int{
    return int(math.Pow(float64(n),2));
}
func SquareOfSum(n int) int {
    result:=0;
	for i:=1;i<=n;i++{
        result+=i;
    }
 return byExponent2(result);
}

func SumOfSquares(n int) int {
	  result:=0;
	for i:=1;i<=n;i++{
        result+=byExponent2(i);
    }
return result
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n);
}
