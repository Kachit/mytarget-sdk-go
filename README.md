# MyTarget API SDK GO (Unofficial)
[![Build Status](https://travis-ci.org/Kachit/mytarget-sdk-go.svg?branch=master)](https://travis-ci.org/Kachit/mytarget-sdk-go)
[![codecov](https://codecov.io/gh/Kachit/mytarget-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Kachit/mytarget-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/kachit/mytarget-sdk-go)](https://goreportcard.com/report/github.com/kachit/mytarget-sdk-go)
[![Release](https://img.shields.io/github/v/release/Kachit/mytarget-sdk-go.svg)](https://github.com/Kachit/mytarget-sdk-go/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/mytarget-sdk-go/blob/master/LICENSE)

## Description
Unofficial MyTarget API Client for Golang

## API documentation
https://target.my.com/help/partners/management_api/ru

https://target.my.com/adv/api-marketing/doc/stat-v2

https://target.my.com/help/partners/reporting_api_statistics/ru

## Download
```shell
go get -u github.com/kachit/mytarget-sdk-go
```

## Usage
```go
package main

import (
    "fmt"
    "net/http"
    mytarget_sdk "github.com/kachit/mytarget-sdk-go"
    mytarget_sdk_marketing "github.com/kachit/mytarget-sdk-go/marketing"
)

func yourFuncName(){ 
    accessToken := "Access token"
    client := mytarget_sdk.NewClientFromCredentials(accessToken, &http.Client{})

    filter := &mytarget_sdk_marketing.StatsFilter{}
    filter.DateFrom = time.Date(2020, time.Month(9), 1, 0, 0, 0, 0, time.UTC)
    filter.DateTo = time.Date(2020, time.Month(9), 2, 0, 0, 0, 0, time.UTC)
    response, err := client.Marketing().Statistics().GetCampaignStatsDaily(filter)

    if response.IsSuccess() {
        var stats marketing.StatsCollection
        response.Unmarshal(&stats)
        fmt.Println(stats.Items[0].Id)
        fmt.Println(stats.Total.Base.Clicks)
    } else {
        errorRes, _ := response.GetError()
        fmt.Println(errorRes.Error.Code)
        fmt.Println(errorRes.Error.Message)
    }

    fmt.Println(err)
}

```
