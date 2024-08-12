package wakatimeapi

import (
	"encoding/json"
	"fmt"
	"github.com/AmadoMuerte/WakaTimeModule/internal/models"
	"io"
	"net/http"
)

type Api struct {
	AuthToken string
}

func (a *Api) GetData() models.Response {
	url := "https://api.wakatime.com/api/v1/users/current/status_bar/today"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return models.Response{}
	}

	req.Header.Set("Authorization", "Basic "+a.AuthToken)
	httpResp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return models.Response{}
	}

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		fmt.Println(err)
		return models.Response{}
	}

	var data models.Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return models.Response{}
	}

	return data
}
