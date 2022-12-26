// 性别
export enum GENDER {
  male = 1,
  female = 2,
}

// 角色类型
export enum USER_ROLE {
  common = 1,
  manager = 2,
}

// 服务端状态码
export enum STATUS {
  ok = 200,
  error = 500,
}

// 响应体结构
export type Result<T> = {
  status: STATUS;
  message: string;
  data: T;
};

export function checkStatus(status: STATUS): boolean {
  if (status === STATUS.ok) return true;
  return false;
}
