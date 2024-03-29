package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Allenxuxu/mogutouERP/models"
	"github.com/Allenxuxu/mogutouERP/pkg/token"
	"github.com/gin-gonic/gin"
	config "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	// "github.com/pkg/browser"
)

func main() {
	fmt.Println("main task start running")
	path := flag.String("c", "/etc/conf", "配置文件夹路径")
	flag.Parse()
	fmt.Println(*path)

	token.InitConfig(*path+"/jwt.json", "jwt-key")
	//read config
	fileSource := file.NewSource(
		file.WithPath(*path + "/conf.json"),
	)
	conf := config.NewConfig()
	err := conf.Load(fileSource)
	if err != nil {
		log.Fatal(err)
	}

	var info models.DBInfo
	err = conf.Get("mysql").Scan(&info)
	fmt.Println(*path)
	if err != nil {
		log.Fatal(err)
	}
	models.Init(&info)

	gin.DisableConsoleColor()
	// gin.SetMode(gin.ReleaseMode)
	r := initRouter()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Open: http://127.0.0.1:5000")
	}()

	r.Run(":5000") // listen and serve on 0.0.0.0:8080
}
