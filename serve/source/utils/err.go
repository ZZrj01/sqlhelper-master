/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-30 14:15:13
 * @LastEditTime: 2022-03-22 13:30:12
 * @Description:
 */
package utils

import "fmt"

var (
	ERR_PARAMETER           = New(1001, "参数错误")
	ERR_USER_DOES_NOT_EXIST = New(1002, "用户不存在")
	ERR_USER_DOES_EXIST     = New(1003, "用户已存在")
	ERR_USER_DOES_KNOWN     = New(1004, "未知请求")
	ERR_USER_PASSWORD_WRONG = New(1005, "用户名或密码错误")
	ERR_SYSTEM_INSTALLED    = New(1006, "系统已初始化")
	ERR_PATH_NOT_EXIST      = New(2001, "非有效路径")
	ERR_PATH_EXISTSED       = New(2002, "路径已存在")
	ERR_PATH_EXISTS_CHILD   = New(2003, "该目录含有子目录不能删除")
	ERR_PATH_ROOT_ERROR     = New(2004, "仓库不能修改名字")
	ERR_AUTH_NOT_EXIST      = New(3001, "用户未登录")
	ERR_AUTH_NOT_VALID      = New(3002, "用户已过期请重新登陆")
	ERR_AUTH_OTHER_LOGIN    = New(3003, "账号已在其他地方登陆")
	ERR_MAX_BUILD           = New(4001, "已超过最大数")
	ERR_NOT_VALID_NUMBER    = New(4002, "请输入有效手机号码")
	ERR_HAS_ROOM            = New(4003, "该房间已处在")
	ERR_ROOM_MAX            = New(4004, "楼层不能大于建筑楼层数")
	ERR_NAME_EXISTSED       = New(4005, "该名字已存在")
	ERR_UNBIND_ROOM         = New(4006, "已有房间绑定了该模板，请先解绑再删除！")
)

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 实现 error 接口
func (e *Err) Error() string {
	return fmt.Sprintf("%d:%s", e.Code, e.Message)
}

func New(code int, msg string) (err *Err) {
	return &Err{
		Code:    code,
		Message: msg,
	}
}
