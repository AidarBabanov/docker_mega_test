package main

import (
	"fmt"

	sp "github.com/AidarBabanov/docker_mega_test/sub_pack"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")
	router := mux.NewRouter()
	fmt.Println(router)
	fmt.Println(sp.Get())
}
