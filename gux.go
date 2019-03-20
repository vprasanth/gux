package main

import (
	"fmt"
	"github.com/vprasanth/gux/cmds"
	"github.com/vprasanth/gux/spec"
	"io/ioutil"
	"log"
)

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	content, err := ioutil.ReadFile("./.gux.yaml")
	checkErr(err)

	var data spec.GuxConfig

	err = data.Parse(content)
	checkErr(err)

	fmt.Printf("Read: %+v\n", data.Session)
	//.for i := 0; i < len(data.Windows[0].Panes); i++ {
	//.	fmt.Println(data.Windows[0].Panes[i].Name)
	//.}

	//cmds.Start(data.Windows[0].Name)
	//cmds.CreateVerticalSplitLayout(data.Windows[0].Name, data.Windows[0])
	cmds.Init(data)
}
