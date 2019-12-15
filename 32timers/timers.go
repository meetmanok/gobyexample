package main

/*
We often want to execute Go code at some point in the future, or repeatedly at some interval.
Go’s built-in timer and ticker features make both of these tasks easy.
We’ll look first at timers and then at tickers.
 */
import (
	"fmt"
	"time"
)

func main(){
	//Timers represent a single event in the future.
	//You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.
	//This timer will wait 2 seconds.
	timer1 := time.NewTimer(time.Second*2)

	//The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer expired.
	<-timer1.C
	fmt.Println("timer1 expired")

	//If you just wanted to wait, you could have used time.Sleep.
	//One reason a timer may be useful is that you can cancel the timer before it expires. Here’s an example of that.
	timer2 := time.NewTimer(time.Second*2)
	go func() {
		timer2.Reset(time.Second)
		<-timer2.C
		fmt.Println("timer2 expired")
	}()

	isStopped := timer2.Stop()
	if isStopped {
		fmt.Println("timer2 stopped")
	}

	//The first timer will expire ~2s after we start the program,
	//but the second should be stopped before it has a chance to expire.
}
