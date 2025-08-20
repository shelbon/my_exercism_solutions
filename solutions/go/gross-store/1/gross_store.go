package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
        "quarter_of_a_dozen":3,
        "half_of_a_dozen":6,
        "dozen":12,
        "small_gross":120,
        "gross":144,
        "great_gross":1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
         unitValue, isUnitInUnits:=units[unit];
	 	if isUnitInUnits {
            _, isItemInBill:=bill[item]
            if(isItemInBill){
                bill[item]+=unitValue
            }else{
            	bill[item]=unitValue
            }
            return true
        }
       return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	  unitValue, isUnitInUnits:=units[unit];
	 	if isUnitInUnits {
            value, isItemInBill:=bill[item]
            if(isItemInBill){
                newQuantity:= value - unitValue
                switch {
                   case newQuantity==0:delete(bill,item)
                   case newQuantity>0:bill[item]-=unitValue
                   default:return false
                }
                return true
            }
        }
       return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	   itemValue, isItemInBill:=bill[item]
    if isItemInBill {
        return itemValue,true
    }
	return 0,false
}
