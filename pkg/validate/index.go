package validate

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"reflect"
	"strings"
)

func MapValidate(r *http.Request, data map[string][]string, validateMsg ...map[string]string) (bool, string) {
	var messages map[string][]string
	if len(validateMsg) > 0 {
		messages = getChMessage(data, validateMsg[0])
	} else {
		messages = getChMessage(data, nil)
	}
	return validate(r, data, messages)
}

func StructValidate(r *http.Request, data interface{}, language ...string) (bool, string) {
	validateMap, validateNameMap := getValidateMap(data)
	var messages map[string][]string
	caseLanguage := "zh"
	if len(language) > 0 {
		caseLanguage = language[0]
	}
	switch caseLanguage {
	case "zh":
		messages = getChMessage(validateMap, validateNameMap)
	}
	return validate(r, validateMap, messages)
}

func MapDataForStruct(data interface{}) (map[string][]string, map[string]string) {
	return getValidateMap(data)
}

func validate(r *http.Request, validateMap map[string][]string, messages map[string][]string) (bool, string) {
	rules := govalidator.Options{
		Request:         r,
		Rules:           validateMap,
		Messages:        messages,
		RequiredDefault: false,
	}
	v := govalidator.New(rules)
	e := v.Validate()
	if len(e) == 0 {
		return true, ""
	}
	for _, v := range e {
		return false, v[0]
	}
	return true, ""
}

func getValidateMap(data interface{}) (map[string][]string, map[string]string) {
	typeOf := reflect.TypeOf(data)
	validateMap := make(map[string][]string)
	validateNameMap := make(map[string]string)
	getMap(validateMap, validateNameMap, typeOf)
	return validateMap, validateNameMap
}

func getMap(validateMap map[string][]string, validateNameMap map[string]string, typeOf reflect.Type) {
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		if field.Type.Kind().String() == "struct" {
			getMap(validateMap, validateNameMap, field.Type)
		}
		key := field.Name
		if field.Tag.Get("json") != "" {
			key = field.Tag.Get("json")
		}
		if field.Tag.Get("validate") == "" {
			continue
		}
		validateMap[key] = strings.Split(field.Tag.Get("validate"), "||")
		if field.Tag.Get("fieldName") == "" {
			continue
		}
		validateNameMap[key] = field.Tag.Get("fieldName")
	}
}
