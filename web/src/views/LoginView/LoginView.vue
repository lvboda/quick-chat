<script setup lang="ts">
import LoginBox from "./LoginBox.vue";
import RegisterBox from "./RegisterBox.vue";

let showLogin = $ref(true);

function cutComponent() {
  showLogin = !showLogin;
}
</script>

<template>
  <div class="login-box">
    <div class="login-warp">
      <div class="login-warp__logo">
        <svg-icon icon-name="chat-logo" className="icon" />
        <span>QuickChat</span>
      </div>
      <div class="login-warp__content">
        <login-box v-if="showLogin" />
        <register-box v-else :cut-component="cutComponent" />
        <el-button type="primary" link @click="cutComponent">
          {{ showLogin ? "快速注册" : "返回登陆" }}
        </el-button>
      </div>
    </div>
  </div>
</template>

<style scoped lang="less">
.login-box {
  width: 100%;
  height: 100%;
  background: #f5f5f5;
  display: 1 1;
}

.login-warp {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  margin: auto;
  width: 472px;
  height: 500px;
  border-radius: 12px 12px 12px 12px;
  align-items: center;
  background: #fff;
  overflow: hidden;
  box-shadow: 0 0 6px 0 rgba(0, 0, 0, 0.2);
}

.login-warp__logo {
  width: 100%;
  height: 30%;
  font-size: 20px;
  color: #ffffff;
  box-shadow: 0 2px 7px 0 #211b4f;
  background-color: #201b4f;
  background-image: linear-gradient(180deg, #404040 0, #1e1e1e 100%),
    linear-gradient(126deg, #201b4f 15%, #0b0930 100%);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  > .icon {
    width: 60px;
    height: 80px;
  }

  > span {
    margin-bottom: 10px;
  }
}

.login-warp__content {
  width: 60%;
  height: 70%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
</style>
