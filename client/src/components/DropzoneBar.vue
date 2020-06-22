<template>
  <div>
    <VueDropzone
      ref="dropzone"
      id="dropzone"
      class="dropzone"
      :useCustomSlot=true
      :options="dropzoneOptions"
      :include-styling="false"
      @vdropzone-success="dropzoneOnSuccess"
      @vdropzone-upload-progress="dropzoneUploadProgress"
      @vdropzone-canceled="fileUploadCanceled"
    >
      <div class="hint">
        Drag and drop your file here
      </div>
    </VueDropzone>
    <div class="button-section">
      <button class="get-button" v-on:click="onSubmit" v-bind:disabled="uploading">
        {{ uploading ? "Uploading" : "Get URL"}}
      </button>
    </div>
  </div>
</template>

<script>
import { Component, Vue } from 'vue-property-decorator'
import VueDropzone from 'vue2-dropzone'

@Component({
  components: {
    VueDropzone
  }
})
export default class DropzoneBar extends Vue {
  data () {
    return {
      dropzoneOptions: {
        url: 'https://httpbin.org/post',
        previewTemplate: this.template(),
        thumbnailWidth: 200,
        addRemoveLinks: true,
        autoProcessQueue: false
      },
      progress: 0,
      uploading: false
    }
  }

  onSubmit () {
    this.$refs.dropzone.processQueue()
    this.uploading = true
  }

  dropzoneOnSuccess (file, response) {
    this.$emit('success', response)
  }

  dropzoneUploadProgress (file, progress, bytesSent) {
    this.progress = progress
    document.querySelector('.progress').style.width = progress + '%'
    document.querySelector('.dz-upload').textContent = Math.round(progress) + '%'
  }

  fileUploadCanceled (file) {
    this.uploading = false
    this.progress = 0
  }

  template () {
    return `
      <div class="dz-file">
        <div class="status">
        <span class="file-name" data-dz-name></span>
        <span class="dz-upload" data-dz-uploadprogress></span>
        </div>
        <div class="progress"></div>
      </div>
    `
  }
}

</script>

<style lang="scss">
  #dropzone {
    position: relative;
    border: none;
    border-radius: 18px;
    box-shadow:5px 4px 35px rgba(131, 131, 131, 0.25);
    width: 60vw;
    min-height: 80px;
    height: 80px;
    padding: 0;
    overflow: hidden;
    margin: auto;
    font-size: 1.5rem;

    .hint {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      margin: 0;
      color: #E5E5E5;
    }
  }
  .dz-file {
      text-align: center;
      height: 100%;
      width: 100%;
      background: #C0E9FE;
      position: relative;

      .status {
        position: absolute;
        z-index: 10;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
      }

      .progress {
        position: absolute;
        top: 0;
        left: 0;
        background: #00A6FB;
        width: 0%;
        height: 100%;
        transition: 0.3s all linear;
      }

      .dz-remove {

      }
  }
  .button-section {
    margin: 2rem;
    text-align: center;
  }
  .get-button {
    cursor: pointer;
    width: 200px;
    padding: 0.5rem 0;
    font-size: 1.5rem;
    color: #fff;
    background: #00A6FB;
    border: none;
    border-radius: 18px;
    transition: 0.3s all linear;

    &:hover {
      background: #0582ca;
    }
  }

  // .get-button-disabled {
  // }

  #dropzone .dz-preview {
    display: block;
    width: 100%;
  }

</style>
