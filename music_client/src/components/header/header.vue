<template>
  <div class="header">
    <span>
      <img v-if="websiteInfo.logo" :src="websiteInfo.static + websiteInfo.logo" class="header-logo" />
      <img v-else src="../../assets/logo.png" class="header-logo" />
    </span>
    <span v-if="websiteInfo">
      {{ websiteInfo.name }}
    </span>
  </div>
</template>

<script>
export default {
  name: "Header",
  props:{
      websiteInfo:{
          type:Object,
          default:() => ({})
      }
  }
};
</script>

<style scoped>
.header {
  background-color:#222;
  height: 55px;
  display: flex;
  align-items: center;
  justify-content: left;
}

.header span:nth-child(1) {
  margin-right: 10px;
  margin-left: 20px;
}

.header span:nth-child(2) {
  color: white;
  font-size: 14px;
}

.header-logo {
  width: 30px;
  height: 30px;
}</style>
