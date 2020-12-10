package test_init

import (
	"log"
	config "web-service/src/config"

	containers "web-service/src/storage_container"
	utils "web-service/src/utils"
)

var (
	errorLogger *log.Logger
	debugLogger *log.Logger
)

func ConfigureTestingEnvironmentForS3Support() {
	if err := config.ReadConfig(config.ConfigPath); err != nil {
		log.Fatalln("Unable to read program config from", config.ConfigPath, err)
	}

	if err := utils.InitializeEmptyDir(config.Internal.TempFilesDir); err != nil {
		log.Fatalln("Unable to initialize temporary directory at", config.Internal.TempFilesDir, err)
	}

	if err := utils.InitializeLogger(config.Internal.LoggingDir); err != nil {
		log.Fatalln("Unable to initialize logging at", config.Internal.LoggingDir, err)
	}

	errorLogger = utils.ErrorLogger
	debugLogger = utils.DebugLogger

	db := containers.NewDB()
	if err := db.Initialize(config.Internal.DbPath); err != nil {
		errorLogger.Fatalln("Unable to open db at", config.Internal.DbPath, err)
	}

	// api.InitializeControllers(db)
}
