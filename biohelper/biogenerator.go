package biohelper

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const quotesEndpoint = "https://geek-jokes.sameerkumar.website/api"
const emptyString = ""

//GetNewUserBio Get random bio string
func GetNewUserBio() string {
	response, err := http.Get(quotesEndpoint)

	if err != nil {
		fmt.Println(err.Error())
		return emptyString
	}

	if response.StatusCode == http.StatusOK {
		randomQuote := getStringFromResponseBody(response.Body)
		return randomQuote
	}

	return emptyString
}

func getStringFromResponseBody(body io.ReadCloser) string {
	bodyAsBytes, err := ioutil.ReadAll(body)
	body.Close()

	if err != nil {
		fmt.Println(err.Error())
		return emptyString
	}

	return string(bodyAsBytes)
}
