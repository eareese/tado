package canvas

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// this package-level variable is not ideal, but will make it easy to use for all our queries
var token string

type coursesResponse []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Init sets the access token found in env variable
func Init() error {
	// expect to find the token in env var TODO_TOKEN
	viper.SetEnvPrefix("todo")
	viper.BindEnv("token")
	token = viper.GetString("token")
	if len(token) <= 0 {
		return errors.New("need an access token")
	}
	return nil
}

func QueryCourses() (coursesResponse, error) {
	var data coursesResponse

	query := fmt.Sprintf("https://temp.acme.instructure.com/api/v1/courses?access_token=%s", token)
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}
