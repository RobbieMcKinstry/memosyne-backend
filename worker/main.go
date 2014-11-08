package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	twilio "github.com/carlosdp/twiliogo"
)

/*
AllMemosByTime(time.Time, time.Time)
AllMemos()
SendMemo()
*/

func main() {
	var demo bool
	flag.BoolVar(&demo, "demo", false, "For use during the stage demonstration at the hackathon. Delete thereafter.")

	if demo {
		log.Println("Running in demo mode.")
		SendAllMemos()
		os.Exit(0)
	}

	endTime := time.Now()
	startTime := time.Now().Add(-1*time.Minute)
	memosToSend := AllMemosByTime(startTime, endTime)

	for _, memo := range memosToSend {
		if memo.IsApproved() {
			SendMemo(memo)
		}
		DeleteMemo(memo)
	}

	// store the command line flag
	// if flag:
	// 	sent all messages in the database
	// exit sys(0)
	//
	// search database for messages that need to be sent
	// for each message that needs to be sent:
	// 	if contact has a status of approved
	// 		send that message
	//	delete that message
	//
}

func SendAllMemos() {
	memosToSend := AllMemos()
	for _, memo := range memosToSend {
		if memo.IsApproved() {
			SendMemo(memo)
		}
		DeleteMemo(memo)
	}
	// go into the database
	// get all memos out
	// send all memos
	// delete all memos
}

func SendMemo(memo *Memo) {
	sid := "ACda518e51210a44b39d89fc196116e229"
	auth := "d5ab11efae16e40cd74216b78c8b3fd6"
	client := twilio.NewClient(sid, auth)
	message, err := twilio.NewMessage(client, "4124453191", "6238504947", twilio.Body("Hello World!"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message.Status)
	}
}
