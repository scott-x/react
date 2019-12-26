/*
* @Author: scottxiong
* @Date:   2019-12-26 22:35:32
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-12-26 23:49:37
 */
package engine

import (
	_ "fmt"
	"github.com/scott-x/gutils/cl"
	"os"
)

func Run() {
	handleArgs()
}

func handleArgs() {
	//args == 1
	if len(os.Args) == 1 {
		printUseage()
	}

	//args == 2
	if len(os.Args) == 2 {
		if os.Args[1] == "init" || os.Args[1] == "i" {
			initProject()
		} else if os.Args[1] == "g" {
			argsNotMatched("react g 【c/p】 ")
		} else {
			printUseage()
		}

	}

	//args == 2
	if len(os.Args) == 3 {
		if os.Args[1] == "g" {
			if os.Args[2] == "c" {
				genComponents()
			} else if os.Args[2] == "p" {
				genPages()
			} else {
				argsNotMatched("react g 【c/p】 ")
			}
		} else {
			invalidArgs("react g 【c/p】 ")
		}

		if os.Args[1] == "add" {
			if os.Args[2] == "c" {
				addComponent()
			} else {
				argsNotMatched("react add c")
			}
		}

	}

}

func printUseage() {
	cl.BoldGreen.Printf("Full Commonds Are: \n")
	cl.BoldCyan.Println("react 【i/init】")
	cl.BoldCyan.Println("react g c")
	cl.BoldCyan.Println("react g p")
}
func argsNotMatched(tip string) {
	cl.BoldRed.Printf("arguments not matched, please use ")
	cl.BoldCyan.Printf("%s \n", tip)
}

func invalidArgs(tip string) {
	cl.BoldRed.Printf("invalid argument, please use ")
	cl.BoldCyan.Printf("%s \n", tip)
}

func missing_necessary_args(tip string) {
	cl.BoldRed.Printf("missing necessary arguments, please use ")
	cl.BoldCyan.Printf("%s \n", tip)
}
