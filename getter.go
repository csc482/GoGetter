/* TODO: 
   >> periodically collect formatted data via API requests 
   >> display this data on the console
   >> report results (or errors) to Loggly.
*/

///////////////////////////////[LET'S GET THIS BREAD]/////////////////////////////////

package main

import "net/http"
import "fmt"

func main(){

  resp, err := http.Get("https://api.darksky.net/forecast/1f960db8bf1129b90c3ee6e265c92924/47.8267,-122.4233")

  if err != nil {

    fmt.Println("%s", resp)
    fmt.Println("hi")

  } else {
    
    fmt.Println("bye")

  }

//TODO: get the JSON object and put it in a struct

}
