<script setup lang="ts">
type Props = {
  click: any;
};

const props = defineProps<Props>();
</script>

<template>
  <el-button color="#fed04f" round @click="props.click()" class="button-box">
    <slot />
  </el-button>
</template>

<style scoped lang="less">
.button-box {
  width: 100% !important;
  font-size: 16px;
  font-weight: bold;
  color: #ffffff;
  background-image: linear-gradient(
    to right,
    #ffb54d 0%,
    #ffcc33 51%,
    #ffb54d 100%
  );
}
</style>
