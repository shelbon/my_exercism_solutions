package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var totalBirdCount int =0;
	for i:=0;i<len(birdsPerDay);i++{
         totalBirdCount+=birdsPerDay[i]
    }
  return totalBirdCount; 
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var totalBirdCount int =0;
    var index int =7*(week-1);
	for i:=0;i<7;i++{
         totalBirdCount+=birdsPerDay[index]
        index++
    }
  return totalBirdCount; 
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i:=0;i<len(birdsPerDay);i++{
        if i%2==0 {
            birdsPerDay[i]++
        }
    }
  return birdsPerDay; 
}
