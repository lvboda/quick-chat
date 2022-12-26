<script setup lang="ts">
import { queryFriend, checkStatus } from "@/api/user";
import { sendFriendValidate } from "@/api/contact";
import useUserStore from "@/stores/user";
import useContactStore from "@/stores/contact";

import ItemRecord from "../ItemRecord/ItemRecord.vue";
import ButtonBox from "@/components/ButtonBox/ButtonBox.vue";
import { Search, Plus } from "@element-plus/icons-vue";

import type { UserInfo } from "@/api/user";

const { userInfo } = storeToRefs(useUserStore());
const { flushContacts } = useContactStore();

let modalVisible = $ref(false);
let queryId = $ref("");
let selectedData = $ref<UserInfo[]>([]);

function closeModal() {
  queryId = "";
  selectedData = [];
  modalVisible = false;
}

async function query() {
  if (!queryId) ElMessage.warning("请输入要查询的id");

  const { status, data } = await queryFriend(queryId);
  if (checkStatus(status) && data) {
    selectedData = [data];
    return;
  }
  selectedData = [];
  ElMessage.warning("未找到好友/群, 请检查id是否正确");
}

async function sendValidate(friendId: string) {
  const { status } = await sendFriendValidate({
    userId: userInfo.value.userId,
    friendId,
    memo: "1",
  });
  if (checkStatus(status)) {
    ElMessage.success("已发送好友验证申请");
    flushContacts();
    closeModal();
    return;
  }
  ElMessage.error("好友验证申请发送失败");
}
</script>

<template>
  <button-box :click="() => (modalVisible = true)">添加好友/群聊</button-box>
  <el-dialog v-model="modalVisible" width="400px" @close="closeModal">
    <div class="add-modal-search">
      <el-input v-model="queryId" placeholder="输入好友/群聊id进行搜索">
        <template #append>
          <el-button class="but" :icon="Search" @click="query" />
        </template>
      </el-input>
    </div>
    <div class="add-modal-content">
      <div v-if="selectedData.length === 0" class="empty">暂无查询结果</div>
      <item-record
        :key="item.nickName"
        v-for="item in selectedData"
        :name="item.nickName"
        :face="item.face"
      >
        <template #append>
          <el-button
            class="but"
            :icon="Plus"
            @click="sendValidate(item.userId)"
          >
            申请添加好友
          </el-button>
        </template>
      </item-record>
    </div>
  </el-dialog>
</template>

<style scoped lang="less">
.add-modal-search {
  display: flex;
  .but {
    border-top-left-radius: 0;
    border-bottom-left-radius: 0;
    font-size: 16px;
    font-weight: bold;
    color: #ffffff;
    background-image: linear-gradient(
      to right,
      #ffb54d 0%,
      #ffcc33 51%,
      #ffb54d 100%
    );
  }
}

.add-modal-content {
  margin-top: 40px;
  div {
    width: auto;
  }
  .empty {
    height: 35px;
    text-align: center;
    font-size: 10px;
    font-weight: 100;
    color: #a7a7a7;
    border-bottom: 1px solid #d8d8d8;
  }
}
</style>
