package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if(kind=="car" || kind=="truck"){
        return true
    }
return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
    choice:=""
     result:=" is clearly the better choice."
	if option1<option2{
    choice +=option1
    }  else {
      choice +=option2
    }
return  choice+result
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    pourcentage:=0.0
    if age >=10 {
        pourcentage=0.5
    } else if age>=3 && age<10 {
        pourcentage=0.7
    } else {
        pourcentage=0.8
    }
    return originalPrice * pourcentage
}
