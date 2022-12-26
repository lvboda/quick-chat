import axios, { AxiosResponse, AxiosRequestConfig } from "axios";

import { getFromLocal } from "@/utils/storage";

import type { UserInfo } from "@/api/user";

const router = useRouter();

const request = axios.create({
  baseURL: `${import.meta.env.BASE_URL}apis`,
});

export enum STATUS_CODE {
  success = 200,
  error = 500,
}

async function handleRequestSuccess(
  request: AxiosRequestConfig
): Promise<AxiosRequestConfig> {
  const userInfo = getFromLocal<UserInfo>("userInfo");

  if (request.headers && userInfo) {
    request.headers.Authorization = `Bearer ${userInfo.token}`;
  }

  return request;
}

async function handleResponseSuccess(
  response: AxiosResponse
): Promise<AxiosResponse> {
  if (response.status === 401) {
    router.replace({ path: "/login" });
  }

  if (response.status !== STATUS_CODE.success) {
    console.error("http-request:response.status not the success code.");
  }

  return response;
}

async function handleResponseError(error: unknown): Promise<unknown> {
  console.error("http-request:responseErr: ", error);
  return Promise.reject(error);
}

request.interceptors.request.use(handleRequestSuccess);
request.interceptors.response.use(handleResponseSuccess, handleResponseError);

export default request;
