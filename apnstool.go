package main

import (
	"flag"
	"fmt"
	"github.com/zhengying/apns"
	"os"
	"path/filepath"
)

var (
	debug       = flag.Bool("debug", true, "use debug server or release server, default is debug")
	pem         = flag.String("pem", "", "cert & key in one pem file , give pem path")
	alert       = flag.String("alert", "hello for test", "alert text for send")
	badge       = flag.Int("badge", 1, "badge count for send")
	sound       = flag.String("sound", "", "sound path for send")
	devicetoken = flag.String("devicetoken", "", "device token")
)

func main() {

	flag.Parse()

	fmt.Println("%v", flag.Args())

	curdir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	var server string

	if *debug == true {
		server = "gateway.sandbox.push.apple.com:2195"
	} else {
		server = "gateway.push.apple.com:2195"
	}

	if *pem == "" {
		filecomps := filepath.SplitList(*pem)
		if len(filecomps) == 1 {
			*pem = fmt.Sprintln(curdir, "/", *pem)
		}
	}

	if *devicetoken == "" {
		fmt.Println("devicetoken must be given")
		return
	}

	fmt.Println("devicetoken:", *devicetoken)
	fmt.Println("server:", server)
	fmt.Println("pem path:", *pem)

	payload := apns.NewPayload()
	payload.Alert = *alert
	payload.Badge = *badge
	payload.Sound = *sound

	pn := apns.NewPushNotification()
	pn.DeviceToken = *devicetoken
	pn.AddPayload(payload)

	client := apns.ComboPEMClient(server, *pem)

	resp := client.Send(pn)

	alert, _ := pn.PayloadString()
	fmt.Println("  Alert:", alert)
	fmt.Println("Success:", resp.Success)
	fmt.Println("  Error:", resp.Error)

}
