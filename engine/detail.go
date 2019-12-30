/*
* @Author: scottxiong
* @Date:   2019-12-26 23:53:42
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-12-30 18:00:47
 */
package engine

import (
	"strconv"
	"strings"

	"github.com/scott-x/gutils/cmd"
	"github.com/scott-x/gutils/fs"
	"github.com/scott-x/gutils/model"
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
	if fs.IsExist("./" + myApp) {
		cmd.Warning(myApp + " already exists")
		return
	}
	err := fs.CopyFolder(template+"/project", "./"+myApp)
	if err != nil {
		panic(err)
	}
	rep := make(map[string]string, 0)
	rep["jianshu"] = myApp
	err = fs.ReadAndReplace("./"+myApp+"/package.json", rep)
	if err != nil {
		panic(err)
	}
	cmd.Info("cd " + myApp + " && yarn install")
}

//page
func p_task1(page string) {
	p := strings.Trim(page, " ")
	if fs.IsExist("./src/pages/" + strings.ToLower(p)) {
		cmd.Warning("page " + strings.ToLower(p) + " already exists")
		return
	}
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
	cmd.Info("src/pages/" + strings.ToLower(p) + "/index.js was created")
	cmd.Info("src/pages/" + strings.ToLower(p) + "/style.js was created")
	//modify src/store/reducer.js
	modifyReducer("pages", strings.ToLower(p))
}

func p_task2(page string) {
	p := strings.Trim(page, " ")
	if fs.IsExist("./src/pages/" + strings.ToLower(p)) {
		cmd.Warning("page " + strings.ToLower(p) + " already exists")
		return
	}
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
	//modify src/store/reducer.js
	modifyReducer("pages", strings.ToLower(p))
}

func p_task3(page string) {
	p := strings.Trim(page, " ")
	if fs.IsExist("./src/pages/" + strings.ToLower(p)) {
		cmd.Warning("page " + strings.ToLower(p) + " already exists")
		return
	}
	err := fs.CopyFolder(template+"/src/pages/detail", "./src/pages/"+strings.ToLower(p))
	if err != nil {
		panic(err)
	}
	rep := make(map[string]string, 0)
	rep["Detail"] = str.FirstLetterToUpper(p, 1)
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
	//modify src/store/reducer.js
	modifyReducer("pages", strings.ToLower(p))
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
	rep["List"] = str.FirstLetterToUpper(c, 1)
	component := strings.Trim(c, " ")
	num, _ := strconv.Atoi(option)
	fs.CreateDirIfNotExist("./src/pages/" + folders[num-1] + "/components")
	fs.CopyAndReplace(template+"/src/pages/home/components/List.js", "./src/pages/"+folders[num-1]+"/components/"+str.FirstLetterToUpper(component, 1)+".js", rep)
	cmd.Info("src/pages/" + folders[num-1] + "/" + str.FirstLetterToUpper(component, 1) + ".js was created")
}

//common
func c_task1(c string) {
	p := strings.Trim(c, " ")
	if fs.IsExist("./src/common/" + strings.ToLower(p)) {
		cmd.Warning("common component " + strings.ToLower(p) + " already exists")
		return
	}
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
	cmd.Info("src/common/" + strings.ToLower(p) + "/index.js was created")
	cmd.Info("src/common/" + strings.ToLower(p) + "/style.js was created")
}

func c_task2(c string) {
	p := strings.Trim(c, " ")
	if fs.IsExist("./src/common/" + strings.ToLower(p)) {
		cmd.Warning("common component " + strings.ToLower(p) + " already exists")
		return
	}
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
	modifyReducer("common", strings.ToLower(p))
}

func c_task3(c string) {
	p := strings.Trim(c, " ")
	if fs.IsExist("./src/common/" + strings.ToLower(p)) {
		cmd.Warning("common component " + strings.ToLower(p) + " already exists")
		return
	}
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

	fs.Copy(template+"/src/pages/detail/loadable.js", "./src/common/"+strings.ToLower(p)+"/loadable.js")
	cmd.Info("src/common/" + strings.ToLower(p) + "/loadable.js was created")
	modifyReducer("common", strings.ToLower(p))

}

func projectDetails() {
	// files := []string{
	// 	"package.json",
	// 	"yarn.lock",
	// 	"yarn-error.log",
	// 	"README.md",
	// 	".gitignore",
	// 	"public",
	// }
	// for _, v := range files {
	// 	cmd.Info(v)
	// }
}

func modifyReducer(destination string, modifiedTo string) {
	//modify src/store/reducer.js
	i := &model.Insert{
		File:     "./src/store/reducer.js",
		NewLine:  "import { reducer as homeReducer } from '../" + destination + "/home/store';",
		Keywords: "'redux-immutable'",
		Replace: model.Replace{
			Old: "home",
			New: modifiedTo,
		},
	}
	fs.InsertAfter(i)
	i = &model.Insert{
		File:     "./src/store/reducer.js",
		NewLine:  fs.Tab(4) + "login: loginReducer,",
		Keywords: "header: headerReducer",
		Replace: model.Replace{
			Old: "login",
			New: modifiedTo,
		},
	}
	fs.InsertBefore(i)
	cmd.Warning("src/store/reducer.js was modified")
}
