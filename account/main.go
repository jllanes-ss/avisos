package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/jllanes-ss/avisos/account/config"
	// "github.com/jllanes-ss/avisos/account/config"
	//_ "github.com/go-sql-driver/mysql"
	//remove esta chingada
	// config "github.com/jllanes-ss/avisos/account/config"
	//datasources "github.com/jllanes-ss/avisos/account/user/repository"
	//_userRepository "github.com/jllanes-ss/avisos/account/user/repository/mysql"
)

//main function
//@todo clean main and mov the functionality to an injection function
func main() {
	log.Println("Starting app...")

	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Unable to load .env file: %v\n", err)
	}

	ds, err := config.GetDS()

	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	app, err := config.Inject(ds)

	if err != nil {
		log.Fatalf("Failure to inject data sources: %v\n", err)
	}

	config.SetMiddleware(app)

	//app := fiber.New()

	// //log.SetFormatter(&log.JSONFormatter{})
	// timeoutContext := time.Duration(viper.GetInt("HANDLER_TIMEOUT")) * time.Second

	// ur := repository.GetRepository(ds.DB)
	// //ur := _userRepository.NewMysqlUserRepository(ds.DB)
	// uuc := _userUseCase.NewUserUseCase(ur, timeoutContext)

	// _userHttpDelivery.NewUserHandler(app, uuc)
	// ...

	//closing connections and shutdown pplication
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("todo send mail when panic...")
		_ = app.Shutdown()
	}()

	log.Println("Starting server...")
	if err := app.Listen(":8088"); err != nil {
		log.Panic(err)
	}

	// shutdown data sources
	log.Println("Shutting down DS...")
	if err := ds.Close(); err != nil {
		log.Fatalf("A problem occurred gracefully shutting down data sources: %v\n", err)
	}
	// Your cleanup tasks go here
}
