<template>
  <div>
    <div v-if="menus">
      <el-menu
        :default-active="activeIndex"
        class="el-menu-demo"
        mode="horizontal"
        @select="handleSelect"
        background-color="#222"
        text-color="#fff"
        active-text-color="#ffd04b"
      >
        <el-menu-item
          v-for="(menu, i) in menus"
          :key="i"
          :index="String(menu.sort)"
          >{{ menu.title }}</el-menu-item
        >
      </el-menu>
    </div>
    <div v-else>
      <el-menu
        :default-active="activeIndex"
        class="el-menu-demo"
        mode="horizontal"
        @select="handleSelect"
        background-color="#222"
        text-color="#fff"
        active-text-color="#ffd04b"
      >
        <el-menu-item index="1">loading...</el-menu-item>
      </el-menu>
    </div>
  </div>
</template>

<script>
export default {
  name: "tabbar",
  data() {
    return {
      activeIndex: "1",
    };
  },
  props: {
    menus: {
      type: Array,
      default: () => [],
    },
  },
  methods: {
    handleSelect(key) {
      var index = key -1
      var toPath = '/home' + this.menus[index].path;
      this.$router.push(toPath);
    },
  },
};
</script>

<style scoped></style>
