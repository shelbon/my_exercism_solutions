package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	arrayLength:=len(slice)-1
    if index >=0 && index<=arrayLength{
        return slice[index],true
    }
	return 0,false
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	 arrayLength:=len(slice)-1
    if index >=0 && index<=arrayLength{
       slice[index]=value
        return slice
    }
	return append(slice,value)
    
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
    if(length<0){
        return make([]int,0)
    }
	slice:=make([]int,length)
    for i:=0 ;i<length;i++{
        slice[i]=value
    }
  return slice;
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
        sliceLength:=len(slice)
    	if index >= 0 && index <= sliceLength {
            return append(slice[:index],slice[index+1:]...);
        }
	return slice
	  
}
