<template>
  <div class="workbench">
    <el-tabs :tab-position="tabPosition">
      <!--      我的企业-->
      <el-tab-pane label="我的企业">
        <el-row>我的企业信息</el-row>
        <el-divider></el-divider>
        <el-form ref="formRef" :model="form" label-width="80px" style="width: 60% ;margin: 0 60px" :rules="formRules">
          <el-form-item label="企业名称" prop="Name">
            <el-input v-model="form.Name" maxlength="15" placeholder="请输入企业名称" show-word-limit></el-input>
          </el-form-item>
          <el-form-item label="创始人" prop="Founder">
            <el-input v-model="form.Founder" maxlength="10" placeholder="请输入创始人" show-word-limit></el-input>
          </el-form-item>
          <el-form-item label="联系方式" prop="PhoneNumber">
            <el-input v-model="form.PhoneNumber" maxlength="20" placeholder="请输入联系方式（号码）" show-word-limit></el-input>
          </el-form-item>
          <el-form-item label="创立时间" required>
            <el-col :span="11">
              <el-form-item prop="FoundDate">
                <el-date-picker type="date" placeholder="选择日期" v-model="form.FoundDate"
                                style="width: 100%;" value-format="yyyy-MM-dd"></el-date-picker>
              </el-form-item>
            </el-col>
          </el-form-item>
          <el-form-item label="总部地点" prop="Region">
            <el-select v-model="form.Region" placeholder="请选择">
              <el-option
                  v-for="item in regions"
                  :key="item"
                  :label="item"
                  :value="item">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="所属行业" prop="Business">
            <el-select v-model="form.Business" placeholder="请选择">
              <el-option
                  v-for="item in businesses"
                  :key="item"
                  :label="item"
                  :value="item">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="公司logo" prop="Logo">
            <el-upload
                class="avatar-uploader"
                action="https://api.superbed.cn/upload"
                :data="{token:'1000766339bd4c248a8ad625a87f687d'}"
                :show-file-list="false"
                :on-success="handleAvatarSuccess"
            >
              <img v-if="form.Logo" :src="form.Logo" class="avatar">
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
            </el-upload>

          </el-form-item>
          <el-form-item label="企业类型" prop="Type">
            <el-radio-group v-model="form.Type">
              <el-radio label="国有企业"></el-radio>
              <el-radio label="私有企业"></el-radio>
              <el-radio label="个体工商户"></el-radio>
              <el-radio label="有限责任公司"></el-radio>
              <el-radio label="股份有限公司"></el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="公司简介" prop="Desc">
            <el-input type="textarea" v-model="form.Desc" maxlength="500"
                      show-word-limit rows="5"></el-input>
          </el-form-item>
          <el-form-item label="招聘简章" prop="Recruitment">
            <el-input type="textarea" v-model="form.Recruitment" maxlength="1000"
                      show-word-limit rows="8"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="editInfo">保存修改</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <!--      添加岗位对话框-->
      <el-dialog title="添加岗位" :visible.sync="addDialogVisible" width="50%"
                 @close="addDialogClosed">
        <el-form ref="jobFormRef" :model="jobForm" label-width="80px" style="width: 80% ;margin: 0 60px"
                 :rules="editFormRules">
          <el-form-item label="岗位名称" prop="Name">
            <el-input v-model="jobForm.Name" maxlength="15" placeholder="请输入岗位名称"
                      show-word-limit></el-input>
          </el-form-item>
          <el-form-item label="薪资" prop="Wage">
            <el-select v-model="jobForm.Wage" placeholder="请选择">
              <el-option
                  v-for="item in wages"
                  :key="item"
                  :label="item"
                  :value="item">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="职位分类" prop="Category">
            <el-select v-model="jobForm.Category" placeholder="请选择">
              <el-option
                  v-for="item in categories"
                  :key="item"
                  :label="item"
                  :value="item">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="工作性质" prop="Type">
            <el-radio-group v-model="jobForm.Type">
              <el-radio label="全职"></el-radio>
              <el-radio label="实习"></el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="福利待遇" prop="Benefit">
            <el-input v-model="jobForm.Benefit" maxlength="50" placeholder="请输入福利待遇"
                      show-word-limit></el-input>
          </el-form-item>
          <el-form-item label="职位描述" prop="Introduction">
            <el-input type="textarea" v-model="jobForm.Introduction" maxlength="500"
                      show-word-limit rows="5"></el-input>
          </el-form-item>
          <el-form-item label="任职要求" prop="Requirement">
            <el-input type="textarea" v-model="jobForm.Requirement" maxlength="1000"
                      show-word-limit rows="8"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="addJob">添加岗位</el-button>
            <el-button @click="addDialogVisible = false">取 消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
      <!--      岗位管理-->
      <el-tab-pane label="岗位管理">
        <el-row>岗位信息</el-row>
        <el-divider></el-divider>
        <el-row :gutter="20">
          <!--    岗位列表-->
          <el-card class=" mg">
            <el-row :gutter="5">
              <el-col :span="2" style="font-size: 20px;font-weight: bold;">岗位列表</el-col>
              <el-col :span="22" style="margin-bottom: 20px">
                <el-button @click="addDialogVisible = true" type="primary" size="medium"
                           round icon="el-icon-plus">添加岗位
                </el-button>
              </el-col>
            </el-row>
            <Table :table-data="jobList" :columns="columns">
              <template #operation="scope">
                <el-button size="mini" type="success" icon="el-icon-view" round
                           @click="showDetails(scope.row)">详情
                </el-button>
                <el-button size="mini" type="primary" icon="el-icon-edit" round
                           @click="showEditDialog(scope.row.ID)">编辑
                </el-button>
                <el-button size="mini" type="danger" icon="el-icon-delete" round
                           @click="removeJobById(scope.row.ID)">删除
                </el-button>
              </template>
            </Table>
            <!--    分页器-->
            <Pagination :total="total" :query-info="queryInfo"
                        @page-size-change="handlePageSizeChange"
                        @page-change="handlePageChange"></Pagination>
          </el-card>
        </el-row>
      </el-tab-pane>
    </el-tabs>
    <!--    修改岗位对话框-->
    <el-dialog title="修改岗位" :visible.sync="editDialogVisible" width="50%"
               @close="editDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="editFormRef" :model="editForm" label-width="80px" style="width: 80% ;margin: 0 60px"
               :rules="editFormRules">
        <el-form-item label="岗位名称" prop="Name">
          <el-input v-model="editForm.Name" disabled></el-input>
        </el-form-item>
        <el-form-item label="薪资" prop="Wage">
          <el-select v-model="editForm.Wage">
            <el-option
                v-for="item in wages"
                :key="item"
                :label="item"
                :value="item">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="职位分类" prop="Category">
          <el-select v-model="editForm.Category">
            <el-option
                v-for="item in categories"
                :key="item"
                :label="item"
                :value="item">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作性质" prop="Type">
          <el-radio-group v-model="editForm.Type">
            <el-radio label="全职"></el-radio>
            <el-radio label="实习"></el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="福利待遇" prop="Benefit">
          <el-input v-model="editForm.Benefit" maxlength="50"
                    show-word-limit></el-input>
        </el-form-item>
        <el-form-item label="职位描述" prop="Introduction">
          <el-input type="textarea" v-model="editForm.Introduction" maxlength="500"
                    show-word-limit rows="5"></el-input>
        </el-form-item>
        <el-form-item label="任职要求" prop="Requirement">
          <el-input type="textarea" v-model="editForm.Requirement" maxlength="1000"
                    show-word-limit rows="8"></el-input>
        </el-form-item>
        <!--        <el-form-item label="岗位名称" prop="Position">-->
        <!--          <el-input v-model="editForm.Position" disabled></el-input>-->
        <!--        </el-form-item>-->
        <!--        <el-form-item label="密码" prop="Password">-->
        <!--          <el-input v-model="editForm.Password"></el-input>-->
        <!--        </el-form-item>-->
        <!--        <el-form-item label="手机号" prop="PhoneNumber">-->
        <!--          <el-input v-model="editForm.PhoneNumber"></el-input>-->
        <!--        </el-form-item>-->
        <!--        <el-form-item label="昵称" prop="NickName">-->
        <!--          <el-input v-model="editForm.NickName"></el-input>-->
        <!--        </el-form-item>-->
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="editDialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="editJobInfo">确 定</el-button>
  </span>
    </el-dialog>
    <!--    抽屉-->
    <Drawer :drawer="detailsDrawer" :title="drawerTitle" @closed="drawerClosed">
      <template #details="scope">
        <el-descriptions direction="vertical" :column="2" border class="mg">
          <el-descriptions-item label="岗位名称">{{ editForm.Name }}</el-descriptions-item>
          <el-descriptions-item label="薪资">{{ editForm.Wage }}</el-descriptions-item>
          <el-descriptions-item label="职位分类">{{ editForm.Category }}</el-descriptions-item>
          <el-descriptions-item label="工作性质">{{ editForm.Type }}</el-descriptions-item>
          <el-descriptions-item label="所属公司">{{ editForm.Enterprise.Name }}</el-descriptions-item>
          <el-descriptions-item label="福利待遇">
            <el-tag size="small">{{ editForm.Benefit }}</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="职位描述" :span="2">
            <el-input type="textarea" v-model="editForm.Introduction" :rows="5" readonly></el-input>
          </el-descriptions-item>
          <el-descriptions-item label="任职要求" :span="2">
            <el-input type="textarea" v-model="editForm.Requirement" :rows="5" readonly></el-input>
          </el-descriptions-item>
        </el-descriptions>

      </template>
    </Drawer>
  </div>

