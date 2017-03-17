package controllers

import (
	"gopkg.in/mgo.v2"
	"taskmanager/common"
)

type Context struct {
	MongoSession *mgo.Session
}

//关闭mongodb连接
func (c *Context) Close()  {
	c.MongoSession.Close()
}
//获取集合
func (c *Context) GetCollection(tableName string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.DataBase).C(tableName)
}
//获取一个新的Context
func NewContext() *Context  {
	session:=common.GetSession().Copy()
	return &Context{session}
}
