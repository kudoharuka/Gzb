<template>
  <div class="block">
    <el-carousel style="margin: 0 auto; width: 100%">
      <el-carousel-item v-for="(item, index) in images" :key="index" style="margin: 0 auto">
        <a :href="item.url" target="_blank"><img :src="item.image"></a>
      </el-carousel-item>
    </el-carousel>
    <div class="block1">
      <el-row :gutter="20">
        <el-col :span="16">
          <el-card class="box-card card">
            <div style="font-size: 18px;font-weight: bold"> 企业指南</div>
            <el-card class="box-card" style="margin: 10px">
              <ul class="list-group list-group-striped">
                <li class="list-group-item">
                  <div><a target="_blank" href="http://fjrclh.fzu.edu.cn/cmss/cksyxx">查看生源信息</a></div>
                </li>
                <li class="list-group-item">
                  <div><a target="_blank" href="http://fjrclh.fzu.edu.cn/cmss/tzgg/53826">招聘入校指南</a></div>
                </li>
                <li class="list-group-item">
                  <div><a target="_blank" href="http://fjrclh.fzu.edu.cn/cmss/tzgg/51560.html">校园招聘流程须知 </a></div>
                </li>
              </ul>
            </el-card>

            <div style="font-size: 18px;font-weight: bold"> 下载中心</div>
            <el-card class="box-card" style="margin: 10px">
              <ul >
                <li class="list-group-item">
                  <div><a target="_blank" href="http://fjrclh.fzu.edu.cn/cmss/comp/52961">全国高校毕业生毕业去向登记与网上签约平台操作指南 </a></div>
                </li>
                <li class="list-group-item">
                  <div><a target="_blank" href="http://fjrclh.fzu.edu.cn/cmss/comp/52741">就业管理系统用人单位操作指南 </a></div>
                </li>
                <li class="list-group-item">
                  <div><a target="_blank" href="http://fjrclh.fzu.edu.cn/cmss/comp/56210">2024届用人单位招聘手册 </a></div>
                </li>
                <li class="list-group-item">
                  <div><a target="_blank" href="http://fjrclh.fzu.edu.cn/cmss/comp/52742">用人单位进校材料  </a></div>
                </li>

              </ul>
            </el-card>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card class="box-card card">
            <div style="font-size: 18px;font-weight: bold; margin: 10px"> 招聘日历</div>
            <el-calendar v-model="value">
            </el-calendar>
          </el-card>
        </el-col>
      </el-row>
      <!--      友链-->
      <el-row>
        <el-card class="box-card card">
          <div style="font-size: 18px;font-weight: bold">友链</div>
          <el-card class="box-card" style="margin: 10px">
            <el-row :gutter="20">
              <el-col :span="12"><a href="http://fjrclh.fzu.edu.cn/">福建人才联合网</a></el-col>
              <el-col :span="12"><a href="https://www.fzu.edu.cn/">福州大学</a></el-col>
            </el-row>
          </el-card>
        </el-card>
      </el-row>
    </div>

  </div>

</template>

<script>

export default {
  name: "Carousel",
  data() {
    return {
      value: new Date(),
      images: [
        {image: 'http://fjrclh.fzu.edu.cn/Uploads/image/2024-01-02/6593b62d631b9.png', url: ''},
        {
          image: 'http://fjrclh.fzu.edu.cn/Uploads/image/2023-09-14/6502fa8eaff77.jpg',
          url: 'https://jobone.51job.com/'
        },
        {
          image: 'http://fjrclh.fzu.edu.cn/Uploads/image/2022-07-19/62d617e29df51.png',
          url: 'https://fzu.ncss.cn/student/index.html'
        },
        {
          image: 'http://fjrclh.fzu.edu.cn/Uploads/image/2024-01-03/65951fd38036a.jpg',
          url: 'https://www.fj99.org.cn/bys/'
        }
      ]
    }
  },
  methods: {}
}
</script>

<style scoped>
.block {
  margin: 10px auto;
  background: #ffffff;
  width: 80%;
}

.list-group-item {
  line-height: 50px;
  width: 290px;

}
a{
  color: inherit;
  text-decoration: none;
}
.card {
  margin: 10px;
  background: #f8eee5;
}

.el-carousel__item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.el-carousel__item h3 {
  color: #475669;
  font-size: 14px;
  opacity: 0.75;
  line-height: 150px;
  margin: 0;
}

</style>