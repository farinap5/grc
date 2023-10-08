package main

import (
	"flag"
	"log"
)

func main() {
  var host = flag.String("a","0.0.0.0:8080","Local listener address.")
  var vrbt = flag.Int("v",1,"Verbosiy: 0=Silent 2=log all")
  var lstn = flag.Bool("l",false,"Listen for connection.")
  var dbgg = flag.Bool("d",false,"Debugging client.")
  flag.Parse()

  if *vrbt > 0 {
    log.Printf("### GRC - Golang Relay Chat Starting ###")
  }

  if *lstn {
    StartServer(*host, *vrbt)
  } else {
    if *dbgg {
      ConnectRemoteForDebugging(*host) 
    } else {
      ConnectRemoteHost(*host)
    }
  }
}






