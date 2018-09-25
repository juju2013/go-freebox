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

  callEntries, err := fbx.GetCallEntries()
  if err != nil {
		logrus.Errorf("fbx.GetCallEntries(): %v", err)
  } else {
	  fmt.Printf("***** : %#v\n", callEntries)
    if err := fbx.MarkAllRead(); nil != err {
  		logrus.Errorf("fbx.MarkAllRead(): %v", err)
    }
    callEntry, err := fbx.GetCallEntrie(523)
    if err != nil {
  		logrus.Errorf("fbx.GetCallEntrie(523): %v", err)
    } else {
  	  fmt.Printf("***** : %#v\n", callEntry)
      fbx.MarkRead(callEntry.ID)
    }
  }
  
  Contacts, err := fbx.GetContacts()
  if err != nil {
		logrus.Errorf("fbx.GetContacts(): %v", err)
  } else {
	  fmt.Printf("***** : %#v\n", Contacts)
    contact, err := fbx.GetContact(4)
    if nil != err {
  		logrus.Errorf("fbx.GetContact(4): %v", err)
    } else {
  	  fmt.Printf("***** : %#v\n", contact)
    }
  }

}
