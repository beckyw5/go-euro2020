package euro2020

import (
	"fmt"
	"net/http"
	"time"
)

type Matches struct {
	Count   int `json:"count"`
	Filters struct {
	} `json:"filters"`
	Competition struct {
		ID   int `json:"id"`
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Name        string    `json:"name"`
		Code        string    `json:"code"`
		Plan        string    `json:"plan"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"competition"`
	Matches []struct {
		ID     int `json:"id"`
		Season struct {
			ID              int    `json:"id"`
			StartDate       string `json:"startDate"`
			EndDate         string `json:"endDate"`
			CurrentMatchday int    `json:"currentMatchday"`
		} `json:"season"`
		UtcDate     time.Time `json:"utcDate"`
		Status      string    `json:"status"`
		Matchday    int       `json:"matchday"`
		Stage       string    `json:"stage"`
		Group       string    `json:"group"`
		LastUpdated time.Time `json:"lastUpdated"`
		Odds        struct {
			Msg string `json:"msg"`
		} `json:"odds"`
		Score struct {
			Winner   string `json:"winner"`
			Duration string `json:"duration"`
			FullTime struct {
				HomeTeam int `json:"homeTeam"`
				AwayTeam int `json:"awayTeam"`
			} `json:"fullTime"`
			HalfTime struct {
				HomeTeam int `json:"homeTeam"`
				AwayTeam int `json:"awayTeam"`
			} `json:"halfTime"`
			ExtraTime struct {
				HomeTeam interface{} `json:"homeTeam"`
				AwayTeam interface{} `json:"awayTeam"`
			} `json:"extraTime"`
			Penalties struct {
				HomeTeam interface{} `json:"homeTeam"`
				AwayTeam interface{} `json:"awayTeam"`
			} `json:"penalties"`
		} `json:"score"`
		HomeTeam struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"homeTeam"`
		AwayTeam struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"awayTeam"`
		Referees []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Role        string `json:"role"`
			Nationality string `json:"nationality"`
		} `json:"referees"`
	} `json:"matches"`
}

type euroTeamsResponse struct {
    Matches []Matches `json:"matches"`
}

func GetAllTeams() ([]Match, error) {
    res, err := http.Get(fmt.Sprintf("%s/id", baseURL))
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    var response euroTeamsResponse
    err = json.NewDecoder(res.Body).Decode(&response)
    if err != nil {
        return nil, err
    }

    return response.Match, nil
}