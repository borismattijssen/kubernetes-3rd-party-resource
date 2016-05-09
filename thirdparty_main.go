package main

import (
	"fmt"
	"net"

	"github.com/borismattijssen/kubernetes-3rd-party-resource/client"

	"k8s.io/kubernetes/pkg/client/restclient"
)

// host and port for kubie api
const (
	HOST = "localhost"
	PORT = "8080"
)

// func TestWatch(t *testing.T) {
func main() {
	thirdPartyClient, err := client.NewThirdparty(&restclient.Config{
		Host: "http://" + net.JoinHostPort(HOST, PORT),
	})
	if err != nil {
		fmt.Println("Couldn't create 3rd-party client: ", err)
		return
	}
	list, err := thirdPartyClient.List()
	if err != nil {
		fmt.Println("Couldn't list workflows: ", err)
		return
	}
	for _, v := range list.Items {
		fmt.Println(v)
	}

	return
}

// func TestWatch(t *testing.T) {
// 	thirdPartyClient, err := NewThirdparty(&restclient.Config{
// 		Host: "http://" + net.JoinHostPort(HOST, PORT),
// 	})
// 	if err != nil {
// 		fmt.Println("Couldn't create 3rd-party client: ", err)
// 	}
// 	w, werr := thirdPartyClient.Watch()
// 	if werr != nil {
// 		fmt.Println("Couldn't watch workflows: ", werr)
// 		return
// 	}
// 	fmt.Println("Watching..")
// 	for res := range w.ResultChan() {
// 		fmt.Println("Something")
// 		fmt.Println(res)
// 	}
// 	return
// }
