package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"log"
	"os"
)

func main() {
	//var definition
	sub := "root"
	obj := "/articles/123"
	act := "GET"

	//load conf
	//e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")

	//generateModelForCode
	m := generateModelForCode()
	a := fileadapter.NewAdapter("config/policy.csv")
	e, err := casbin.NewEnforcer(m, a)
	//testData := e.GetAllSubjects()
	//log.Printf("%+v", testData)

	//fmt.Println(e.GetAllNamedSubjects("p"))
	//fmt.Println(e.GetPolicy())
	//fmt.Println(e.GetFilteredPolicy(0, "role1"))
	fmt.Println(e.HasPolicy("role2", "/users/:id", "GET"))
	os.Exit(0)

	if err != nil {
		log.Println(err)
	}

	if ok, err := e.Enforce(sub, obj, act); ok {
		fmt.Println("allow!!")
	} else {
		fmt.Println("deny!!")
		fmt.Println(err)
	}

}

func generateModelForCode() model.Model {
	//Initialize the model from Co code
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "r.sub == p.sub && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act) || r.sub == 'root'")

	return m
}
