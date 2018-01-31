package canvas

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

// this package-level variable is not ideal, but will make it easy to use for all our queries
var token string

type CoursesResponse []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AssignmentsResponse []struct {
	ID                      int       `json:"id"`
	Description             string    `json:"description"`
	DueAt                   time.Time `json:"due_at"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	CourseID                int       `json:"course_id"`
	Name                    string    `json:"name"`
	HasSubmittedSubmissions bool      `json:"has_submitted_submissions"`
	NeedsGradingCount       int       `json:"needs_grading_count"`
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

func QueryCourses() (data CoursesResponse, err error) {
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

// QueryAssignmentsUngraded retrieves the list of ungraded assignments for the given course ID
func QueryAssignmentsUngraded(courseID int) (data AssignmentsResponse, err error) {
	query := fmt.Sprintf("https://temp.acme.instructure.com/api/v1/courses/%d/assignments?access_token=%s", courseID, token)
	query = fmt.Sprintf("%s&bucket=ungraded", query)

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
