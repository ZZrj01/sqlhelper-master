/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-30 14:15:13
 * @LastEditTime: 2022-03-22 13:35:40
 * @Description:
 */
package utils

import (
	"fmt"
	"strings"

	"com.fanyunai.sqlHelper/source/config"
	"com.fanyunai.sqlHelper/source/model"
)

/**
 * @description: 获取数据库连接字符串
 * @param {model.ItemConnect} model
 * @return {*}
 */
func GetDbConnectStr(model model.Connect) (sql string, str string) {
	if model.Type == 0 {
		sql = "mysql"
		str = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", model.UserName, model.Password, model.Host, model.Port, model.DbName, model.Charset)
	} else {
		sql = "sqlite3"
		str = model.DbPath
	}
	return
}

/**
 * @description: 字母转换
 * @param {string} s
 * @return {*}
 */
func UpperString(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return s
}

/**
 * @description: 下划线分割字符串
 * @param {string} s
 * @return {*}
 */
func ToField(s string) string {
	temp := make([]string, 0)
	strs := strings.Split(s, "_")
	for _, st := range strs {
		temp = append(temp, UpperString(st))
	}
	return strings.Join(temp, "")
}

/**
 * @description: 首字母小写
 * @param {string} s
 * @return {*}
 */
func LeftLower(s string) string {
	if len(s) > 0 {
		return strings.ToLower(string(s[0])) + s[1:]
	}
	return s
}

/**
 * @description: 字段添加后缀
 * @param {*}
 * @return {*}
 */
func FieldSuffix(strs []string, field string) string {
	temp := make([]string, 0)
	for _, v := range strs {
		if v == "gorm" {
			temp = append(temp, fmt.Sprintf("%s:\"%s\"", v, field))
		} else {
			temp = append(temp, fmt.Sprintf("%s:\"%s\"", v, LeftLower(ToField(field))))
		}
	}
	return strings.Join(temp, " ")
}

/**
 * @description: 获取数据类型
 * @param {string} t
 * @return {*}
 */
func FieldType(t string) string {
	for _, v := range config.Conf.SqlField {
		if v.Key == t {
			return v.Value
		}
	}
	return "string"
}

/**
 * @description: 去除表前缀
 * @param {*}
 * @return {*}
 */
func RemoveTablePrefix(tablename string, prefix string) string {
	if len(prefix) > 0 {
		return tablename[len(prefix):]
	}
	return tablename
}
