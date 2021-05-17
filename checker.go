package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
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
		timeNow := int(time.Now().Unix())
		accounts := cache.GetAll()
		for i, d := range accounts {
			attack, err := getJsonAttack(i, d.lastNotifyAttack+1, timeNow)
			if err != nil {
				log.Printf("%+v\n",err)
				log.Println(errors.New("response:" + attack))
				continue
			}

			var at []Attack

			err = json.Unmarshal([]byte(attack), &at)
			if err!=nil{
				log.Printf("%+v\n",err)
				log.Println(errors.New("response:" + attack))
				continue
			}

			if len(at)>0{
				err  = updateAttackTime(db,i, timeNow)
				if err!=nil{
					log.Printf("%+v\n",errors.Wrap(err, "update attack time in database error"))
				}else{
					err = SendNotification(i,"атака","на вас была совершена атака")
					if err!=nil{log.Printf("%+v\n",errors.Wrap(err, "send attack push error"))}
				}
			}

			av, err := getJsonAV(i, d.lastNotifyAV+1, timeNow)
			if err != nil {
				log.Printf("%+v\n",err)
				log.Println(errors.New("response:" + av))
				continue
			}

			var antiv []AV

			err = json.Unmarshal([]byte(av), &antiv)
			if err!=nil{
				log.Printf("%+v\n",err)
				log.Println(errors.New("response:" + av))
				continue
			}

			if len(antiv)>0 {
				err = updateAVTime(db,i,timeNow)
				if err!=nil{
					log.Printf("%+v\n",errors.Wrap(err, "update av time in database error"))
				}else {
					err = SendNotification(i, "мониторинг файлов", "событие файловой системы")
					if err!=nil{log.Printf("%+v\n", errors.Wrap(err, "send antivirus push error"))}
				}

			}

		}
		cache.Refresh()
		time.Sleep(10 * time.Second)
	}
}
