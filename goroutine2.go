package main
import "fmt"
func channel2 (){
    const	N=10
	done:=make(chan bool,N)
	
	for i:=0;i<N;i++{

		go func (i int){
	
			fmt.Println(i,"你好！并发")
			done <-true
		}(i)
		//循环开启goroutine，把他们一个一个缓存起来,一个一个堵塞channel
	
		
	}
	//循环接收值，打开channel
    for i:=1;i<N;i++{
			<-done
		}
        
}