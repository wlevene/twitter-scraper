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

	likebody := LikeBody{
		Variables: struct {
			TweetId string `json:"tweet_id"`
		}{
			TweetId: tweet_id, // You should have the tweet_id variable defined somewhere
		},
		QueryId: queryId,
	}

	jsonData, err := json.Marshal(likebody)
	if err != nil {
		return
	}

	req, err := s.newRequest("POST", "https://twitter.com/i/api/graphql/lI07N6Otwv1PhnEgXILM7A/FavoriteTweet")

	req.Body = io.NopCloser(bytes.NewBuffer(jsonData))

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
