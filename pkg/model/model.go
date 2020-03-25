package model

import (
	"github.com/jinzhu/gorm"
	"strings"
)

const ProblemBaseDir = "https://atcoder.jp/contests/"

type Problem struct {
	Id        string `json:"id"`
	ContestId string `json:"contest_id"`
	Title     string `json:"title"`
	Url       string
	Used      bool
}

func AddProblemList(ps []Problem) error {
	for _, p := range ps {
		if strings.Index(p.ContestId, "abc") != -1 && strings.Index(p.Id, "_a") == -1 {
			p.Url = ProblemBaseDir + p.ContestId + "/tasks/" + p.Id
			p.Used = false
			db.Create(p)
		}
	}
	return nil
}

func GetRandomProblemData() *Problem {
	var p Problem
	db.Where("used = ?", "0").Order(gorm.Expr("random()")).First(&p)
	p.Used = true
	db.Save(p)
	return &p
}
