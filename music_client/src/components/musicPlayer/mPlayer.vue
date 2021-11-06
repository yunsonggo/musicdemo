<template>
  <div class="player" v-show="playList.length > 0" ref="player">
    <transition
      name="normal"
      @enter="enter"
      @after-enter="afterEnter"
      @leave="leave"
      @after-leave="afterLeave"
    >
      <div class="normalPlayer" v-show="fullScreen" ref="normalPlayer">
        <div class="background" v-if="playList.length > 0">
          <img
            width="100%"
            height="100%"
            :src="websiteInfo.static + playList[currentIndex].pic"
            alt=""
          />
        </div>
        <div class="top">
          <div class="back">
            <van-icon
              class-prefix="music-icon"
              name="igw-l-arrow-down"
              size="22"
              @touchstart="handleMiniPlayer"
            ></van-icon>
          </div>
          <h1 class="title" v-if="playList.length > 0">{{ playList[currentIndex].name }}</h1>
          <h2 class="subtitle" v-if="playList.length > 0">{{ playList[currentIndex].singer_name }}</h2>
        </div>
        <div class="middle">
          <div class="middle-l">
            <div class="cd-wrapper" ref="cdWrapper">
              <div class="cd" v-if="playList.length > 0" :class="cdCls">
                <img
                  :src="websiteInfo.static + playList[currentIndex].pic"
                  alt=""
                  class="image"
                />
              </div>
            </div>
          </div>
        </div>
        <div class="bottom">
          <div class="operators">
            <div class="icon i-left">
              <van-icon
                class-prefix="music-icon"
                name="shunxubofang"
                size="18"
              ></van-icon>
            </div>
            <div class="icon i-left">
              <van-icon
                class-prefix="music-icon"
                name="yduishangyiqu"
                size="20"
                @touchstart="prevSong"
              ></van-icon>
            </div>
            <div class="icon i-center" v-if="songState">
              <van-icon
                class-prefix="music-icon"
                name="zanting"
                size="28"
                @touchstart="songState = false"
              ></van-icon>
            </div>
            <div class="icon i-center" v-else>
              <van-icon
                class-prefix="music-icon"
                name="bofang"
                size="30"
                @touchstart="songState = true"
              ></van-icon>
            </div>
            <div class="icon i-right">
              <van-icon
                class-prefix="music-icon"
                name="yduixiayiqu"
                size="20"
                @touchstart="nextSong"
              ></van-icon>
            </div>
            <div class="icon i-right">
              <van-icon
                class-prefix="music-icon"
                name="xihuan"
                size="18"
              ></van-icon>
            </div>
          </div>
        </div>
      </div>
    </transition>
    <transition name="mini">
      <div class="miniPlayer" v-show="!fullScreen">
        <div class="miniPlayer-icon" @touchstart="handleToPlayer" v-if="playList.length > 0" :class="miniCdCls">
          <van-image
            width="40"
            height="40"
            round
            :src="websiteInfo.static + playList[currentIndex].pic"
            alt=""
          />
        </div>
        <div class="miniPlayer-text" @touchstart="handleToPlayer">
          <div class="miniPlayer-name" v-if="playList.length > 0">{{ playList[currentIndex].name }}</div>
          <div class="miniPlayer-name-desc" v-if="playList.length > 0">
            {{ playList[currentIndex].singer_name }}
          </div>
        </div>
        <div class="miniPlayer-control">
          <van-icon
            class-prefix="music-icon"
            name="shunxubofang"
            size="18"
          ></van-icon>
        </div>
        <div class="miniPlayer-control">
          <van-icon
            class-prefix="music-icon"
            name="yduishangyiqu"
            size="20"
            @touchstart="prevSong"
          ></van-icon>
        </div>
        <div class="miniPlayer-control" v-if="songState">
          <van-icon
            class-prefix="music-icon"
            name="zanting"
            size="28"
            @touchstart="songState = false"
          ></van-icon>
        </div>
        <div class="miniPlayer-control" v-else>
          <van-icon
            class-prefix="music-icon"
            name="bofang"
            size="30"
            @touchstart="songState = true"
          ></van-icon>
        </div>
        <div class="miniPlayer-control">
          <van-icon
            class-prefix="music-icon"
            name="yduixiayiqu"
            size="20"
            @touchstart="nextSong"
          ></van-icon>
        </div>
        <div class="miniPlayer-control">
          <van-icon
            class-prefix="music-icon"
            name="xihuan"
            size="18"
          ></van-icon>
        </div>
      </div>
    </transition>
    <audio :src="currentSong.link" ref="audio"></audio>
  </div>
</template>

