package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/kataras/iris"
)

const command = "./core/predictor.py"

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./pages", ".html").Reload(true))
	app.StaticWeb("/statics", "./statics")

	app.Get("/", func(c iris.Context) {
		c.View("index.html")
	})
	app.Post("/eval_photo", func(c iris.Context) {

		file, _, err := c.FormFile("image")
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error": err.Error(),
			})
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error": err.Error(),
			})
		}
		file.Close()

		err = ioutil.WriteFile("tmp/image.jpg", data, os.ModeDir)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error": err.Error(),
			})
		}

		cmd := exec.Command("python", command, "./tmp/image.jpg")
		out, err := cmd.Output()
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error": err.Error(),
			})
		}

		log.Println(string(out))
	})

	app.Run(iris.Addr(":3600"))
}
