<script setup lang="ts">
import { login, checkStatus } from "@/api/user";
import { getFromLocal, setToLocal, removeFromLocal } from "@/utils/storage";

import type { FormInstance, FormRules } from "element-plus";

const router = useRouter();

const savedLoginInfo = getFromLocal<{ userId: string; password: string }>(
  "savedLoginInfo"
);

const loginQuery = reactive(
  savedLoginInfo ? savedLoginInfo : { userId: "", password: "" }
);

const isRemember = $ref(!!savedLoginInfo);

const formRef = $ref<FormInstance>();
const formRules = reactive<FormRules>({
  userId: [
    { min: 6, max: 12, message: "请输入6-12位英文或数字账号", trigger: "blur" },
    { required: true, message: "请输入账号", trigger: "blur" },
  ],
  password: [
    { min: 6, max: 12, message: "请输入6-12位密码", trigger: "blur" },
    { required: true, message: "请输入密码", trigger: "blur" },
  ],
});

async function signIn() {
  formRef.validate(async (valid) => {
    if (!valid) return;

    const { status, message, data } = await login(loginQuery);
    let type = "error" as "error" | "success";
    let msg = `登陆失败: ${message}`;

    if (checkStatus(status)) {
      type = "success";
      msg = "登陆成功";
      setToLocal("userInfo", data);

      if (isRemember) setToLocal("savedLoginInfo", loginQuery);
      if (!isRemember) removeFromLocal("savedLoginInfo");

      router.push({ path: "/home" });
    }
    ElMessage({ type, message: msg });
  });
}
</script>

<template>
  <el-form
    ref="formRef"
    :model="loginQuery"
    :rules="formRules"
    class="login-form"
  >
    <el-form-item prop="userId">
      <el-input
        v-model="loginQuery.userId"
        placeholder="账号"
        class="login-input"
      />
    </el-form-item>
    <el-form-item prop="password" style="margin-bottom: 10px">
      <el-input
        type="password"
        v-model="loginQuery.password"
        placeholder="密码"
        show-password
        class="login-input"
      />
    </el-form-item>
    <el-form-item label-position="right" style="margin-bottom: 10px">
      <el-checkbox v-model="isRemember" label="记住账号密码" size="small" />
    </el-form-item>
    <el-form-item>
      <el-button color="#fed04f" round @click="signIn()" class="login-btn">
        登录
      </el-button>
    </el-form-item>
  </el-form>
</template>

<style scoped lang="less">
.login-form {
  width: 100%;

  .login-input {
    margin-bottom: 5px;
    font-size: 16px;
    color: #757575;
    font-weight: bold;
  }

  .login-btn {
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
}
</style>
