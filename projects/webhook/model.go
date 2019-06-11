package main

import (
	"time"
)

type Alert struct {
	Labes		map[string]string	`json:"lables"`
	Annotations	map[string]string	`json:"annotations"`
	StartsAt	time.Time			`json:"startAt"`
	EndsAt		time.Time			`json:"endsAt"`
}

type Notification struct {
	Version				string				`json:"version"`
	Groupkey			string				`json:"grouKey"`
	Status				string            	`json:"status"`
	Receiver			string            	`json:"receiver"`
	GroupLabels     	map[string]string 	`json:"groupLabels"`
	CommonLabels		map[string]string 	`json:"commonLabels"`
	CommonAnnotations 	map[string]string 	`json:"commonAnnotations"`
	ExternalURL       	string            	`json:"externalURL"`
	Alerts            	[]Alert           	`json:"alerts"`
}

type At struct {
    AtMobiles []string `json:"atMobiles"`
    IsAtAll   bool     `json:"isAtAll"`
}

type DingTalkMarkdown struct {
    MsgType  string    `json:"msgtype"`
    At       *At       `json:at`
    Markdown *Markdown `json:"markdown"`
}

type Markdown struct {
    Title string `json:"title"`
    Text  string `json:"text"`
}