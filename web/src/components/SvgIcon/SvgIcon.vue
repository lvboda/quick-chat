<script setup lang="ts">
type Props = {
  iconName: string;
  className?: string;
  color?: string;
};

const props = defineProps<Props>();

const iconName = computed(() => `#icon-${props.iconName}`);
const svgClass = computed(() =>
  props.className ? `svg-icon ${props.className}` : `svg-icon`
);
</script>

<template>
  <svg :class="svgClass" aria-hidden="true">
    <use :xlink:href="iconName" :fill="color" />
  </svg>
</template>

<style scoped lang="less">
.svg-icon {
  width: 1em;
  height: 1em;
  position: relative;
  fill: currentColor;
  vertical-align: -2px;
}
</style>
