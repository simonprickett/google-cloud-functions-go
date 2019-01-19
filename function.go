package function

import (
	"io/ioutil"
	"net/http"
	"time"
)

func F(w http.ResponseWriter, r *http.Request) {
	randomUserClient := http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequest(http.MethodGet, "https://randomuser.me/api/", nil)
	res, err2 := randomUserClient.Do(req)
	body, err3 := ioutil.ReadAll(res.Body)
}
