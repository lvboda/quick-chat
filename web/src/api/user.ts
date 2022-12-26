import request from "@/utils/http-request";
export { checkStatus } from "@/api/common";
import type { GENDER, USER_ROLE, Result } from "@/api/common";

export type UserInfo = {
  id: string;
  userId: string;
  password: string;
  nickName: string;
  gender: GENDER;
  mobile: number;
  userRole: USER_ROLE;
  signature: string;
  face: string;
  token: string;
};

export type LoginParams = Pick<UserInfo, "userId" | "password">;

export type RegisterParams = Required<
  Pick<UserInfo, "userId" | "password" | "nickName" | "userRole">
>;

export function getDefaultUserInfo(userInfo?: any): UserInfo {
  return {
    id: "",
    userId: "",
    password: "",
    nickName: "",
    gender: NaN,
    mobile: NaN,
    userRole: NaN,
    signature: "",
    face: "",
    token: "",
    ...userInfo,
  };
}

export async function register(params: RegisterParams) {
  const { data } = await request.post<Result<null>>(
    "/api/v1/user/register",
    params
  );

  return data;
}

export async function login(params: LoginParams) {
  const { data } = await request.post<Result<UserInfo>>(
    "/api/v1/user/login",
    params
  );

  return data;
}

export async function queryFriend(friendId: string) {
  const { data } = await request.get<Result<UserInfo>>(
    `/api/v1/user/${friendId}`
  );

  return data;
}
