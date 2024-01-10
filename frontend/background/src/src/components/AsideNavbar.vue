<template>
  <el-container style="height: 100%; border: 1px solid #eee">
    <!--    侧边菜单导航-->
    <el-aside width="250px" style="background-color: #004578">
      <el-menu background-color="#004578" text-color="#fff" active-text-color="#409EFF" default-active="1"
               :router="true">

        <div class="logo">
          <img src="../assets/logo.png" style="width: 40px;padding:10px">
          <span style="font-weight: bold;">工作帮管理系统</span>
        </div>
        <!--    用户管理子菜单-->
        <el-menu-item index="adminHome"><i class="el-icon-s-home"/> <span class="el-menu-text">首页</span></el-menu-item>
        <el-submenu index="1">
          <template slot="title">
            <i class="el-icon-s-custom"></i>
            <span>用户管理</span>
          </template>
          <el-menu-item-group>
            <el-menu-item index="user">
              <i class="el-icon-user"/> <span class="el-menu-text">普通用户管理</span>
            </el-menu-item>
            <el-menu-item index="enterprise">
              <i class="el-icon-office-building"/> <span class="el-menu-text">企业管理</span>
            </el-menu-item>
          </el-menu-item-group>
        </el-submenu>
        <!--    内容管理子菜单-->
        <el-submenu index="2">
          <template slot="title">
            <i class="el-icon-location"></i>
            <span>内容管理</span>
          </template>
          <el-menu-item-group>
            <el-menu-item index="post">
              <i class="el-icon-tickets"/> <span class="el-menu-text">帖子管理</span>
            </el-menu-item>
            <el-menu-item index="news">
              <i class="el-icon-discover"/> <span class="el-menu-text">资讯管理</span>
            </el-menu-item>
            <el-menu-item index="comment">
              <i class="el-icon-s-comment"/> <span class="el-menu-text">评论管理</span>
            </el-menu-item>
            <el-menu-item index="feedback">
              <i class="el-icon-message-solid"/> <span class="el-menu-text">反馈中心</span>
            </el-menu-item>
          </el-menu-item-group>
        </el-submenu>
        <el-menu-item index="about">
          <i class="el-icon-info"/> <span class="el-menu-text">关于</span>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <!--    主体头部-->
    <el-container>
      <el-header style="text-align: right; font-size: 12px;">
        <el-dropdown trigger="click">
          <div class="top_right">
            <el-avatar :size="'large'" :src="avatar" :fit="'cover'"
                       style="margin-right: 10px"></el-avatar>
            <span class="el-dropdown-link">{{ account }}<i class="el-icon-caret-bottom el-icon--right"></i></span></div>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item class="clearfix" @click.native="toProfile">个人页面</el-dropdown-item>
            <el-dropdown-item class="clearfix" @click.native="logout">退出</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </el-header>
      <el-main class="el-main">
        <router-view></router-view>
      </el-main>
      <el-footer class="top_right">Copyright © 2023-2024 工作帮 版权所有.</el-footer>
    </el-container>
  </el-container>
</template>
<script>
import avatar from "../assets/images/profile.jpg"
import {resetToRoutes} from "@/router";

export default {
  data() {
    return {
      avatar: avatar,
      account: window.sessionStorage.getItem("account"),
      id: window.sessionStorage.getItem("id"),
    }
  },
  methods: {
    logout() {

      this.$confirm('确定要退出登录？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        window.sessionStorage.clear();
        this.$router.push('/login');
        resetToRoutes();
      })
    },
    toProfile() {
      this.$router.push('/profile');
    }
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

.el-header {
  background-color: #ffffff;
  color: #333;
  line-height: 60px;

}
.el-footer{
  background: #ffffff;
  color: #004578;
  font-size: 10px;
  line-height: 60px;
}

.el-aside {
  color: #ffffff;
}

.el-dropdown-link {
  font-size: 20px;

}

.el-main {
  background-color: #e5f1f8;
}

</style>