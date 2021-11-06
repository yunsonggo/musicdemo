<template>
  <div>
    <Mheader :websiteInfo="websiteInfo"></Mheader>
    <Mtabbar :menus="menus"></Mtabbar>
    <router-view></router-view>
    <MPlayer></MPlayer>
  </div>
</template>

<script>
import Mheader from "@/components/header/mHeader.vue";
import Mtabbar from "@/components/tabbar/mTabbar.vue";
import { reqWebsite,reqMenus} from "@/api";
import MPlayer from "@/components/musicPlayer/mPlayer.vue"
export default {
  name: "Mhome",
  components: {
    Mheader,
    Mtabbar,
    MPlayer,
  },
  data() {
    return {
      websiteInfo: {},
      menus: [],
    };
  },
  mounted() {
    this.getWebsiteInfo();
    this.getMenus();
  },
  methods: {
    async getWebsiteInfo() {
      const result = await reqWebsite();
      if (result.code === 0) {
        this.$store.dispatch("setWebsite", result.data);
        this.websiteInfo = result.data;
      }
    },
    async getMenus() {
      const result = await reqMenus();
      if (result.code === 0) {
        this.$store.dispatch("setMenu", result.data);
        this.menus = result.data;
      }
    },
  },
};
</script>

<style scoped></style>
