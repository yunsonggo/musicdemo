<template>
  <div class="swiper">
    <el-carousel indicator-position="outside" style="width:1080px;">
      <el-carousel-item v-for="(swiper, index) in swiperArr" :key="index">
        <img
          v-lazy="websiteInfo.static + swiper.image"
          @click="handleTouch(swiper)"
        />
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
import { reqSwiper } from "@/api";
export default {
  name: "swiper",
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
    display: flex;
    justify-content: center;
    background-color:#f4f4f4
}

</style>
