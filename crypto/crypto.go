package crypto

import "encoding/hex"
import "encoding/base64"
import "log"

func HexToBase64(hexStringToConvert string) string{
  decoded, err := hex.DecodeString(hexStringToConvert)
  if err != nil {
    log.Fatal(err)
  }
  encodedBase64String := base64.StdEncoding.EncodeToString(decoded)
  return encodedBase64String
}
