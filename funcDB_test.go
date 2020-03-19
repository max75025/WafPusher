package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"fmt"
	"reflect"
)





func TestAddApiKey(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	//mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO apiKey")
	mock.ExpectExec("INSERT INTO apiKey").WithArgs("test", 13 , 13).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectPrepare("INSERT INTO apiKey")
	mock.ExpectExec("INSERT INTO apiKey").WithArgs("test2", 13 , 13).WillReturnError(fmt.Errorf("some error"))



	if _,err = addApiKey(db, "test",13,13); err != nil {
		t.Error("error:", err)
	}


	if _,err = addApiKey(db, "test2",13,13); err == nil {
		t.Error("wait a error, but err = nil")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAddToken(t *testing.T){
	db, mock, err:= sqlmock.New()
	if err!=nil{t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO token")
	mock.ExpectExec("INSERT INTO token").WithArgs( "test", 13).WillReturnResult(sqlmock.NewResult(13,13))

	mock.ExpectPrepare("INSERT INTO token")
	mock.ExpectExec("INSERT INTO token").WithArgs( "test2", 13).WillReturnError(fmt.Errorf("some error"))

	if _,err:=addToken(db, 13, "test");err!=nil{
		t.Error("err:", err)
	}

	if _,err:=addToken(db, 13, "token2");err==nil{
		t.Error("wait a error, but err=nil")

	}
}

func TestDeleteApiKey(t *testing.T){
	db, mock, err:= sqlmock.New()
	if err!=nil{t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)}
	defer db.Close()

	mock.ExpectPrepare("DELETE FROM apiKey")
	mock.ExpectExec("DELETE FROM apiKey").WithArgs("test").WillReturnResult(sqlmock.NewResult(1,1))

	mock.ExpectPrepare("DELETE FROM apiKey")
	mock.ExpectExec("DELETE FROM apiKey").WithArgs("test").WillReturnError(fmt.Errorf("some error"))

	if err:=deleteApiKey(db, "test");err!=nil{
		t.Error("err:", err)
	}

	if err:=deleteApiKey(db, "test");err==nil{
		t.Error("wait a error, but err=nil")

	}
}


func TestDeleteToken(t *testing.T){
	db, mock, err:= sqlmock.New()
	if err!=nil{t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)}
	defer db.Close()

	mock.ExpectPrepare("DELETE FROM token")
	mock.ExpectExec("DELETE FROM token").WithArgs("test",13).WillReturnResult(sqlmock.NewResult(1,1))

	mock.ExpectPrepare("DELETE FROM apiKey")
	mock.ExpectExec("DELETE FROM apiKey").WithArgs("test",13).WillReturnError(fmt.Errorf("some error"))

	if err:=deleteToken(db, "test",13);err!=nil{
		t.Error("err:", err)
	}

	if err:=deleteToken(db, "test",13);err==nil{
		t.Error("wait a error, but err=nil")

	}
}

func TestUpdateAttackTime(t *testing.T){
	db, mock, err:= sqlmock.New()
	if err!=nil{t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)}
	defer db.Close()

	mock.ExpectPrepare("UPDATE apiKey")
	mock.ExpectExec("UPDATE apiKey").WithArgs(13, "test").WillReturnResult(sqlmock.NewResult(1,1))

	mock.ExpectPrepare("UPDATE apiKey")
	mock.ExpectExec("UPDATE apiKey").WithArgs(13, "test").WillReturnError(fmt.Errorf("some error"))

	if err:=updateAttackTime(db, "test",13);err!=nil{
		t.Error("err:", err)
	}

	if err:=updateAttackTime(db, "test",13);err==nil{
		t.Error("wait a error, but err=nil")

	}
}

func TestUpdateAV(t *testing.T){
	db, mock, err:= sqlmock.New()
	if err!=nil{t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)}
	defer db.Close()

	mock.ExpectPrepare("UPDATE apiKey")
	mock.ExpectExec("UPDATE apiKey").WithArgs(13, "test").WillReturnResult(sqlmock.NewResult(1,1))

	mock.ExpectPrepare("UPDATE apiKey")
	mock.ExpectExec("UPDATE apiKey").WithArgs(13, "test").WillReturnError(fmt.Errorf("some error"))

	if err:=updateAVTime(db, "test",13);err!=nil{
		t.Error("err:", err)
	}

	if err:=updateAVTime(db, "test",13);err==nil{
		t.Error("wait a error, but err=nil")

	}
}

func TestGetAllTokens(t *testing.T){
	testResult:=make(map[int][]string)
	testResult[1] = []string{"test1", "test2"}
	testResult[2] = []string{"test3", "test4"}

	db, mock, err:= sqlmock.New()
	if err!=nil{t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)}
	defer db.Close()

	testRows:= sqlmock.NewRows([]string{"idApiKey", "token"}).
		AddRow(1,"test1").
		AddRow(1,"test2").
		AddRow(2, "test3").
		AddRow(2, "test4")

	mock.ExpectQuery("SELECT idApiKey, token FROM token").WillReturnRows(testRows)

	mock.ExpectQuery("SELECT idApiKey, token FROM token ").WillReturnError(fmt.Errorf("some error"))

	if result,err:=getAllTokens(db);err!=nil || !reflect.DeepEqual(result, testResult){
		t.Error("result:",result,"err:", err)
	}

	if _,err:=getAllTokens(db);err==nil {
		t.Error("err:", err)
	}
}

func TestGetAllApiKey(t *testing.T){

	//fmt.Println([]byte("SELECT * FROM apiKey;"))
	//fmt.Println([]byte())
	apiKeys:= make(map[string]fcmStruct)
	apiKeys["ApiKey1"] = fcmStruct{
		idApiKey:         1,
		tokens:          []string{"test1", "test2"},
		lastNotifyAttack: 13,
		lastNotifyAV:     13,
	}
	apiKeys["ApiKey2"] = fcmStruct{
		idApiKey:         2,
		tokens:          []string{"test3", "test4"},
		lastNotifyAttack: 13,
		lastNotifyAV:     13,
	}



	db, mock, err:= sqlmock.New()
	if err!=nil{t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)}
	defer db.Close()

	apiKeyRows:= sqlmock.NewRows([]string{"ID", "apiKey", "lastNotifyAttack", "lastNotifyAV"}).
		AddRow(1,"ApiKey1",13,13).
		AddRow(2,"ApiKey2",13,13)

	tokensRows:= sqlmock.NewRows([]string{"idApiKey", "token"}).
		AddRow(1,"test1").
		AddRow(1,"test2").
		AddRow(2, "test3").
		AddRow(2, "test4")

	mock.ExpectQuery("SELECT \\* FROM apiKey;").WillReturnRows(apiKeyRows)
	mock.ExpectQuery("SELECT idApiKey, token FROM token").WillReturnRows(tokensRows)

	//mock.ExpectQuery("SELECT idApiKey, token FROM token ").WillReturnError(fmt.Errorf("some error"))

	if result,err:=getAllApiKey(db); err!=nil || !reflect.DeepEqual(result, apiKeys){
		t.Error("result:",result,"err:", err)
	}
/*
	if _,err:=getAllTokens(db);err==nil {
		t.Error("err:", err)
	}*/


}






