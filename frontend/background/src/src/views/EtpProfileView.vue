<template>
  <div class="app-container" style="margin-top: 20px">
    <el-row :gutter="20" class="setcenter">
      <el-col :span="10" :xs="24">
        <el-card class="box-card " >
          <div slot="header" class="clearfix">
            <span>企业信息</span>
          </div>
          <div>
            <div class="text-center">
              <el-avatar :size="100" :src=" etp.Logo||avatar" :fit="'cover'"
                         style="margin-right: 10px"></el-avatar>
            </div>
            <div class="setcenter">
              <ul class="list-group list-group-striped">
                <li class="list-group-item">
                  <i class="el-icon-user"></i>用户名
                  <div class="pull-right">{{account}}</div>
                </li>
                <li class="list-group-item">
                  <i class="el-icon-edit-outline"></i>密码
                  <div class="pull-right">{{etp.Password}}</div>
                </li>
                <li class="list-group-item">
                  <i class="el-icon-reading"></i>联系号码
                  <div class="pull-right">{{ etp.PhoneNumber }}</div>
                </li>
                <li class="list-group-item">
                  <i class="el-icon-school"></i>角色
                  <div class="pull-right">企业用户</div>
                </li>
                <li class="list-group-item">
                  <i class="el-icon-edit"></i>简介
                  <div class="pull-right">无</div>
                </li>
              </ul>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>


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
    }
  },
  methods:{

  }
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
