package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/midoblgsm/myFirstProtogRPC/resources"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myClient := resources.Client{}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = proto.Unmarshal(data, &myClient)
		if err != nil {
			fmt.Println(err)
		}

		println(myClient.Id, ":", myClient.Name, ":", myClient.Email, ":", myClient.Country)
		for _, mail := range myClient.Inbox {
			fmt.Println(mail.RemoteEmail, ":", mail.Body)
		}
	})

	http.ListenAndServe(":3000", nil)
}
