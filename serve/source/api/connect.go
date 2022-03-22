/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-30 14:15:13
 * @LastEditTime: 2022-03-22 13:35:06
 * @Description:
 */
package api

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"com.fanyunai.sqlHelper/source/config"
	"com.fanyunai.sqlHelper/source/model"
	"com.fanyunai.sqlHelper/source/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/yaml.v2"
)

/**
 * @description: 初始化
 * @param {*gin.Context} c
 * @return {*}
 */
func Initialize(c *gin.Context) {
	// 返回基础数据
	utils.Success(c, map[string]interface{}{
		"serve": config.Conf.Server,
		"sqls":  config.DataBase,
	})
}

/**
 * @description: 测试连接
 * @param {*gin.Context} c
 * @return {*}
 */
func TestConnect(c *gin.Context) {
	var connect model.Connect
	err := c.ShouldBindJSON(&connect)
	if err != nil {
		utils.Error(c, err)
		return
	}
	// 获取连接字符串
	sqlType, str := utils.GetDbConnectStr(connect)
	db, err := gorm.Open(sqlType, str)
	if err != nil {
		utils.Error(c, err)
		return
	}
	defer db.Close()
	utils.Success(c, nil)
}

/**
 * @description: 添加连接
 * @param {*gin.Context} c
 * @return {*}
 */
func AddConnect(c *gin.Context) {
	// 解析对像
	var connect model.Connect
	err := c.ShouldBindJSON(&connect)
	if err != nil {
		utils.Error(c, err)
		return
	}
	connect.ID = time.Now().Unix()

	// 追加到数组
	config.DataBase = append(config.DataBase, connect)
	err = saveData(config.DataBase)
	if err != nil {
		utils.Error(c, err)
		return
	}
	utils.Success(c, connect.ID)
}

/**
 * @description: 查询所有数据
 * @param {*gin.Context} c
 * @return {*}
 */
func GetConnect(c *gin.Context) {
	// 返回所有数据
	utils.Success(c, config.DataBase)
}

/**
 * @description: 查询所有表
 * @param {*gin.Context} c
 * @return {*}
 */
func GetTables(c *gin.Context) {
	var connect model.Connect
	err := c.ShouldBindJSON(&connect)
	if err != nil {
		utils.Error(c, err)
		return
	}
	// 获取连接字符串
	sqlType, str := utils.GetDbConnectStr(connect)
	db, err := gorm.Open(sqlType, str)
	if err != nil {
		utils.Error(c, err)
		return
	}
	// 延迟关闭数据库
	defer db.Close()
	var rows *sql.Rows
	// 查询所有表
	switch connect.Type {
	case 0: // mysql
		rows, err = db.Raw("select table_name from information_schema.tables where table_schema = ?;", connect.DbName).Rows()
	case 1: // sqlite

	}
	if err == nil {
		// 表数组
		tableNames := make([]string, 0)
		// 遍历结果
		for rows.Next() {
			var tableName string
			err = rows.Scan(&tableName)
			if err == nil {
				tableNames = append(tableNames, tableName)
			}
		}
		utils.Success(c, tableNames)
	} else {
		utils.Error(c, err)
	}
}

/**
 * @description: 删除连接
 * @param {*gin.Context} c
 * @return {*}
 */
func DeleteConnect(c *gin.Context) {
	var connect model.Connect
	err := c.ShouldBindJSON(&connect)
	if err != nil {
		utils.Error(c, err)
		return
	}
	index := -1
	for i, v := range config.DataBase {
		if v.ID == connect.ID {
			index = i
			break
		}
	}
	if index != -1 {
		// 删除数据
		config.DataBase = append(config.DataBase[:index], config.DataBase[index+1:]...)
		// 保存数据
		saveData(config.DataBase)
		utils.Success(c, nil)
	} else {
		utils.Error(c, fmt.Errorf("删除失败"))
	}
}

/**
 * @description: 到出模型
 * @param {*gin.Context} c
 * @return {*}
 */
func ExportModel(c *gin.Context) {
	var connect model.Connect
	err := c.ShouldBindJSON(&connect)
	if err != nil {
		utils.Error(c, err)
		return
	}
	// 获取连接字符串
	sqlType, str := utils.GetDbConnectStr(connect)
	db, err := gorm.Open(sqlType, str)
	// db.Debug()
	db.LogMode(true)
	if err != nil {
		utils.Error(c, err)
		return
	}
	// 延迟关闭数据库
	defer db.Close()
	var rows *sql.Rows
	// 查询所有表
	switch connect.Type {
	case 0: // mysql
		rows, err = db.Raw("select COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT,TABLE_NAME from information_schema.columns where table_name in (?);", connect.TableName).Rows()
	case 1: // sqlite

	}
	if err == nil {
		// 表数组
		fields := make([]model.Table, 0)
		// 遍历结果
		for rows.Next() {
			var f1, f2, f3, f4 sql.NullString
			var item model.Table
			err = rows.Scan(&f1, &f2, &f3, &f4)
			if err == nil {
				item.Field = f1.String
				item.Type = f2.String
				item.Desc = f3.String
				item.Name = f4.String
				fields = append(fields, item)
			}
		}
		// 导出model
		for _, tablename := range connect.TableName {
			exportModelFile(fields, connect, tablename)
		}
		utils.Success(c, fields)
	} else {
		utils.Error(c, err)
	}
}

/**
 * @description: 加载data.yml
 * @param {*}
 * @return {*}
 */
func loadData() (data []model.Connect, err error) {
	by, err := os.ReadFile(config.DataFilePath)
	if err != nil {
		fmt.Println("加载data.yml 文件错误", err.Error())
		return
	}
	err = yaml.Unmarshal(by, &data)
	if err != nil {
		fmt.Println("加载data.yml 文件错误")
		return
	}
	return
}

/**
 * @description: 存储文件
 * @param {[]model.Connect} data
 * @return {*}
 */
func saveData(data []model.Connect) (err error) {
	by, err := yaml.Marshal(data)
	if err == nil {
		err = os.WriteFile(config.DataFilePath, by, 0777)
	}
	return
}

/**
 * @description: 导出模型
 * @param {*}
 * @return {*}
 */
func exportModelFile(table []model.Table, connect model.Connect, tablename string) {
	className := utils.UpperString(utils.RemoveTablePrefix(tablename, connect.TablePrefix))
	str := fmt.Sprintf("package %s \n\n", connect.Model.Name)
	str += fmt.Sprintf("type %s struct{ \n", className)
	for _, item := range table {
		if item.Name == tablename {
			str += fmt.Sprintf("\t%s %s `%s` // %s\n", utils.ToField(item.Field), utils.FieldType(item.Type), utils.FieldSuffix(connect.Model.Fields, item.Field), item.Desc)
		}
	}
	str += "}\n\n"
	// table
	str += fmt.Sprintf("func (%s) TableName() string {\n", className)
	str += fmt.Sprintf("\treturn \"%s\" \n", tablename)
	str += "}\n"

	file := path.Join(connect.Model.Path, fmt.Sprintf("%s.go", className))
	ioutil.WriteFile(file, []byte(str), 0755)
}

/**
 * @description: 导出service
 * @param {*}
 * @return {*}
 */
func exportService(table []model.Table) {

}

/**
 * @description: 导出dao
 * @param {model.Table} table
 * @return {*}
 */
func exportDao(table []model.Table) {

}
