<template>
  <div class="login">
    <el-form ref="loginForm" :model="loginForm" :rules="loginRules" class="login-form">
      <h3 class="title">工作帮管理系统</h3>
      <el-form-item prop="account">
        <el-input
            v-model="loginForm.account"
            type="text"
            auto-complete="off"
            placeholder="账号"
            prefix-icon="el-icon-user-solid"
        >
        </el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input
            v-model="loginForm.password"
            type="password"
            auto-complete="off"
            placeholder="密码"
            show-password
            prefix-icon="el-icon-lock"
        >
        </el-input>
      </el-form-item>
      <el-row :gutter="20" style="margin-bottom: 10px">
        <!--        <el-radio-group v-model="loginForm.role" class="radio-group">-->
        <!--          <el-radio label="企业" class="radio"></el-radio>-->
        <!--          <el-radio label="管理员" class="radio"></el-radio>-->
        <!--        </el-radio-group>-->
      </el-row>
      <el-row :gutter="20" style="margin-bottom: 10px" class="radio-group">
        <el-col :span="6" class="radio">
          <el-button type="text">
            <router-link :to="{ path: '/register' }" target="_blank">企业注册</router-link>
          </el-button>
        </el-col>
        <el-col :span="6" class="radio">
          <el-button type="text" @click="dialogFormVisible = true">忘记密码？</el-button>
        </el-col>
      </el-row>

      <!--      忘记密码对话框-->
      <el-dialog title="忘记密码" :visible.sync="dialogFormVisible">
        <el-form :model="form" ref="formRef" :rules="formRules">
          <el-form-item label="企业账号" :label-width="formLabelWidth" prop="account">
            <el-input v-model="form.account" autocomplete="off"></el-input>
          </el-form-item>
          <el-row>
            <el-col :span="15">
              <el-form-item label="验证码" :label-width="formLabelWidth" prop="code">
                <el-input v-model="form.code" autocomplete="off"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8" offset="1">
              <el-button type="primary" v-count-down="60" @click="getCode">发送邮箱验证</el-button>
            </el-col>
          </el-row>
          <el-form-item label="新的密码" :label-width="formLabelWidth" prop="password">
            <el-input v-model="form.password" autocomplete="off"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="DialogClosed">取 消</el-button>
          <el-button type="primary" @click="resetPwd">重置密码</el-button>
        </div>
      </el-dialog>


      <el-form-item style="width:100%;">
        <el-row>
          <el-button
              :loading="loading"
              size="medium"
              type="primary"
              style="width:100%; background-image:linear-gradient(83deg,#4394c5,#1585ff)"
              @click.native.prevent="handleLogin"
              round
          >
            <span v-if="!loading">管理员登录</span>
            <span v-else>登 录 中...</span>
          </el-button>
        </el-row>

        <el-row style="margin-top: 10px">
          <el-button
              :loading="loading"
              size="medium"
              type="primary"
              style="width:100%; background-image:linear-gradient(83deg,#f1e4c9,#1585ff)"
              @click.native.prevent="handleLogin1"
              round
          >
            <span v-if="!loading">企业 登录</span>
            <span v-else>登 录 中...</span>
          </el-button>
        </el-row>
      </el-form-item>
    </el-form>
    <!--  底部  -->
    <div class="el-login-footer">
      <span>Copyright © 2023-2024 工作帮 版权所有.</span>
    </div>
  </div>
</template>

<script>

