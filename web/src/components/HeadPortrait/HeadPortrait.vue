<script setup lang="ts">
type Props = {
  src: string;
  showDetails: boolean;
};

const props = defineProps<Props>();
const defaultFaceURL = `${window.location.origin}${
  import.meta.env.BASE_URL
}apis/assets/fulltime/default-face.png`;
</script>

<template>
  <div class="head-portrait-box">
    <img :src="props.src ? props.src : defaultFaceURL" />
    <slot name="point" />
  </div>
</template>

<style scoped lang="less">
.head-portrait-box {
  position: relative;
  width: 50px;
  height: 50px;
  border-radius: 10%;
  img {
    width: 100%;
    height: 100%;
    border-radius: 10%;
  }
}
</style>
