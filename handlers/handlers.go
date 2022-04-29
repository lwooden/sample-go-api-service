package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"sample-go-api-service/models"

	"github.com/gin-gonic/gin"
)

// In the Go GIN WebFramework, controllers are more like like Handlers;

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Sample Go service is healthy!",
	})
}

func GenerateMessage(c *gin.Context) {

	val, doesExist := os.LookupEnv("ENV")

	if doesExist == false {
		c.JSON(http.StatusOK, gin.H{
			"message": "Looks like your ENV enviromnet variable is not set! I don't know wehre I am!",
		})
	} else if val == "C2S" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Looks like you don't have access to the internet! All is working here!",
		})

	} else if val == "Public" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Looks like you have access to the internet! Head over to the /public internet and get some Pokemon facts!",
		})

	}

}

func FetchCatFacts(c *gin.Context) {
	println("Preparing HTTP Client")
	url := "https://cat-fact.herokuapp.com/facts/random"

	// Set up the client
	client := http.Client{}

	// Prep the request
	println("Preparing the request")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		println(err)
	}

	// Invoke the request
	println("Sending the request")
	res, getErr := client.Do(req)

	if getErr != nil {
		println(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}
	// Response comes back in the form of an address to a location in memory -- Response => 0xc000506000
	println("Response =>", res)

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	// Same for the body -- Body => (0x14502f0,0xc000474d00)
	println("Body =>", res.Body)

	// Instantiate empty CatFact object
	payload := models.CatFact{}

	// Unmarshall JSON body into CatFact object
	jsonErr := json.Unmarshal(body, &payload)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(payload.ID)

	c.IndentedJSON(http.StatusOK, payload)

}
