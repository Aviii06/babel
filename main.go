package main

import "os"
import "babel/pkg/repl"

func main() {
    repl.Start(os.Stdin, os.Stdout)
}


