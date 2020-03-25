package collector

import (
	"contest-daily-bot/pkg/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	ProblemApiPrefix = "https://kenkoooo.com/atcoder/resources/problems.json"
)

func CollectProblems() []model.Problem {
	resp, _ := http.Get(ProblemApiPrefix)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	var data []model.Problem
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil
	}
	return data
}

