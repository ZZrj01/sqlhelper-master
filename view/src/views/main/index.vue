/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-30 14:15:13
 * @LastEditTime: 2022-03-22 13:29:43
 * @Description:
 */
 
<template>
  <div
    class="main-contain"
    v-loading="loading"
    element-loading-text="加载中..."
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.5)"
  >
    <div class="left">
      <el-button
        circle
        size="mini"
        type="primary"
        icon="el-icon-plus"
        class="add-btn"
        @click="addHandler"
      ></el-button>
      <template v-for="(item, index) in dataConnects">
        <div
          class="div-item"
          :class="{ active: index == checkLeftId }"
          :key="index"
          @click="checkItem(item, index)"
        >
          <div>
            <h4>{{ item.name }}</h4>
            <h6>
              {{ getSqlType(item.type) }} -
              {{ item.host }}
            </h6>
          </div>
          <i class="el-icon-arrow-right"></i>
        </div>
      </template>
    </div>
    <div class="right">
      <MainRight
        :rightData="rightData"
        @checkLoading="checkLoading"
        @deleteConnect="deleteConnect"
      ></MainRight>
    </div>
    <!-- 添加配置弹出 -->
    <el-dialog
      title="配置"
      :visible.sync="dialogVisible"
      :modal="false"
      :fullscreen="true"
      :show-close="false"
    >
      <el-form
        status-icon
        :model="addForm"
        ref="addForm"
        :rules="formRules"
        label-width="100px"
      >
        <el-tabs type="border-card">
          <el-tab-pane label="基本配置">
            <el-row :gutter="20">
              <el-col :span="11">
                <el-form-item label="项目名称" prop="name">
                  <el-input
                    v-model="addForm.name"
                    placeholder="请输入项目名称"
                  ></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="11">
                <el-form-item label="数据库类型">
                  <el-select v-model="addForm.type">
                    <el-option
                      v-for="item in options"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value"
                    >
                    </el-option> </el-select
                ></el-form-item>
              </el-col>
            </el-row>
            <template v-if="addForm.type == 0">
              <el-row :gutter="20">
                <el-col :span="11"
                  ><el-form-item label="host" prop="host">
                    <el-input
                      v-model="addForm.host"
                      placeholder="请输入数据库地址"
                    ></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="11">
                  <el-form-item label="port" prop="port">
                    <el-input
                      v-model="addForm.port"
                      placeholder="请输入数据库端口号"
                    ></el-input>
                  </el-form-item>
                </el-col>
              </el-row>
              <el-row :gutter="20">
                <el-col :span="11"
                  ><el-form-item label="dbname" prop="dbname">
                    <el-input
                      v-model="addForm.dbname"
                      placeholder="请输入数据库名字"
                    ></el-input> </el-form-item
                ></el-col>
                <el-col :span="11">
                  <el-form-item label="charset" prop="charset">
                    <el-input
                      v-model="addForm.charset"
                      placeholder="请输入字符集"
                    ></el-input> </el-form-item
                ></el-col>
              </el-row>
              <el-row :gutter="20">
                <el-col :span="11"
                  ><el-form-item label="username" prop="username">
                    <el-input
                      v-model="addForm.username"
                      placeholder="请输入用户名"
                    ></el-input> </el-form-item
                ></el-col>
                <el-col :span="11"
                  ><el-form-item label="password" prop="password">
                    <el-input
                      v-model="addForm.password"
                      placeholder="请输入用户密码"
                    ></el-input> </el-form-item
                ></el-col>
              </el-row>
              <el-row :gutter="20">
                <el-col :span="11"
                  ><el-form-item label="表前缀">
                    <el-input
                      v-model="addForm.tableprefix"
                      placeholder="请输入表前缀"
                    ></el-input> </el-form-item
                ></el-col>
              </el-row>
            </template>
            <template v-if="addForm.type == 1">
              <el-form-item label="数据库地址">
                <el-button size="small">选择文件</el-button>
              </el-form-item>
            </template>
            <el-form-item label="">
              <el-button @click="testConnect()" size="small"
                >测试连接</el-button
              >
            </el-form-item>
          </el-tab-pane>
          <!-- model 配置 -->
          <el-tab-pane label="model配置">
            <el-row :gutter="20">
              <el-col :span="11">
                <el-form-item label="包名">
                  <el-input
                    v-model="addForm.model.name"
                    placeholder="请输入包名"
                  ></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row :gutter="20">
              <el-col :span="11">
                <el-form-item
                  label="路径"
                  prop="model.path"
                  :rules="{
                    required: true,
                    message: '不能为空',
                    trigger: 'blur',
                  }"
                >
                  <el-input
                    v-model="addForm.model.path"
                    placeholder="请输入导出路径"
                  ></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item label="字段">
              <el-checkbox-group v-model="addForm.model.fields">
                <el-checkbox label="json"></el-checkbox>
                <el-checkbox label="form"></el-checkbox>
                <el-checkbox label="gorm"></el-checkbox>
                <el-checkbox label="yaml"></el-checkbox>
                <el-checkbox label="bson"></el-checkbox>
              </el-checkbox-group>
            </el-form-item>
          </el-tab-pane>
          <!-- service配置 -->
          <el-tab-pane label="service配置">
            <el-row>
              <el-col :span="11">
                <el-form-item label="接口包名">
                  <el-input
                    v-model="addForm.service.interName"
                    placeholder="请输入包名"
                  ></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="11">
                <el-form-item label="实现包名">
                  <el-input
                    v-model="addForm.service.relName"
                    placeholder="请输入包名"
                  ></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row :gutter="20">
              <el-col :span="11">
                <el-form-item label="路径">
                  <el-input
                    v-model="addForm.service.path"
                    placeholder="请输入导出路径"
                  ></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item label="方法">
              <el-checkbox-group v-model="addForm.service.methods">
                <el-checkbox
                  :label="key"
                  v-for="key in methods"
                  :key="key"
                ></el-checkbox>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item label="新增参数类型">
              <el-input
                v-model="addForm.service.params"
                placeholder="model.BaseModel、string"
              ></el-input>
            </el-form-item>
          </el-tab-pane>
          <!-- dao配置 -->
          <el-tab-pane label="dao配置">
            <el-row>
              <el-col :span="11">
                <el-form-item label="接口包名">
                  <el-input
                    v-model="addForm.dao.interName"
                    placeholder="请输入包名"
                  ></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="11">
                <el-form-item label="实现包名">
                  <el-input
                    v-model="addForm.dao.relName"
                    placeholder="请输入包名"
                  ></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row :gutter="20">
              <el-col :span="11">
                <el-form-item label="路径">
                  <el-input
                    v-model="addForm.dao.path"
                    placeholder="请输入导出路径"
                  ></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item label="方法">
              <el-checkbox-group v-model="addForm.dao.methods">
                <el-checkbox
                  :label="key"
                  v-for="key in methods"
                  :key="key"
                ></el-checkbox>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item label="新增参数类型">
              <el-input
                v-model="addForm.dao.params"
                placeholder="model.BaseModel、string"
              ></el-input>
            </el-form-item>
          </el-tab-pane>
        </el-tabs>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false" size="small">取消</el-button>
        <el-button
          type="primary"
          @click="commitAddHandler('addForm')"
          size="small"
          >添加</el-button
        >
      </span>
    </el-dialog>
  </div>
