package main

import(
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	counts := make(map[string]int)
	for _,file := range os.Args[1:]{
		text,err := ioutil.ReadFile(file)
		if err != nil{
			fmt.Print("error")
			continue
		}
		for _,line := range strings.Split(string(text),"\n") {
			fmt.Printf("%s\n",line)
			counts[line]++
		}
	}

	for temp,n := range counts {
		fmt.Printf("%d\t%s\n",n,temp)
	}
}