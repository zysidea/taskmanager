package controllers

import (
	"gopkg.in/mgo.v2"
	"taskmanager/common"
)

type MongoContext struct {
	MongoSession *mgo.Session
	User string
}

//关闭mongodb连接
func (c *MongoContext) Close()  {
	c.MongoSession.Close()
}
//获取集合
func (c *MongoContext) GetCollection(tableName string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.DataBase).C(tableName)
}
//获取一个新的Context
func NewMongoContext() *MongoContext {
	session:=common.GetSession().Copy()
	return &MongoContext{
		session,
		"",
	}
}
