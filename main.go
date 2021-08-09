package main

import (
	"Blog/model"
	router "Blog/routes"
)

func main() {
	//引用数据库
	model.InitDb()

	router.InitRouter()
}
