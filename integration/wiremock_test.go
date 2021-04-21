// +build integration

package integration

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestShouldReturnMockedResponseFromWiremock(t *testing.T) {
	url := "http://localhost:" + appPort + "/planets/1"
	fmt.Println("calling application on : ", url)

	var planet Planet

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fail()
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	bytes, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bytes, &planet)
	fmt.Println(planet)
	assert.Nil(t, err)

}

type Planet struct {
	Name            string
	Rotation_period string
	Orbital_period  string
	Diameter        string
	Climate         string
	Gravity         string
	Terrain         string
	Surface_water   string
	Population      string
}
