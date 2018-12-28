package main

import (
	"sync"
)
/*
type fcmStruct struct{
	lastLiveTime 			int //time when last get this token
	lastNotifyAttack		int
	lastNotifyAV			int
}

type syncFcmStruct struct{
	Name 		string
	ApiKey 		string
}*/




type fcmStruct struct{
	idApiKey				int
	tokens 					[]string	//[token]lastLiveTime
	lastNotifyAttack		int
	lastNotifyAV			int
}

type syncFcmStruct struct{
	Name 		string
	ApiKey 		string
}

type cacheStorage struct {
	mx sync.RWMutex
	data map[string]fcmStruct //[apikey]fcmStruct
}

func (c *cacheStorage) GetID(apiKey string)(int){
	c.mx.RLock()
	defer c.mx.RUnlock()
	val:=c.data[apiKey].idApiKey
	return val
}


func (c *cacheStorage) Get(apiKey string)(fcmStruct, bool){
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok:=c.data[apiKey]
	return val, ok
}

func (c *cacheStorage) GetAll()(map[string]fcmStruct){
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.data
}

func (c *cacheStorage) Set(data map[string]fcmStruct){
	c.mx.Lock()
	defer c.mx.Unlock()
	c.data = data
}

func (c *cacheStorage)Refresh(){

	d,err := getAllApiKey(db)
	if err != nil {
		saveLog(err)
		return
	}
	cache.Set(d)
}

/*

func (c *cacheStorage) UpdateFcmStruct(apiKey string, token string,  lastLiveTime int, lastNotifyAttack int, lastNotifyAV int){
	c.mx.Lock()
	defer c.mx.Unlock()
	st:=c.data[apiKey][token]
	st.lastLiveTime = lastLiveTime
	st.lastNotifyAttack = lastNotifyAttack
	st.lastNotifyAV = lastNotifyAV
	c.data[apiKey][token] = st
}
*/
