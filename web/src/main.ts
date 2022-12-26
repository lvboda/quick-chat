import { createApp } from "vue";
import Lib from "@/plugins/lib";
import App from "@/App.vue";

import "virtual:svg-icons-register";
import "@/styles/index.less";

const app = createApp(App);
app.use(Lib);
app.mount("#app");
