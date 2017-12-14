package main

import "os"
import "fmt"
import "crypto_pals/crypto"

func main() {
  hexStringToConvert := os.Args[1]

  if len(hexStringToConvert) == 0{
    fmt.Printf("please include a hex string to convert\n")
  }else{
    fmt.Printf("%v", crypto.HexToBase64(hexStringToConvert))
  }

}
