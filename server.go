package main

import (
    "encoding/json"
    "time"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {

fmt.Println("Starting...")
    type Currently struct{
        Summary string `json:"summary"`
        PType string `json:"precipType"`
        Temp float32 `json:"temperature"`
        FeelsLike float32 `json:"apparentTemperature"`
        DewPt float32 `json:"dewPoint"`
        Humidity float32 `json:"humidity"`
        Pressure float32 `json:"pressure"`
        WindSpd float32 `json:"windSpeed"`
        WindDir int `json:"windBearing"`
    }

    type Forecast struct{
        Time time.Time
        Latitude float32 `json:"latitude"`
        Longitude float32 `json:"longitude"`
        Timezone string `json:"timezone"`
        Currently Currently `json:"currently"`
    }
    
    r := mux.NewRouter()

    sess := session.Must(session.NewSession(&aws.Config{
       Region: aws.String("us-east-1"),
    }))

    svc := dynamodb.New(sess)

    r.HandleFunc("/agoldste/all", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("all handler")
        params := &dynamodb.ScanInput{
		    TableName: aws.String("DarkSky"),
        }
        result, err := svc.Scan(params)
	    if err != nil {
	        fmt.Println(err)
		}
		items := []Forecast{}
		err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
	    if err != nil {
		    fmt.Println(err)
        }
        json.NewEncoder(w).Encode(items)
        fmt.Printf("%+v\n", items)
    })

    r.HandleFunc("/agoldste/status", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("status handler")
        req := &dynamodb.DescribeTableInput{
            TableName: aws.String("DarkSky"),
        }
        result, err := svc.DescribeTable(req)
        if err != nil {
            fmt.Printf("%s", err)
        }
        table := result.Table
        fmt.Fprintf(w, "", table)
    })

        http.ListenAndServe(":8080", r)
}
