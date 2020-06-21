<template>
  <div class="home">
    <div class="logo-section">
      <img alt="Airfile logo" src="../assets/logo.png">
      <h1>Airfile</h1>
    </div>
    <div class="file-upload-section">
      <Bar />
    </div>

    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import Bar from '@/components/Bar'

export default {
  name: 'Home',
  components: {
    Bar
  }
}
</script>

<style lang="scss" scoped>
.logo-section {
  text-align: center;
  img {
    display: inline-block;
    height: 80px;
    vertical-align: middle;
    margin-right: 12px;
  }

  h1 {
    display: inline-block;
    vertical-align: middle;
    font-size: 60px;
  }

}
</style>
