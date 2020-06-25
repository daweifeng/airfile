<template>
    <div>
        <div class="url-section">
        <div class="url-bar">
          <input class="url-input" ref="input"  type="text" :value="url" @focus="selectAll()">
          <transition name="fade">
            <div class="copy-alert" v-if="copied">URL Copied</div>
          </transition>
        </div>
      <div class="button" @click="copy()">
        Copy
      </div>
      </div>
      <div class="new-file-section">
        <div class="button" @click="newFile()">
          New File
        </div>
      </div>
    </div>
</template>

<script>
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component
export default class UrlBar extends Vue {
    @Prop() url
    copied = false

    mounted () {
      this.focusInput()
    }

    selectAll () {
      this.$refs.input.select()
    }

    focusInput () {
      this.$refs.input.focus()
    }

    copy () {
      this.focusInput()
      this.$copyText(this.url).then((e) => {
        this.copied = true
        setTimeout(() => { this.copied = false }, 3000)
      }, function (e) {
        alert('Can not copy')
      })
    }

    newFile () {
      this.$emit('newfile')
    }
}
</script>

<style lang="scss" scoped>
    .new-file-section {
      margin: 2rem 0;
      text-align: center;
    }
    .url-section {
        margin: auto;
        text-align: center;
    }
    .url-bar {
        position: relative;
        display: inline-block;
        border: none;
        border-radius: 18px;
        box-shadow:5px 4px 35px rgba(131, 131, 131, 0.25);
        width: 60vw;
        max-width: 1400px;
        min-height: 80px;
        height: 80px;
        padding: 0;
        overflow: hidden;
        font-size: 1.5rem;
        flex-wrap: wrap;
        vertical-align: middle;
        .url-input {
            display: block;
            text-align: center;
            width: 100%;
            height: 100%;
            margin: auto;
            border: none;
            background: #e5e5e5;
            font-size: 1.5rem;
            padding: 0 2rem;
            outline: none;
        }

        .copy-alert {
          position: absolute;
          top: 0;
          left: 0;
          background: #3BB273;
          color: #fff;
          text-align: center;
          line-height: 80px;
          width: 100%;
          height: 100%;
        }
    }

    .button {
            display: inline-block;
            text-align: center;
            width: 200px;
            height: 60px;
            line-height: 60px;
            font-size: 1.5rem;
            padding: 0.5rem 0;
            margin-left: 1rem;
            color: #fff;
            background: #00A6FB;
            border: none;
            border-radius: 18px;
            transition: 0.3s all linear;
            vertical-align: middle;
            &:hover {
                cursor: pointer;
                background: #0582ca;
            }
    }

    .fade-enter-active, .fade-leave-active {
      transition: opacity .3s;
    }
    .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
      opacity: 0;
    }
</style>
