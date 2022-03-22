<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
html,
body {
  height: 100%;
  padding: 0;
  margin: 0;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  height: 100%;
}
.el-dialog__body {
  padding: 0 20px !important;
  height: 460px;
}
.el-form-item {
  margin-bottom: 20px !important;
}
.el-tabs__content {
  height: 380px;
}
</style>
