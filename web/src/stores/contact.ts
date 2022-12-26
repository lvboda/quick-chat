import { queryList } from "@/api/contact";

import useUserStore from "@/stores/user";

import type { Contacts, Relation } from "@/api/contact";

const useContactStore = defineStore("contact", () => {
  const { userInfo } = storeToRefs(useUserStore());

  const validateList = ref<Relation[]>([]);
  const friendList = ref<Relation[]>([]);
  const groupList = ref<Relation[]>([]);

  function setContacts(data: Contacts) {
    validateList.value = data.validateList;
    friendList.value = data.friendList;
    groupList.value = data.groupList;
  }

  function flushContacts() {
    queryList(userInfo.value.userId).then(setContacts);
  }

  watch(userInfo, flushContacts);

  return { validateList, friendList, groupList, flushContacts };
});

export default useContactStore;
