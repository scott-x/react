/*
* @Author: scottxiong
* @Date:   2019-12-26 22:49:46
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-12-27 00:50:41
 */
package engine

import (
	"fmt"
	"github.com/scott-x/gutils/cmd"
)

func genPages() {
	cmd.AddQuestion("page", "your page(default with store): ", "Please input correct page name: ", "^[a-zA-Z]+")
	answers := cmd.Exec()
	option := cmd.AddTask("your close", 7, " simple", " with store")
	switch option {
	case "1":
		p_task1()
	//anycode here ...
	case "2":
		p_task2()
	default:
		p_task2()
	}
	fmt.Println(answers)
}

func genComponents() {
	cmd.AddQuestion("common_component", "your component name: ", "Please input correct component name: ", "^[a-zA-Z]+")
}

func addComponent() {
	cmd.AddQuestion("component", "your component name: ", "Please input correct component name: ", "^[a-zA-Z]+")
}

func initProject() {
	cmd.AddQuestion("project", "your project name: ", "Please input correct project name: ", "^[a-zA-Z]+")
}
