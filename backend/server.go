package main

import (
	"io/ioutil"
	"os/exec"
	"strings"

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
			return
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error": err.Error(),
			})
			return
		}

		err = ioutil.WriteFile("tmp/image.jpg", data, 0644)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error": err.Error(),
			})
			return
		}

		cmd := exec.Command("python", command, "./tmp/image.jpg")
		out, err := cmd.Output()
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error": err.Error(),
			})
			return
		}
		from := strings.Index(string(out), "[[")
		to := strings.Index(string(out), "]]")
		score := string(out)[from+2 : to]
		score = strings.Trim(score, " ")
		c.StatusCode(iris.StatusOK)
		c.JSON(iris.Map{
			"score": score,
		})
	})

	app.Run(iris.Addr(":3600"))
}
