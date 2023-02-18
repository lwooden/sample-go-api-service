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

func GetEcho(c *gin.Context) {

	term := c.Query("term") // access a query string parameter
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("You said %s", term),
	})
}

func GetEnvironment(c *gin.Context) {

	val, doesExist := os.LookupEnv("ENV")

	if doesExist {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Looks like your are running in a %s environment", val),
		})
	} else {

		c.JSON(http.StatusNotFound, "No environment has been set!")

	}
}

func Sum(c *gin.Context) {

	var values models.Sum // initialize Sum object that I am going to bind the payload data to
	c.Bind(&values)

	sum := values.Val_1 + values.Val_2

	c.JSON(http.StatusOK, gin.H{
		"sum": sum,
	})

}

func GetPodInfo(c *gin.Context) {

	nodeName, exist := os.LookupEnv("NODE_NAME")
	podName, _ := os.LookupEnv("POD_NAME")
	podIP, _ := os.LookupEnv("POD_IP")
	podNamespace, _ := os.LookupEnv("POD_NAMESPACE")
	podServiceAccount, _ := os.LookupEnv("POD_SERVICE")

	if !exist {

		nodeName = "Not in a Kubernetes Environment. Can't tell ya!"
		podName = "Not in a Kubernetes Environment. Can't tell ya!"
		podIP = "Not in a Kubernetes Environment. Can't tell ya!"
		podNamespace = "Not in a Kubernetes Environment. Can't tell ya!"
		podServiceAccount = "Not in a Kubernetes Environment. Can't tell ya!"

	}

	pod := models.PodMetadata{nodeName, podName, podIP, podNamespace, podServiceAccount}

	c.JSON(http.StatusOK, pod)

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

func GetPokemom(c *gin.Context) {

	id := c.Param("id")
	println("Pokemon ID =>", id)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		println(err)
	}

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

	payload := models.Pokemon{}

	jsonErr := json.Unmarshal(body, &payload)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(payload.ID)

	c.IndentedJSON(http.StatusOK, payload)

}
