//start outScope OMIT
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/howetze/slides/api/dht" // HL
)

//end outScope OMIT

//start inScope OMIT

func main() {
	//start response OMIT
	response, err := http.Get("http://api.luftdaten.info/v1/sensor/6940/") // HL

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()
	//end response OMIT
	//start tmp OMIT
	var tmp dht.Result // HL

	json.NewDecoder(response.Body).Decode(&tmp)
	//end tmp OMIT

	//start print OMIT
	fmt.Println("Temperatur auf meinem Balkon:", tmp[1].Sensordatavalues[0].Value, "°C") // HL
	//end print OMIT
}

//end inScope OMIT

//fmt.Printf("%T\n%v\n\n", response, response)
//fmt.Printf("%T\n%v\n\n", response.Body, response.Body)

// body, _ := ioutil.ReadAll(response.Body)
//
// var r result
// json.Unmarshal(body, &r)
// fmt.Println(r.Data.Amount)
