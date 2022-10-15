import { getDefaultUserInfo } from "@/api/user";
import type { UserInfo } from "@/api/user";

const useUserStore = defineStore("user", () => {
  const userInfo = ref<UserInfo>(getDefaultUserInfo());

  return { userInfo };
});

export default useUserStore;
