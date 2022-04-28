package main

/*
// Author: Aliasgar Khimani (NovusEdge)
// Project: github.com/NovusEdge/buraq
//
// Copyright: GNU General Public License v3.0
// See the LICENSE file for more info.
//
// All Rights Reserved
*/

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	src "github.com/NovusEdge/buraq/src"
	utils "github.com/NovusEdge/buraq/utils"
)

// ENV is the environment variables for the program
var ENV = utils.GetEnv()

func main() {
	checkCommand(os.Args)
	executeCommand(os.Args)
}

func checkCommand(args []string) {
	var ok bool
	if len(args) < 2 {
		help()
		os.Exit(0)
	}

	cmd := args[1]
	for _, c := range src.ValidCommands() {
		if cmd == c {
			ok = true
		}
	}

	if !ok {
		fmt.Println(utils.ColorIt(utils.ColorRed, "[E]: Invalid command!\nUse \"buraq help\" to show the usage for buraq."))
		os.Exit(1)
	}

}

func executeCommand(args []string) {
	var out bytes.Buffer

	command := args[1]
	commandString := filepath.Join(ENV["BURAQCMDBIN"], command)

	cmd := exec.Command(commandString, args...)
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if s := out.String(); s != "" {
		fmt.Println(s)
	}
}

func env() {
	var out bytes.Buffer
	cmd := exec.Command(fmt.Sprintf("%s/env", ENV["BURAQCMDBIN"]))
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
}

func help() {
	var out bytes.Buffer
	cmd := exec.Command(fmt.Sprintf("%s/help", ENV["BURAQCMDBIN"]))
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
}
