package main

import (
	"fmt"
	
)

func main() {
   a := map[string]string{
      "asfaf": "asds",
      "tams": "azam",
      "misha": "cyka",
      "Linda": "lin",
   }

   for name, val := range a{
      fmt.Println(name, val)
   }
}
