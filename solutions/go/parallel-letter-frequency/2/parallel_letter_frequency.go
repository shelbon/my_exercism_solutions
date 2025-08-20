package letter
import (
    "sync"
)
// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string ) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
    frequencies := FreqMap{}
    wg := sync.WaitGroup{}
    wg.Add(len(texts))
	for _,text :=range texts{
		go func(text string){
            defer wg.Done()
            var freq = Frequency(text)
            if frequencies==nil{
                frequencies=freq
            }else{
				for k,v:= range freq{
                    (frequencies)[k]+=v
                }
                
            }
        }(text)
    }
    wg.Wait()
    return  frequencies
     
}
