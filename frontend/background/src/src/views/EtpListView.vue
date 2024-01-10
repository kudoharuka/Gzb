<template>
  <div>

    <!--    企业查询-->
    <el-card class=" mg">
      <div style="font-size: 20px;font-weight: bold;">企业查询</div>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入企业名" v-model="queryInfo.query">
            <el-button slot="append" icon="el-icon-search" @click="getEtpList"></el-button>
          </el-input>
        </el-col>
      </el-row>
    </el-card>
    <!--    企业列表-->
    <el-card class=" mg">
      <div style="font-size: 20px;font-weight: bold"> 企业列表</div>
      <Table :table-data="etpList" :columns="columns">
        <template #operation="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="showDetails(scope.row)">详情
          </el-button>
        </template>
      </Table>
    </el-card>

    <!--    分页器-->
    <Pagination :total="total" :query-info="queryInfo"
                @page-size-change="handlePageSizeChange"
                @page-change="handlePageChange"></Pagination>
    <!--    抽屉-->
    <Drawer :drawer="detailsDrawer" :title="drawerTitle" @closed="drawerClosed">
      <template #details="scope">
        <div class="mg setcenter">
          <el-image :src="editForm.Logo" style="height: 150px"></el-image>
        </div>
        <el-descriptions direction="vertical" :column="2" border class="mg">
          <el-descriptions-item label="企业ID"> {{ editForm.ID }}</el-descriptions-item>
          <el-descriptions-item label="企业名称"> {{ editForm.Name }}</el-descriptions-item>
          <el-descriptions-item label="联系方式">{{ editForm.PhoneNumber }}</el-descriptions-item>
          <el-descriptions-item label="创始人">{{ editForm.Founder }}</el-descriptions-item>
          <el-descriptions-item label="创立时间">{{ editForm.FoundDate }}</el-descriptions-item>
          <el-descriptions-item label="总部地址">{{ editForm.Region }}</el-descriptions-item>
          <el-descriptions-item label="企业类型">{{ editForm.Type }}</el-descriptions-item>
          <el-descriptions-item label="所属行业">{{ editForm.Business }}</el-descriptions-item>
          <el-descriptions-item label="公司简介" :span="2">
            <el-input type="textarea" v-model="editForm.Desc" :rows="5" readonly></el-input>
          </el-descriptions-item>
          <el-descriptions-item label="招聘简章" :span="2">
            <el-input type="textarea" v-model="editForm.Recruitment" :rows="5" readonly></el-input>
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
    return {
      title: "企业一览",
      drawerTitle: "企业详情",
      total: 0,
      //企业数据集合
      etpList: [],
      //查询到的企业信息
      editForm: {},
      //控制详情抽屉可见否
      detailsDrawer: false,

      //列表配置
      columns: [
        {prop: 'Account', label: '企业账号', width: '150px'},
        {prop: 'Name', label: '企业名', width: '150px'},
        {prop: 'PhoneNumber', label: '联系方式', width: '150px'},
        {prop: 'Business', label: '企业业务', width: '150px'},
        {prop: 'Scale', label: '规模', width: '100px', sortable: true},
        // {prop: 'College', label: '大学', width: '100px'},
        // {prop: 'Balance', label: '学币', width: '100px', sortable: true},
      ],

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
      this.editForm = row;
      this.detailsDrawer = true;
    },
    //关闭详情
    drawerClosed() {
      this.detailsDrawer = false;
    },
    //获取焦点事件
    focus(event) {
      event.enable(false);
    }
  }
}
</script>

<style scoped>

</style>