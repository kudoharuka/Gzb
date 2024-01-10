<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class=" pd whiteback mg" :title="title"></Breadcrumb>
    <!--    企业查询-->
    <el-card class=" mg">
      <div style="font-size: 20px;font-weight: bold;">企业查询</div>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入内容" v-model="queryInfo.query">
            <el-button slot="append" icon="el-icon-search" @click="getEtpList"></el-button>
          </el-input>
        </el-col>
        <el-col :span=2>
          <el-button @click="addDialogVisible = true" type="primary"
                     round icon="el-icon-plus">添加企业
          </el-button>
        </el-col>
      </el-row>
    </el-card>
    <!--    企业列表-->
    <el-card class=" mg">
      <div style="font-size: 20px;font-weight: bold"> 企业列表</div>
      <Table :table-data="etpList" :columns="columns" >
        <template #operation="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="showDetails(scope.row)">详情
          </el-button>
          <el-button size="mini" type="primary" icon="el-icon-edit" round
                     @click="showEditDialog(scope.row.ID)">编辑
          </el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" round
                     @click="removeEtpById(scope.row.ID)">删除
          </el-button>
        </template>
      </Table>
    </el-card>
    <!--    分页器-->
    <Pagination :total="total" :query-info="queryInfo"
                @page-size-change="handlePageSizeChange"
                @page-change="handlePageChange"></Pagination>
    <!--    添加企业的对话框-->
    <el-dialog title="添加企业" :visible.sync="addDialogVisible" width="30%"
               @close="addDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="addFormRef" :model="addForm" label-width="80px"
               :rules="addFormRules">
        <el-form-item label="企业账号" prop="account">
          <el-input v-model="addForm.account"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="addForm.password"></el-input>
        </el-form-item>
        <el-form-item label="联系号码" prop="phoneNumber">
          <el-input v-model="addForm.phoneNumber"></el-input>
        </el-form-item>
        <el-form-item label="企业名称" prop="name">
          <el-input v-model="addForm.name"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
         <el-button @click="addDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="addEtp">确 定</el-button>
      </span>
    </el-dialog>
    <!--    修改企业对话框-->
    <el-dialog title="修改企业" :visible.sync="editDialogVisible" width="30%"
               @close="editDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="editFormRef" :model="editForm" label-width="80px"
               :rules="editFormRules">
        <el-form-item label="企业账号" prop="Account">
          <el-input v-model="editForm.Account" disabled></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="Password">
          <el-input v-model="editForm.Password"></el-input>
        </el-form-item>
        <el-form-item label="联系号码" prop="PhoneNumber">
          <el-input v-model="editForm.PhoneNumber"></el-input>
        </el-form-item>
        <el-form-item label="企业名称" prop="Name">
          <el-input v-model="editForm.Name"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="editDialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="editEtpInfo">确 定</el-button>
  </span>
    </el-dialog>
    <!--    抽屉-->
    <Drawer  :drawer="detailsDrawer" :title="drawerTitle" @closed="drawerClosed">
      <template #details="scope" >
        <div class="mg setcenter"><el-image :src="editForm.Logo" style="height: 150px"></el-image></div>
        <el-descriptions  direction="vertical" :column="2" border class="mg">
          <el-descriptions-item label="企业ID" > {{editForm.ID}}</el-descriptions-item>
          <el-descriptions-item label="企业账号" > {{editForm.Account}}</el-descriptions-item>
          <el-descriptions-item label="企业名称">{{editForm.Name}}</el-descriptions-item>
          <el-descriptions-item label="联系号码">{{editForm.PhoneNumber}}</el-descriptions-item>
<!--          <el-descriptions-item label="地址">{{editForm.Area}}</el-descriptions-item>-->
<!--          <el-descriptions-item label="性别">{{editForm.Sex}}</el-descriptions-item>-->
<!--          <el-descriptions-item label="目标院校">{{editForm.TargetCollege}}</el-descriptions-item>-->
          <el-descriptions-item label="企业地区">{{editForm.Region}}</el-descriptions-item>
          <el-descriptions-item label="创始人">{{editForm.Founder}}</el-descriptions-item>
<!--          <el-descriptions-item label="专业">{{editForm.Major}}</el-descriptions-item>-->
<!--          <el-descriptions-item label="规模"><el-tag size="small">{{editForm.Year}}</el-tag></el-descriptions-item>-->
          <el-descriptions-item label="企业简介" :span="2">
            <el-input type="textarea" v-model="editForm.Desc" :rows="5" readonly ></el-input>
          </el-descriptions-item>
          <el-descriptions-item label="招聘简章" :span="2">
            <el-input type="textarea" v-model="editForm.Recruitment" :rows="5" readonly ></el-input>
          </el-descriptions-item>
        </el-descriptions>

      </template>
    </Drawer>
  </div>
</template>

