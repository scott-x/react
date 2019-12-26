/*
* @Author: scottxiong
* @Date:   2019-12-26 22:35:32
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-12-27 03:42:10
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
		} else if os.Args[1] == "add" {
			argsNotMatched("react add c ")
		} else if os.Args[1] == "d" {
			argsNotMatched("react d i ")
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
		} else if os.Args[1] == "add" {
			if os.Args[2] == "c" {
				addComponent()
			} else {
				argsNotMatched("react add c")
			}
		} else if os.Args[1] == "d" {
			if os.Args[2] == "i" {
				dealWithIconfont()
			} else {
				argsNotMatched("react d i")
			}
		} else {
			invalidArgs("react g 【c/p】 ")
		}

	}

}

func printUseage() {
	cl.BoldRed.Printf("Command\tFunc\tType\tDesc\n")
	cl.BoldCyan.Printf("react\ti/init\tnull\tinit react project\n")
	cl.BoldCyan.Printf("react\tg\tc\tgenerator component\n")
	cl.BoldCyan.Printf("react\tg\tp\tgenerator page\n")
	cl.BoldCyan.Printf("react\tadd\tc\tpage's component\n")
	cl.BoldCyan.Printf("react\td\ti\tdeal with iconfont\n")
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
