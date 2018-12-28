package main

import (
	"database/sql"
	"os"
	_"github.com/max75025/go-sqlite3"

)

const dbFilePath  = "./db.db"


func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func openDB(filePath string)(*sql.DB, error){
	ex, err:= exists(filePath)
	if err!=nil{
		return nil,err
	}
	if !ex{
		_, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return nil,err
		}
	}

	db, err:= sql.Open("sqlite3", filePath)
	if err!= nil{
		return nil, err
	}
	if !ex{
		_,err = db.Exec("CREATE TABLE IF NOT EXISTS apiKey (ID INTEGER PRIMARY KEY, apiKey TEXT UNIQUE, lastNotifyAttack INTEGER, lastNotifyAV INTEGER );")
		if err!=nil{return nil, err}
		_,err = db.Exec("CREATE TABLE IF NOT EXISTS token (ID INTEGER PRIMARY KEY, token TEXT , idApiKey INTEGER );")
		if err!=nil{return nil, err}
	}
	return db, nil
}



/*
type DBser interface {
	addApiKey(apiKey string, lastNotifyAttack int, lastNotifyAV int)( sql.Result,error)
	addToken(token string, idApiKey int)(sql.Result, error)

	deleteApiKey(apiKey string)error
	deleteToken(token string, idApiKey int)error

	updateAttackTime(apiKey string,  lastNotifyTime int)error
	updateAVTime(apiKey string, lastNotifyTime int)error

	getAllApiKey()(map[string]fcmStruct, error)
	getAllTokens()(map[int][]string, error)
}

type dbStruct struct {
	db    *sql.DB
}


func (d dbStruct)addApiKey(apiKey string, lastNotifyAttack int, lastNotifyAV int)( sql.Result,error){
	stat,err:=d.db.Prepare("INSERT INTO apiKey( apiKey, lastNotifyAttack, lastNotifyAV) VALUES (?,?,?);")
	if err!=nil{ return nil,err}
	result,err := stat.Exec(  apiKey, lastNotifyAttack, lastNotifyAV)

	return result,err
}
func (d dbStruct)addToken(token string, idApiKey int)(sql.Result, error){
	stat,err:=d.db.Prepare("INSERT INTO token(token, idApiKey) VALUES (?,?)")
	if err!=nil{ return nil,err}
	result,err := stat.Exec(token,  idApiKey)
	return result,err
}

func (d dbStruct) deleteApiKey( apiKey string)error{
	stat,err:=d.db.Prepare("DELETE FROM apiKey WHERE apiKey=?;")
	if err!=nil{ return err}
	_,err = stat.Exec(apiKey)
	return err
}
func (d dbStruct) deleteToken(token string, idApiKey int)error{
	stat,err:=d.db.Prepare("DELETE FROM token WHERE token=? AND idApiKey=?;")
	if err!=nil{ return err}
	_,err = stat.Exec(token, idApiKey)
	return err
}

func (d dbStruct) updateAttackTime(  apiKey string,  lastNotifyTime int) error{
	stat,err:=db.Prepare("UPDATE apiKey SET   lastNotifyAttack=?  WHERE  apiKey=?;")
	if err!=nil{ return err}
	_,err = stat.Exec( lastNotifyAttack, apiKey)
	return err
}
func (d dbStruct) updateAVTime(  apiKey string,  lastNotifyTime int) error{
	stat,err:=d.db.Prepare("UPDATE apiKey SET   lastNotifyAV=?  WHERE  apiKey=?;")
	if err!=nil{ return err}
	_,err = stat.Exec( lastNotifyAV, apiKey)
	return err
}


func (d dbStruct) getAllApiKey()(map[string]fcmStruct, error){

	rows,err:= d.db.Query("SELECT * FROM apiKey;")
	if err!=nil{return nil,err}
	defer rows.Close()
	allTokens,err:= getAllTokens(db)
	if err!=nil{return nil,err}
	//fmt.Println(allTokens)
	var apiKey string
	data:=  make( map[string]fcmStruct)
	for rows.Next(){
		addFcmStruct := fcmStruct{}
		rows.Scan( &addFcmStruct.idApiKey,   &apiKey, &addFcmStruct.lastNotifyAttack, &addFcmStruct.lastNotifyAV)
		addFcmStruct.tokens = allTokens[addFcmStruct.idApiKey]
		data[apiKey]= addFcmStruct

	}

	return data, nil
}
func (d dbStruct) getAllTokens()(map[int][]string, error){
	result := make(map[int][]string)
	inRows,err:= d.db.Query("SELECT idApiKey, token FROM token ;")
	if err!=nil{return nil,err}
	defer inRows.Close()
	var idApi int
	var t	  string
	for inRows.Next(){
		inRows.Scan(&idApi, &t)
		result[idApi] = append(result[idApi], t)
	}
	return result, nil
}*/
func addApiKey(db *sql.DB, apiKey string, lastNotifyAttack int, lastNotifyAV int)( sql.Result,error){
	stat,err:=db.Prepare("INSERT INTO apiKey( apiKey, lastNotifyAttack, lastNotifyAV) VALUES (?,?,?);")
	if err!=nil{ return nil,err}
	result,err := stat.Exec(  apiKey, lastNotifyAttack, lastNotifyAV)

	return result,err
}