export default {
  name: "Login",
  // 自定义指令
  directives: {
    'count-down': {
      bind(el, binding) {
        let timer = null
        let count = binding.value

        function updateCount() {
          el.textContent = `剩余 ${count} 秒`
          count--
          if (count < 0) {
            clearInterval(timer)
            el.removeAttribute('disabled')
            el.textContent = '发送验证码'
          }
        }

        el.addEventListener('click', () => {
          if (timer) {
            return
          }
          el.setAttribute('disabled', 'disabled')
          updateCount()
          timer = setInterval(updateCount, 1000)
        })
      }
    }
  },

  data() {
    // 验证邮箱的规则
    var checkEmail = (rule, value, cb) => {
      const regPhoneNumber = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/
      if (regPhoneNumber.test(value)) {
        return cb()
      }
      cb(new Error('请输入合法的邮箱'))
    };
    return {
      dialogFormVisible: false,
      formLabelWidth: '120px',
      // 忘记密码表单
      form: {
        account: '',
        password: '',
        code: '',
      },
      //表单数据绑定对象
      loginForm: {
        account: '',
        password: '',
        // role: '企业',
      },

      //忘记密码验证规则
      formRules: {
        account: [
          {required: true, trigger: "blur", message: "请输入您的账号"},
          {min: 3, max: 20, trigger: "blur", message: "长度在3到20个字符"},
          {validator: checkEmail, trigger: 'blur'}
        ],
        password: [
          {required: true, trigger: "blur", message: "请输入您的密码"},
          {min: 6, max: 20, trigger: "blur", message: "长度在6到20个字符"}
        ],
        code: {required: true, trigger: "blur", message: "请输入您的验证码"},
      },

      //表单验证规则
      loginRules: {
        account: [
          {required: true, trigger: "blur", message: "请输入您的账号"},
          {min: 3, max: 20, trigger: "blur", message: "长度在3到20个字符"}
        ],
        password: [
          {required: true, trigger: "blur", message: "请输入您的密码"},
          {min: 6, max: 20, trigger: "blur", message: "长度在6到20个字符"}
        ],
      },
      loading: false,

    };
  },
  methods: {
    resetPwd() {
      let mes;
      this.$refs.formRef.validate(async valid => {
        console.log(valid)
        if (!valid) return
        // 发起修改用户信息的数据请求
        try {
          const {data: res} = await this.axios.post('/resetPwd', {
            'account': this.form.account,
            'password': this.form.password,
            'code': this.form.code
          })
          mes = res.message;
          if (res.code == 200) {
            this.$message.success('重置密码成功！')
            // 关闭对话框
            this.DialogClosed()
          } else {
            console.log(res.message)
            this.$message.error(res.message);
          }
        } catch (error) {
          console.error(error);
          if (mes) {
            this.$message.error(mes.message);
          } else {
            this.$message.error('重置失败,验证码错误或已失效！');
          }
        }
      })

    },
    async getCode() {
      if (!this.form.account) this.$message({
        message: "请先输入邮箱",
        type: "error"
      })
      else {
        // 获取验证码
        const {data: res} = await this.axios.post('/getCode', {
          'account': this.form.account
        }).then(() => {
          this.$message({
            message: "已发送验证码，5分钟内有效！",
            type: "success"
          });
        }).catch(() => {
          this.$message({
            message: "请求超时，请检查网络连接",
            type: "error"
          });
        });
        console.log(res)
        if (res.code !== 200) return this.$message.error('获取验证码失败！')
      }
    },

    DialogClosed() {
      this.dialogFormVisible = false,
          this.$refs.formRef.resetFields()
    },

    handleLogin() {
      this.loading = true;
      let res;
      this.$refs.loginForm.validate(async valid => {
        try {
          const response = await this.axios.post('/admin/login', this.loginForm);
          res = response.data;
          if (res.code === 200) {
            this.$message.success('登录成功');
            console.log(res)
            window.sessionStorage.setItem('token', res.token);
            window.sessionStorage.setItem('account', res.admin.Account);
            window.sessionStorage.setItem('id', res.admin.ID);
            window.sessionStorage.setItem('admin', JSON.stringify(res.admin));
            this.$router.push({path: '/adminHome'});
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
            this.$message.error('登录失败！');
          }
        }
      });
    },
    handleLogin1() {
      this.loading = true;
      let res;
      this.$refs.loginForm.validate(async valid => {
        try {
          const response = await this.axios.post('/enterprise/login', this.loginForm);
          res = response.data;
          if (res.code === 200) {
            this.$message.success('登录成功');
            console.log(res)
            window.sessionStorage.setItem('token', res.token);
            window.sessionStorage.setItem('account', res.enterprise.Account);
            window.sessionStorage.setItem('id', res.enterprise.ID);
            window.sessionStorage.setItem('enterprise', JSON.stringify(res.enterprise));
            this.$router.push({path: '/enterpriseHome'});
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
            this.$message.error('登录失败，请检查角色、账号密码是否匹配！');
          }
        }
      });
    }
  }
}

</script>
<style scoped>
.radio-group {
  display: flex;
  justify-content: center;
}

.radio {
  margin: 0 10px;
}

.login {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  background-size: 100% 100%;
  background-image: url("../assets/images/login-background.jpg");

}

.title {
  margin: 0 auto 30px auto;
  text-align: center;
  font-weight: bold;
  color: #1585ff;
}

.login-form {
  border-radius: 15px;
  background: rgba(255, 255, 255, 0.80);
  width: 400px;
  padding: 25px 25px 5px 25px;
  box-shadow: 0 0 20px #3e3e3f;
}

.el-login-footer {
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

