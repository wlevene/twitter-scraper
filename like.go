package twitterscraper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type (
	LikeBody struct {
		Variables struct {
			TweetId string `json:"tweet_id"`
		} `json:"variables"`

		QueryId string `json:"queryId"`
	}
)

func (s *Scraper) Like(tweet_id string, queryId string) (err error) {

	if tweet_id == "" {
		return
	}
	params := make(map[string]interface{})
	params["queryId"] = queryId
	variables := make(map[string]string)
	variables["tweet_id"] = tweet_id
	params["variables"] = variables

	jsonData, err := json.Marshal(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("jsonData:", string(jsonData))

	req, err := s.newRequest("POST", "https://twitter.com/i/api/graphql/lI07N6Otwv1PhnEgXILM7A/FavoriteTweet")

	req.Body = io.NopCloser(bytes.NewBuffer(jsonData))

	fmt.Println("req:", req)

	if err != nil {
		return err
	}

	curBearerToken := s.bearerToken
	if curBearerToken != bearerToken2 {
		s.setBearerToken(bearerToken2)
	}

	var result map[string]interface{}

	err = s.RequestAPI(req, &result)

	if curBearerToken != bearerToken2 {
		s.setBearerToken(curBearerToken)
	}

	fmt.Println("like result: ", result)

	return
}
