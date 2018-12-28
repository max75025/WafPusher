package main

import (
	"fmt"
	"log"
	"github.com/max75025/fcm"
	"net/http"
	"errors"
	"os"
	"database/sql"
)

var cache cacheStorage
var db	  *sql.DB



func init() {
	var err error
	db,err = openDB(dbFilePath)
	if err!=nil{panic(err)}
	cache.Refresh()
	//fmt.Println(cache)
	return
}

func saveLog(s error) {
	log.Println(s)
	f, err := os.OpenFile("error.log", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		f, err = os.Create("error.log")
		if err != nil {
			panic(err)
		}
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(s)
}

func SendNotification(apiKey string, title string, body string) error {
	fmt.Println("sendNotification")
	allFcmDevices, ok := cache.Get(apiKey)
	if !ok {
		return errNotFindDataForApiKey
	}

	if len(allFcmDevices.tokens)==0{
		err:=deleteApiKey(db, apiKey)
		if err!=nil{
			return err
		}
		cache.Refresh()
		return nil
	}

	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}
	c := fcm.NewFCM(fcm_server_api_key)


	response, err := c.Send(fcm.Message{
		Data:             data,
		RegistrationIDs:  allFcmDevices.tokens,
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Notification: fcm.Notification{
			Title: title,
			Body:  body,
			Sound: "notification.mp3",
		},
	})
	if err != nil {
		return err
	}

	fmt.Println("Status Code   :", response.StatusCode)
	fmt.Println("Success       :", response.Success)
	fmt.Println("Fail          :", response.Fail)
	fmt.Println("Canonical_ids :", response.CanonicalIDs)
	fmt.Println("Topic MsgId   :", response.MsgID)



	if response.Fail > 0 {
		f, err := os.OpenFile("badNotify.log", os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			f, err = os.Create("badNotify.log")
			if err != nil {
				panic(err)
			}
		}
		defer f.Close()
		log.SetOutput(f)
		log.Println("apiKey		  :", apiKey)
		log.Println("tokens		  :", allFcmDevices.tokens)
		log.Println("Status Code   :", response.StatusCode)
		log.Println("Success       :", response.Success)
		log.Println("Fail          :", response.Fail)
		log.Println("Canonical_ids :", response.CanonicalIDs)
		log.Println("Topic MsgId   :", response.MsgID)
		log.Println("result		  :", response.Results)

		//delete NotRegistered token
		idApiKey:= cache.GetID(apiKey)

		for i,r:= range response.Results{
			//fmt.Println(r.Error, " ", i)
			if r.Error == "NotRegistered"{
				//fmt.Println(allFcmDevices.tokens[i])
				err = deleteToken(db, allFcmDevices.tokens[i], idApiKey)
				if err!=nil{saveLog(err)}
				fmt.Println("delete not registred tokens")
				cache.Refresh()
			}
		}

	}

	return nil

}



func deleteFCM(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		token := r.FormValue("FCMtoken")
		apiKey := r.FormValue("ApiKey")

		idApiKey:= cache.GetID(apiKey)
		err:=deleteToken(db, token, idApiKey)
		if err!=nil{
			saveLog(err)
			return
		}
		fmt.Fprintf(w,"true")
	} else {
		fmt.Fprintf(w, "sorry only POST");
	}
}

func addNewFCM(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		token := r.FormValue("FCMtoken")
		apiKey := r.FormValue("ApiKey")
		//timeUnix := int(time.Now().Unix())
		//var id int

		if d, ok := cache.Get(apiKey); !ok {
			result, err := addApiKey(db, apiKey, 0, 0)
			if err != nil {
				saveLog(err)
				return
			}
			i64, err := result.LastInsertId()
			if err != nil {
				saveLog(err)
				return
			}
			id := int(i64)
			_,err = addToken(db, id, token)
			if err != nil {
				saveLog(err)
				return
			}
		} else {
			contain := isContain(token, d.tokens)
			if !contain{
				_,err := addToken(db , d.idApiKey, token)
				if err != nil {
					saveLog(err)
					return
				}
			}
		}


		cache.Refresh()
		fmt.Fprintf(w, "true")
		fmt.Println("add new fcm")
	} else {
		fmt.Fprintf(w, "sorry only POST");
	}
}

func isContain(ss string, contain []string)bool{
	for _,d:= range contain{
		if d==ss{
			return true
		}
	}
	return false
}


func main() {
	saveLog(errors.New("start server..."))
	go checker()

	http.HandleFunc("/deleteFCM/", deleteFCM)
	http.HandleFunc("/addNewFCM/", addNewFCM)
	fmt.Println("listen and serve...")
	log.Fatal(http.ListenAndServe(":1313", nil))

}