<script>
export default {
  name: "EtpView",

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
      title: "企业管理",
      drawerTitle:"企业详情",
      total: 0,
      //企业数据集合
      etpList: [],
      //查询到的企业信息
      editForm: {},
      //添加企业对话框
      addDialogVisible: false,
      //控制详情抽屉可见否
      detailsDrawer: false,
      //控制修改企业对话框的显示与隐藏
      editDialogVisible: false,
      //列表配置
      columns: [
        {prop: 'Account', label: '企业账号', width: '150px'},
        {prop: 'Name', label: '企业名', width: '150px'},
        {prop: 'PhoneNumber', label: '联系号码', width: '150px'},
        {prop: 'Business', label: '企业业务', width: '150px'},
        {prop: 'Scale', label: '规模', width: '100px', sortable: true},
        // {prop: 'College', label: '大学', width: '100px'},
        // {prop: 'Balance', label: '学币', width: '100px', sortable: true},
      ],
      //添加企业的表单数据
      addForm: {
        account: '',
        password: '',
        phoneNumber: '',
        name:''
      },
      // 添加企业的规则
      addFormRules: {
        account: [
          {required: true, message: '请输入企业账号', trigger: 'blur'},
          {min: 3, max: 20, message: '企业账号的长度在3~20个字符间', trigger: 'blur'},
          {validator: checkEmail, trigger: 'blur'}
        ],
        name: [
          {required: true, message: '请输入企业名称', trigger: 'blur'},
          {min: 1, max: 15, message: '企业名称的长度在1~15个字符间', trigger: 'blur'}
        ],
        password: [
          {required: true, message: '请输入密码', trigger: 'blur'},
          {min: 6, max: 20, message: '密码的长度在6~20个字符间', trigger: 'blur'}
        ],
        phoneNumber: [
          {required: true, message: '请输入手机号', trigger: 'blur'},
          {validator: checkPhoneNumber, trigger: 'blur'}
        ]
      },
      //编辑企业的规则
      editFormRules: {
        PhoneNumber:[
          {required: true, message: '请输入手机号', trigger: 'blur'},
          {validator: checkPhoneNumber, trigger: 'blur'}
        ],
        Name: [
          {required: true, message: '请输入企业名称', trigger: 'blur'},
          {min: 1, max: 15, message: '企业名称的长度在1~15个字符间', trigger: 'blur'}
        ],
        Password: [
          {required: true, message: '请输入密码', trigger: 'blur'},
          {min: 6, max: 20, message: '密码的长度在6~20个字符间', trigger: 'blur'}
        ],
      },
      //假数据
      // 获取企业列表的参数对象
      queryInfo: {
        query: '',
        pageNum: 1,
        pageSize: 10
      },
    }
  },
  created() {
    this.getEtpList()
  },
  methods: {
    // 查询方法
    async getEtpList() {
      const {data: res} = await this.axios.post('enterprise/list', {
        query: this.queryInfo.query,
        pageSize: this.queryInfo.pageSize,
        pageNum: this.queryInfo.pageNum
      }, {
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      });
      console.log(res)
      if (res.code !== 200) return this.$message.error('获取企业列表失败！')
      this.etpList = res.data.enterprises
      this.total = res.data.total
      console.log(this.etpList);
    },
    //处理每页显示数量变化
    handlePageSizeChange(newSize) {
      this.queryInfo.pageSize = newSize;
      this.getEtpList()
    },
    //处理页码变化
    handlePageChange(newPage) {
      this.queryInfo.pageNum = newPage;
      this.getEtpList()
    },

    // 显示详情
    showDetails(row) {
      console.log(row);
      this.editForm=row;
      this.detailsDrawer = true;
    },
    //关闭详情
    drawerClosed() {
      this.detailsDrawer = false;
    },

    //添加企业对话框关闭
    addDialogClosed() {
      this.$refs.addFormRef.resetFields()
    },

    //点击按钮添加企业
    addEtp() {
      this.$refs.addFormRef.validate(async valid => {
        if (!valid) return
        console.log(this.addForm)
        // 发起添加企业网络请求
        try{
          const {data: res} = await this.axios.post('enterprise/add', this.addForm, {
            headers: {
              'Authorization': window.sessionStorage.getItem("token")
            }
          })
          if (res.code !== 200) {
            this.$message.error('添加企业失败！')
          } else this.$message.success('添加企业成功！')
          //隐藏对话框
          this.addDialogVisible = false
          //刷新企业列表
          await this.getEtpList()
        }catch (err){
          this.$message.error(err.response.data.message);
        }
      })
    },

    //展示编辑企业的对话框
    async showEditDialog(id) {
      console.log(id);
      const {data: res} = await this.axios.get('enterprise/searchById', {
        params: {'id': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) {
        return this.$message.error('查询企业信息失败！')
      }
      this.editForm = res.data
      this.editDialogVisible = true
    },

    // 修改企业对话框关闭
    editDialogClosed() {
      this.$refs.editFormRef.resetFields()
    },

    // 修改信息并提交
    editEtpInfo() {
      this.$refs.editFormRef.validate(async valid => {
        console.log(valid)
        if (!valid) return
        console.log(this.editForm.Account)
        console.log(this.editForm.PhoneNumber)

        // 发起修改企业信息的数据请求
        const {data: res} = await this.axios.patch('enterprise/update', {
          'account': this.editForm.Account,
          'phoneNumber': this.editForm.PhoneNumber,
          'name': this.editForm.Name,
          'passWord': this.editForm.Password
        },{
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          return this.$message.error('更新企业信息失败！')
        }
        // 关闭对话框
        this.editDialogVisible = false
        // 刷新数据列表
        await this.getEtpList()
        this.$message.success('更新企业信息成功！')
      })
    },

    //删除企业
    async removeEtpById(id) {
      console.log(id)
      const result = await this.$confirm('确定要删除该企业？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).catch(err => err)
      console.log(result)
      if (result !== 'confirm') {
        return this.$message.info('已取消删除')
      }
      console.log(id);
      const {data: res} = await this.axios.delete('enterprise/delete', {
        params:{'id': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) return this.$message.error('删除企业失败！')
      this.$message.success('删除企业成功！')
      await  this.getEtpList()
    },

    //获取焦点事件
    focus(event){
      event.enable(false);
    }
  }
}
</script>

<style scoped>

</style>