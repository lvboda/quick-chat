<script setup lang="ts">
import useUserStore from "@/stores/user";
import useContactStore from "@/stores/contact";
import { addFriend, checkStatus } from "@/api/contact";

import HeadPortrait from "@/components/HeadPortrait/HeadPortrait.vue";
import ButtonBox from "@/components/ButtonBox/ButtonBox.vue";

import type { Relation } from "@/api/contact";

type Props = {
  data: Relation | null;
  close: () => void;
};

const props = defineProps<Props>();

const router = useRouter();

const { userInfo } = storeToRefs(useUserStore());
const { flushContacts } = useContactStore();

async function handleClick() {
  if (!props.data) return;
  const { relationType, friendInfo } = props.data;

  if (relationType === 1) {
    const { status } = await addFriend(
      userInfo.value.userId,
      friendInfo.userId
    );
    if (checkStatus(status)) {
      ElMessage.success("好友添加成功");
      flushContacts();
      props.close();
    } else ElMessage.error("好友添加失败");
    return;
  }

  router.push({ path: "/home/chat", query: { friendId: friendInfo.userId } });
}
</script>

<template>
  <svg-icon v-if="!props.data" class="logo" icon-name="chat-logo"></svg-icon>
  <div v-else class="detail-box">
    <div class="detail-box__top">
      <head-portrait
        class="face"
        :src="props.data.friendInfo.face"
        :show-details="false"
      />
      <div class="content">
        <span>
          {{ props.data.friendInfo.nickName }}
          <svg-icon :icon-name="'gender-' + props.data.friendInfo.gender" />
        </span>
      </div>
    </div>
    <div class="detail-box__middle">
      <div>
        用户id: <span>{{ props.data.friendInfo.userId }}</span>
      </div>
      <div>
        电话号码: <span>{{ props.data.friendInfo.mobile }}</span>
      </div>
      <div>
        个性签名: <span>{{ props.data.friendInfo.signature }}</span>
      </div>
      <div v-if="props.data.relationType === 1">
        验证信息:
        <span>用户{{ props.data.friendInfo.nickName }}请求添加您为好友</span>
      </div>
    </div>
    <div class="detail-box__bottom">
      <div class="but">
        <button-box :click="handleClick">
          {{ props.data.relationType === 1 ? "通过好友验证" : "发送消息" }}
        </button-box>
      </div>
    </div>
  </div>
</template>

<style scoped lang="less">
.logo {
  width: 100px;
  height: 100px;
}

.detail-box {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  flex-direction: column;
  flex-wrap: nowrap;
  align-items: center;
}

.detail-box__top {
  width: 50%;
  height: 20%;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  border-bottom: 1px solid #d8d8d8;

  .face {
    width: 100px;
    height: 100px;
  }

  .content {
    margin-left: 20px;
    font-size: 25px;
    display: flex;
    flex-direction: column;
  }
}

.detail-box__middle {
  width: 50%;
  height: 40%;
  display: flex;
  flex-direction: column;
  flex-wrap: nowrap;
  align-items: flex-start;
  border-bottom: 1px solid #d8d8d8;

  div {
    margin-top: 25px;
    font-size: 18px;
  }

  span {
    font-size: 14px;
    font-weight: 100;
    color: #a7a7a7;
  }
}

.detail-box__bottom {
  width: 50%;
  height: 20%;
  display: flex;
  align-items: center;
  justify-content: center;

  .but {
    width: 50%;
  }
}
</style>
