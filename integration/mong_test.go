// +build integration

package integration

import (
	"encoding/json"
	"fmt"
	"github.com/mahjadan/go-integration-test/domain"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestShouldReturnMockedResponseFromMongoDB(t *testing.T) {
	url := "http://localhost:" + appPort + "/people/1"
	fmt.Println("calling application on : ", url)

	var people domain.People

	request,err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fail()
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	bytes,_ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bytes, &people)
	assert.NotEmpty(t, people)
	assert.Nil(t, err)

}
