package utils

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"strings"
)

func getChMessage(data map[string][]string) map[string][]string {
	messages := make(map[string][]string)
	for key, v := range data {
		var item []string
		for _, k := range v {
			switch k {
			case "alpha":
				item = append(item, fmt.Sprintf("%s:%s 字段只能是字母。", k, key))
			case "alpha_dash":
				item = append(item, fmt.Sprintf("%s:%s 字段只能包含字母数字字符、破折号和下划线。", k, key))
			case "alpha_space":
				item = append(item, fmt.Sprintf("%s:%s 字段只能包含字母数字字符、破折号、下划线和空格。", k, key))
			case "alpha_num":
				item = append(item, fmt.Sprintf("%s:%s 字段必须包含字母和数字。", k, key))
			case "numeric":
				item = append(item, fmt.Sprintf("%s:%s 字段必须为数字。", k, key))
			case "bool":
				item = append(item, fmt.Sprintf("%s:%s 字段必须能够转换为布尔值。如：0，1，false，true。", k, key))
			case "coordinate":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是有效坐标的值。", k, key))
			case "css_color":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是一个有效的CSS颜色值。如：#909，#00aaff, rgb(255,122,122)", k, key))
			case "date":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是有效日期。如：yyyy-mm-dd或yyyy/mm/dd。", k, key))
			case "email":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是有效的电子邮件。", k, key))
			case "float":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是一个有效的浮点数。", k, key))
			case "mac_address":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是一个有效的Mac地址。", k, key))
			case "ip":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是一个有效的ip地址。", k, key))
			case "ip_v4":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是一个有效的IP V4地址。", k, key))
			case "ip_v6":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是有效的IP V6地址。", k, key))
			case "json":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是一个有效的json字符串。", k, key))
			case "lat":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是有效的纬度。", k, key))
			case "lon":
				item = append(item, fmt.Sprintf("%s:%s 段必须是有效的经度。", k, key))
			case "required":
				item = append(item, fmt.Sprintf("%s:%s 字段不能为空", k, key))
			case "url":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是有效的url。", k, key))
			case "uuid":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是有效的uuid。", k, key))
			case "uuid_v3":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是有效的UUID V3。", k, key))
			case "uuid_v4":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是有效的UUID V4。", k, key))
			case "uuid_v5":
				item = append(item, fmt.Sprintf("%s:%s 字段必须是有效的UUID V5。", k, key))
			case "regex":
				item = append(item, fmt.Sprintf("%s:%s 字段不是正确的数据格式", k, key))
			default:
				n := strings.Index(k, ":")
				if n > 0 {
					rlt := strings.Split(k, ":")
					switch rlt[0] {
					case "size":
						item = append(item, fmt.Sprintf("%s:%s 文件大小不能超过 %s。", rlt[0], key, rlt[1]))
					case "ext":
						item = append(item, fmt.Sprintf("%s:%s 文件扩展名只能是：%s", rlt[0], key, rlt[1]))
					case "digits:int":
						item = append(item, fmt.Sprintf("%s:%s 字段必须是数字，并且长度为：%s。", rlt[0], key, rlt[1]))
					case "in":
						item = append(item, fmt.Sprintf("%s:%s 字段只能是（%s）中的一个。", rlt[0], key, rlt[1]))
					case "not_in":
						item = append(item, fmt.Sprintf("%s:%s 字段不能能是（%s）中的一个。", rlt[0], key, rlt[1]))
					case "min":
						item = append(item, fmt.Sprintf("%s:%s 字段的长度不能小于：%s", rlt[0], key, rlt[1]))
					case "max":
						item = append(item, fmt.Sprintf("%s:%s 字段的长度不能大于：%s", rlt[0], key, rlt[1]))
					case "len":
						item = append(item, fmt.Sprintf("%s:%s 字段的长度必须是：%s", rlt[0], key, rlt[1]))
					case "between":
						num := strings.Split(rlt[1], ",")
						item = append(item, fmt.Sprintf("%s:%s 字段长度必须在 %s 和 %s 之间。", rlt[0], key, num[0], num[1]))
					case "numeric_between":
						num := strings.Split(rlt[1], ",")
						item = append(item, fmt.Sprintf("%s:%s 字段必须是数字，且值范围只能在 %s 和 %s 之间。", rlt[0], key, num[0], num[1]))
					case "digits_between":
						num := strings.Split(rlt[1], ",")
						item = append(item, fmt.Sprintf("%s:%s 字段必须是数字，且值长度只能在 %s 和 %s 之间。", rlt[0], key, num[0], num[1]))
					}
				}
			}
		}
		messages[key] = item
	}
	return messages
}
func Validator(r *http.Request, data map[string][]string) []string {
	messages := getChMessage(data)
	rules := govalidator.Options{
		Request:         r,
		Rules:           data,
		Messages:        messages,
		RequiredDefault: true,
	}
	v := govalidator.New(rules)
	e := v.Validate()
	if len(e) == 0 {
		return nil
	}
	for _, v := range e {
		return v
	}
	return nil
}
