# MyTarget API SDK GO

## Description
MyTarget API Client for Go

## Service documentation
https://target.my.com/help/partners/reporting_api_statistics/ru

## Usage
```go
package main

import (
    "fmt"
	"net/http"
    mytarget_sdk "github.com/Kachit/mytarget-sdk-go"
)

func yourFuncName(){ 
    cfg := mytarget_sdk.NewConfig()

    client := mytarget_sdk.NewClient(cfg, http.Client{})

    fmt.Print(response)
    fmt.Print(err)
}

```
