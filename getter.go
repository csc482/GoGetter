//golang theme: Railscasts Black Improved

///////////////////////////////[LET'S GET THIS BREAD]/////////////////////////////////

package main

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "fmt"
  //"os"
  loggly "github.com/jamespearly/loggly"
  "time"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main(){


  type Currently struct{

    //Time int `json:"time"`
    Summary string `json:"summary"`
    //Icon string `json:"icon"`
    //NearestStorm int `json:"nearestStormDistance"`
    //PrecipIntensity float32 `json:"precipIntensity"`
    //PrecipIntensityErr float32 `json:"precipIntensityError"`
    //PrecipProb float32 `json:"precipProbability"`
    PType string `json:"precipType"`
    Temp float32 `json:"temperature"`
    FeelsLike float32 `json:"apparentTemperature"`
    DewPt float32 `json:"dewPoint"`
    Humidity float32 `json:"humidity"`
    Pressure float32 `json:"pressure"`
    WindSpd float32 `json:"windSpeed"`
    //WindGst float32 `json:"windGust"`
    WindDir int `json:"windBearing"`
    //CloudCvr float32 `json:"cloudCover"`
    //UV int `json:"uvIndex"`
    //Visibility float32 `json:"visibility"`
    //Ozone float32 `json:"ozone"`

  }

  type Forecast struct{

    Time time.Time
    Latitude float32 `json:"latitude"`
    Longitude float32 `json:"longitude"`
    Timezone string `json:"timezone"`
    Currently Currently `json:"currently"`

  }

//[SETUP DYNAMODB]////////////////////////////////////////////////////////////

  sess := session.Must(session.NewSession(&aws.Config{
       Region: aws.String("us-east-1"),
  }))

  svc := dynamodb.New(sess)

  for {

     f1:= new(Forecast)

//[GET API KEY]///////////////////////////////////////////////////////////////

    //key := os.Getenv("DARKSKY") //doesn't work?? Need system restart??

    resp, err := http.Get("https://api.darksky.net/forecast/1f960db8bf1129b90c3ee6e265c92924/47.8267,-122.4233")

    if err == nil{

//[GET RESPONSE]/////////////////////////////////////////////////////////////

      body, err := ioutil.ReadAll(resp.Body)
      defer resp.Body.Close()

      if err == nil {

//[UNMARSHALLING]////////////////////////////////////////////////////////////

        err := json.Unmarshal(body, &f1)
        if err == nil {

        } else { fmt.Println(err) }

        fmt.Printf("%+v\n", f1)

//[LOGGLY]///////////////////////////////////////////////////////////////////
      
        var tag string
        tag = "GoGetter"
        breadGetter := loggly.New(tag)
        echo := breadGetter.EchoSend("info", "Successful API Pull!")
        fmt.Println(echo)

      }

    }

//[WRITING TO DYNAMODB]//////////////////////////////////////////////////////

  item := Forecast{
       Time: time.Now(),
       Latitude: f1.Latitude,
       Longitude: f1.Longitude,
  }

  mm, err := dynamodbattribute.MarshalMap(item)

  input := &dynamodb.PutItemInput{
        Item: mm,
        TableName: aws.String("DarkSky"),
  }

  result, err := svc.PutItem(input)
  if err != nil {
    fmt.Println(err)
  } else {
  fmt.Println(result)
  }

  time.Sleep(10 * time.Second)
  }

}
