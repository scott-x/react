/*
* @Author: scottxiong
* @Date:   2019-12-26 22:49:46
* @Last Modified by:   sottxiong
* @Last Modified time: 2019-12-30 08:24:21
 */
package engine

import (
	"fmt"

	"github.com/scott-x/gutils/cmd"
)

func genPages() {
	cmd.AddQuestion("page", "your page(default with store): ", "Please input correct page name: ", "^[a-zA-Z]+")
	answers := cmd.Exec()
	option := cmd.AddTask("which kind of page do you want?", 6, " simple", " with store", " store with loadable")
	switch option {
	case "1":
		p_task1(answers["page"])
	//anycode here ...
	case "2":
		p_task2(answers["page"])
	case "3":
		p_task3(answers["page"])
	default:
		p_task2(answers["page"])
	}
	fmt.Println(answers)
}

func genComponents() {
	cmd.AddQuestion("common_component", "your component name: ", "Please input correct component name: ", "^[a-zA-Z]+")
	answers := cmd.Exec()
	option := cmd.AddTask("which kind of component do you want?", 6, " simple", " with store", " store with loadable")
	switch option {
	case "1":
		c_task1(answers["common_component"])
	//anycode here ...
	case "2":
		c_task2(answers["common_component"])
	case "3":
		c_task3(answers["common_component"])
	default:
		c_task2(answers["common_component"])
	}
}

func addComponent() {
	cmd.AddQuestion("component", "your component name: ", "Please input correct component name: ", "^[a-zA-Z]+")
	answers := cmd.Exec()
	option := cmd.AddTask("which kind of component do you want?", 6, " simple", " with store")
	switch option {
	case "1":
		a_task1(answers["component"])
	//anycode here ...
	case "2":
		a_task2(answers["component"])
	default:
		a_task2(answers["component"])
	}
}
