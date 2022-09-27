import type { App } from "vue";

import { createPinia } from "pinia";
import router from "@/router";

import SvgIcon from "@/components/SvgIcon/SvgIcon.vue";

export default {
  install(app: App) {
    app.use(createPinia());
    app.use(router);
    // svg组件
    app.component("svg-icon", SvgIcon);
  },
};
