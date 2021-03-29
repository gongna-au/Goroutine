package main

import (
	"fmt"
	"sync"
)

func main() {
	const N = 10
	var wg sync.WaitGroup
	    for i := 0; i < N; i++ {
			wg.Add(1)
			//Add必须要在go前面
		    go func(i int) {
			    fmt.Println(i, "你好！并发")
				defer wg.Done()
				//defer要与Done连用，用于给计数器减一
		    }(i)
	    }
	wg.Wait()
	//当计数器为0时不再阻塞main协程

}
