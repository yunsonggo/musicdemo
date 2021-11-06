<template>
  <div class="mrecommend">
      <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list v-model="loading" :finished="finished" @load="onLoad">
        <Mswiper></Mswiper>
        <div class="more">
          <span></span>
          <span>推荐歌单</span>
          <span @touchstart="categoryList">更多...</span>
        </div>
        <div class="top">
          <div v-for="(cate, index) in topCategories" :key="index">
            <div class="cate">
              <div>
                <img :src="websiteInfo.static + cate.image" />
              </div>
              <div class="comment">
                <span>{{ cate.title }}</span>
                <span>描述:{{ cate.desc }}</span>
              </div>
            </div>
          </div>
        </div>
      </van-list>
    </van-pull-refresh>
  </div>
</template>

<script>
import Mswiper from "@/components/swiper/mSwiper.vue";
import Scroll from "@/components/scroll/scroll.vue";
import { reqTopCategory } from "@/api";
export default {
  name: "Mrecommend",
  components: {
    Mswiper,
    Scroll
  },
  computed: {
    websiteInfo() {
      return this.$store.getters.website;
    },
  },
  data() {
    return {
      loading: false,
      finished: false,
      refreshing: false,
      count: 6,
      start: 1,
      topCategories: [],
    };
  },
  mounted() {
    this.getTopCategories();
  },
  methods: {
    async getTopCategories() {
      const result = await reqTopCategory(this.count, this.start);
      console.log(result);
      if (result.code === 0) {
        this.topCategories = result.data;
      }
    },
    onRefresh() {
      // 清空列表数据
      this.finished = false;
      // 重新加载数据
      // 将 loading 设置为 true，表示处于加载状态
      this.loading = true;
      this.onLoad();
    },
    onLoad() {
      setTimeout(() => {
        if (this.refreshing) {
          this.refreshing = false;
        }
        this.loading = false;
        this.finished = true;
      }, 1000);
    },
    categoryList() {
      console.log("list");
    }
  },
};
</script>

<style scoped>
.mrecommend {
  height:550px;
  margin-top:110px;
}
.more {
  height: 35px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.more span {
  flex: 1;
}
.more span:nth-child(2) {
  text-align: center;
  font-size: 16px;
  padding-bottom: 5px;
  border-bottom: #d43c33 1px solid;
}
.more span:nth-child(3) {
  text-align: center;
  font-size: 14px;
  color:#666
}
.top {
  display: flex;
  flex-flow: column;
  width: 100%;
  background-color: #f4f4f4;
  margin-top: 10px;
}
.cate {
  margin-bottom: 5px;
  display: flex;
  background-color: #fff;
}
.cate div:nth-child(1) {
  width: 88px;
  height: 88px;
  display: flex;
  justify-content: center;
  margin-left: 5px;
  background-color: #fff;
}
.cate img {
  padding: 5px;
}
.comment {
  display: flex;
  flex-flow: column;
  justify-content: center;
}
.comment span:nth-child(1) {
  font-size: 16px;
  margin-bottom: 5px;
}
.comment span:nth-child(2) {
  font-size: 14px;
  color: #5f6368;
}
</style>
