<template>
  <div class="home">
    <Header :websiteInfo="websiteInfo"></Header>
    <Tabbar :menus="menus"></Tabbar>
    <router-view></router-view>
  </div>
</template>

<script>
import Header from "@/components/header/header.vue";
import Tabbar from "@/components/tabbar/tabbar.vue"
import { reqWebsite,reqMenus} from "@/api";
export default {
  name: "home",
  components: {
    Header,
    Tabbar
  },
  data() {
    return {
      websiteInfo: {},
      menus:[],
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
          this.$store.dispatch("setWebsite",result.data)
          this.websiteInfo = result.data
      }
    },
    async getMenus() {
        const result = await reqMenus();
        if (result.code === 0) {
            this.$store.dispatch("setMenu",result.data)
            this.menus = result.data
        }
    }
  },
};
</script>

<style scoped>

</style>