</template>

<script>
export default {
  name: "EtpWorkbenchView",
  data() {
    return {
      title: "工作台",
      drawerTitle: "岗位详情",
      tabPosition: 'left',
      total: 0,
      // 企业信息
      form: JSON.parse(window.sessionStorage.getItem("enterprise")),
      //岗位数据集合
      jobList: [],
      //查询到的岗位信息
      editForm: {
        Enterprise: {},
      },
      // 获取岗位列表的参数对象
      queryInfo: {
        query: window.sessionStorage.getItem('id'),
        pageNum: 1,
        pageSize: 10
      },
      //控制详情抽屉可见否
      detailsDrawer: false,
      //控制修改岗位对话框的显示与隐藏
      editDialogVisible: false,
      //添加对话框
      addDialogVisible: false,
      // imageUrl: '',

      //列表配置
      columns: [
        {prop: 'Name', label: '岗位名称', width: '150px'},
        {prop: 'Category', label: '职位分类', width: '150px'},
        {prop: 'Wage', label: '薪资', width: '150px', sortable: true},
        {prop: 'Type', label: '工作性质', width: '100px'},
      ],

      // 地区
      regions: [
        '北京', '福建', '天津', '上海', '重庆',
        '内蒙古', '广西', '西藏', '宁夏', '新疆', '山西', '辽宁', '吉林',
        '黑龙江', '江苏', '浙江', '安徽', '江西', '山东', '河北', '河南', '湖北',
        '湖南', '广东', '海南', '四川', '贵州', '云南', '陕西', '甘肃', '青海', '台湾'
      ],
      // 行业
      businesses: ['信息传输、计算机服务和软件业', '制造业', '建筑业', '金融、保险业', '教育', '其它'],
      // 薪资
      wages: ['0k~5k', '5k~10k', '10k~15k', '15k~20k', '20k以上'],
      //职位分类
      categories: ['IT/互联网/软件', '通讯/电子', '艺术/设计', '建筑/土木', '金融/财务', '教育/培训'],
      // 岗位信息
      jobForm: {
        // name: '',
        Category: 'IT/互联网/软件',
        Type: '全职',
        EnterpriseID: window.sessionStorage.getItem('id'),
        // firm: '',
        Wage: '0k~5k',
        // requirement: '',
        // introduction: '',
        // benefit: '',
      },
      // 企业信息
      // form: this.etpInfo,
      // form: {
      //   // name: '',
      //   // founder: '',
      //   // region: '北京',
      //   // phoneNumber :'',
      //   // foundDate: '',
      //   // logo: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
      //   // etpType: "国有企业",
      //   // business: '信息传输、计算机服务和软件业',
      //   // desc: '',
      //   // recruitment: '',
      // },
      // 企业修改规则
      formRules: {
        Name: {required: true, message: '企业名称不能为空', trigger: 'blur'},
        Founder: {required: true, message: '创始人不能为空', trigger: 'blur'},
        Desc: {required: true, message: '公司简介不能为空', trigger: 'blur'},
        Recruitment: {required: true, message: '招聘简章不能为空', trigger: 'blur'},
        PhoneNumber: {required: true, message: '联系方式不能为空', trigger: 'blur'},
      },
      // 岗位新增规则
      jobFormRules: {
        name: {required: true, message: '岗位名称不能为空', trigger: 'blur'},
        requirement: {required: true, message: '任职要求不能为空', trigger: 'blur'},
        jobDesc: {required: true, message: '职位描述不能为空', trigger: 'blur'},
      },
      //修改岗位的规则
      editFormRules: {
        Name: {required: true, message: '岗位名称不能为空', trigger: 'blur'},
        Requirement: {required: true, message: '任职要求不能为空', trigger: 'blur'},
        Introduction: {required: true, message: '职位描述不能为空', trigger: 'blur'},
      },
    }
  },

  created() {
    this.getJobList()

    // this.showEtp(this.form.ID)
  },

  methods: {
    // 查询企业方法

    // async getEtpList() {
    //   const {data: res} = await this.axios.post('enterprise/list', {
    //     query: this.queryInfo.query,
    //     pageSize: this.queryInfo.pageSize,
    //     pageNum: this.queryInfo.pageNum
    //   }, {
    //     headers: {
    //       'Authorization': window.sessionStorage.getItem("token")
    //     }
    //   });
    //   console.log(res)
    //   if (res.code !== 200) return this.$message.error('获取企业列表失败！')
    //   this.etpList = res.data.enterprises
    //   this.total = res.data.total
    //   console.log(this.etpList);
    // },

    //展示编辑企业的对话框
    // async showEtp(id) {
    //   console.log('根据id查找企业信息'+id);
    //   const {data: res} = await this.axios.get('enterprise/searchById', {
    //     params: {'id': id},
    //     headers: {
    //       'Authorization': window.sessionStorage.getItem("token")
    //     }
    //   })
    //   if (res.code !== 200) {
    //     return this.$message.error('查询企业信息失败！')
    //   }
    //   this.form = res.data
    // },


    // 修改企业信息并提交
    editInfo() {
      this.$refs.formRef.validate(async valid => {
        console.log(valid)
        if (!valid) return
        console.log(this.form)
        // 发起修改企业信息的数据请求
        const {data: res} = await this.axios.patch('enterprise/update', {
          'account': this.form.Account,
          'name': this.form.Name,
          'founder': this.form.Founder,
          'region': this.form.Region,
          'phoneNumber': this.form.PhoneNumber,
          'foundDate': this.form.FoundDate,
          'logo': this.form.Logo,
          'type': this.form.Type,
          'business': this.form.Business,
          'desc': this.form.Desc,
          'recruitment': this.form.Recruitment,
        }, {
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          return this.$message.error('更新企业信息失败！')
        }
        // 刷新数据列表
        this.$message.success('更新企业信息成功！')
        sessionStorage.setItem('enterprise', JSON.stringify(this.form));
        // 触发 user-updated 事件
        this.$emit('etp-updated');
        // await this.showEtp()
      })
    },

    // 查询岗位方法
    async getJobList() {
      const {data: res} = await this.axios.post('job/list', {
        query: this.queryInfo.query,
        pageSize: this.queryInfo.pageSize,
        pageNum: this.queryInfo.pageNum
      }, {
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      });
      console.log(res)
      if (res.code !== 200) return this.$message.error('获取岗位列表失败！')
      this.jobList = res.data.jobs
      this.total = res.data.total
      console.log(this.jobList);
    },

    //处理每页显示数量变化
    handlePageSizeChange(newSize) {
      this.queryInfo.pageSize = newSize;
      this.getJobList()
    },

    //处理页码变化
    handlePageChange(newPage) {
      this.queryInfo.pageNum = newPage;
      this.getJobList()
    },

    // 显示详情
    showDetails(row) {
      console.log(row);
      this.editForm = row;
      this.detailsDrawer = true;
    },

    //关闭详情
    drawerClosed() {
      this.detailsDrawer = false;
    },

    //添加对话框关闭
    addDialogClosed() {
      this.$refs.jobFormRef.resetFields()
    },

    //点击按钮添加岗位
    addJob() {
      this.$refs.jobFormRef.validate(async valid => {
        if (!valid) return
        console.log(this.jobForm)
        // 发起添加岗位网络请求
        const {data: res} = await this.axios.post('job/add', this.jobForm
          , {
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          this.$message.error('添加岗位失败！')
        } else this.$message.success('添加岗位成功！')
        //隐藏对话框
        this.addDialogVisible = false
        //刷新岗位列表
        await this.getJobList()
      })
    },

    //展示编辑岗位的对话框
    async showEditDialog(id) {
      console.log(id);
      const {data: res} = await this.axios.get('job/searchById', {
        params: {'id': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) {
        return this.$message.error('查询岗位信息失败！')
      }
      this.editForm = res.data
      console.log(this.editForm)
      this.editDialogVisible = true
    },

    // 修改岗位对话框关闭
    editDialogClosed() {
      this.$refs.editFormRef.resetFields()
    },

    // 修改岗位信息并提交
    editJobInfo() {
      this.$refs.editFormRef.validate(async valid => {
        console.log(valid)
        if (!valid) return
        console.log(this.editForm)
        // 发起修改岗位信息的数据请求
        const {data: res} = await this.axios.patch('job/update', {
          'id': this.editForm.ID,
          'name': this.editForm.Name,
          'category': this.editForm.Category,
          'type': this.editForm.Type,
          'enterpriseID': window.sessionStorage.getItem('id'),
          'wage': this.editForm.Wage,
          'requirement': this.editForm.Requirement,
          'introduction': this.editForm.Introduction,
          'benefit': this.editForm.Benefit
        }, {
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          return this.$message.error('更新岗位信息失败！')
        }
        // 关闭对话框
        this.editDialogVisible = false
        // 刷新数据列表
        await this.getJobList()
        this.$message.success('更新岗位信息成功！')
      })
    },

    //删除岗位
    async removeJobById(id) {
      const result = await this.$confirm('确定要删除该岗位？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).catch(err => err)
      console.log(result)
      if (result !== 'confirm') {
        return this.$message.info('已取消删除')
      }
      console.log(id);
      const {data: res} = await this.axios.delete('job/delete', {
        params: {'id': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) return this.$message.error('删除岗位失败！')
      this.$message.success('删除岗位成功！')
      await this.getJobList()
    },

    // 上传图片
    handleAvatarSuccess(res, file) {
      // this.form.logo = URL.createObjectURL(file.raw);
      this.form.Logo = res.url;
      console.log(this.form.Logo);
    },

    //获取焦点事件
    focus(event) {
      event.enable(false);
    }
  },

}
</script>

<style scoped>
.workbench {
  background: #ffffff;
  margin: 20px 0px;
  padding: 20px 0px;
}

/deep/ .avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

/deep/ .avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}

/deep/ .avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}

.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>