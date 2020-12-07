package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/juju2013/go-freebox"
)

func main() {
	switch strings.ToUpper(os.Getenv("GOFBX_LOGLEVEL")) {
	case "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	fbx := freebox.New()

	err := fbx.Connect()
	if err != nil {
		logrus.Fatalf("fbx.Connect(): %v", err)
	}

	err = fbx.Authorize()
	if err != nil {
		logrus.Fatalf("fbx.Authorize(): %v", err)
	}

	{
		title := "AirMedia"
		err = fbx.Login()
		if err != nil {
			logrus.Fatalf("fbx.Login(): %v", err)
		}
		ret, err := fbx.GetAirMediaReceivers()
		if err != nil {
			logrus.Errorf("Error %v: %v", title, err)
		} else {
			dump(title, ret)
		}
	}
	{
		title := "CallEntries"
		ret, err := fbx.GetCallEntries()
		if err != nil {
			logrus.Errorf("Error %v: %v", title, err)
		} else {
			dump(title, ret)

			title = "CallEntry"
			callEntry, err := fbx.GetCallEntrie(555)
			if err != nil {
				fmt.Printf("fbx.GetCallEntrie(555): %v", err)
			} else {
				fbx.MarkRead(callEntry.ID)
				dump(title, callEntry)
			}
			title = "Contacts"
			Contacts, err := fbx.GetContacts()
			if err != nil {
				logrus.Errorf("Error %v: %v", title, err)
			} else {
				dump(title, Contacts)

				title = "Contact"
				contact, err := fbx.GetContact(2)
				if nil != err {
					logrus.Errorf("fbx.GetContact(2): %v", err)
				} else {
					dump(title, contact)
				}
			}
		}
	}
	{
		title := "LanConfig"
		ret, err := fbx.GetLanConfig()
		if err != nil {
			logrus.Errorf("Error %v: %v", title, err)
		} else {
			dump(title, ret)
		}
	}
	{
		title := "LanBrowserInterfaces"
		ret, err := fbx.GetLanBrowserInterfaces()
		if err != nil {
			logrus.Errorf("Error %v: %v", title, err)
		} else {
			dump(title, ret)
		}
	}
	{
		title := "LanBrowserInterface"
		ret, err := fbx.GetLanBrowserInterface("pub")
		if err != nil {
			logrus.Errorf("Error %v: %v", title, err)
		} else {
			dump(title, ret)
		}
	}
	{
		title := "DhcpConfig"
		ret, err := fbx.GetDhcpConfig()
		if err != nil {
			logrus.Errorf("fbx.GetDhcpConfig(): %v", err)
		} else {
			dump(title, ret)
		}
	}
	{
		title := "DhcpStaticLease"
		ret, err := fbx.GetDhcpStaticLease()
		if err != nil {
			logrus.Errorf(title, err)
		} else {
			dump(title, ret)
		}
	}
	{
		title := "DhcpDynamicLease"
		ret, err := fbx.GetDhcpDynamicLease()
		if err != nil {
			logrus.Errorf(title, err)
		} else {
			dump(title, ret)
		}
	}
}

func dump(title string, obj interface{}) {
	j, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		logrus.Errorf("%v", err)
	}
	fmt.Printf("========%s========\n%s\n\n", title, string(j))
}
