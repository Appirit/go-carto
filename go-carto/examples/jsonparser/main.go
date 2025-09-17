package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Appirit/go-carto/internal/jsonparser"
)

func main() {
	log.Println("*************")
	log.Println("*** Manip ***")
	log.Println("*************")

	os.Chdir(must(filepath.Abs(`../../../testdata`))) // tous les outils nécessaires sont dans ce dossier
	var data = must(os.ReadFile("D0576.json"))
	log.Printf("taille=%d", len(data))

	members, _ := jsonparser.ParseJSON(data)
	log.Println("---end---")

	for _, f := range jsonparser.GetFeaturesNode(members) {
		fmt.Printf("%s - level=%d terminal=%v len=%d = %s\n",
			f.This.ParentChain(), f.This.Level, f.This.Terminal, len(f.Descendants), f.GetGeometryType())
		start, end := f.This.GetOffsets(data)
		fmt.Printf("\tdata[%d, %d]=%s ... %s \n", start, end, f.This.Value[0:20], f.This.Value[len(f.This.Value)-20:len(f.This.Value)])
		fmt.Printf("\tdata[%d, %d]=%s ... %s \n", start, end, data[start:start+20], data[end-20:end])
		for _, n := range f.Descendants {
			fmt.Printf("%s%s%s- leaf=%v len=%d\n",
				strings.Repeat("  ", n.Level), n.Name, strings.Repeat(" ", 20-len(n.Name)-2*n.Level), n.Terminal, len(n.Value))
		}
	}
}

func must[T any](val T, err error) T {
	if err != nil {
		log.Println("***********************************************************")
		log.Fatal(err)
	}
	return val
}
