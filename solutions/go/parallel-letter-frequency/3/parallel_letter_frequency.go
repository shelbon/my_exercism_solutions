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
    mutex:=sync.Mutex{}
    wg := sync.WaitGroup{}
    wg.Add(len(texts))
	for _,text :=range texts{
		go func(text string,mutex *sync.Mutex){
            defer wg.Done()
            var calculedFrequency = Frequency(text)
            mutex.Lock()
            defer mutex.Unlock()
            if frequencies==nil{
                frequencies=calculedFrequency
            }else{
				for k,v:= range calculedFrequency{
                    (frequencies)[k]+=v
                }
                
            }
        }(text,&mutex)
    }
    wg.Wait()
    return  frequencies
     
}