<script>
import animations from "create-keyframe-animation";
export default {
  name: "mPlayer",
  data() {
    return {
      currentSong:{},
      songState:true, // 歌曲播放状态
    }
  },
  watch: {
    currentSong() {
      this.$nextTick(() =>{
        this.$refs.audio.play()
        this.showFullScreen()
        this.songState = true
      })
    },
    songState(newState) {
      const audio = this.$refs.audio
      this.$nextTick(() => {
        newState ? audio.play() : audio.pause()
      })
      
    }
  },
  computed: {
    websiteInfo() {
      return this.$store.getters.website;
    },
    playList() {
      return this.$store.getters.playList;
    },
    fullScreen() {
      return this.$store.getters.fullScreen;
    },
    currentIndex() {
      var ci = this.$store.getters.currentIndex;
      if (this.playList) {
        this.currentSong = this.playList[ci]
      }
      return ci;
    },
    cdCls() {
      return this.songState ? 'play' : 'pause'
    },
    miniCdCls() {
      return this.songState ? 'miniPlayer-icon-play' : 'miniPlayer-icon-pause'
    }
  },
  methods: {
    // 转换 MINI player
    handleMiniPlayer() {
      this.$store.dispatch("setFullScreen", false);
      this.showMiniScreen()
    },
    // 转换 PLAYER
    handleToPlayer() {
      this.$store.dispatch("setFullScreen", true);
      this.showFullScreen()
    },
    // 显示播放页
    showFullScreen() {
      this.$refs.player.style.zIndex = "150"
      this.$refs.normalPlayer.style.zIndex = "150"
    },
    // 隐藏播放页
    showMiniScreen() {
      this.$refs.player.style.zIndex = "-1"
      this.$refs.normalPlayer.style.zIndex = "-1"
    },
    // 下一曲
    nextSong() {
      var i = this.currentIndex + 1
      if (i < this.playList.length) {
        this.$store.dispatch('setCurrentIndex',i)
      } 
    },
    // 上一曲
    prevSong() {
      var i = this.currentIndex - 1
      if (i >= 0) {
        this.$store.dispatch('setCurrentIndex',i)
      } 
    },
    //  动画钩子
    // JS 写 CSS3 动画 三方库 https://github.com/HenrikJoreteg/create-keyframe-animation
    enter(el, done) {
      // 回调函数 done 执行时 跳到 afterEnter()钩子函数
      // 1. 获取坐标及缩放比例
      const { x, y, scale } = this.getPosAndScale();
      // 2. 定义动画
      let animation = {
        0: {
          transform: `translate3d(${x}px,${y}px,0) scale(${scale})`,
        },
        60: {
          transform: `translate3d(0,0,0) scale(1.1)`,
        },
        100: {
          transform: `translate3d(0,0,0) scale(1)`,
        },
      };
      // 3. 注册动画
      animations.registerAnimation({
        name: "move", // 动画名称
        animation, // 上边定义的动画
        presets: {
          duration: 400,
          easing: "line",
        },
      });
      // 4. 执行动画
      animations.runAnimation(this.$refs.cdWrapper, "move", done);
    },
    afterEnter() {
      animations.unregisterAnimation("move");
      this.$refs.cdWrapper.style.animation = "";
    },
    leave(el, done) {
      // 回调函数 done 执行时 跳到 leaveEnter()钩子函数
      this.$refs.cdWrapper.style.transition = "all 0.4s";
      // const { x, y, scale } = this.getPosAndScale();
    },
    afterLeave() {},
    // 获取唱片图片 相对 屏幕位置
    getPosAndScale() {
      const targetWidth = 40; // mini 图片宽
      const paddingLeft = 40; // mini 图片中心到屏幕左边距
      const paddingBottom = 30; // mini 图片中心到屏幕底边距
      const paddingTop = 80; // 全屏图片容器到屏幕顶边距
      const width = window.innerWidth * 0.8; // window.innerWidth 屏幕宽度 的 0.8 因为图片容器的宽度为 80%
      const scale = targetWidth / width; //  初始缩放比例 从mini -> width = 40 缩放到 屏幕的 80%宽度 的比例
      const x = -(window.innerWidth / 2 - paddingLeft); // 初始 X 坐标 mini 状态 X 坐标 相对于全屏图片的 X 坐标 因此为负数
      const y = window.innerHeight - paddingTop - width / 2 - paddingBottom;
      return {
        x,
        y,
        scale,
      };
    },
  },
};
</script>

