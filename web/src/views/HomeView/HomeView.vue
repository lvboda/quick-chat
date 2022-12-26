<script setup lang="ts">
import useChat from "@/hooks/use-common";

import NavigationBar from "@/layouts/NavigationBar/NavigationBar.vue";

useChat();
</script>
<template>
  <div class="home-box">
    <navigation-bar />
    <router-view v-slot="{ Component }">
      <keep-alive>
        <component :is="Component" />
      </keep-alive>
    </router-view>
  </div>
</template>

<style scoped lang="less">
.home-box {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
}
</style>
