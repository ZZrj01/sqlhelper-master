/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-30 14:15:13
 * @LastEditTime: 2022-03-22 13:29:43
 * @Description:
 */
package config

import "com.fanyunai.sqlHelper/source/model"

var (
	Conf         model.Config
	DataBase     []model.Connect // 数据库对象
	DataFilePath string          // 数据库文件路径
)
