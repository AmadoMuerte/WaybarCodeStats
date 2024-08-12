package main

import (
	"fmt"
	"github.com/AmadoMuerte/WakaTimeModule/internal/timeconv"
	"github.com/AmadoMuerte/WakaTimeModule/internal/wakatimeapi"
)

func main() {

	api := wakatimeapi.Api{
		AuthToken: "TOKEN",
	}

	res := api.GetData()
	dailySeconds := res.Data.DailyAverage
	time := timeconv.SecondToTime(dailySeconds)
	fmt.Println(time)
}
