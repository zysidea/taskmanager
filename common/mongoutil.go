package common

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
)

var session *mgo.Session

func GetSession() *mgo.Session  {
	if session==nil {
		var err error
		info:=&mgo.DialInfo{
			Addrs:[]string{AppConfig.MongoDBHost},
			Username:AppConfig.DBUser,
			Password:AppConfig.DBPwd,
			Timeout:60*time.Second,
		}
		session,err=mgo.DialWithInfo(info)
		if err!=nil {
			log.Fatalf("[GetSession]: %s\n",err)
		}
	}
	return session
}
func createDbSession() {
	var err error
	info:=&mgo.DialInfo{
		Addrs:[]string{AppConfig.MongoDBHost},
		Username:AppConfig.DBUser,
		Password:AppConfig.DBPwd,
		Timeout:60*time.Second,
	}
	session,err=mgo.DialWithInfo(info)
	if err!=nil {
		log.Fatalf("[CreateDbSession]: %s\n",err)
	}
}

//添加索引
func addIndexes()  {
	var err error
	userIndex:=mgo.Index{
		Key:[]string{"email"},
		Unique:true,
		Background:true,
		Sparse:true,
	}
	taskIndex:=mgo.Index{
		Key:[]string{"createdby"},
		Unique:false,
		Background:true,
		Sparse:true,
	}
	noteIndex:=mgo.Index{
		Key:[]string{"taskid"},
		Unique:false,
		Background:true,
		Sparse:true,
	}

	session:=GetSession().Copy()
	defer session.Close()
	userCol:=session.DB(AppConfig.DataBase).C("users")
	taskCol:=session.DB(AppConfig.DataBase).C("tasks")
	noteCol:=session.DB(AppConfig.DataBase).C("notes")

	err=userCol.EnsureIndex(userIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n",err)
	}
	err=taskCol.EnsureIndex(taskIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n",err)
	}
	err=noteCol.EnsureIndex(noteIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n",err)
	}
}
