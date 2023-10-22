package main

import (
	"flag"
	"grc/tuichat"
	"log"
)

func main() {
  var host = flag.String("a","0.0.0.0:8080","Local listener address.")
  var intf = flag.Bool("i", true, "Graphical Interface");
  var vrbt = flag.Int("v",1,"Verbosiy: 0=Silent 2=log all")
  var lstn = flag.Bool("l",false,"Listen for connection.")
  var robo = flag.Bool("r",false,"Bot mode");
  flag.Parse()


  if *robo {
    ConnectRemoteForDebugging(*host) 
    
  } else if *intf && !*lstn {
    tuichat.Start()

  } else { 
    if *vrbt > 0 {
      log.Printf("### GRC - Golang Relay Chat Starting ###")
    }
    if *lstn {
      StartServer(*host, *vrbt)
    } else {
      ConnectRemoteHost(*host)
    }
  }
}



