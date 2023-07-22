package main

import (
	"fmt"
)

type Logger struct {
	message map[string]int
 }
 
 
 func Constructor() Logger {
	 return Logger{
		 message: make(map[string]int),
	 }
 }
 
 
 func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	 if ts, ok := this.message[message]; ok && timestamp < ts {
		 return false
	 }
	 this.message[message] = timestamp + 10
	 return true
 }
 
 
 /**
  * Your Logger object will be instantiated and called as such:
  * obj := Constructor();
  * param_1 := obj.ShouldPrintMessage(timestamp,message);
  */

func main(){
	
}