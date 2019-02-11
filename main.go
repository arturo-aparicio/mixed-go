package main

import (
	"fmt"
        ellipse "github.com/jfrog-aparicio/mixed-simple-mod"
        "golang.org/x/net/context"
)

func main() {
 //Initalise the Init function with value of A,B
 e := ellipse.Init{
  9, 2,
 }
 //this will give answer as 0.9749960430435691
 fmt.Println(e.GetEccentricity())
 // Just to avoid go complaining that context is not being used 
 type Context = context.Context
 type CancelFunc = context.CancelFunc
}