<style scoped>
.player .normalPlayer {
  position: fixed;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  z-index: 150;
  background-color: #222;
  color: #fff;
}
.background {
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
  opacity: 0.6;
  filter: blur(20px);
}
.top {
  position: relative;
  margin-bottom: 25px;
}
.back {
  position: absolute;
  z-index: 50;
  color: #ffcd32;
  padding: 10px;
}
.iconBack {
  display: block;
  padding: 9px;
  font-size: 22px;
  color: #ffcd32;
  transform: rotate(-90deg);
}
.title {
  width: 70%;
  margin: 0 auto;
  line-height: 40px;
  text-align: center;
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
  font-size: 18px;
  color: #fff;
}
.subtitle {
  line-height: 20px;
  text-align: center;
  font-size: 14px;
  color: #fff;
}
.middle {
  position: fixed;
  width: 100%;
  top: 80px;
  bottom: 170px;
  white-space: nowrap;
  font-size: 0;
}
.middle-l {
  display: inline-block;
  vertical-align: top;
  position: relative;
  width: 100%;
  height: 0;
  padding-top: 80%;
}
.cd-wrapper {
  position: absolute;
  left: 10%;
  top: 0;
  width: 80%;
  height: 100%;
}
.cd {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  border: 10px solid rgba(255, 255, 255, 0.1);
  border-radius: 50%;
}
.play {
  animation: rotate 20s linear infinite;
}
.pause {
  animation-play-state: paused;
}
.image {
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  border-radius: 50%;
}
.playing-lyric-wrapper {
  width: 80%;
  margin: 30px auto 0 auto;
  overflow: hidden;
  text-align: center;
}
.playing-lyric {
  height: 20px;
  line-height: 20px;
  font-size: 14px;
  color: #222;
}
.middle-r {
  display: inline-block;
  vertical-align: top;
  width: 100%;
  height: 100%;
  overflow: hidden;
}
.lyric-wrapper {
  width: 80%;
  margin: 0 auto;
  overflow: hidden;
  text-align: center;
}
.text {
  line-height: 32px;
  color: rgba(255, 255, 255, 0.5);
  font-size: 14px;
}
.text .current {
  color: #222;
}
.bottom {
  position: absolute;
  bottom: 50px;
  width: 100%;
}
.dot-wrapper {
  text-align: center;
  font-size: 0;
}
.dot {
  display: inline-block;
  vertical-align: middle;
  margin: 0 4px;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.5);
}
.dot .active {
  width: 20px;
  border-radius: 5px;
  background: rgba(255, 255, 255, 0.8);
}
.progress-wrapper {
  display: flex;
  align-items: center;
  width: 80%;
  margin: 0px auto;
  padding: 10px 0;
}
.time {
  color: #222;
  font-size: 12px;
  flex: 0 0 30px;
  line-height: 30px;
  width: 30px;
}
.time .time-l {
  text-align: left;
}
.time .time-r {
  text-align: right;
}
.progress-bar-wrapper {
  flex: 1;
}
.operators {
  display: flex;
  align-items: center;
}
.icon {
  flex: 1;
  color: #ffcd32;
}
.icon .disable {
  color: rgba(255, 205, 49, 0.5);
}
.icon .disable i {
  font-size: 30px;
}
.i-left {
  text-align: right;
}
.i-center {
  padding: 0 20px;
  text-align: center;
}
.i-center i {
  font-size: 40px;
}
.i-right {
  text-align: left;
}
.icon-favorite {
  color: #ffcd32;
}
.miniPlayer {
  display: flex;
  align-items: center;
  position: fixed;
  left: 0;
  bottom: 0;
  z-index: 180;
  width: 100%;
  height: 60px;
  background: #333;
  color: #fff;
}
.miniPlayer .mini-enter-active,
.mini-leave-active {
  transition: all 0.4s;
}
.miniPlayer.mini-enter,
.mini-leave-to {
  opacity: 0;
}
.miniPlayer-icon {
  flex: 0 0 40px;
  width: 40px;
  padding: 10px;
}
.miniPlayer-icon-play {
  animation: rotate 10s linear infinite;
}
.miniPlayer-icon-pause {
  animation-play-state: paused;
}
.miniPlayer-text {
  display: flex;
  flex-direction: column;
  justify-content: center;
  flex: 1;
  overflow: hidden;
  line-height: 20px;
}
.miniPlayer-name {
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
  font-size: 16px;
  color: #fff;
  padding: 3px;
}
.miniPlayer-name-desc {
  padding: 0 3px;
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
  font-size: 12px;
  color: rgba(255, 255, 255, 0.3);
}

.miniPlayer-control {
  flex: 0 0 30px;
  width: 30px;
  padding: 0 10px;
}

.miniPlayer-icon-play-mini,
.miniPlayer-icon-pause-mini,
.miniPlayer-icon-playlist {
  font-size: 30px;
  color: rgba(255, 205, 49, 0.5);
}

.miniPlayer-icon-mini {
  font-size: 32px;
  position: absolute;
  left: 0;
  top: 0;
}

.normal-enter-active,
.normal-leave-active {
  transition: all 0.4s;
}
.normal-enter-active,
.normal-leave-active .top,
.bottom {
  transition: all 0.4s cubic-bezier(0.86, 0.18, 0.82, 1.32);
}
.normal-enter,
.normal-leave-to {
  opacity: 0;
}
.normal-enter,
.normal-leave-to .top {
  transform: translate3d(0, -60px, 0);
}
.normal-enter,
.normal-leave-to .bottom {
  transform: translate3d(0, 60px, 0);
}
.mini-enter-active,
.mini-leave-active {
  transition: all 0.4s;
}
.mini-enter,
.mini-leave-to {
  opacity: 0;
}
</style>
