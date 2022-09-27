import request from "@/utils/http-request";
export { checkStatus } from "@/api/common";

import type { Result } from "@/api/common";
import type { UserInfo } from "@/api/user";

export type Relation = {
  id: string;
  userId: string;
  friendId: string;
  relationType: 1 | 2 | 3;
  roleType: 1 | 2;
  memo: string;
  friendInfo: UserInfo;
  proposerInfo: UserInfo;
};

async function commonQueryList(
  params: Pick<Relation, "friendId" | "relationType" | "roleType">
) {
  const { data } = await request.post<Result<Relation[]>>(
    "/api/v1/relation/list",
    params
  );

  return data;
}

// 添加好友
export async function addFriend(userId: string, friendId: string) {
  const { data } = await request.post<Result<null>>("/api/v1/relation", {
    userId,
    friendId,
    roleType: 1,
  });

  return data;
}

// 发送添加好友验证
export async function sendFriendValidate(
  params: Pick<Relation, "userId" | "friendId" | "memo">
) {
  const { data } = await request.post<Result<null>>(
    "/api/v1/relation/validate",
    { ...params, roleType: 1 }
  );

  return data;
}

// 好友验证列表
export async function queryFriendValidateList(userId: string) {
  return commonQueryList({
    friendId: userId,
    relationType: 1,
    roleType: 1,
  });
}

// 好友列表
export async function queryFriendList(userId: string) {
  return commonQueryList({
    friendId: userId,
    relationType: 2,
    roleType: 1,
  });
}

// 被删除好友列表
export async function queryBeDeletedFriendList(userId: string) {
  return commonQueryList({
    friendId: userId,
    relationType: 3,
    roleType: 1,
  });
}

// 发送添加群验证
export async function sendGroupValidate(
  params: Pick<Relation, "userId" | "friendId" | "memo">
) {
  const { data } = await request.post<Result<null>>(
    "/api/v1/relation/validate",
    { ...params, roleType: 2 }
  );

  return data;
}

// 添加群
export async function addGroup(userId: string, friendId: string) {
  const { data } = await request.post<Result<null>>("/api/v1/relation", {
    userId,
    friendId,
    roleType: 1,
  });

  return data;
}

// 群验证列表
export async function queryGroupValidateList(userId: string) {
  return commonQueryList({
    friendId: userId,
    relationType: 1,
    roleType: 2,
  });
}

// 群列表
export async function queryGroupList(userId: string) {
  return commonQueryList({
    friendId: userId,
    relationType: 2,
    roleType: 2,
  });
}

// 被移除群列表
export async function queryBeDeletedGroupList(userId: string) {
  return commonQueryList({
    friendId: userId,
    relationType: 2,
    roleType: 2,
  });
}

export type Contacts = {
  validateList: Relation[];
  friendList: Relation[];
  groupList: Relation[];
};

// 查所有
export async function queryList(userId: string): Promise<Contacts> {
  const { data: friendValidateList } = await queryFriendValidateList(userId);
  const { data: groupValidateList } = await queryGroupValidateList(userId);

  const { data: friendList } = await queryFriendList(userId);
  const { data: beDeletedFriendList } = await queryBeDeletedFriendList(userId);

  const { data: groupList } = await queryGroupList(userId);
  const { data: beDeletedGroupList } = await queryBeDeletedGroupList(userId);

  return {
    validateList: [...friendValidateList, ...groupValidateList],
    friendList: [...friendList, ...beDeletedFriendList],
    groupList: [...groupList, ...beDeletedGroupList],
  };
}
