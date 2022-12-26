import { getFromLocal } from "@/utils/storage";
import useUserStore from "@/stores/user";
import useContactStore from "@/stores/contact";

import type { UserInfo } from "@/api/user";

function useCommon() {
  const router = useRouter();

  const localUserInfo = getFromLocal<UserInfo>("userInfo");

  if (!localUserInfo) {
    ElMessage.error("用户信息获取失败, 请重新登录");
    router.push({ path: "/login" });
    return;
  }

  const { userInfo } = storeToRefs(useUserStore());
  userInfo.value = localUserInfo;

  const { flushContacts } = useContactStore();
  flushContacts();
}

export default useCommon;
