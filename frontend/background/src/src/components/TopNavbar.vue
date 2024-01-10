<template>
  <el-container style="height: 100%">
    <el-header class="el-header1">
      <el-row :gutter="20">
        <el-col :span="18">
          <span class="el-dropdown-link1">工作帮企业前台</span>
        </el-col>
        <el-col :span="3" :offset="3">
          <el-dropdown trigger="click">
            <div class="top_right">
              <el-avatar :size="'large'" :src=" etp.Logo||avatar" :fit="'cover'"
                         style="margin-right: 10px"></el-avatar>
              <span class="el-dropdown-link1" style="overflow: hidden;text-overflow:ellipsis; width: 100px">{{ etp.Name }}</span></div>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item class="clearfix" @click.native="toProfile">企业信息</el-dropdown-item>
              <el-dropdown-item class="clearfix" @click.native="logout">退出</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </el-col>
      </el-row>
    </el-header>
    <el-container>
      <!--      <el-aside width="200px" style="background: #0050b3">Aside</el-aside>-->
      <el-main>
<!--        首页导航-->
        <el-menu :default-active="activeIndex" class="el-menu-demo " mode="horizontal" @select="handleSelect" :router="true">
          <el-row gutter="20">
            <el-col :span="6" class="top_right">
              <el-menu-item index="enterpriseHome">首页</el-menu-item>
            </el-col>
            <el-col :span="6" class="top_right">
              <el-menu-item index="etpList">企业一览</el-menu-item>
            </el-col>
            <el-col :span="6" class="top_right">
              <el-menu-item index="etpWorkbench">工作台</el-menu-item>
            </el-col>

            <el-col :span="6" class="top_right">
              <el-menu-item index="etpAbout">关于</el-menu-item>
            </el-col>
          </el-row>
        </el-menu>
        <router-view></router-view>
      </el-main>
      <el-footer class="top_right el-footer1">Copyright © 2023-2024 工作帮 版权所有.</el-footer>
    </el-container>
  </el-container>
</template>
<script>
import {resetToRoutes} from "@/router";

export default {
  mounted() {
    // 监听etp-updated事件，并更新数据
    this.$on('etp-updated', () => {
      this.etp = JSON.parse(window.sessionStorage.getItem("enterprise"));
    });
  },
  data() {
    return {
      etp: JSON.parse(window.sessionStorage.getItem("enterprise")),
      avatar:'https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg' ,
      account:window.sessionStorage.getItem("account"),
      // id: this.data().etp.ID,
    }
  },
  methods: {
    logout() {
      this.$confirm('确定要退出登录？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$router.push('/login');
        window.sessionStorage.clear();
        resetToRoutes();
      })
    },
    toProfile() {
      this.$router.push('/etpProfile');
    },

  }
};
</script>

<style>
.top_right {
  display: flex;
  justify-content: center;
  align-items: center;
}

.logo {
  height: 60px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.el-header1 {
  background: #262626;
  color: #ffffff;
  line-height: 60px;

}
.el-footer1{
  background: #262626;
  color: #ffffff;
  font-size: 10px;
  line-height: 60px;
}

.el-aside {
  color: #ffffff;
}

.el-dropdown-link1 {
  font-size: 20px;
  color: #fcf7f7;
}

.el-main {
  background-color:#e5f1f8;
}

</style>