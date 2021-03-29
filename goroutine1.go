package main
import "fmt"
func  channel1 ()  {
	done:= make(chan bool)
	
	
	go func ()  {
		fmt.Println("你好！并发")
		//管道有数据时打印完成
		done <- true

	}()
	
        
	<- done
	//接收数据，完美退出
}




	
