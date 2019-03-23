package main

import (
	"fmt"
)

func main() {
	ch:=make(chan int) //create channel before making a call
	for st:=0;st<=10;st++{
		go evenOrOddWithCommunication(st,ch)  // pass channel to calling method
	}

	//the below snippet handles messaage from channel and it makes execution continous

	/*for ra:= range ch{
      go func(ra1 int){
		time.Sleep(2 * time.Second)
		  evenOrOddWithCommunication(ra1,ch)
	  }(ra)
	}*/
	//alternative option to control the execution as well and stop when count reaches 10.
	for i:=0;i<=10;i++{
		<-ch
	}
}

func evenOrOddWithCommunication(val int,chl chan int){
	if val%2==0{
		fmt.Println("Iam Even : ",val)
	}else{
		fmt.Println("Iam Odd : ",val)
	}
	chl<-val
}

