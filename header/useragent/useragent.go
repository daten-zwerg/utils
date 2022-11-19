package useragent

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"google.golang.org/api/idtoken"
)

type UserAgent struct {
	Value    string  `json:"value" bson:"value" `
	Browser  Browser `json:"browser" bson:"browser"`
	OS       string  `json:"os" bson:"os"`
	Hardware string  `json:"hardware" bson:"hardware"`
}

type Browser struct {
	Name    string `json:"name" bson:"name"`
	Version int    `json:"version" bson:"version"`
}

func GetRandomUserAgent(audience string) (UserAgent, error) {
	var ua UserAgent
	ctx := context.Background()
	ts, err := idtoken.NewTokenSource(ctx, audience)

	if err != nil {
		return ua, err
	}

	token, err := ts.Token()
	if err != nil {
		return ua, err
	}

	req, _ := http.NewRequest(http.MethodGet, audience, nil)
	if err != nil {
		return ua, err
	}
	token.SetAuthHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ua, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return ua, err
	}

	err = json.Unmarshal(resBody, &ua)
	if err != nil {
		return ua, err
	}

	return ua, err
}
