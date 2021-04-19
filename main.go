package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mahjadan/go-integration-test/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var dbClient *mongo.Client

func main() {
	var err error
	dbClient, err = setupDB()
	if err != nil {
		panic(fmt.Sprintf("can not connect to mongoDB: %s", err.Error()))
	}
	e := echo.New()

	e.GET("/resource-status", getResourceStatus)
	e.GET("/planets/:id", getPlanetFromAPI)
	e.GET("/people/:id", getPeopleFromDB)

	e.Logger.Fatal(e.Start(":8080"))
}

func getResourceStatus(c echo.Context) error {
	return c.JSON(http.StatusOK,echo.Map{"status":"OK"})
}

func getPeopleFromDB(c echo.Context) error {
	id := c.Param("id")
	collection := dbClient.Database("test_db").Collection("people")
	var result domain.People
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "resource not found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "db error: " + err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func getPlanetFromAPI(c echo.Context) error {
	id := c.Param("id")

	url := getStartWarURI() + id
	ctx, _ := context.WithTimeout(c.Request().Context(), 5*time.Second)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "can not create request: " + err.Error()})
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("can not make request to %s: %s", url, err.Error())})
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "can not response body: " + err.Error()})
	}

	defer response.Body.Close()
	var result interface{}
	json.Unmarshal(bytes, &result)
	return c.JSON(http.StatusOK, result)
}

func getStartWarURI() string {
	uri := os.Getenv("STAR_WAR_URI")
	if uri == "" {
		return "http://swapi.dev/api/planets/"
	}
	return uri
}

