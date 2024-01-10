<template>
  <div class="home">
    <AsideNavbar></AsideNavbar>
  </div>
</template>

<script>
// @ is an alias to /src


import AsideNavbar from "@/components/AsideNavbar";

export default {
  name: 'HomeView',
  components: {
    AsideNavbar

  }
}
</script>
<style>
.home{
  height: 100%;
}
</style>