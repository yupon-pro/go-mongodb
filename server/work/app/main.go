package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"work/infrastructure"
	"work/interface"
	"work/usecase"

	"github.com/labstack/echo/v4"
)


func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mydb := new(infrastructure.MyMongoDB)
	if err := mydb.Connect(ctx); err != nil{
		log.Fatal(err)
	}
	defer mydb.Disconnect(ctx)

	if err := mydb.Ping(ctx); err != nil{
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context)error{return c.String(http.StatusOK, "hello")})

	receiptRepository := infrastructure.NewReceiptRepositoryInfrastructure(mydb, ctx)
	receiptUsecase := usecase.NewReceiptUsecase(receiptRepository)
	receiptController := interfaces.NewReceiptController(receiptUsecase)
	controllers := interfaces.NewControllers(receiptController)
	controllers.Mount(e)

	e.Logger.Fatal(e.Start(":8080"))

}

