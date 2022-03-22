/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-30 14:15:13
 * @LastEditTime: 2022-03-22 13:29:58
 * @Description:
 */
package model

type Config struct {
	Server   Server     `yaml:"server"`
	SqlField []SqlField `yaml:"sqlfield"`
}

type Server struct {
	Port         int    `yaml:"port"`
	SerInterName string `yaml:"serInterName"`
	SerRelName   string `yaml:"serRelName"`
	DaoInterName string `yaml:"daoInterName"`
	DaoRelName   string `yaml:"daoRelName"`
}

type SqlField struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
