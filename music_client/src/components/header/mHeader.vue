<template>
  <div class="m-header">
    <span>
      <img v-if="websiteInfo.logo" :src="websiteInfo.static + websiteInfo.logo" class="m-header-logo" />
      <img v-else src="../../assets/logo.png" class="m-header-logo" />
    </span>
    <span v-if="websiteInfo">
      {{ websiteInfo.name }}
    </span>
  </div>
</template>

<script>
export default {
  name: "Mheader",
  props: {
    websiteInfo:{
          type:Object,
          default:() => ({})
      }
  },
};
</script>

<style scoped>
.m-header {
  background-color: #d43c33;
  height: 60px;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: left;
  color: white;
  font-size: 16px;
  position:fixed;
  top:0;
  z-index: 100;
}

.m-header span:nth-child(1) {
  margin-left:10px;
  margin-right: 5px;
}

.m-header-logo {
  width: 20px;
  height: 20px;
}
</style>
