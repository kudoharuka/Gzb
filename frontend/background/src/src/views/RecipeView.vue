<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class=" pd whiteback mg" :title="title"></Breadcrumb>
    <!--    企业查询-->
    <el-card class=" mg">
      <div style="font-size: 20px;font-weight: bold;"> 企业查询</div>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入企业"  v-model="queryInfo.query">
            <el-button slot="append" icon="el-icon-search" @click="getRecipeList"></el-button>
          </el-input>
        </el-col>
        <el-col :span="5">
        </el-col>
        <el-col :span=2>
          <el-button @click="addDialogVisible = true" type="primary"
                     round icon="el-icon-plus">添加企业
          </el-button>
        </el-col>
      </el-row>
    </el-card>
    <!--    动态列表-->
    <el-card class=" mg">
      <div style="font-size: 20px;font-weight: bold"> 企业列表</div>
      <Table :table-data="recipeList" :columns="columns" >
        <!--        操作区-->
        <template #operation="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="showDetails(scope.row)">详情
          </el-button>
          <el-button size="mini" type="primary" icon="el-icon-edit" round
                     @click="showEditDialog(scope.row.ID)">编辑
          </el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" round
                     @click="removeRecipeById(scope.row.ID)">删除
          </el-button>
        </template>
      </Table>
    </el-card>
    <!--    分页器-->
    <Pagination :total="total" :query-info="queryInfo"
                @page-size-change="handlePageSizeChange"
                @page-change="handlePageChange"></Pagination>
    <!--    添加企业的对话框-->
    <el-dialog title="添加企业" :visible.sync="addDialogVisible" width="50%"
               @close="addDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="addFormRef" :model="addForm" label-width="80px"
               :rules="addFormRules">
        <el-form-item label="作者" prop="author">
          <el-input v-model="addForm.author"></el-input>
        </el-form-item>
        <el-form-item label="标题" prop="title">
          <el-input v-model="addForm.title"></el-input>
        </el-form-item>
        <el-form-item label="院校代码" prop="code">
          <el-input v-model="addForm.code"></el-input>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <el-input type="textarea" :rows="5" placeholder="请输入内容" v-model="addForm.content">
          </el-input>
<!--          <quill-editor v-model="addForm.content" :options="editorOption"></quill-editor>-->
<!--          <Editor :value.sync="addForm.content"></Editor>-->
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
         <el-button @click="addDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="addRecipe">确 定</el-button>
      </span>
    </el-dialog>
    <!--    修改企业对话框-->
    <el-dialog title="修改企业" :visible.sync="editDialogVisible" width="50%"
               @close="editDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="editFormRef" :model="editForm" label-width="80px"
               :rules="editFormRules">
        <el-form-item label="作者" prop="Author">
          <el-input v-model="editForm.Author" disabled></el-input>
        </el-form-item>
        <el-form-item label="标题" prop="Title">
          <el-input v-model="editForm.Title"></el-input>
        </el-form-item>
        <el-form-item label="院校代码" prop="Code">
          <el-input v-model="editForm.Code"></el-input>
        </el-form-item>
        <el-form-item label="内容" prop="Content">
          <el-input type="textarea" :rows="5" placeholder="请输入内容" v-model="editForm.Content">
          </el-input>
<!--          <quill-editor v-model="editForm.Content" :options="editorOption"></quill-editor>-->
<!--          <Editor :value.sync="editForm.Content" ref="editor"></Editor>-->
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="editDialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="editRecipeInfo">确 定</el-button>
  </span>
    </el-dialog>
    <!--    抽屉-->
    <Drawer :drawer="detailsDrawer" :title="drawerTitle" @closed="drawerClosed">
      <template #details="scope">
        <div class="mg setcenter"><el-image :src="editForm.PageUrl" style="width: 70%"></el-image></div>
        <el-descriptions direction="vertical" :column="2" border class="mg">
          <el-descriptions-item label="作者"> {{ editForm.Author }}</el-descriptions-item>
          <el-descriptions-item label="标题" > {{ editForm.Title }}</el-descriptions-item>
          <el-descriptions-item label="院校代码" > {{ editForm.Code }}</el-descriptions-item>
          <el-descriptions-item label="收藏数" > {{ editForm.Favorite }}</el-descriptions-item>
          <el-descriptions-item label="发布时间"> {{ formattedPublishTime }}</el-descriptions-item>
        </el-descriptions>
        <div class="mg">内容</div>
        <quill-editor v-model="editForm.Content" @focus="focus($event)" class="mg"></quill-editor>
      </template>
    </Drawer>
  </div>
</template>

<script>

import {QuillWatch} from "quill-image-extend-module";

