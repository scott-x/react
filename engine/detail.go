/*
* @Author: scottxiong
* @Date:   2019-12-26 23:53:42
* @Last Modified by:   sottxiong
* @Last Modified time: 2019-12-30 08:35:12
 */
package engine

import (
	"strconv"
	"strings"

	"github.com/scott-x/gutils/cl"
	"github.com/scott-x/gutils/cmd"
	"github.com/scott-x/gutils/fs"
	"github.com/scott-x/gutils/str"
)

var template string

func init() {
	str := fs.GetEnv("GOPATH")
	template = str + "/src/github.com/scott-x/react-templates"
}

func initProject() {
	cmd.AddQuestion("project", "your project name: ", "Please input correct project name: ", "^[a-zA-Z]+")
	answers := cmd.Exec()
	myApp := strings.Trim(answers["project"], " ")
	err := fs.CopyFolder(template, "./"+myApp)
	if err != nil {
		panic(err)
	}
	rep := make(map[string]string, 0)
	rep["jianshu"] = myApp
	err = fs.ReadAndReplace("./"+myApp+"/package.json", rep)
	if err != nil {
		panic(err)
	}
	projectDetails()
	cl.BoldGreen.Printf("cd %s && yarn install\n", myApp)
}

//page
func p_task1(page string) {
	p := strings.Trim(page, " ")
	err := fs.CopyFolder(template+"/src/pages/write", "./src/pages/"+strings.ToLower(p))
	if err != nil {
		panic(err)
	}
	rep := make(map[string]string, 0)
	rep["Write"] = str.FirstLetterToUpper(p, 1)
	err = fs.ReadAndReplace("./src/pages/"+strings.ToLower(p)+"/index.js", rep)
	if err != nil {
		panic(err)
	}
	err = fs.ReadAndReplace("./src/pages/"+strings.ToLower(p)+"/style.js", rep)
	if err != nil {
		panic(err)
	}
}

func p_task2(page string) {
	p := strings.Trim(page, " ")
	err := fs.CopyFolder(template+"/src/pages/login", "./src/pages/"+strings.ToLower(p))
	if err != nil {
		panic(err)
	}
	rep := make(map[string]string, 0)
	rep["Login"] = str.FirstLetterToUpper(p, 1)
	err = fs.ReadAndReplace("./src/pages/"+strings.ToLower(p)+"/index.js", rep)
	if err != nil {
		panic(err)
	}

	err = fs.ReadAndReplace("./src/pages/"+strings.ToLower(p)+"/style.js", rep)
	if err != nil {
		panic(err)
	}
	cmd.Info("src/pages/" + strings.ToLower(p) + "/index.js was created")
	cmd.Info("src/pages/" + strings.ToLower(p) + "/style.js was created")
	cmd.Info("src/pages/" + strings.ToLower(p) + "/store/actionCreator.js was created")
	cmd.Info("src/pages/" + strings.ToLower(p) + "/store/actionType.js was created")
	cmd.Info("src/pages/" + strings.ToLower(p) + "/store/reducer.js was created")
	cmd.Info("src/pages/" + strings.ToLower(p) + "/store/index.js was created")
}

func p_task3(page string) {
	p := strings.Trim(page, " ")
	err := fs.CopyFolder(template+"/src/pages/detail", "./src/pages/"+strings.ToLower(p))
	if err != nil {
		panic(err)
	}
	rep := make(map[string]string, 0)
	rep["Login"] = str.FirstLetterToUpper(p, 1)
	err = fs.ReadAndReplace("./src/pages/"+strings.ToLower(p)+"/index.js", rep)
	if err != nil {
		panic(err)
	}

	err = fs.ReadAndReplace("./src/pages/"+strings.ToLower(p)+"/style.js", rep)
	if err != nil {
		panic(err)
	}
	cmd.Info("src/pages/" + strings.ToLower(p) + "/index.js was created")
	cmd.Info("src/pages/" + strings.ToLower(p) + "/style.js was created")
	cmd.Info("src/pages/" + strings.ToLower(p) + "/loadable.js was created")
	cmd.Info("src/pages/" + strings.ToLower(p) + "/store/actionCreator.js was created")
	cmd.Info("src/pages/" + strings.ToLower(p) + "/store/actionType.js was created")
	cmd.Info("src/pages/" + strings.ToLower(p) + "/store/reducer.js was created")
	cmd.Info("src/pages/" + strings.ToLower(p) + "/store/index.js was created")
}

