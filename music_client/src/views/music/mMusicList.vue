<template>
  <div style="margin-top:110px;">
    <div class="singerInfo">
      <div class="avatar">
        <img :src="websiteInfo.static + singer.avatar" />
        <div class="desc">
          <h3>{{ singer.name }}</h3>
          <p>别名: {{ singer.nickname }} 国籍: {{ singer.area }}</p>
          <div class="playall" @touchstart="playingAll">
            <van-icon style="color:#fff" name="play-circle-o"></van-icon>
            <span>一键播放</span>
          </div>
        </div>
      </div>
    </div>
    <div>
      <van-list
        v-model="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
        class="musiclist"
      >
        <van-cell
          v-for="(item, index) in list"
          :key="index"
          class="musicitem"
          center
          :label="item.singer_name + ' 专辑:[' + item.category_title + ']'"
          :title="item.name"
        >
          <template #icon>
            <van-image class="pic" :src="websiteInfo.static + item.pic" radius="5"></van-image>
          </template>
          <van-icon name="play-circle-o" color="#07c160" size="30" @touchstart="toPlayer(index)"></van-icon>
        </van-cell>
      </van-list>
    </div>
  </div>
</template>

<script>
import { reqMusicBySinger } from "@/api";
export default {
  name: "mMusicList",
  beforeRouteEnter(to, from, next) {
    if (to.params.singer) {
      next((vm) => {
        vm.singer = to.params.singer;
      });
    } else {
      this.$router.back();
    }
  },
  computed: {
    websiteInfo() {
      return this.$store.getters.website;
    },
  },
  data() {
    return {
      singer: {},
      list: [],
      loading: false,
      finished: false,
      count: 5,
      start: 1,
    };
  },
  //   mounted() {
  //     this.getMusicBySinger();
  //   },
  methods: {
    // 获取歌手歌曲
    async getMusicBySinger() {
      this.loading = true;
      var s = (this.start - 1) * this.count;
      const result = await reqMusicBySinger(this.singer.id, this.count, s);
      if (result.code === 0) {
        if (result.data == null || result.data.length < this.count) {
          this.finished = true;
          this.loading = false;
          return;
        }
        var res = result.data;
        this.loading = false;
        if (this.list.length > 0) {
          this.list.push(...res);
        } else {
          this.list = res;
        }
      } else {
        this.loading = false;
        this.finished = true;
      }
      this.start += 1;
    },
    playingAll() {
      console.log("all");
    },
    onLoad() {
      setTimeout(() => {
        if (!this.finished) {
          this.getMusicBySinger();
        }
      }, 1000);
    },
    toPlayer(index) {
      this.$store.dispatch('setPlayList', this.list) 
      this.$store.dispatch('setFullScreen', true) 
      this.$store.dispatch('setCurrentIndex',index)
    }
  },
};
</script>

<style scoped>
.singerInfo {
  display: flex;
  background-color: #f4f4f4;
}
.avatar {
  display: flex;
  height: 128px;
  align-items: center;
}
.avatar img {
  width: 108px;
  height: 108px;
  margin-left: 15px;
  border-radius: 15px;
}
.desc {
  margin-left: 15px;
  display: flex;
  flex-flow: column;
}
.desc p {
  font-size: 14px;
}
.playall {
  width: 105px;
  height: 25px;
  background-color: #07c160;
  margin: 15px 0;
  display: flex;
  align-items: center;
  border-radius: 10px;
  color: #fff;
  justify-content: center;
}
.playall span {
  font-size: 14px;
  margin-left: 10px;
}
.musiclist {
  margin-top: 15px;
  height: 100%;
}
.musicitem {
  background-color: #fff;
  height: 95px;
}
.pic {
  width: 68px;
  height:68px;
}
.van-cell__title {
  font-size:16px;
  font-weight: 700;
  color:#18180c;
  margin-left: 25px;
  display: flex;
  flex-flow: column;
}
.van-cell__value{
  display: flex;
  justify-content: center;
}
</style>
