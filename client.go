package fitbit

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/zrm625/fitbit/model"

	"golang.org/x/oauth2"
)

const (
	authURI      = "https://www.fitbit.com/oauth2/authorize"
	refreshToken = "https://api.fitbit.com/oauth2/token"

	baseURI = "https://api.fitbit.com"

	// start-time	The start of the period, in the format HH:mm. Optional.
	// end-time	The end of the period, in the format HH:mm. Optional.
	// detail-level	Number of data points to include. Either 1sec or 1min. Optional.
	heartRate = "/1/user/-/activities/heart/date/%s/%s/%s.json"

	// date	The date in the format yyyy-MM-dd.
	// base-date	The end date when period is provided, in the format yyyy-MM-dd; range start date when a date range is provided.
	// end-date	Range end date when date range is provided. Note: The period must not be longer than 31 days.
	weight = "/1/user/-/body/log/weight/date/%s/%s.json"
)

type startEnd struct {
	start time.Time
	end   time.Time
}

func divideTimes(start, end time.Time, durInDays int) []startEnd {
	var times []startEnd
	s := start
	e := end
	done := false
	for {
		e = s.Add(time.Hour * 24 * time.Duration(durInDays))
		if e.After(end) {
			e = end
			done = true
		}

		times = append(times, startEnd{
			start: s,
			end:   e,
		})
		if done {
			break
		} else {
			s = e.Add(time.Hour * 24)
			e = e.Add(time.Hour * 24 * time.Duration(durInDays))
		}
	}
	return times
}

func (f *Fitbit) GetWeights(start, end time.Time) (weights []model.Weight, err error) {
	// Need to get in multiple calls if the period is greater than 31 days
	for _, times := range divideTimes(start, end, 31) {
		url := fmt.Sprintf(baseURI+weight, times.start.Format("2006-01-02"), times.end.Format("2006-01-02"))
		resp, err := f.client.Get(url)
		if err != nil {
			return nil, fmt.Errorf("getting wieghts %w", err)
		}
		defer resp.Body.Close()
		tmp := struct {
			Weight []model.Weight
		}{}

		if err := json.NewDecoder(resp.Body).Decode(&tmp); err != nil {
			return nil, fmt.Errorf("decoding weight response: %w", err)
		}
		weights = append(weights, tmp.Weight...)
	}
	return weights, nil
}

type Fitbit struct {
	client *http.Client
}

func New(code, clientID, clientSecret string) (*Fitbit, error) {
	conf := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"activity", "heartrate", "location", "nutrition", "profile", "settings", "sleep", "social", "weight"},
		RedirectURL:  "https://localhost",
		Endpoint: oauth2.Endpoint{
			AuthURL:   authURI,
			TokenURL:  refreshToken,
			AuthStyle: oauth2.AuthStyleInHeader,
		},
	}

	var token *oauth2.Token
	var err error
	f, err := os.Open("token")
	if err == nil {
		err := gob.NewDecoder(f).Decode(&token)
		if err != nil {
			return nil, err
		}
	} else {
		token, err = conf.Exchange(context.Background(), code)
		if err != nil {
			return nil, fmt.Errorf("could not get token: %w", err)
		}
		f, err = os.Create("token")
		if err != nil {
			return nil, err
		}

		err = gob.NewEncoder(f).Encode(token)
		if err != nil {
			return nil, err
		}
	}

	return &Fitbit{client: conf.Client(context.Background(), token)}, nil
}

func (f *Fitbit) Get() (model.User, error) {
	resp, err := f.client.Get("https://api.fitbit.com/1/user/-/profile.json")
	if err != nil {
		return model.User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.User{}, errors.New("non-200")
	}
	var u model.User
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return model.User{}, fmt.Errorf("decoding user: %w", err)
	}
	return u, nil
}
