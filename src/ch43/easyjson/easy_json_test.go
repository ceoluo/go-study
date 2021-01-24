package easyjson

import (
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills": ["java", "Go", "C"]
	}
}`

func TestEmbeddedJson(t *testing.T)  {
	//e := new(Employee)
	e := Employee{}
	//err := json.Unmarshal([]byte(jsonStr), &e)
	err := e.UnmarshalJSON([]byte(jsonStr))
	if err != nil{
		t.Error(err)
	}
	fmt.Println(e)
	fmt.Println(e.BasicInfo.Age)
	fmt.Println(e.JobInfo.Skills)
	if v,err := e.MarshalJSON(); err == nil{
		//fmt.Println(v)
		fmt.Println(string(v))
	}else {
		t.Error(err)
	}
}
