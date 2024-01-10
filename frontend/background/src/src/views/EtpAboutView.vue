<template>
  <div class="app-container" style="margin-top: 20px">
    <el-row :gutter="20">
      <el-col :span="24" :xs="24">
        <el-card>
          <div slot="header" class="clearfix">
            <span>关于我们</span>
          </div>
          <div class="info">
            <div class="img-center">
              <el-image :src="require('../assets/images/GZB.png')" style="width: 290px;"
                        fit="fill"></el-image>
            </div>
            <el-row style="line-height: 42px;" align="middle">
              <el-col :span="16" offset="8">应用简介：由来自福州大学20级软件工程专业的学生制作</el-col>
            </el-row>
            <el-row style="line-height: 42px;" align="middle">
              <el-col :span="16" offset="8">主要功能：提供各企业各岗位工作资讯</el-col>
            </el-row>
            <el-row style="line-height: 42px;" align="middle">
              <el-col :span="16" offset="8">开发目的：方便用户寻找心仪的工作岗位的求职APP</el-col>
            </el-row>
            <el-row style="line-height: 42px;" align="middle">
              <el-col :span="16" offset="8">涉及方向：Vue、uniapp、Go、Gin</el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>


export default {
  components: {},

  data() {
    return {}
  },
  methods: {}
}
</script>
<style scoped>
.text-center {
  /*flex 布局*/
  display: flex;
  /*实现垂直居中*/
  align-items: center;
  /*实现水平居中*/
  justify-content: center;

  text-align: justify;
  width: 120px;
  height: 120px;
  background: rgb(255, 255, 255);
  margin: 0 auto;
  color: rgb(0, 0, 0);
}

.img-center {
  /*flex 布局*/
  display: flex;
  /*实现垂直居中*/
  align-items: center;
  /*实现水平居中*/
  justify-content: center;

  text-align: justify;
  margin: 0 auto;
}

li {
  list-style: none;
}

.list-group-item {
  line-height: 60px;
  width: 290px;
  border-bottom: 1px solid rgb(70, 70, 70)
}

.list-group {
  margin-left: -40px;
}

.pull-right {
  float: right;
}

.info {
  font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
  font-size: 18px;
  color: rgb(75, 75, 75);
}
</style>
