/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-30 14:15:13
 * @LastEditTime: 2022-03-22 13:30:49
 * @Description:
 */
package source

import (
	"fmt"
	"path/filepath"
	"testing"

	"com.fanyunai.sqlHelper/source/utils"
)

func Test_Main(t *testing.T) {
	str := "acsdsda_sda_dra"

	fmt.Println(utils.ToField(str))

	file, err := filepath.Abs("")
	if err != nil {
		fmt.Println("问" + file)
	}
	// api.ExportModelFile()
	fmt.Println(file)
}
