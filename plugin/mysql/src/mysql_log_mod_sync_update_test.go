package src_test

import (
	"github.com/brokercap/Bifrost/sdk/pluginTestData"
	"testing"
)

func TestUpdateSyncAndChekcData(t *testing.T){
	beforeTest()
	initDBTable(true)
	conn := getPluginConn("LogUpdate")
	e := pluginTestData.NewEvent()
	insertdata := e.GetTestInsertData()
	conn.Insert(insertdata)
	conn.Insert(e.GetTestInsertData())
	updateData := e.GetTestUpdateData()
	conn.Update(updateData)
	deleteData := e.GetTestDeleteData()
	conn.Del(deleteData)
	_,err2 := conn.Commit()
	if err2 != nil{
		t.Fatal(err2)
	}

	n , err := getTableCount()
	if err != nil{
		t.Fatal(err)
	}

	if n != 2{
		t.Fatal("append result count != 1")
	}
	t.Log("success")
}

