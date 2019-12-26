/*
* @Author: scottxiong
* @Date:   2019-12-27 03:33:42
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-12-27 03:44:20
 */
package engine

import (
	"fmt"
	"github.com/scott-x/gutils/cl"
	"github.com/scott-x/gutils/cmd"
	"github.com/scott-x/gutils/fs"
	"os"
	"strings"
)

func dealWithIconfont() {
	cmd.AddQuestion("path", "The iconfont folder path: ", "Please drag the folder here: ", "^/.*font_.*")
	answers := cmd.Exec()
	path := strings.Trim(answers["path"], " ")
	flag := runTask(path)
	if flag {
		cl.BoldCyan.Printf("Done, you can use iconfont in your React project now...\n")
	} else {
		cl.BoldRed.Printf("The file is incompleted, please double check, then try agagin...\n")
	}

}

func runTask(path string) bool {
	fs.CreateDirIfNotExist("./src/statics/iconfont")
	var flag bool = true
	if fs.IsExist(path + "/iconfont.css") {
		fs.Copy(path+"/iconfont.css", "./src/statics/iconfont/iconfont.css")
	} else {
		flag = false
		fmt.Printf("not found: iconfont.css \n")
	}
	if fs.IsExist(path + "/iconfont.eot") {
		fs.Copy(path+"/iconfont.eot", "./src/statics/iconfont/iconfont.eot")
	} else {
		flag = false
		fmt.Printf("not found: iconfont.eot \n")
	}
	if fs.IsExist(path + "/iconfont.svg") {
		fs.Copy(path+"/iconfont.svg", "./src/statics/iconfont/iconfont.svg")
	} else {
		flag = false
		fmt.Printf("not found: iconfont.svg \n")
	}
	if fs.IsExist(path + "/iconfont.ttf") {
		fs.Copy(path+"/iconfont.ttf", "./src/statics/iconfont/iconfont.ttf")
	} else {
		flag = false
		fmt.Printf("not found: iconfont.ttf \n")
	}
	if fs.IsExist(path + "/iconfont.woff") {
		fs.Copy(path+"/iconfont.woff", "./src/statics/iconfont/iconfont.woff")
	} else {
		flag = false
		fmt.Printf("not found: iconfont.woff \n")
	}
	if fs.IsExist(path + "/iconfont.woff2") {
		fs.Copy(path+"/iconfont.woff2", "./src/statics/iconfont/iconfont.woff2")
	} else {
		flag = false
		fmt.Printf("not found: iconfont.woff2 \n")
	}

	content, err := fs.ReadFile1("./src/statics/iconfont/iconfont.css")
	if err != nil {
		panic(err)
	}
	var newContent []string
	newContent = append(newContent, "import { createGlobalStyle } from 'styled-components';")
	newContent = append(newContent, "")
	newContent = append(newContent, "export const GlobalIcon = createGlobalStyle`")
	arr := strings.Split(content, "\n")
	for x, c := range arr {
		if x > 15 {
			break
		}
		if strings.Contains(c, "url('iconfont") {
			c = strings.Replace(c, "url('iconfont", "url('./iconfont", -1)
		}
		newContent = append(newContent, c)

	}
	newContent = append(newContent, "`")
	fs.WriteString("./src/statics/iconfont/iconfont.js", strings.Join(newContent, "\n"))
	//delete iconfont.css
	err = os.Remove("./src/statics/iconfont/iconfont.css")
	if err != nil {
		panic(err)
	}
	return flag
}
