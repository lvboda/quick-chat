import moment from "moment";
import { v4 as UUID } from "uuid";

import { getFromLocal } from "@/utils/storage";
import { wait } from "@/utils/wait";

import type { UserInfo } from "@/api/user";

const SOCKET_URL = new URL(
  `${import.meta.env.BASE_URL}ws`,
  `wss://${location.host}`
);

export type Message = {
  id: string; // id
  senderId: string; // 发送者id
  receiverId: string; // 接受者id
  content: string; // 消息内容
  extra: string; // 附加信息
  contentType: number; // 消息类型 1文本类型 2ice-offer 3ice-candidate 4ice-answer
  processType: number; // 处理类型 1单聊 2群聊 3关闭ws连接
  sendTime: string; // 发送时间
  resource: any[]; // 源数据
};

export function createMessage(
  rid: string,
  content: string,
  contentType?: number,
  processType?: number
): Message {
  const userInfo = getFromLocal<UserInfo>("userInfo");
  return {
    id: UUID(),
    senderId: userInfo?.userId || "",
    receiverId: rid,
    content,
    extra: "",
    contentType: contentType || 1,
    processType: processType || 1,
    sendTime: moment().format("YYYY-MM-DD HH:mm:ss"),
    resource: [],
  };
}

class WebsocketClient {
  private client: WebSocket | null = null;

  private path = "/";

  private uid: string | null = null;

  private closeFlag = false;

  private reconnectionMaxTry = 3;

  private waitTime = 3000;

  private listenersMap: Map<string, Set<(event?: Event) => void>> = new Map();

  constructor(path: string, uid: string) {
    if (!path || !uid) return;

    this.path = path;
    this.uid = uid;

    this.setupClient();
  }

  private setupClient(): void {
    const client = new WebSocket(`${SOCKET_URL}${this.path}/${this.uid}`);

    client.addEventListener("open", () => {
      console.info(`ws:用户${this.uid}开启websocket`);

      this.call("open");
      this.initHeart();
    });

    client.addEventListener("message", (event) => {
      this.call("message", event);
    });

    client.addEventListener("error", (error) => {
      console.error(
        `ws:用户${this.uid}连接websocket发生错误: ${JSON.stringify(error)}`
      );
      this.call("error", error);
    });

    client.addEventListener("close", async () => {
      console.info(`ws:用户${this.uid}websocket连接断开`);

      if (this.closeFlag) return;

      if (this.reconnectionMaxTry === 0) {
        console.error(`ws:用户${this.uid}websocket重新连接失败`);
        this.call("error");
        return;
      }

      this.call("close");
      await wait(this.waitTime);
      this.reconnectionMaxTry--;
      this.setupClient();
    });

    this.client = client;
  }

  private initHeart() {
    setInterval(() => this.client?.send("heart"), 1000 * 10);
  }

  private call(type: string, event?: Event): void {
    const listeners = this.listenersMap.get(type);
    if (!listeners) return;

    listeners.forEach((listener) => listener(event));
  }

  on(type: string, listener: (event?: Event | undefined) => void): void {
    if (!this.listenersMap.get(type)) {
      this.listenersMap.set(type, new Set());
    }

    const listeners = this.listenersMap.get(type);
    listeners?.add(listener);
  }

  off(type: string, listener: (event?: Event) => void): void {
    const listeners = this.listenersMap.get(type);
    if (listeners === undefined || listeners === null) return;
    listeners.delete(listener);
  }

  offAll(): void {
    this.listenersMap = new Map();
  }

  close() {
    this.closeFlag = true;
    this.client?.close();
  }

  send(msg: Message) {
    this.client?.send(JSON.stringify(msg));
  }
}

export function createChatClient(uid: string) {
  return new WebsocketClient("/api/v1/chat", uid);
}

export default WebsocketClient;
