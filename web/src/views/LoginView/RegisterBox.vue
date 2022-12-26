<script setup lang="ts">
import { register, checkStatus } from "@/api/user";
import type { FormInstance, FormRules } from "element-plus";

type Props = {
  cutComponent: () => void;
};

const props = defineProps<Props>();

const registerQuery = reactive({
  userId: "",
  password: "",
  nickName: "",
  userRole: 1,
});

const formRef = $ref<FormInstance>();
const formRules = reactive<FormRules>({
  userId: [
    {
      min: 6,
      max: 12,
      message: "账号格式不正确,请输入6-12位英文或数字",
      trigger: "blur",
    },
    { required: true, message: "请输入账号", trigger: "blur" },
  ],
  nickName: [
    { min: 1, max: 6, message: "请输入1-6位昵称", trigger: "blur" },
    { required: true, message: "请输入昵称", trigger: "blur" },
  ],
  password: [
    { min: 6, max: 12, message: "请输入6-12位密码", trigger: "blur" },
    { required: true, message: "请输入密码", trigger: "blur" },
  ],
});

async function registerFn() {
  formRef.validate(async (valid) => {
    if (!valid) return;

    const { status, message } = await register(registerQuery);
    let type = "error" as "error" | "success";
    let msg = `注册失败: ${message}`;

    if (checkStatus(status)) {
      type = "success";
      msg = "注册成功";
      props.cutComponent();
    }

    ElMessage({ type, message: msg });
  });
}
</script>

<template>
  <h2 class="register-title">注册</h2>
  <el-form
    ref="formRef"
    :model="registerQuery"
    :rules="formRules"
    class="register-form"
  >
    <el-form-item prop="userId">
      <el-input
        v-model="registerQuery.userId"
        placeholder="账号"
        class="register-input"
      />
    </el-form-item>
    <el-form-item prop="nickName">
      <el-input
        v-model="registerQuery.nickName"
        placeholder="昵称"
        class="register-input"
      />
    </el-form-item>
    <el-form-item prop="password">
      <el-input
        type="password"
        v-model="registerQuery.password"
        placeholder="密码"
        show-password
        class="register-input"
      />
    </el-form-item>
    <el-form-item>
      <el-button color="#fed04f" round @click="registerFn" class="register-btn">
        注册
      </el-button>
    </el-form-item>
  </el-form>
</template>

<style scoped lang="less">
.register-title {
  margin-top: -15px;
  font-size: 20px;
  font-weight: 600;
  color: #333;
}
.register-form {
  width: 100%;

  .register-input {
    margin-bottom: 5px;
    font-size: 16px;
    color: #757575;
    font-weight: bold;
  }

  .register-btn {
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
