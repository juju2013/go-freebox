package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/juju2013/go-freebox"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	fbx := freebox.New()

	err := fbx.Connect()
	if err != nil {
		logrus.Fatalf("fbx.Connect(): %v", err)
	}

	err = fbx.Authorize()
	if err != nil {
		logrus.Fatalf("fbx.Authorize(): %v", err)
	}

	err = fbx.Login()
	if err != nil {
		logrus.Fatalf("fbx.Login(): %v", err)
	}

  airMedia, err := fbx.GetAirMediaReceivers()
  if err != nil {
		logrus.Errorf("fbx.getAirMediaReceivers(): %v", err)
  } else {
	  fmt.Printf("***** : %#v\n", airMedia)
  }

  callEntry, err := fbx.GetCallEntries()
  if err != nil {
		logrus.Errorf("fbx.GetCallEntries(): %v", err)
  } else {
	  fmt.Printf("***** : %#v\n", callEntry)
  }
}
