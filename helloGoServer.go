package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type HelloController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world\n")
	this.Ctx.WriteString("FUCK YOU ALL!")
}

func (this *HelloController) Get() {
	this.Ctx.WriteString("Hello, World")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Router("/hello", &HelloController{})
	beego.Run()
}
