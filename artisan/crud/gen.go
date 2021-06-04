package crud

import "fmt"

type Command struct {
	Name string
}

func (curd Command) GetDir() string {
	return ""
}

func (curd *Command) Run() error {
	structData := &jsonStruct{}
	if err := jsonData(structData, curd.Name); err != nil {
		fmt.Println("app/http/", curd.Name, "/model.json 不存在，请使用")
		return err
	}
	if err := setService(structData); err != nil {
		return err
	}
	if err := setModel(structData); err != nil {
		return err
	}
	if err := setController(structData); err != nil {
		return err
	}
	return nil
}
