package tools

import (
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

func BuildQuery(params map[string]string) string {
	if params == nil {
		return ""
	}

	query := &strings.Builder{}
	hasParam := false

	for k, v := range params {
		if v != "" {
			if hasParam {
				query.WriteString("&")
			} else {
				hasParam = true
			}

			query.WriteString(k)
			query.WriteString("=")
			query.WriteString(url.QueryEscape(v))
		}
	}
	return query.String()
}

var getname = func(v string) string {
	return strings.Split(v, ",")[0]
}

func Signbuf(val interface{}) string {
	nameArr := make([]string, 0)
	keyVal := make(map[string]string, 0)
	v := reflect.ValueOf(val).Elem()
	st := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.CanInterface() {
			if f.Type().Kind() == reflect.Interface {
				continue
			}
			if f.Type().Kind() == reflect.Struct {
				sf := f.Type()
				for j := 0; j < sf.NumField(); j++ {
					name := getname(sf.Field(j).Tag.Get("post"))
					value := f.Field(j).String()
					if len(value) > 0 && name != "signature" {
						nameArr = append(nameArr, name)
						keyVal[name] = value
					}
				}
				continue
			}
			name := getname(st.Field(i).Tag.Get("post"))
			value := f.String()
			if len(value) > 0 && name != "signature" {
				nameArr = append(nameArr, name)
				keyVal[name] = value
			}
		}

	}
	sort.Strings(nameArr)
	var signBuf bytes.Buffer
	for _, name := range nameArr {
		signBuf.WriteString(fmt.Sprintf("%s=%s&", name, keyVal[name]))
	}
	signBuf.Truncate(signBuf.Len() - 1)
	return signBuf.String()
}

/*通用组请求函数*/
func RequestBuf(val interface{}) string {
	v := reflect.ValueOf(val).Elem()
	st := v.Type()
	u := url.Values{}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.CanInterface() {
			if f.Type().Kind() == reflect.Interface {
				continue
			}
			name := getname(st.Field(i).Tag.Get("post"))
			value := f.String()
			if len(value) > 0 {
				u.Add(name, value)
			}
		}
	}
	return u.Encode()
}
