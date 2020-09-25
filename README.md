# MyTarget API SDK GO
[![Build Status](https://travis-ci.org/Kachit/mytarget-sdk-go.svg?branch=master)](https://travis-ci.org/Kachit/mytarget-sdk-go)
[![codecov](https://codecov.io/gh/Kachit/mytarget-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Kachit/mytarget-sdk-go)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/mytarget-sdk-go/blob/master/LICENSE)

## Description
MyTarget API Client for Go

## API documentation
https://target.my.com/help/partners/management_api/ru

https://target.my.com/help/partners/reporting_api_statistics/ru

## Download
```shell
go get github.com/kachit/mytarget-sdk-go
```

## Usage
```go
package main

import (
    "fmt"
    "net/http"
    mytarget_sdk "github.com/kachit/mytarget-sdk-go"
)

func yourFuncName(){ 
    cfg := mytarget_sdk.NewConfig()

    client := mytarget_sdk.NewClient(cfg, &http.Client{})

    fmt.Print(client)
}

```
