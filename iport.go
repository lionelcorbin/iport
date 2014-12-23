package iport

import (
        "os"
        "fmt"
)

func Ipetport() string {
  s := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
  return s
}



