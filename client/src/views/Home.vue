<template>
  <div class="home">
    <div class="logo-section">
      <img alt="Airfile logo" src="../assets/logo.png">
      <h1>Airfile</h1>
    </div>
    <div class="file-upload-section">
      <Bar />
    </div>
    <div class="features-section">
      <h1 style="text-align:center">
      Features
      </h1>
      <div class="features">
        <div class="item">
          <div class="icon">
            <img src="../assets/unique-link.png" />
          </div>
          <div class="title">Get a unique link per file</div>
        </div>
        <div class="item">
          <div class="icon">
            <img src="../assets/download.png" />
          </div>
          <div class="title">Download from anywhere</div>
        </div>
        <div class="item">
          <div class="icon">
            <img src="../assets/remove.png" />
          </div>
          <div class="title">File auto-destoried after 10 minutes </div>
        </div>
      </div>
    </div>
    <div class="footer">
      <a href="https://github.com/daweifeng/airfile">
        <div class="item">
          <div class="icon">
            <img src="../assets/git.png" />
          </div>
          <div class="title">Open source on Github</div>
        </div>
      </a>

      <p>Disclaimer: Please do not upload any classified documents! This project is for personal daily use. All files are not encrypted on cloud.</p>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
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

.features-section {
  margin-top: 2rem;
  padding: 1rem 0 3rem 0;
  background: #00A6FB;
  h1 {
    color: #fff;
  }
}

.features {
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
}

.item {
  display: inline-block;
  border-radius: 18px;
  box-shadow:5px 4px 35px rgba(131, 131, 131, 0.25);
  min-height: 30px;
  padding: 1rem 2rem;
  background: #fff;
  margin: 1rem;

  .icon {
    display: inline-block;
    vertical-align: middle;
    height: 30px;

    img {
      height: 100%;
    }
  }

  .title {
    display: inline-block;
    line-height: 30px;
    vertical-align: middle;
    font-size: 1rem;
    padding-left: 1rem;
  }
}

.footer {
  padding: 2rem;
  text-align: center;
}
</style>
