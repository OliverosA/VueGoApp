<template>
  <div class="mx-3 mt-10 max-h-4/5 w-1/2">
    <h2
      class="text-7xl w-full text-center text-white"
      v-if="showEmail.id == ''"
    >
      Waiting for a selected email...
    </h2>
    <div v-else class="max-w-full h-full bg-gray-50 rounded">
      <div class="mx-3 text-start flex flex-col space-y-4 overflow-y-auto h-44">
        <h2 class="font-bold mt-3">Subject: {{ showEmail?.subject }}</h2>
        <hr />
        <h2 class="font-bold">From: {{ showEmail.from }}</h2>
        <hr />
        <h2 class="font-bold">To: {{ showEmail.to }}</h2>
      </div>
      <hr />
      <div
        class="mx-3 pt-5 text-ellipsis text-justify leading-7 overflow-y-auto max-h-96"
      >
        {{ showEmail.content }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
export default defineComponent({
  props: {
    selectedEmail: Object,
  },
  setup(props) {
    const showEmail = ref(props.selectedEmail);

    return { showEmail };
  },
});
</script>

<style scoped>
hr {
  background-color: #fff;
  border-top: 2px dashed #8c8b8b;
}

.LoadingMessage {
  font-size: 6vw;
  text-align: center;
  width: 100%;
}

.contentTitle {
  font-size: 2vw;
  text-align: center;
  width: 100%;
}
</style>
