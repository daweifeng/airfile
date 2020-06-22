<template>
    <div class="url-section">
        <div class="url-bar">
        <input class="url-input" ref="input"  type="text" :value="url" @focus="selectAll()">
    </div>
    <div class="copy-button" @click="copy()">
            Copy
        </div>
    </div>
</template>

<script>
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component
export default class UrlBar extends Vue {
    @Prop() url

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
      this.$copyText(this.url).then(function (e) {
        alert('Copied')
      }, function (e) {
        alert('Can not copy')
      })
    }
}
</script>

<style lang="scss" scoped>
    .url-section {
        margin: auto;
        text-align: center;
    }
    .url-bar {
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
            width: 100%;
            height: 100%;
            margin: auto;
            border: none;
            background: #e5e5e5;
            font-size: 1.5rem;
            padding: 0 2rem;
            outline: none;
        }
    }

    .copy-button {
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
</style>