//add
func a_task1(c string) {
	folders := fs.List("./src/pages")
	option := cmd.AddTask("Please select the page:", 7, folders...)

	rep := make(map[string]string, 0)
	rep["Writer"] = str.FirstLetterToUpper(c, 1)
	component := strings.Trim(c, " ")
	num, _ := strconv.Atoi(option)
	fs.CreateDirIfNotExist("./src/pages/" + folders[num-1] + "/components")
	fs.CopyAndReplace(template+"/src/pages/home/components/Writer.js", "./src/pages/"+folders[num-1]+"/components/"+str.FirstLetterToUpper(component, 1)+".js", rep)
	cmd.Info("src/pages/" + folders[num-1] + "/" + str.FirstLetterToUpper(component, 1) + ".js was created")
}

func a_task2(c string) {
	folders := fs.List("./src/pages")
	option := cmd.AddTask("Please select the page:", 7, folders...)

	rep := make(map[string]string, 0)
	rep["Writer"] = str.FirstLetterToUpper(c, 1)
	component := strings.Trim(c, " ")
	num, _ := strconv.Atoi(option)
	fs.CreateDirIfNotExist("./src/pages/" + folders[num-1] + "/components")
	fs.CopyAndReplace(template+"/src/pages/home/components/List.js", "./src/pages/"+folders[num-1]+"/components/"+str.FirstLetterToUpper(component, 1)+".js", rep)
	cmd.Info("src/pages/" + folders[num-1] + "/" + str.FirstLetterToUpper(component, 1) + ".js was created")
}

//common
func c_task1(c string) {
	p := strings.Trim(c, " ")
	err := fs.CopyFolder(template+"/src/common/loading", "./src/common/"+strings.ToLower(p))
	if err != nil {
		panic(err)
	}
	rep := make(map[string]string, 0)
	rep["Loading"] = str.FirstLetterToUpper(p, 1)
	err = fs.ReadAndReplace("./src/common/"+strings.ToLower(p)+"/index.js", rep)
	if err != nil {
		panic(err)
	}
	err = fs.ReadAndReplace("./src/common/"+strings.ToLower(p)+"/style.js", rep)
	if err != nil {
		panic(err)
	}
}

func c_task2(c string) {
	p := strings.Trim(c, " ")
	err := fs.CopyFolder(template+"/src/common/header", "./src/common/"+strings.ToLower(p))
	if err != nil {
		panic(err)
	}
	rep := make(map[string]string, 0)
	rep["Header"] = str.FirstLetterToUpper(p, 1)
	err = fs.ReadAndReplace("./src/common/"+strings.ToLower(p)+"/index.js", rep)
	if err != nil {
		panic(err)
	}

	err = fs.ReadAndReplace("./src/common/"+strings.ToLower(p)+"/style.js", rep)
	if err != nil {
		panic(err)
	}
	cmd.Info("src/common/" + strings.ToLower(p) + "/index.js was created")
	cmd.Info("src/common/" + strings.ToLower(p) + "/style.js was created")
	cmd.Info("src/common/" + strings.ToLower(p) + "/store/actionCreator.js was created")
	cmd.Info("src/common/" + strings.ToLower(p) + "/store/actionType.js was created")
	cmd.Info("src/common/" + strings.ToLower(p) + "/store/reducer.js was created")
	cmd.Info("src/common/" + strings.ToLower(p) + "/store/index.js was created")
}

func c_task3(c string) {
	c_task2(c)
	p := strings.Trim(c, " ")
	cmd.Info("src/common/" + strings.ToLower(p) + "/loadable.js was created")
}

func projectDetails() {
	files := []string{
		"package.json",
		"yarn.lock",
		"yarn-error.log",
		"README.md",
		".gitignore",
		"public",
	}
	for _, v := range files {
		cmd.Info(v)
	}
}
