<template>
  <div v-if="uploaded">
    <UrlBar :url="url" @newfile="onNewFile()" />
  </div>
  <div v-else>
    <DropzoneBar
      @success="onSuccess"
    />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import DropzoneBar from '@/components/DropzoneBar.vue'
import UrlBar from '@/components/UrlBar.vue'

interface Response {
  id: string;
}

@Component({
  components: {
    DropzoneBar,
    UrlBar
  }
})
export default class Bar extends Vue {
  data () {
    return {
      uploaded: false,
      url: null
    }
  }

  onSuccess (response: Response) {
    (this as any).url = `${process.env.VUE_APP_BACKEND_DOWNLOAD_API}?key=${response.id}`;
    (this as any).uploaded = true
  }

  onNewFile () {
    (this as any).uploaded = false
  }
}
</script>
