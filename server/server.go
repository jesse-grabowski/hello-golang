package server

import (
	"fmt"
	"log"
)

func Start(port string) {
	log.Panicln(NewRouter().Run(fmt.Sprintf(":%s", port)))
}
