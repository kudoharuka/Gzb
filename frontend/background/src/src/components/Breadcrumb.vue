<template>
  <el-breadcrumb separator-class="el-icon-arrow-right">
    <el-breadcrumb-item :to="{ path: '/adminHome' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item>{{title}}</el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script>
export default {
  name: "Breadcrumb",
  props:{
    title:{
      type:String,
    }
  }
}
</script>

<style scoped>

</style>