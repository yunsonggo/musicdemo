<template>
  <div class="swiper">
    <van-swipe class="my-swipe" :autoplay="3000" indicator-color="#d43c33">
      <van-swipe-item v-for="(swiper, index) in swiperArr" :key="index">
        <img
          style="height:250px;width:100%"
          v-lazy="websiteInfo.static + swiper.image"
          @touchstart="handleTouch(swiper)"
        />
      </van-swipe-item>
    </van-swipe>
  </div>
</template>

<script>
import { reqSwiper } from "@/api";
export default {
  name: "Mswiper",
  data() {
    return {
      swiperArr: [],
    };
  },
  computed: {
      websiteInfo() {
        return this.$store.getters.website
      }
  },
  mounted() {
    this.getSwiperArr();
  },
  methods: {
    // 获取轮播图数据
    async getSwiperArr() {
      const result = await reqSwiper();
      if (result.code === 0) {
        this.swiperArr = result.data;
      }
    },
    handleTouch(data) {
      console.log(data);
    },
  },
};
</script>

<style scoped>
.swiper {
  height: 280px;
}
.my-swipe .van-swipe-item {
  color: #fff;
  font-size: 14px;
  line-height: 70px;
  text-align: center;
  background-color: #fff;
}
</style>
