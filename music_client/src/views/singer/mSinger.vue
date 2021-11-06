<template>
  <div style="margin-top:110px;">
    <van-index-bar style="font-size:20px">
      <fragment v-for="(singeritem,i) in singerList" :key="i">
        <van-index-anchor  :index="String(singeritem.index)" />
        <fragment v-if="singeritem.singers">
             <van-cell center v-for="(singer,k) in singeritem.singers" :key="k" :title="singer.name" @click="handelSinger(singer)">
               <template #icon>
                 <van-image class="avatar" round :src="websiteInfo.static + singer.avatar"></van-image>
               </template>
             </van-cell>
        </fragment>
      </fragment>
    </van-index-bar>
  </div>
</template>

<script>
import { reqSinger } from "@/api";
export default {
  name: "mSinger",
  computed: {
    websiteInfo() {
      return this.$store.getters.website;
    },
  },
  data() {
    return {
      singerList: [],
    };
  },
  mounted() {
    this.getSingerList();
  },
  methods: {
    async getSingerList() {
      const result = await reqSinger();
      if (result.code === 0) {
        this.singerList = result.data;
      }
    },
    handelSinger(data) {
        this.$router.push({
            name:"Mmusiclist",
            params:{
                singer:data
            }
        })
        console.log(data);
    }
  },
};
</script>

<style scoped>
.avatar {
  width: 38px;
}
.van-cell__title {
  font-size:14px;
  font-weight: 500;
  color:#18180c;
  margin-left: 35px;
}
</style>