export default {
  name: "RecipeView",
  computed:{
    formattedPublishTime() {
      const publishTime = this.editForm.PublishTime;
      return this.$moment(publishTime).format('YYYY-MM-DD HH:mm:ss');
    },
  },
  data(){
    return{
      title: '企业管理',
      drawerTitle: "企业详情",
      //添加对话框
      addDialogVisible: false,
      //控制详情抽屉可见否
      detailsDrawer: false,
      //控制修改对话框的显示与隐藏
      editDialogVisible: false,
      //数据总数
      total: 0,
      //配置项
      // editorOption: {
      //   modules: {
      //     ImageExtend: {
      //       loading: true,
      //       name: 'img',//图片参数名
      //       size: 5, // 可选参数 图片大小，单位为M，1M = 1024kb
      //       action: 'https://api.superbed.cn/upload',//上传的服务器地址，如果action为空，则采用base64插入图片
      //       // response 为一个函数，用来获取服务器返回的具体图片地址
      //       response: res => {
      //         console.log(res)
      //         return res.url
      //       },
      //       success: () => {
      //         console.log('ImageExtend中的success：上传成功')
      //       }, // 可选参数  上传成功触发的事件
      //       change: (xhr, formData) => {
      //         // xhr.setRequestHeader('myHeader','myValue')
      //         formData.append('token', '1000766339bd4c248a8ad625a87f687d')
      //       } // 可选参数 每次选择图片触发，也可用来设置头部，但比headers多了一个参数，可设置formData
      //     },
      //     // 如果不上传图片到服务器，此处不必配置
      //     toolbar: {
      //       // container为工具栏，此次引入了全部工具栏，也可自行配置
      //       container: [
      //         ["bold", "italic", "underline", "strike"],
      //         ["blockquote", "code-block"],
      //         [{ header: 1 }, { header: 2 }],
      //         [{ list: "ordered" }, { list: "bullet" }],
      //         [{ script: "sub" }, { script: "super" }],
      //         [{ indent: "-1" }, { indent: "+1" }],
      //         [{ direction: "rtl" }],
      //         [{ size: ["small", false, "large", "huge"] }],
      //         [{ header: [1, 2, 3, 4, 5, 6, false] }],
      //         [{ color: [] }, { background: [] }],
      //         [{ font: [] }],
      //         [{ align: [] }],
      //         ["image"]
      //       ],
      //       // 上传成功，回显图片（会进入如上面ImageExtend的各过程，返回<img src="http://xx.xx.xx.xx:xxxx/file/xxx.jpg">）
      //       handlers: {
      //         image: function() {
      //           // 劫持原来的图片点击按钮事件
      //           QuillWatch.emit(this.quill.id)
      //         }
      //       }
      //     }
      //   }
      // },
      //获取列表的参数对象
      queryInfo: {
        query: '',
        pageNum: 1,
        pageSize: 10
      },
      //查询到的信息
      editForm: {

      },
      //添加的表单数据
      addForm: {
        author: '',
        title: '',
        code:'',
        content: '',
      },
      // 添加规则
      addFormRules: {
        author: [
          {required: true, message: '请输入作者名', trigger: 'blur'},
          {min: 2, max: 10, message: '作者名的长度在2~10个字符间', trigger: 'blur'}
        ],
        title: {required: true, message: '标题不能为空', trigger: 'blur'},
        content: {required: true, message: '内容不能为空', trigger: 'blur'},
        code:{required: true, message: '院校代码不能为空', trigger: 'blur'},
      },
      // 修改规则
      editFormRules: {
        Title: {required: true, message: '标题不能为空', trigger: 'blur'},
        Content:{required: true, message: '内容不能为空', trigger: 'blur'},
        Code:{required: true, message: '院校代码不能为空', trigger: 'blur'},
      },
      //企业数据集合
      recipeList:[],
      columns: [
        {prop: 'Author', label: '作者', width: '150px'},
        {prop: 'PublishTime', label: '发布时间', width: '200px', sortable: true,
          formatter: (row, column) => {
            const publishTime = row[column.property];
            return this.$moment(publishTime).format('YYYY-MM-DD HH:mm:ss');
          }},
        {prop: 'Title', label: '标题', width: '180px', showOverflowTooltip: true},
        {prop: 'Favorite', label: '收藏数', width: '150px', sortable: true},
      ],
    }
  },
  created() {
    this.getRecipeList()
  },
  methods:{

    // 查询方法
    async getRecipeList() {
      const {data: res} = await this.axios.post('recipe/list', {
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
      this.recipeList = res.data.recipes
      this.total = res.data.total
      console.log(this.recipeList);
    },
    //处理每页显示数量变化
    handlePageSizeChange(newSize) {
      this.queryInfo.pageSize = newSize;
      this.getRecipeList()
    },
    //处理页码变化
    handlePageChange(newPage) {
      this.queryInfo.pageNum = newPage;
      this.getRecipeList()
    },
    //删除企业
    async removeRecipeById(id) {
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
      const {data: res} = await this.axios.delete('recipe/delete', {
        params: {'id': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) return this.$message.error('删除企业失败！')
      this.$message.success('删除企业成功！')
      await this.getRecipeList()
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
    //添加对话框关闭
    addDialogClosed() {
      this.$refs.addFormRef.resetFields()
    },
    //点击按钮添加
    addRecipe() {
      this.$refs.addFormRef.validate(async valid => {
        if (!valid) return
        console.log(this.addForm)
        // 发起添加网络请求
        const {data: res} = await this.axios.post('recipe/add', this.addForm, {
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          this.$message.error('添加企业失败！')
        } else this.$message.success('添加企业成功！')
        //隐藏对话框
        this.addDialogVisible = false
        //刷新用户列表
        await this.getRecipeList()
      })
    },
    //展示编辑的对话框
    async showEditDialog(id) {
      console.log(id);
      const {data: res} = await this.axios.get('recipe/searchById', {
        params: {'id':id},
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
    // 修改对话框关闭
    editDialogClosed() {
      this.$refs.editFormRef.resetFields()

    },
    // 修改信息并提交
    editRecipeInfo() {
      this.$refs.editFormRef.validate(async valid => {
        console.log(valid)
        if (!valid) return
        console.log(this.editForm)
        // 发起修改信息的数据请求
        const {data: res} = await this.axios.patch('recipe/update', {
          'id':this.editForm.ID,
          'author': this.editForm.Author,
          'title': this.editForm.Title,
          'content': this.editForm.Content,
          'code':this.editForm.Code
        },{
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          return this.$message.error('更新企业失败！')
        }
        // 关闭对话框
        this.editDialogVisible = false
        // 刷新数据列表
        await this.getRecipeList()
        this.$message.success('更新企业成功！')
      })
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