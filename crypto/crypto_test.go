package crypto

import "testing"

func TestHexToBase64(t *testing.T){

  str := HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

  if str == "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
    //success
  }else{
    t.Log("strings do not match sample")
    t.Fail()
  }

}
