<template>
  <div ref="wrapper">
    <slot></slot>
  </div>
</template>

<script>
import BScroll from "@better-scroll/core";
export default {
  name: "scroll",
  props: {
    probeType: {
      type: Number,
      default: 1,
    },
    click: {
      type: Boolean,
      default: true,
    },
    dataArr: {
      type: Array,
      default: () => [],
    },
  },
  watch:{
      // 数据变化 重新计算高度
      dataArr() {
          setTimeout(() => {
              this.refresh()
          },20)
      }
  },
  mounted() {
    setTimeout(() => {
      this.initScroll();
    }, 20);
  },
  methods: {
    initScroll() {
      if (!this.$refs.wrapper) {
        return;
      }
      this.scroll = new BScroll(this.$refs.wrapper, {
        probeType: this.probeType,
        click: this.click,
      });
    },
    enable() {
        this.scroll && this.scroll.enable()
    },
    disable() {
        this.scroll && this.scroll.disable()
    },
    // 重新计算高度
    refresh() {
        this.scroll && this.scroll.refresh()
    }
  }, 
};
</script>

<style scoped></style>
