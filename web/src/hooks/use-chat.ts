import useUserStore from "@/stores/user";
import {
  // onReceiveICEOffer,
  // onReceiveICECandidate,
  // onReceiveICEAnswer,
  onSignal,
} from "@/utils/web-rtc";
import WebsocketClient, {
  createChatClient,
  createMessage,
} from "@/utils/websocket";

import type { Message } from "@/utils/websocket";

export let chatClient: WebsocketClient | null = null;

function useChat() {
  const { userInfo } = storeToRefs(useUserStore());

  // 创建ws连接
  if (chatClient === null) chatClient = createChatClient(userInfo.value.userId);

  const onMessage = $ref<{ cb: (msg: Message) => void }>({
    cb: () => {},
  });

  // 存储全局聊天记录 key: friendId, value: 聊天记录list
  // 劫持发送和接收函数, 把数据存进来
  const globalChatMap = $ref<Map<string, Message[]>>(new Map());
  // get
  function getChatBy(uid: string): Message[] {
    return globalChatMap.get(uid) || [];
  }
  // set
  function setChat(uid: string, msg: Message) {
    if (!globalChatMap.has(uid)) globalChatMap.set(uid, []);
    globalChatMap.get(uid)?.push(msg);
  }

  chatClient.on("error", () =>
    ElMessage.error("连接服务端失败, 请尝试切换网络环境")
  );

  chatClient.on("message", (event: any) => {
    if (!event || !event.data) return;

    const msg = JSON.parse(event.data) as Message;

    switch (msg.contentType) {
      case 1:
        onMessage.cb(msg);
        setChat(msg.senderId, msg);
        break;
      case 2:
        onMessage.cb(msg);
        onSignal(msg);
        // onReceiveICEOffer(msg);
        break;
      case 3:
        // onReceiveICECandidate(msg);
        break;
      case 4:
        // onReceiveICEAnswer(msg);
        break;
    }
  });

  function sendMessage(rid: string, content: string, contentType?: number) {
    const msg = createMessage(rid, content, contentType);
    setChat(msg.receiverId, msg);
    chatClient?.send(msg);
  }

  return { chatClient, getChatBy, onMessage, sendMessage };
}

export default useChat;
