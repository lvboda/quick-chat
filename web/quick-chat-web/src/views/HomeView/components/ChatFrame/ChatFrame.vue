<script setup lang="ts">
import useUserStore from "@/stores/user";

import HeadPortrait from "@/components/HeadPortrait/HeadPortrait.vue";

import type { Relation } from "@/api/contact";
import type { Message } from "@/utils/websocket";

type Props = {
  currentInfo: Relation;
  currentChatList: Message[];
  // eslint-disable-next-line no-unused-vars
  sendMessage: (rid: string, content: string, contentType?: number) => void;
  // eslint-disable-next-line no-unused-vars
  onFocus: (fid: string) => void;
};

const props = defineProps<Props>();

const chatFrameRef = $ref<Element | null>(null);
function scrollToBottom() {
  if (!chatFrameRef) return;
  nextTick(() => chatFrameRef.scrollTo(0, chatFrameRef.scrollHeight));
  showNewMessage = false;
}

let showNewMessage = $ref(false);
watch(
  () => props.currentChatList.length,
  (newLen, oldLen) => {
    if (
      newLen > oldLen &&
      chatFrameRef &&
      chatFrameRef.scrollHeight > chatFrameRef.clientHeight &&
      chatFrameRef.scrollTop + chatFrameRef.clientHeight + 20 <
        chatFrameRef.scrollHeight
    ) {
      showNewMessage = true;
    }
  }
);

let content = $ref("");
const { userInfo } = storeToRefs(useUserStore());

function send() {
  if (!content.replaceAll("\n", "") || !props.currentInfo.userId) {
    ElMessage.warning("请输入消息内容");
    return;
  }

  props.sendMessage(props.currentInfo.userId, content.replaceAll("\n", ""));
  content = "";
  scrollToBottom();
}

function sendVideo() {
  props.sendMessage(props.currentInfo.userId, "", 2);
}

onMounted(() => {
  if (!chatFrameRef) return;

  chatFrameRef.addEventListener("scroll", () => {
    if (
      chatFrameRef.scrollTop + chatFrameRef.clientHeight + 30 >=
      chatFrameRef.scrollHeight
    ) {
      showNewMessage = false;
    }
  });
});

onUpdated(scrollToBottom);
</script>

<template>
  <svg-icon
    v-if="!props.currentInfo"
    class="logo"
    icon-name="chat-logo"
  ></svg-icon>
  <div v-else class="chat-frame">
    <div class="chat-frame__top">
      <span>{{ props.currentInfo.friendInfo.nickName }}</span>
    </div>
    <div ref="chatFrameRef" class="chat-frame__middle">
      <div class="item-chat" :key="chat.id" v-for="chat in currentChatList">
        <div
          v-if="chat.senderId === props.currentInfo.friendInfo.userId"
          class="item-chat--left"
        >
          <head-portrait
            :src="props.currentInfo.friendInfo.face"
            :showDetails="true"
          />
          <div class="content">
            <span>{{ chat.content }}</span>
          </div>
        </div>
        <div v-else class="item-chat--right">
          <div class="content">{{ chat.content }}</div>
          <head-portrait :src="userInfo.face" :showDetails="true" />
        </div>
      </div>
    </div>
    <div class="chat-frame__bottom">
      <div v-if="showNewMessage" class="arrows" @click="scrollToBottom">
        ↓ 新的消息
      </div>
      <div class="info-bar">
        <div @click="sendVideo">视频</div>
      </div>
      <div class="input">
        <el-input
          :rows="9"
          type="textarea"
          placeholder="请输入聊天内容"
          v-model="content"
          @keyup.enter="send"
          @focus="props.onFocus(props.currentInfo.userId)"
        />
      </div>
    </div>
  </div>
</template>

<style scoped lang="less">
.logo {
  width: 100px;
  height: 100px;
}
.chat-frame {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  flex-direction: column;
  flex-wrap: nowrap;
  align-items: center;
}

.chat-frame__top {
  width: 100%;
  height: 10%;
  display: flex;
  align-items: center;
  span {
    margin-left: 25px;
    font-size: 20px;
  }
}

.chat-frame__middle {
  width: calc(100% - 8%);
  height: 60%;
  padding: 0 4% 0 4%;
  background: #ffffff;
  overflow-y: auto;
  .item-chat {
    width: 100%;
    margin: 10px 0 10px 0;
  }

  .item-chat--left {
    width: 100%;
    display: flex;
    flex-direction: row;
    .content {
      padding: 15px;
      max-width: 50%;
      margin-left: 10px;
      border-radius: 10px;
      background: #b5b5b5;
      display: flex;
      align-items: center;
      span {
        width: 100%;
        word-wrap: break-word;
        word-break: normal;
      }
    }
  }

  .item-chat--right {
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    .content {
      padding: 10px;
      max-width: 50%;
      margin-right: 10px;
      border-radius: 10px;
      background: #fdcf5b;
      color: #ffffff;
      display: flex;
      align-items: center;

      span {
        width: 100%;
        word-wrap: break-word;
        word-break: normal;
      }
    }
  }
}

.chat-frame__bottom {
  position: relative;
  width: 100%;
  height: 30%;
  border-top: 2px solid #8c8c8e;

  .info-bar {
    height: 40px;
    div {
      cursor: pointer;
    }
  }

  .arrows {
    top: -50px;
    left: 30px;
    position: absolute;
    padding: 5px;
    border-radius: 10px;
    background: #fcce65;
    color: #ffffff;
    cursor: pointer;
  }
}
</style>
