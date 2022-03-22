/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-30 14:15:13
 * @LastEditTime: 2022-03-22 13:19:54
 * @Description:
 */
package model

type Connect struct {
	ID          int64          `json:"id" yaml:"id"`
	Type        int            `json:"type" yaml:"type"`
	Host        string         `json:"host" yaml:"host"`
	Port        int            `json:"port" yaml:"port"`
	DbName      string         `json:"dbname" yaml:"dbname"`
	UserName    string         `json:"username" yaml:"username"`
	Password    string         `json:"password" yaml:"password"`
	TablePrefix string         `json:"tableprefix" yaml:"tableprefix"`
	Charset     string         `json:"charset" yaml:"charset"`
	DbPath      string         `json:"dbpath" yaml:"dbpath"`
	Name        string         `json:"name" yaml:"name"`
	Model       ConnetModel    `json:"model" yaml:"model"`
	Service     ConnectService `json:"service" yaml:"service"`
	Dao         ConnectService `json:"dao" yaml:"dao"`
	TableName   []string       `json:"tablename" yaml:"-"`
}

type ConnetModel struct {
	Name   string   `json:"name" yaml:"name"`
	Fields []string `json:"fields" yaml:"fields"`
	Path   string   `json:"path" yaml:"path"`
}

type ConnectService struct {
	InterName string   `json:"intername" yaml:"intername"`
	RelName   string   `json:"relname" yaml:"relname"`
	Methods   []string `json:"methods" yaml:"methods"`
	Params    string   `json:"params" yaml:"params"`
	Path      string   `json:"path" yaml:"path"`
}