func addToken(db *sql.DB, idApiKey int, token string)(sql.Result,error){
	stat,err:=db.Prepare("INSERT INTO token(token, idApiKey) VALUES (?,?)")
	if err!=nil{ return nil,err}
	result,err := stat.Exec(token,  idApiKey)
	return result,err
}

/*func updateToken( token string)error{
	stat,err:=db.Prepare("UPDATE token SET   lastLiveTime=? WHERE  token=?")
	if err!=nil{ return err}
	_,err = stat.Exec(lastLiveTime, token)
	db.Close()
	return err
}*/


func deleteApiKey(db *sql.DB, apiKey string)error{
	stat,err:=db.Prepare("DELETE FROM apiKey WHERE apiKey=?;")
	if err!=nil{ return err}
	_,err = stat.Exec(apiKey)
	return err
}


func deleteToken(db *sql.DB,token string, idApiKey int)error{
	stat,err:=db.Prepare("DELETE FROM token WHERE token=? AND idApiKey=?;")
	if err!=nil{ return err}
	_,err = stat.Exec(token, idApiKey)
	return err
}

func updateApiKey(db *sql.DB,  apiKey string,  lastNotifyAttack int, lastNotifyAV int) error{
	stat,err:=db.Prepare("UPDATE apiKey SET   lastNotifyAttack=?, lastNotifyAV=?  WHERE  apiKey=?;")
	if err!=nil{ return err}
	_,err = stat.Exec( lastNotifyAttack,lastNotifyAV, apiKey)
	return err
}



func updateAttackTime(db *sql.DB,  apiKey string,  lastNotifyAttack int) error{
	stat,err:=db.Prepare("UPDATE apiKey SET   lastNotifyAttack=?  WHERE  apiKey=?;")
	if err!=nil{ return err}
	_,err = stat.Exec( lastNotifyAttack, apiKey)
	return err
}

func updateAVTime(db *sql.DB,  apiKey string,  lastNotifyAV int) error{
	stat,err:=db.Prepare("UPDATE apiKey SET   lastNotifyAV=?  WHERE  apiKey=?;")
	if err!=nil{ return err}
	_,err = stat.Exec( lastNotifyAV, apiKey)
	return err
}


func getAllApiKey(db *sql.DB)(map[string]fcmStruct, error){

	rows,err:= db.Query("SELECT * FROM apiKey;")
	if err!=nil{return nil,err}
	defer rows.Close()
	allTokens,err:= getAllTokens(db)
	if err!=nil{return nil,err}
	//fmt.Println(allTokens)
	var apiKey string
	data:=  make( map[string]fcmStruct)
	for rows.Next(){
		addFcmStruct := fcmStruct{}
		rows.Scan( &addFcmStruct.idApiKey,   &apiKey, &addFcmStruct.lastNotifyAttack, &addFcmStruct.lastNotifyAV)
		addFcmStruct.tokens = allTokens[addFcmStruct.idApiKey]
		data[apiKey]= addFcmStruct

	}

	return data, nil
}


func getAllTokens(db *sql.DB)(map[int][]string, error){
	result := make(map[int][]string)
	inRows,err:= db.Query("SELECT idApiKey, token FROM token ;")
	if err!=nil{return nil,err}
	defer inRows.Close()
	var idApi int
	var t	  string
	for inRows.Next(){
		inRows.Scan(&idApi, &t)
		result[idApi] = append(result[idApi], t)
	}
	return result, nil
}

