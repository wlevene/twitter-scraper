package twitterscraper

import (
	"encoding/json"
	"fmt"
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

	requestBody := struct {
		Variables struct {
			TweetID string `json:"tweet_id"`
		} `json:"variables"`
		QueryID string `json:"queryId"`
	}{
		Variables: struct {
			TweetID string `json:"tweet_id"`
		}{
			TweetID: tweet_id,
		},
		QueryID: queryId,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("jsonData:", string(jsonData))

	req, err := s.newRequest("POST",
		"https://twitter.com/i/api/graphql/lI07N6Otwv1PhnEgXILM7A/FavoriteTweet",
		jsonData)

	fmt.Println("req:", req)

	fmt.Println("req.body:", req.Body)

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
