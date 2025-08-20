package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string,averagePreparationTime int) int {
     numberOfLayers:=len(layers)
    if averagePreparationTime == 0 {
        return numberOfLayers * 2
    }
     return numberOfLayers * averagePreparationTime
}
 
func Quantities(layers []string)(int,float64){
    neededNoodles:=0
    neededSauces:=0.0
    for _, layer := range layers {
       if layer == "sauce" {
           neededSauces+=0.2
       }else if layer == "noodles" {
       		neededNoodles +=50
       }
    }
    return neededNoodles,neededSauces
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string,yourList []string){
     lastElementFriendList:=friendsList[len(friendsList) - 1]
     yourList[len(yourList)-1]=lastElementFriendList
}
// TODO: define the ' ScaleRecipe()' function
func  ScaleRecipe(quantities []float64,nbPortions int) []float64 {
    newQuantities:=make([]float64,len(quantities))
    for index,amount := range quantities {
        newQuantities[index]=amount*float64(nbPortions)/2
    }
	return newQuantities
}