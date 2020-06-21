<template>
  <div v-if="uploaded">
    uploaded
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

interface Response {
  url: string;
}

@Component({
  components: {
    DropzoneBar
  }
})
export default class Bar extends Vue {
  data () {
    return {
      uploaded: false,
      displayCopiedBar: false,
      url: null
    }
  }

  onSuccess (response: Response) {
    console.log(response)
  }
}
</script>
