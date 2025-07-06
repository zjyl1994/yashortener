package main

import (
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/yashortener/infra/startup"
)

func main() {
	if err := startup.Start(); err != nil {
		logrus.Fatalln(err.Error())
	}
}
