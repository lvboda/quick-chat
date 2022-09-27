<script setup lang="ts">
import HeadPortrait from "@/components/HeadPortrait/HeadPortrait.vue";

type Props = {
  name: string;
  face: string;
  point?: number;
};

const props = defineProps<Props>();
</script>

<template>
  <div class="item-record-box">
    <head-portrait
      class="item-record__img"
      :src="props.face || ''"
      :show-details="false"
    >
      <template #point>
        <div v-if="point" class="point">{{ point }}</div>
      </template>
    </head-portrait>
    <div class="item-record__content">
      <span class="name">{{ props.name }}</span>
      <span class="info">
        <slot name="info" />
      </span>
    </div>
    <div class="item-record__time">
      <slot name="append" />
    </div>
  </div>
</template>

<style scoped lang="less">
.item-record-box {
  width: 90%;
  padding: 5px;
  padding-top: 15px;
  padding-bottom: 15px;
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  border-top: 1px solid #d8d8d8;
  cursor: pointer;

  .item-record__img {
    width: 40px;
    height: 40px;

    .point {
      top: -5px;
      right: -5px;
      position: absolute;
      min-width: 18px;
      min-height: 18px;
      border-radius: 18px;
      background: #f75356;
      color: #ffffff;
      font-weight: 100;
      font-size: 12px;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }

  .item-record__content {
    width: auto;
    height: auto;
    margin-left: 10px;
    display: flex;
    flex-direction: column;
    flex-wrap: nowrap;
    align-items: flex-start;
    .info {
      font-size: 10px;
      font-weight: 100;
      color: #a7a7a7;
    }
  }

  .item-record__time {
    flex: 1;
    font-size: 10px;
    font-weight: 100;
    color: #a7a7a7;
    text-align: right;
  }
}
</style>
