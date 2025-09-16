package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Appirit/go-carto/internal"
)

func main() {
	log.Println("*************")
	log.Println("*** Manip ***")
	log.Println("*************")

	os.Chdir(must(filepath.Abs(`../../../testdata`))) // tous les outils nécessaires sont dans ce dossier
	log.Println(filepath.Abs("D0576.json"))
	var data = must(os.ReadFile("D0576.json"))
	log.Printf("taille=%d", len(data))

	nodes, _ := internal.ParseJSON(data)

	for _, n := range nodes {
		fmt.Printf("level=%d terminal=%v name=%s value=%d\n",
			n.Level, n.Terminal, n.Name, len(n.Value))
	}
}

func must[T any](val T, err error) T {
	if err != nil {
		log.Println("***********************************************************")
		log.Fatal(err)
	}
	return val
}
