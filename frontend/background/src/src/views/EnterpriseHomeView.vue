<template>
  <div class="home">
    <TopNavbar></TopNavbar>
  </div>
</template>

<script>
// @ is an alias to /src


import TopNavbar from "@/components/TopNavbar";

export default {
  name: 'enterpriseHomeView',
  components: {
    TopNavbar

  }
}
</script>
<style>
.home{
  height: 100%;
}
</style>