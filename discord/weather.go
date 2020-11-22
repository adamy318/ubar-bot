package discord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type OpenWeatherMap struct {
	Coordinates   Coord     `json:"coord"`
	WeatherInfo   []Weather `json:"weather"`
	Base          string    `json:"base"`
	MainInfo      Forecast  `json:"main"`
	Visibility    int       `json:"visibility"`
	WindInfo      Wind      `json:"wind"`
	CloudInfo     Clouds    `json:"clouds"`
	TimeRequested int       `json:"dt"`
	//RainInfo      Rain     `json:"rain"`
	//SnowInfo      Snow     `json:"snow"`
	SysInfo  Sys    `json:"sys"`
	Timezone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type Coord struct {
	Longitude float32 `json:"lon"`
	Latitude  float32 `json:"lat"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Forecast struct {
	Temp     float32 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempLow  float32 `json:"temp_min"`
	TempHigh float32 `json:"temp_max"`
}

type Wind struct {
	Speed  float32 `json:"speed"`
	Degree int     `json:"deg"`
}

type Sys struct {
	Type        int     `json:"type"`
	Id          int     `json:"id"`
	Message     float32 `json:"message"`
	Country     string  `json:"country"`
	SunriseTime int     `json:"sunrise"`
	SunsetTime  int     `json:"sunset"`
}

type Clouds struct {
	CloudinessPercentage int `json:"all"`
}

func GetWeather(s *discordgo.Session, m *discordgo.MessageCreate) {

	api_key := os.Getenv("OPENWEATHERMAP_API_KEY")

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	msg := strings.Fields(m.Content)
	if msg[0] != "!weather" {
		s.ChannelMessageSend(m.ChannelID, "please include !weather to get weather")
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", msg[1], api_key)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var w = new(OpenWeatherMap)
	err = json.Unmarshal(body, &w)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.ChannelMessageSend(m.ChannelID, w.Name)

}
