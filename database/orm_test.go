package database

import (
	"auth-demo/models"
	"github.com/xormplus/xorm"
	"strconv"
	"testing"
)

var E *xorm.Engine

func init() {
	var e error
	E, e = conn()
	if e != nil {
		panic(e)
	}
}
func TestConn(t *testing.T) {
	e := E.Ping()
	t.Log(e)
}
// 当调用QueryBytes时，第一个返回值results为[]map[string][]byte的形式。
func TestQuery1(t *testing.T) {
	sql := "select comment from comment"
	results, err := E.QueryBytes(sql)
	if err != nil {
		t.Log(err)
	}
	for _,v := range results {
		t.Log(string(v["comment"]))
	}
	sql = "select * from comment where id =?"
	results, err = E.SQL(sql, "2").QueryBytes()
	if err != nil {
		t.Log(err)
	}
	for _,v := range results {
		t.Log(string(v["comment"]))
	}
	sql = "select * from comment where id =?"
	results, err = E.SQL(sql, "2").QueryBytes()
	if err != nil {
		t.Log(err)
	}
	for _,v := range results {
		t.Log(string(v["comment"]))
	}
	sql = "select id,comment from comment where id =?id"
	params := map[string]interface{}{"id": "2"}
	results, err = E.SQL(sql, &params).QueryBytes()
	if err != nil {
		t.Log(err)
	}
	for _,v := range results {
		t.Log(string(v["comment"]))
	}
}
// 当调用QueryString时，第一个返回值results为[]map[string]string的形式。
func TestQuery2(t *testing.T) {
	sql := "select comment from comment"
	results, e := E.QueryString(sql)
	t.Log(e)
	t.Log(string(results[0]["comment"]))
	sql = "select comment from comment where id=?"
	results, e = E.SQL(sql,"2").QueryString()
	t.Log(e)
	t.Log(string(results[0]["comment"]))
	sql = "select comment from comment where id=?id"
	params := map[string]interface{}{"id": "2"}
	results, e = E.SQL(sql,&params).QueryString()
	t.Log(e)
	t.Log(string(results[0]["comment"]))
}
// 当调用QueryValue时，第一个返回值results为[]map[string]xorm.Value的形式。xorm.Value类型本质是[]byte，具有一系列类型转换函数。
func TestQuery3(t *testing.T) {
	sql := "select comment from comment"
	results, e := E.QueryValue(sql)
	t.Log(e)
	t.Log(results[0]["comment"].String())
	sql = "select comment from comment where id=?"
	results, e = E.SQL(sql,"2").QueryValue()
	t.Log(e)
	t.Log(results[0]["comment"].String())
	sql = "select comment,create_time from comment where id=?id"
	params := map[string]interface{}{"id": "2"}
	results, e = E.SQL(sql,&params).QueryValue()
	t.Log(e)
	t.Log(results[0]["comment"].String())
	t.Log(results[0]["create_time"].Time("2006-01-02 15:04:05"))
}
// 当调用QueryResult时，第一个返回值results为xorm.Result的形式。
func TestQuery4(t *testing.T) {
	sql := "select comment from comment"
	results, e := E.QueryResult(sql).List()
	t.Log(e)
	t.Log(results[0]["comment"].String())
	sql = "select comment from comment where id=?"
	results, e = E.SQL(sql,"2").QueryResult().List()
	t.Log(e)
	t.Log(results[0]["comment"].String())
	sql = "select comment,create_time from comment where id=?id"
	params := map[string]interface{}{"id": "2"}
	results, e = E.SQL(sql,&params).QueryResult().List()
	t.Log(e)
	t.Log(results[0]["comment"].String())
	t.Log(results[0]["create_time"].Time("2006-01-02 15:04:05"))
}
// 当调用QueryInterface，List或ListPage时，第一个返回值results为[]map[string]interface{}的形式。
func TestQuery5(t *testing.T) {
	sql := "select comment from comment"
	results, e := E.QueryInterface(sql)
	t.Log(e)
	t.Log(string(results[0]["comment"].([]byte)))
	sql = "select comment from comment where id=?"
	results, e = E.SQL(sql,"2").QueryInterface()
	t.Log(e)
	t.Log(results[0]["comment"])
	sql = "select comment,create_time from comment where id=?id"
	params := map[string]interface{}{"id": "2"}
	results, e = E.SQL(sql,&params).QueryInterface()
	t.Log(e)
	t.Log(results[0]["comment"])
	t.Log(results[0]["create_time"])
}
func TestQuery8(t *testing.T) {
	sql := "select * from comment where id =1"
	var comments []models.Comment
	find := E.SQL(sql).Find(&comments)
	t.Log(find)
	t.Log(comments)
	t.Log(comments[0].CreateTime.Format("2006-01-02 15:04:05"))
}
func TestQuery9(t *testing.T) {
	//var comments []models.Comment
	//results, _ := E.Where("id=?", "1").Search(&comments).Xml()
	//t.Log(results)
	//t.Log(comments)
	//comments = []models.Comment{}
	//results, _ = E.Where("id=?", "1").Search(&comments).Json()
	//t.Log(results)
	//t.Log(comments)
	//comments = []models.Comment{}
	//results, _ := E.SQL("select * from comment where id =?", "1").QueryWithDateFormat("20060102 15:04:05").Json()
	//t.Log(results)
	results, _ := E.SQL("select * from comment where id =?", "1").QueryWithDateFormat("20060102 15:04:05").XmlIndent("","  ","comment")
	t.Log(results)
}
func TestSimpleSession(t *testing.T) {
	session := E.NewSession()
	defer session.Close()
	// add Begin() before any action
	err := session.Begin()
	if err!=nil {
		t.Log(err)
		return
	}
	comment1:=models.Comment{
		Id:"20",
		Comment:"comment1",
		UserId:"1",
		PostId:"1",
	}
	i, err := session.Insert(&comment1)
	if err!=nil {
		t.Log(err)
		session.Rollback()
		return
	}
	comment2:=models.Comment{
		Id:"20"+strconv.Itoa(int(i)),
		Comment:"comment1",
		UserId:"1",
		PostId:"1",
	}
	_, err = session.Insert(&comment2)
	if err != nil {
		t.Log(err)
		session.Rollback()
		return
	}
	err = session.Commit()
	if err!=nil {
		t.Log(err)
		return
	}
}
func TestUpdate(t *testing.T)  {
	var comment1 models.Comment
	E.ID("14").Get(&comment1)
	comment1.Comment="efgtgrfertgrewrgre"
	E.ID("14").Update(&comment1)
}
func TestEvent(t *testing.T) {
	comment2:=models.Comment{
		Id:"22",
		Comment:"comment1",
		UserId:"1",
		PostId:"1",
	}
	E.Insert(&comment2)
}