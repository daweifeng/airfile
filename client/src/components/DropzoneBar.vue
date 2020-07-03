<template>
  <div>
    <div class="dropzone-section">
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
        @vdropzone-error="dropzoneUploadError"
      >
        <div class="hint">
          Drag and drop your file here
        </div>
      </VueDropzone>
      <transition name="fade">
        <div class="failed" v-if="failed">
          Upload Failed
        </div>
      </transition>
    </div>
    <div class="hint">
      Max file size: 100MB
    </div>
    <div class="button-section">
      <button class="get-button" v-on:click="onSubmit" v-bind:disabled="uploading">
        {{ buttonTextContent() }}
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
        url: `${process.env.VUE_APP_BACKEND_UPLOAD_API}`,
        previewTemplate: this.template(),
        thumbnailWidth: 200,
        addRemoveLinks: true,
        autoProcessQueue: false,
        maxFilesize: 100 // MB
      },
      progress: 0,
      uploading: false,
      failed: false,
      fileOversized: false
    }
  }

  onSubmit () {
    if (this.$refs.dropzone.getQueuedFiles().length === 0 && !this.failed) {
      return
    }
    if (this.failed) {
      this.failed = false
    }
    this.$refs.dropzone.processQueue()
    this.uploading = true
    this.failed = false
  }

  dropzoneOnSuccess (file, response) {
    this.$emit('success', response)
  }

  dropzoneUploadProgress (file, progress, bytesSent) {
    this.progress = progress
    document.querySelector('.progress').style.width = this.progress + '%'
    document.querySelector('.dz-upload').textContent = Math.round(this.progress) + '%'
  }

  fileUploadCanceled (file) {
    this.uploading = false
    this.progress = 0
  }

  dropzoneUploadError (file, message, xhr) {
    this.uploading = false
    this.progress = 0
    this.failed = true

    if (file.size > 100000000) {
      this.fileOversized = true
      this.$refs.dropzone.removeAllFiles()
      this.failed = false
    } else {
      file.status = 'queued'
    }
    console.log(message)
  }

  buttonTextContent () {
    if (this.failed) {
      return 'Retry'
    }
    if (this.uploading) {
      return 'Uploading'
    } else {
      return 'Get URL'
    }
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
  .dropzone-section {
    position: relative;

    .failed {
      z-index: 11;
      position: absolute;
      top: 0;
      left: 50%;
      transform: translateX(-50%);
      background: #f94144;
      color: rgb(252, 236, 236);
      line-height: 80px;
      text-align: center;
      border-radius: 18px;
      width: 60vw;
      max-width: 1400px;
      min-height: 80px;
      height: 80px;
    }
  }
  #dropzone {
    position: relative;
    border: none;
    border-radius: 18px;
    box-shadow:5px 4px 35px rgba(131, 131, 131, 0.25);
    width: 60vw;
    max-width: 1400px;
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
    margin: 1rem;
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

  .fade-enter-active, .fade-leave-active {
      transition: opacity .3s;
    }
    .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
      opacity: 0;
  }
  .hint {
    text-align: center;
    padding: 1rem;
  }
</style>
