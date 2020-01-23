package main

func StubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLenth :=0

	for i, ch := range []byte(s){
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start{
			start = lastI + 1
		}

		if i - start+1 > maxLenth{
			maxLenth = i-start+1
		}
		lastOccurred[ch] = i
	}
	return maxLenth
}

func main()  {

}