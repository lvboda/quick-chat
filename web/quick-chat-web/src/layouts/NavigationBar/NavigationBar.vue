<script setup lang="ts">
import { routes } from "@/router/index";
import useUserStore from "@/stores/user";

import HeadPortrait from "@/components/HeadPortrait/HeadPortrait.vue";

const router = useRouter();
const { userInfo } = useUserStore();

const barRoutes = $ref(
  routes.find((item) => item.name === "home")?.children || []
);

watch(
  () => router.currentRoute.value.path,
  (newPath) => {
    barRoutes.forEach((item) => {
      if (item.path === newPath) {
        item.name = item.name + "-f";
        return;
      }
      item.name = item.name.replace("-f", "");
    });
  },
  { immediate: true }
);

function jumpPage(path: string) {
  router.push({ path });
}
</script>
<template>
  <div class="navigation-bar-box">
    <head-portrait
      class="navigation-bar-box__face"
      :src="userInfo.face"
      :show-details="false"
    />
    <svg-icon
      class-name="navigation-bar-box__bar"
      :icon-name="route.name"
      :key="route.name"
      v-for="route in barRoutes"
      @click="jumpPage(route.path)"
    />
  </div>
</template>

<style scoped lang="less">
.navigation-bar-box {
  width: 7%;
  height: 100%;
  background: #56565a;
  background-image: linear-gradient(180deg, #404040 0, #1e1e1e 100%),
    linear-gradient(126deg, #201b4f 15%, #0b0930 100%);
  display: flex;
  flex-direction: column;
  flex-wrap: nowrap;
  align-items: center;
}

.navigation-bar-box__face {
  margin-top: 25px;
}

.navigation-bar-box__bar {
  width: 35px;
  height: 35px;
  margin-top: 25px;
  cursor: pointer;
}
</style>
