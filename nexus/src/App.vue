<template>
  <router-view></router-view>
</template>

<script>
import axios from 'axios';

export default {
  name: 'App',
  created: function () {
  axios.interceptors.response.use(undefined, function (err) {
    return new Promise(function (resolve, reject) {
      if (err.status === 401 && err.config && !err.config.__isRetryRequest) {
        // if you ever get an unauthorized, logout the user
        this.$store.dispatch(AUTH_LOGOUT)
        // you can also redirect to /login if needed !
      }
      throw err;
    });
  });
}
}
</script>

<style lang="scss">
  // Import Main styles for this application
  @import 'assets/scss/style';
</style>
