<script setup lang="ts">
import useContactStore from "@/stores/contact";

import ItemRecord from "./components/ItemRecord/ItemRecord.vue";
import AddModal from "./components/AddModal/AddModal.vue";
import DetailBox from "./components/DetailBox/DetailBox.vue";
import { Search } from "@element-plus/icons-vue";

import type { Relation } from "@/api/contact";

const defaultNames = ["3"];

let detailInfo = $ref<Relation | null>(null);

function showDetail(data: Relation | null) {
  if (detailInfo && data && detailInfo.id === data.id) {
    detailInfo = null;
    return;
  }
  detailInfo = data;
}

const { validateList, friendList, groupList } = storeToRefs(useContactStore());
</script>

<template>
  <div class="home-contact-view">
    <div class="home-contact-view__list">
      <el-input class="home-contact-view__list__search" placeholder="搜索">
        <template #append>
          <el-button :icon="Search" />
        </template>
      </el-input>
      <add-modal />
      <div class="home-contact-view__list__content">
        <el-collapse v-model="defaultNames">
          <el-collapse-item title="验证信息" name="1">
            <div v-if="validateList.length === 0" class="empty">
              暂无验证信息
            </div>
            <item-record
              :key="item.id"
              v-for="item in validateList"
              :name="item.proposerInfo.nickName"
              :face="item.proposerInfo.face"
              :class="detailInfo?.id === item.id ? 'clicked' : ''"
              @click="showDetail(item)"
            >
              <template #info><div>申请添加您为好友</div></template>
            </item-record>
          </el-collapse-item>
          <el-collapse-item title="群聊" name="2">
            <div v-if="groupList.length === 0" class="empty">
              暂无群聊, 快去添加吧!
            </div>
          </el-collapse-item>
          <el-collapse-item title="好友" name="3">
            <div v-if="friendList.length === 0" class="empty">
              暂无好友, 快去添加吧!
            </div>
            <item-record
              :key="item.id"
              v-for="item in friendList"
              :name="item.friendInfo.nickName"
              :face="item.friendInfo.face"
              :class="detailInfo?.id === item.id ? 'clicked' : ''"
              @click="showDetail(item)"
            >
            </item-record>
          </el-collapse-item>
        </el-collapse>
      </div>
    </div>
    <div class="home-contact-view__detail">
      <detail-box :data="detailInfo" :close="() => showDetail(null)" />
    </div>
  </div>
</template>

<style scoped lang="less">
.home-contact-view {
  width: 93%;
  height: 100%;
  display: flex;
}

.home-contact-view__list {
  width: 28%;
  height: calc(100% - 50px);
  padding: 25px;
  background: #e4e4e4;
  border-right: 2px solid #8c8c8e;
}

.home-contact-view__list__search {
  margin-bottom: 15px;
}

.home-contact-view__list__content {
  width: 100%;
  height: 95%;
  margin-top: 5px;
  overflow-y: scroll;

  .empty {
    height: 35px;
    text-align: center;
    font-size: 10px;
    font-weight: 100;
    color: #a7a7a7;
    border-bottom: 1px solid #d8d8d8;
  }

  .clicked {
    background: #9c9c9c;
  }

  :deep(.el-collapse) {
    border: none;
  }

  :deep(.el-collapse-item__header) {
    margin-left: 8px;
    background: transparent;
    border: none;
  }

  :deep(.el-collapse-item__content) {
    margin-left: 8px;
  }

  :deep(.el-collapse-item__wrap) {
    background: transparent;
    border: none;
  }
}
::-webkit-scrollbar {
  display: none;
}

.home-contact-view__detail {
  width: 70%;
  height: 100%;
  background: #ededed;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
