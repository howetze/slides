//start outScope OMIT
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/howetze/slides/api/dht"
)

//end outScope OMIT

//start inScope OMIT

func main() {
	//start response OMIT

	response, err := http.Get("http://api.luftdaten.info/v1/sensor/6940/")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()
	//end response OMIT

	//start tmp OMIT

	body, _ := ioutil.ReadAll(response.Body)

	var r dht.Result

	json.Unmarshal(body, &r) //exceptionally without error handling

	//end tmp OMIT

	//start print OMIT

	fmt.Println("Temperatur auf meinem Balkon:", r[1].Sensordatavalues[0].Value, "°C")

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

/* mit Decoder
//var tmp dht.Result // HL
//json.NewDecoder(response.Body).Decode(&tmp)
//fmt.Println("Temperatur auf meinem Balkon:", tmp[1].Sensordatavalues[0].Value, "°C") // HL
*/