</template>
<script>
import MainRight from "./right.vue";
var vm;
export default {
  name: "",
  components: {
    MainRight,
  },
  data() {
    return {
      loading: false, // 界面加载中
      dialogVisible: false, // 添加数据库弹窗
      checkLeftId: -1,
      // 数据连接对象
      dataConnects: [],
      // 右边选中数据
      rightData: {},
      // 支持的方法
      methods: ["add", "delete", "modify", "select", "exist"],
      // 支持的数据类型
      options: [
        { label: "mysql", value: 0 },
        { label: "sqlite", value: 1 },
      ],
      // 默认form
      form: {
        name: "项目1", //项目名称
        type: 0, // 类型
        host: "127.0.0.1", // 数据库地址
        port: 3306, // 端口
        dbname: "", // 数据库名字
        charset: "utf8", // 字符集
        tableprefix: "tab_",
        username: "", // 名称
        password: "", // 密码
        dbPath: "", // sqlite 文件地址
        // model 包
        model: {
          name: "model",
          fields: ["json", "form", "gorm"],
          path: "",
        },
        // service
        service: {
          interName: "service",
          relName: "impl",
          methods: [],
          params: "",
          path: "",
        },
        // dao
        dao: {
          interName: "dao",
          relName: "impl",
          methods: [],
          params: "",
          path: "",
        },
      },
      addForm: null,
      formRules: {
        name: [{ required: true, message: "不能为空", trigger: "blur" }],
        host: [{ required: true, message: "不能为空", trigger: "blur" }],
        port: [{ required: true, message: "不能为空", trigger: "blur" }],
        dbname: [{ required: true, message: "不能为空", trigger: "blur" }],
        username: [{ required: true, message: "不能为空", trigger: "blur" }],
        password: [{ required: true, message: "不能为空", trigger: "blur" }],
        charset: [{ required: true, message: "不能为空", trigger: "blur" }],
        dbPath: [{ required: true, message: "不能为空", trigger: "blur" }],
      },
    };
  },
  //
  created() {
    // 全局
    vm = this;
    // 初始化
    this.initForm();
  },
  methods: {
    // 初始化form
    initForm() {
      // 初始化
      this.addForm = Object.assign({}, this.form);
      this.addForm.service.methods = this.methods;
      this.addForm.dao.methods = this.methods;
      // 获取初始化数据
      this.getInitData();
    },
    // 是否加载中
    checkLoading(val) {
      this.loading = val;
    },
    // 添加事件
    addHandler() {
      this.dialogVisible = true;
    },
    // 测试连接
    testHandler() {},
    // 提交添加事件
    commitAddHandler(formName) {
      this.$refs[formName].validate((valid) => {
        console.log(valid);
        if (valid) {
          console.log(vm.addForm);
          // 提交
          vm.$POST("/addconnect", vm.addForm, function (res) {
            if (res.code == 200) {
              // 获取数据
              vm.getDataConnect();
              vm.dialogVisible = false;
            }
          });
        } else {
          if (vm.addForm.model.path.length == 0) {
            vm.$message.error("model 配置下的路径不能为空");
          }
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    // 获取所有连接
    getDataConnect() {
      vm.$POST("/getconnect", {}, function (res) {
        if (res.code == 200) {
          vm.dataConnects = res.data;
        }
      });
    },
    // 获取初始化
    getInitData() {
      this.loading = true;
      vm.$POST("/initialize", {}, function (res) {
        vm.loading = false;
        if (res.code == 200) {
          vm.dataConnects = res.data.sqls;
        }
      });
    },
    // 测试连接
    testConnect() {
      this.loading = true;
      vm.$POST("/testConnect", vm.addForm, function (res) {
        vm.loading = false;
        if (res.code == 200) {
          vm.$message.success("连接成功");
        } else {
          vm.$message.error("连接失败");
        }
      });
    },
    // 获取数据库类型
    getSqlType(type) {
      for (const item of this.options) {
        if (item.value == type) {
          return item.label;
        }
      }
      return "mysql";
    },
    // 切换数据
    checkItem(item, index) {
      this.checkLeftId = index;
      this.rightData = item;
    },
    // 删除连接
    deleteConnect() {
      this.dataConnects.splice(this.checkLeftId, 1);
      this.rightData = {};
      this.checkLeftId = -1;
    },
  },
};
</script>
<style lang='scss' scoped>
// 宽度
$leftWidth: 200px;
.main-contain {
  position: relative;
  height: 100%;
}
.left {
  background-color: #c0c4cc;
  width: $leftWidth;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  .add-btn {
    margin: 10px;
  }
  .div-item {
    padding-left: 10px;
    height: 60px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #f2f6fc;
    & > i {
      color: white;
    }
    & > div {
      h4,
      h6 {
        margin: 0;
      }
    }
  }
  .active {
    background-color: #a2a7b2;
  }
}
.right {
  width: calc(100% - #{$leftWidth});
  height: 100%;
  position: absolute;
  top: 0;
  left: $leftWidth;
}
</style>