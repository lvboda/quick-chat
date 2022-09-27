// eslint-disable-next-line @typescript-eslint/no-unused-vars
import _ from "webrtc-adapter";

import { chatClient } from "@/hooks/use-chat";
import { createMessage } from "@/utils/websocket";
import { iceServer } from "@/utils/constant";

import type { Message } from "@/utils/websocket";

let pc: RTCPeerConnection | null = null;
const senderList: RTCRtpSender[] = [];

let peer: any = null;
let iceRemoteVideo: HTMLVideoElement | null = null;
let iceLocalVideo: HTMLVideoElement | null = null;
export async function createPeer(isSender: boolean, receiverId: string) {
  peer = new window.SimplePeer({ initiator: isSender, config: iceServer });

  peer.on("signal", (signal: any) => {
    const msg = createMessage(receiverId, JSON.stringify(signal), 2);
    chatClient?.send(msg);
  });

  peer.on("stream", (stream: any) => {
    iceRemoteVideo = document.getElementById("iceRemoteVideo") as HTMLVideoElement;
    if (iceRemoteVideo) iceRemoteVideo.srcObject = stream;
  });

  const webcamStream = await navigator.mediaDevices.getUserMedia({
    video: true,
    audio: true,
  });

  const ms = new MediaStream();
  webcamStream.getVideoTracks().forEach((track) => {
    ms.addTrack(track);
  });

  iceLocalVideo = document.getElementById("iceLocalVideo") as HTMLVideoElement;
  if (iceLocalVideo) iceLocalVideo.srcObject = ms;

  peer.addStream(webcamStream);

  return () => {
    if (iceRemoteVideo) iceRemoteVideo.srcObject = null;
    if (iceLocalVideo) iceLocalVideo.srcObject = null;
    peer.removeStream(webcamStream);
    webcamStream.getTracks().forEach((item) => item.stop());
    ms.getTracks().forEach((item) => item.stop());
    peer.destroy();
  };
}

export async function onSignal(msg: Message) {
  if (!peer) await createPeer(false, msg.senderId);
  peer.signal(JSON.parse(msg.content));
}

export async function createRTCPeerConnection(
  isSender: boolean,
  receiverId: string
) {
  pc = new RTCPeerConnection();
  pc.setConfiguration(iceServer);

  let iceRemoteVideo: HTMLVideoElement | null = null;
  let iceLocalVideo: HTMLVideoElement | null = null;

  pc.addEventListener("icecandidate", (event) => {
    console.log(111);
    if (!event.candidate) return;
    console.log(222);
    const msg = createMessage(receiverId, JSON.stringify(event.candidate), 3);
    chatClient?.send(msg);
  });

  pc.ontrack = (event) => {
    iceRemoteVideo = document.getElementById("iceRemoteVideo") as HTMLVideoElement;
    if (iceRemoteVideo) iceRemoteVideo.srcObject = event.streams[0];
  };

  const webcamStream = await navigator.mediaDevices.getUserMedia({
    video: true,
    audio: true,
  });

  iceLocalVideo = document.getElementById("iceLocalVideo") as HTMLVideoElement;
  if (iceLocalVideo) {
    const ms = new MediaStream();
    webcamStream.getVideoTracks().forEach((track) => {
      ms.addTrack(track);
    });
    iceLocalVideo.srcObject = ms;
  }

  webcamStream.getTracks().forEach((track) => {
    const sender = pc?.addTrack(track, webcamStream);
    if (sender) senderList.push(sender);
  });

  if (isSender) {
    const offer = await pc.createOffer();
    await pc.setLocalDescription(offer);
    const msg = createMessage(
      receiverId,
      JSON.stringify(pc?.localDescription),
      2
    );
    chatClient?.send(msg);
  }

  return () => {
    if (iceRemoteVideo) iceRemoteVideo.srcObject = null;
    if (iceLocalVideo) iceLocalVideo.srcObject = null;
    webcamStream.getTracks().forEach((item) => item.stop());
    senderList.forEach((item) => pc?.removeTrack(item));
    pc?.close();
  };
}

export async function onReceiveICEOffer(msg: Message) {
  await createRTCPeerConnection(false, msg.senderId);
  if (!pc) return;

  await pc.setRemoteDescription(
    new RTCSessionDescription(JSON.parse(msg.content))
  );
  const answer = await pc.createAnswer();
  await pc.setLocalDescription(answer);

  chatClient?.send(createMessage(msg.senderId, JSON.stringify(answer), 4));
}

export async function onReceiveICECandidate(msg: Message) {
  try {
    await pc?.addIceCandidate(new RTCIceCandidate(JSON.parse(msg.content)));
  } catch (err) {
    reportError(err);
  }
}

export async function onReceiveICEAnswer(msg: Message) {
  await pc
    ?.setRemoteDescription(new RTCSessionDescription(JSON.parse(msg.content)))
    .catch(reportError);
}
