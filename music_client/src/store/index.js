import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const types = {
  SET_MENU:"SET_MENU",
  SET_CATEGORY:"SET_CATEGORY",
  SET_WEBSITE:"SET_WEBSITE",
  SET_PLAYING:"SET_PLAYING",
  SET_FULLSCREEN:"SET_FULLSCREEN",
  SET_PLAYLIST:"SET_PLAYLIST",
  SET_SEQUENCELIST:"SET_SEQUENCELIST",
  SET_PLAYMODE:"SET_PLAYMODE",
  SET_CURRENTINDEX:"SET_CURRENTINDEX",
};

const state = {
  menu:{},
  category:{},
  website:{},
  playing:false,    // 播放器状态 true 播放  false 暂停
  fullScreen:false, //  播放器是否全屏
  playList:[], // 播放歌曲默认列表 顺序播放
  sequenceList:[], // 播放歌曲随机列表 随机播放
  playMode:0, // 0 默认顺序播放 1 循环播放 2 随机播放  
  currentIndex:-1, // 当前播放曲目游标
};

const getters = {
  menu:(state) => state.menu,
  category:(state) => state.category,
  website:(state) => state.website,
  playing:(state) => state.playing,
  fullScreen:(state) => state.fullScreen,
  playList:(state) => state.playList,
  sequenceList:(state) => state.sequenceList,
  playMode:(state) => state.playMode, 
  currentIndex:(state) => state.currentIndex,
};

const mutations = {
  [types.SET_MENU](state,menu) {
    if (menu) {
      state.menu = menu;
    } else {
      state.menu = null;
    }
  },
  [types.SET_CATEGORY](state,category) {
    if (category) {
      state.category = category;
    } else {
      state.category = null;
    }
  },
  [types.SET_WEBSITE](state,website) {
    if (website) {
      state.website = website;
    } else {
      state.website = null;
    } 
  },
  [types.SET_PLAYING](state,playing) {
    if (playing == true) {
      state.playing = true
    } else {
      state.playing = false
    }
  },
  [types.SET_FULLSCREEN](state,fullScreen) {
    if (fullScreen == true) {
      state.fullScreen = true
    } else {
      state.fullScreen = false
    }
  },
  [types.SET_PLAYLIST](state,playList) {
    if (playList) {
      state.playList = playList
    } else {
      state.playList = null
    }
  },
  [types.SET_SEQUENCELIST](state,sequenceList) {
    if (sequenceList) {
      state.sequenceList = sequenceList
    } else {
      state.sequenceList = null
    }
  },
  [types.SET_PLAYMODE](state,playMode) {
    if (playMode > 0 ) {
      state.playMode = playMode
    } else {
      state.playMode = 0
    }
  },
  [types.SET_CURRENTINDEX](state,currentIndex) {
    if (currentIndex > -1) {
      state.currentIndex = currentIndex
    } else {
      state.currentIndex = -1
    }
  },
};

const actions = {
  setMenu:({commit},menu) => {
    commit(types.SET_MENU,menu);
  },
  setCategory:({commit},category) => {
    commit(types.SET_CATEGORY,category)
  },
  setWebsite:({commit},website) => {
    commit(types.SET_WEBSITE,website)
  },
  setPlaying:({commit},playing) => {
    commit(types.SET_PLAYING,playing)
  },
  setFullScreen:({commit},fullScreen) => {
    commit(types.SET_FULLSCREEN,fullScreen)
  },
  setPlayList:({commit},playList) => {
    commit(types.SET_PLAYLIST,playList)
  },
  setSequenceList:({commit},sequenceList) => {
    commit(types.SET_SEQUENCELIST,sequenceList)
  },
  setPlayMode:({commit},playMode) => {
    commit(types.SET_PLAYMODE,playMode)
  },
  setCurrentIndex:({commit},currentIndex) => {
    commit(types.SET_CURRENTINDEX,currentIndex)
  },
};

export default new Vuex.Store({
  state,
  getters,
  mutations,
  actions,
});

// this.$store.dispatch("setExample",data)
// this.$store.getters....