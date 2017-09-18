package main

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/midoblgsm/myFirstProtogRPC/resources"
	"net/http"
)

func main() {
	myClient := resources.Client{Id: 526, Name: "John Doe", Email: "johndoe@example.com", Country: "US"}
	clientInbox := make([]*resources.Client_Mail, 0, 20)
	clientInbox = append(clientInbox,
		&resources.Client_Mail{RemoteEmail: "jannetdoe@example.com", Body: "Hello. Greetings. Bye."},
		&resources.Client_Mail{RemoteEmail: "WilburDoe@example.com", Body: "Bye, Greetings, hello."})

	myClient.Inbox = clientInbox

	data, err := proto.Marshal(&myClient)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = http.Post("http://localhost:3000", "", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return
	}
}
