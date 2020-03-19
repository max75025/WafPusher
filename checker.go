package main

import (
	"strconv"
	"net/http"
	"io/ioutil"
	"time"
	"fmt"
)



func getJsonAttack(apiKey string, startTime int, endTime int) (string, error) {
	if startTime<0 || endTime<0{
		return "", errTimeLessZero
	}
	url := "http://2waf.com/eventclient/" + apiKey + "/" + strconv.Itoa(startTime) + "/" + strconv.Itoa(endTime)
	resp, err := http.Get(url)
	if err != nil {
		return "null", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "null", err
	}

	return string(content), nil
}

func getJsonAV(apiKey string, startTime int, endTime int) (string, error) {
	if startTime<0 || endTime<0{
		return "", errTimeLessZero
	}
	url := "http://2waf.com/eventav/" + apiKey + "/" + strconv.Itoa(startTime) + "/" + strconv.Itoa(endTime)
	resp, err := http.Get(url)
	if err != nil {
		return "null", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "null", err
	}

	return string(content), nil
}

func checker() {

	for {
		fmt.Println("check...")
		timeNow := int(time.Now().Unix())
		accounts := cache.GetAll()
		for i, d := range accounts {
			attack, err := getJsonAttack(i, d.lastNotifyAttack+1, timeNow)
			if err != nil {
				saveLog(err)
				continue
			}
			if attack!="null"{
				SendNotification(i,"атака","на вас была совершена атака")
				updateAttackTime(db,i, timeNow)

			}

			av, err := getJsonAV(i, d.lastNotifyAV+1, timeNow)
			if err != nil {
				saveLog(err)
				continue
			}
			if av!="null"{
				SendNotification(i,"антивирус","событие файловой системы")
				updateAVTime(db,i,timeNow)
			}

		}
		cache.Refresh()
		time.Sleep(10 * time.Second)
	}
}
