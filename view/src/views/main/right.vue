<!--
  * @Author: WUMIANHUA
  * @Date: 2021-10-30 14:15:13
  * @LastEditTime: 2022-03-22 13:29:43
  * @Description:
-->
<template>
  <div class="right-contain">
    <el-empty v-if="isEmpty" description="请选择数据库"></el-empty>
    <template v-else>
      <div class="right-top">
        <el-button type="primary" @click="connectSql" size="mini"
          >刷新数据库</el-button
        >
        <span>
          <el-button
            type="primary"
            icon="el-icon-edit"
            size="mini"
            circle
          ></el-button
          ><el-button
            type="danger"
            icon="el-icon-delete"
            size="mini"
            @click="deleteHandler"
            circle
          ></el-button>
        </span>
      </div>
      <template v-if="tablenames.length == 0">
        <el-empty description="该数据库没有表"></el-empty>
      </template>
      <template v-else>
        <div class="right-all-check">
          <el-checkbox
            :indeterminate="isIndeterminate"
            v-model="checkAll"
            @change="handleCheckAllChange"
            >全选</el-checkbox
          >
        </div>
        <div class="right-center">
          <el-checkbox-group
            v-model="checkedTables"
            @change="handleCheckedChange"
          >
            <el-checkbox
              border
              v-for="table in tablenames"
              :label="table"
              :key="table"
              >{{ table }}</el-checkbox
            >
          </el-checkbox-group>
        </div>
      </template>

      <div class="right-bottom">
        <el-button
          type="success"
          size="mini"
          :disabled="checkedTables.length == 0"
          >全部导出</el-button
        >
        <el-button
          type="primary"
          size="mini"
          :disabled="checkedTables.length == 0"
          @click="exportModel"
          >导出模型</el-button
        >
      </div>
    </template>
  </div>
</template>
<script>
export default {
  name: "MainRight",
  props: {
    rightData: {
      type: Object,
    },
  },
  data() {
    return {
      tablenames: [],
      connectData: {},
      isIndeterminate: false,
      checkAll: false,
      checkedTables: [],
    };
  },
  computed: {
    isEmpty() {
      let keys = Object.keys(this.$props.rightData);
      if (keys.length > 0) {
        this.connectSql();
      }
      return keys.length === 0;
    },
  },
  created() {},
  methods: {
    // 单个选择事件
    handleCheckedChange(val) {
      // 是否全选
      let checkedCount = val.length;
      this.checkAll = checkedCount === this.tablenames.length;
      this.isIndeterminate =
        checkedCount > 0 && checkedCount < this.tablenames.length;
    },
    // 全选
    handleCheckAllChange(val) {
      this.checkedTables = val ? this.tablenames : [];
      this.isIndeterminate = false;
    },
    // 删除事件
    deleteHandler() {
      var that = this;
      this.$confirm("此操作将永久删除该连接, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        that.$POST("/deleteconnect", that.$props.rightData, function (res) {
          if (res.code == 200) {
            that.$emit("deleteConnect");
          }
        });
      });
    },
    // 连接数据库
    connectSql() {
      var that = this;
      this.checkedTables = [];
      this.isIndeterminate = false;
      this.checkAll = false;
      // 加载中
      this.$emit("checkLoading", true);
      // 获取所有表
      this.$POST("/gettable", this.$props.rightData, function (res) {
        that.$emit("checkLoading", false);
        if (res.code == 200) {
          that.tablenames = res.data;
        }
      });
    },
    // 导出模型
    exportModel() {
      this.setTablenames();
      var that = this;
      // 加载中
      this.$emit("checkLoading", true);
      // 到出模型
      this.$POST("/exportmodel", this.connectData, function (res) {
        that.$emit("checkLoading", false);
        if (res.code == 200) {
          that.checkedTables = [];
          that.checkAll = false;
        }
      });
    },
    // 全选事件
    handleCheckedCitiesChange(value) {
      let checkedCount = value.length;
      this.checkAll = checkedCount === this.cities.length;
      this.isIndeterminate =
        checkedCount > 0 && checkedCount < this.cities.length;
    },
    // 设置表数据
    setTablenames() {
      this.connectData = Object.assign({}, this.$props.rightData);
      this.connectData.tablename = this.checkedTables;
    },
  },
};
</script>
<style lang='scss' scoped>
.right-contain {
  height: 100%;
  width: 100%;
}
.right-top {
  margin: 0px 10px;
  display: flex;
  justify-content: space-between;
  padding-bottom: 10px;
  padding-top: 10px;
  border-bottom: 1px #ccc solid;
}
.right-all-check {
  height: 22px;
  margin: 10px 10px 0px;
}
.right-center {
  display: flex;
  align-items: flex-start;
  margin: 10px;
  height: calc(100% - 170px);
  overflow-y: auto;
  padding-bottom: 10px;
  border-bottom: 1px #ccc solid;
  .el-checkbox {
    margin-left: 0px !important;
    margin-bottom: 10px;
  }
}
.right-bottom {
  margin: 10px;
  display: flex;
  justify-content: flex-end;
}
</style>