package ds_hw_0

import (
	"bufio"
	"io"
	"strconv"
	"os"
	"sync"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	sum:=0
	for num := range nums{
		sum=sum+num
	}
	out <- sum
	// TODO: implement me
	// HINT: use for loop over `nums`
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	nums := make(chan int, num)
	out := make(chan int, num)
	file,_ := os.Open(fileName)
	elem,_ :=readInts(file)
	waitGroup := sync.WaitGroup{}
	for i:=0;i<num;i++{
		waitGroup.Add(1)
		go func(){
			sumWorker(nums,out)
			waitGroup.Done()
		}()
	}
	waitGroup.Add(1)
	go func(){
		for _, i := range elem{
			nums<-i
		}
		waitGroup.Done()
		close(nums)
	}()
	waitGroup.Wait()
	close(out)
	res:=0
	for i:= range out{
		res=res+i
	}
	return res
	// TODO: implement me
	// HINT: use `readInts` and `sumWorkers`
	// HINT: used buffered channels for splitting numbers between workers
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
