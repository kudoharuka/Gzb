<template>
  <div class="register">
    <el-form ref="registerForm" :model="registerForm" :rules="registerRules" class="register-form">
      <h3 class="title">企业注册页面</h3>
      <el-form-item prop="account">
        <el-input
            v-model="registerForm.account"
            type="text"
            auto-complete="off"
            placeholder="企业账号（邮箱）"
            prefix-icon="el-icon-user-solid"
        >
        </el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input
            v-model="registerForm.password"
            type="password"
            auto-complete="off"
            placeholder="密码"
            show-password
            prefix-icon="el-icon-lock"
        >
        </el-input>
      </el-form-item>
      <el-form-item prop="phoneNumber">
        <el-input
            v-model="registerForm.phoneNumber"
            type="tel"
            auto-complete="off"
            placeholder="联系号码"
            prefix-icon="el-icon-phone"
        >
        </el-input>
      </el-form-item>
      <el-form-item prop="name">
        <el-input
            v-model="registerForm.name"
            type="text"
            auto-complete="off"
            placeholder="企业名称"
            prefix-icon="el-icon-office-building"
        >
        </el-input>
      </el-form-item>

      <el-form-item style="width:100%;">
        <el-button
            :loading="loading"
            size="medium"
            type="primary"
            style="width:100%; background-image:linear-gradient(83deg,#4394c5,#1585ff)"
            @click.native.prevent="handleRegister"
            round
        >
          <span v-if="!loading">注册</span>
          <span v-else>注 册 中...</span>
        </el-button>
      </el-form-item>
    </el-form>
    <!--  底部  -->
    <div class="el-register-footer">
      <span>Copyright © 2023-2024 工作帮 版权所有.</span>
    </div>
  </div>
</template>

<script>

export default {
  name: "register",
  data() {
    // 验证手机号的规则
    var checkPhoneNumber = (rule, value, cb) => {
      const regPhoneNumber = /^1(3\d|4[5-9]|5[0-35-9]|6[567]|7[0-8]|8\d|9[0-35-9])\d{8}$/
      if (regPhoneNumber.test(value)) {
        return cb()
      }
      cb(new Error('请输入合法的手机号'))
    };
    // 验证邮箱的规则
    var checkEmail = (rule, value, cb) => {
      const regPhoneNumber = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/
      if (regPhoneNumber.test(value)) {
        return cb()
      }
      cb(new Error('请输入合法的邮箱'))
    };
    return {

      //表单数据绑定对象
      registerForm: {
        account: '',
        password: '',
        phoneNumber: '',
        name: '',
      },
      //表单验证规则
      registerRules: {
        account: [
          {required: true, trigger: "blur", message: "请输入您的账号"},
          {min: 3, max: 20, trigger: "blur", message: "长度在3到20个字符"},
          {validator: checkEmail, trigger: 'blur'}
        ],
        password: [
          {required: true, trigger: "blur", message: "请输入您的密码"},
          {min: 6, max: 20, trigger: "blur", message: "长度在6到20个字符"}
        ],
        phoneNumber: [
          {required: true, trigger: "blur", message: "请输入您的联系号码"},
          {min: 6, max: 20, trigger: "blur", message: "长度在6到20个字符"},
          {validator: checkPhoneNumber, trigger: 'blur'}
        ],
        name: {required: true, trigger: "blur", message: "请输入企业名称"},
      },
      loading: false,

    };
  },
  methods: {
    handleRegister() {
      this.loading = true;
      let res;
      this.$refs.registerForm.validate(async valid => {
        try {
          const response = await this.axios.post('enterprise/register', this.registerForm);
          res = response.data;
          if (res.code === 200) {
            this.$message.success('注册成功');
            this.$router.push({path: '/login'});
          } else {
            this.loading = false;
            this.$message.error(res.message);
          }
        } catch (error) {
          // 处理请求错误
          console.error(error);
          this.loading = false;
          if (res) {
            this.$message.error(res.message);
          } else {
            this.$message.error('注册失败，此用户已存在！');
          }
        }
      });
    }
  }
}

</script>
<style scoped>
.register {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  background-size: 100% 100%;
  background-image: linear-gradient(165deg,#1c92d2,#f2fcfe);

}

.title {
  margin: 0 auto 30px auto;
  text-align: center;
  font-weight: bold;
  color: #1585ff;
}

.register-form {
  border-radius: 15px;
  background: rgba(255, 255, 255, 0.80);
  width: 45%;
  padding: 25px 25px 5px 25px;
  box-shadow: 0 0 20px #3e3e3f;
}

.el-register-footer {
  height: 40px;
  line-height: 40px;
  position: fixed;
  bottom: 0;
  width: 100%;
  text-align: center;
  color: #fff;
  font-family: Arial;
  font-size: 12px;
  letter-spacing: 1px;
}
</style>

