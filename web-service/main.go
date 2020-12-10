package main

import (
	"net/http"

	api "web-service/src/api/controllers"
	config "web-service/src/config"
	initialize "web-service/src/init"
	utils "web-service/src/utils"
)

func setupRoutes() {
	http.HandleFunc("/api/upload_files", api.UploadFilesHandler)
	http.HandleFunc("/api/view/", api.ViewRoomHandler)
	http.HandleFunc("/api/link", api.GetFileLinkById)
	http.HandleFunc("/api/cmp_files", api.CompareFilesHandler)
}

func main() {
	initialize.Configure()

	setupRoutes()

	utils.DebugLogger.Println("Starting fair online judge service on", config.Server.Port)
	utils.ErrorLogger.Fatalln(http.ListenAndServe(config.Server.Port, nil))
}
