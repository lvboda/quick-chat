<script setup lang="ts">
import useContactStore from "@/stores/contact";
import useChat from "@/hooks/use-chat";
import { createPeer } from "@/utils/web-rtc";

import ButtonBox from "@/components/ButtonBox/ButtonBox.vue";
import ItemRecord from "./components/ItemRecord/ItemRecord.vue";
import ChatFrame from "./components/ChatFrame/ChatFrame.vue";
import { Search } from "@element-plus/icons-vue";

import type { Relation } from "@/api/contact";
import type { Message } from "@/utils/websocket";

type Chat = Relation & { info?: string; time?: string; point?: number };

const router = useRouter();
const { friendList } = storeToRefs(useContactStore());

let detailInfo = $ref<Relation | null>(null);
const chatList = $ref<Chat[]>([]);

const { getChatBy, onMessage, sendMessage } = useChat();
const currentChatList = computed(() => getChatBy(detailInfo?.userId || ""));

let isShowVideo = $ref(false);
function showVideo(is = true) {
  isShowVideo = is;
}
let isLocalVideo = $ref(false);

let closeConn: (() => void) | null = null;

function sendMessageProxy(id: string, content: string, contentType?: number) {
  switch (contentType) {
    case undefined:
      sendMessage(id, content);
      break;
    case 2: {
      showVideo();
      createPeer(true, id).then((cb) => closeConn = cb);
      // createRTCPeerConnection(true, id).then((cb) => closeConn = cb);
      break;
    }
  }
}

function hangup() {
  closeConn?.();
  showVideo(false);
}

onMessage.cb = (msg: Message) => {
  if (msg.contentType === 2) {
    showVideo();
    return;
  }
  const info = friendList.value.find((item) => item.userId === msg.senderId);
  if (!info) return;
  const chatInfo = chatList.find((item) => item.userId === msg.senderId);
  if (!chatInfo) {
    chatList.unshift({
      ...info,
      info: msg.content,
      time: msg.sendTime,
      point: 1,
    });
  } else {
    chatInfo.info = msg.content;
    chatInfo.time = msg.sendTime;
    chatInfo.point = chatInfo.point ? chatInfo.point + 1 : 1;
  }
};

watch(
  () => [router.currentRoute.value, friendList.value],
  () => {
    const { friendId } = router.currentRoute.value.query;
    if (!friendId) return;

    const info = friendList.value.find((item) => item.userId === friendId);
    if (!info) return;

    const chatInfo = chatList.find((item) => item.userId === info.userId);
    if (!chatInfo) {
      chatList.unshift(info);
      showChatFrame(info);
    }
  },
  { immediate: true, deep: true }
);

function showChatFrame(data: Relation | null) {
  if (detailInfo && data && detailInfo.id === data.id) {
    detailInfo = null;
    return;
  }
  detailInfo = data;
}

function onFocus(fid: string) {
  const info = chatList.find((item) => item.userId === fid);
  if (info) info.point = 0;
}
</script>

<template>
  <div class="home-chat-view">
    <div class="home-chat-view__list">
      <el-input class="home-chat-view__list__search" placeholder="搜索">
        <template #append>
          <el-button :icon="Search" />
        </template>
      </el-input>
      <div class="home-chat-view__list__content">
        <div v-if="chatList.length === 0" class="empty">暂无会话</div>
        <item-record
          :key="item.id"
          v-for="item in chatList"
          :name="item.friendInfo.nickName"
          :face="item.friendInfo.face"
          :point="item.point"
          :class="detailInfo?.id === item.id ? 'clicked' : ''"
          @click="showChatFrame(item)"
        >
          <template #info>
            {{ item?.info }}
          </template>
          <template #append>
            {{ item?.time }}
          </template>
        </item-record>
      </div>
    </div>
    <div class="home-chat-view__detail">
      <chat-frame
        v-if="!!detailInfo"
        :current-info="detailInfo"
        :current-chat-list="currentChatList"
        :send-message="sendMessageProxy"
        :on-focus="onFocus"
      />
    </div>
    <div v-if="isShowVideo" class="home-chat-view__video-frame">
      <div class="video-group">
        <video id="iceRemoteVideo" :class="isLocalVideo ? 'iceLocalVideo' : 'iceRemoteVideo'" autoplay playsinline muted @click=" isLocalVideo = !isLocalVideo" />
        <video id="iceLocalVideo" :class="isLocalVideo ? 'iceRemoteVideo' : 'iceLocalVideo'" autoplay playsinline muted @click=" isLocalVideo = !isLocalVideo" />
        <button-box class="but" :click="hangup">挂断</button-box>
      </div>
    </div>
  </div>
</template>

<style scoped lang="less">
.home-chat-view {
  width: 93%;
  height: 100%;
  display: flex;
}

.home-chat-view__list {
  width: 28%;
  height: calc(100% - 50px);
  padding: 25px;
  background: #e4e4e4;
  border-right: 2px solid #8c8c8e;
}

.home-chat-view__list__search {
  margin-bottom: 15px;
}

.home-chat-view__list__content {
  width: 100%;
  height: 95%;
  margin-top: 5px;
  overflow-y: scroll;

  .empty {
    height: 35px;
    text-align: center;
    font-size: 10px;
    font-weight: 100;
    color: #a7a7a7;
    border-bottom: 1px solid #d8d8d8;
  }

  .clicked {
    background: #9c9c9c;
  }
}
::-webkit-scrollbar {
  display: none;
}

.home-chat-view__detail {
  width: 70%;
  height: 100%;
  background: #ededed;
  display: flex;
  align-items: center;
  justify-content: center;
}

.home-chat-view__video-frame {
  position: fixed;
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  .video-group {
    position: relative;
    width: 55%;
    height: 79%;
    .iceRemoteVideo {
      position: absolute;
      width: 100%;
      height: 100%;
    }

    .iceLocalVideo {
      right: 0;
      bottom: 0;
      position: absolute;
      width: 30%;
      height: 30%;
      z-index: 2;
      cursor: pointer;
    }
    .but {
      width: 100px !important;
      left: calc(50% - 50px);
      bottom: 10%;
      position: absolute;
      z-index: 3;
    }
  }
}
</style>
