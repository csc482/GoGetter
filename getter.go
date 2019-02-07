/* TODO: 
   >> periodically collect formatted data via API requests
      [UPDATE: this will be done using linux "watch"]
   >> display this data on the console
   >> report results (or errors) to Loggly.
*/

///////////////////////////////[LET'S GET THIS BREAD]/////////////////////////////////

package main

import "net/http"
//import "log"
import "io/ioutil"
import "encoding/json"
import "fmt"

func main(){

  type Currently struct{

    Time string `json:"time"`
    Summary string `json:"summary"`
    Icon string `json:"icon"`
    NearestStorm string `json:"nearestStormDistance"`
    PrecipIntensity string `json:"precipIntensity"`
    PrecipIntensityErr string `json:"precipIntensityError"`
    PrecipProb string `json:"precipProbability"`
    PType string `json:"precipType"`
    Temp string `json:"temperature"`
    FeelsLike string `json:"apparentTemperature"`
    DewPt string `json:"dewPoint"`
    Humidity string `json:"humidity"`
    Pressure string `json:"pressure"`
    WindSpd string `json:"windSpeed"`
    WindGst string `json:"windGust"`
    WindDir string `json:"windBearing"`
    CloudCvr string `json:"cloudCover"`
    UV string `json:"uvIndex"`
    Visibility string `json:"visibility"`
    Ozone string `json:"ozone"`

  }


  type Forecast struct{

    Latitude string `json:"latitude"`
    Longitude string `json:"longitude"`
    Timezone string `json:"timezone"`
    Currently Currently `json:"currently"`

  }

  var f1 Forecast
  resp, err := http.Get("https://api.darksky.net/forecast/1f960db8bf1129b90c3ee6e265c92924/47.8267,-122.4233")

  if err == nil{

    if resp != nil{

      body, err := ioutil.ReadAll(resp.Body) //get response
      defer resp.Body.Close()

      if err == nil {

        fmt.Println("wow")
        //log.Println(string(body)) //print response
        json.Unmarshal(body, &f1)
        fmt.Printf("%+v\n", f1)

      }

    }

  }

}
