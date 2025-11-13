/**
 * Print & log
 Opções:

     fmt.Println() → stdout
     log.Println() → timestamp + pode-se determinar onde o erro ficará logado
     log.Fatalln() → os.Exit(1) sem defer
     log.Panicln() → println + panic → funcões em defer rodam; dá pra usar recover
     panic()

 Recomendação: use log.
 Código:

         fmt.Println
         log.Println
         log.SetOutput
         log.Fatalln
         log.Panicln
         panic


 */
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	_, err = os.Open("no-file.txt")
	if err != nil {
		// fmt.Println(err) // Print do error
		// log.Println(err) // Print do error com prefox datetime
		// log.Fatalln(err) // Print do error com prefox datetime e mata o programa com exit status 1
		log.Panicln(err)
	}
}
