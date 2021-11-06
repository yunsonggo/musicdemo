<template>
  <div class="mtabbar">
    <van-tabbar v-model="active" active-color="#ee0a24" :fixed="false" @change="onChange">
      <van-tabbar-item v-for="(menu,index) in menus" :key="index" style="font-size:14px">{{menu.title}}</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script>
export default {
  name: "mTabbar",
  props: {
    menus: {
      type: Array,
      default: () => [],
    },
  },
  data() {
      return {
          active:0
      }
  },
  methods:{
      onChange(index) {
          console.log(this.menus[index].path);
          var toPath = "/m/home" + this.menus[index].path
          this.$router.push(toPath)
      }
  }
};
</script>

<style scoped>
.mtabbar {
  position: fixed;
  width: 100%;
  top:60px;
  z-index: 100;
}
</style>
