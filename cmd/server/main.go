package main

import (
	"fmt"
	"github.com/noctivs/smart-community-server/internal/database"
	"github.com/noctivs/smart-community-server/internal/router"
	"github.com/noctivs/smart-community-server/internal/utils"
)

func main() {
	utils.InitConfig()
	database.InitPostgres()

	r := router.InitRouter()
	port := utils.Config.GetInt("server.port")

	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
