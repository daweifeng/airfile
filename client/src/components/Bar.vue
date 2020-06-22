<template>
  <div v-if="uploaded">
    <UrlBar :url="url" />
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
  url: string;
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
      uploaded: true,
      url: 'test.com'
    }
  }

  onSuccess (response: Response) {
    console.log(response);
    (this as any).uploaded = true
  }
}
</script>
