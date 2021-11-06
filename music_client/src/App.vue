<template>
  <keep-alive>
    <router-view />
  </keep-alive>
</template>

<script>
export default {
  mounted() {
    localStorage.removeItem("appType")
    if (this._isMobile()) {
      localStorage.setItem("appType","m")
      this.$router.replace("/m");
    } else {
      localStorage.setItem("appType","pc")
      this.$router.replace("/")
    }
  },
  methods: {
    _isMobile() {
      let flag = navigator.userAgent.match(
        /(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i
      );
      return flag;
    },
  },
};
</script>

<style>
html,
body {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  background-color: #fcfcfd;
  height: 100%;
}
p {
  margin: 0;
}
</style>
