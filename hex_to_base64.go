package main

import "os"
import "fmt"
import "encoding/hex"
import "encoding/base64"
import "log"

func main() {
  hexStringToConvert := os.Args[1]

  if len(hexStringToConvert) == 0{
    fmt.Printf("please include a hex string to convert\n")
  }else{
    decoded, err := hex.DecodeString(hexStringToConvert)
  	if err != nil {
  		log.Fatal(err)
  	}
    // fmt.Printf("%v", decoded)
    // fmt.Printf("%s", decoded)
    encodedBase64String := base64.StdEncoding.EncodeToString(decoded)

    fmt.Printf("%v", encodedBase64String)

  }

}
