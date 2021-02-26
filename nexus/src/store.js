import Vue from 'vue'
import Vuex from 'vuex'
import user from "./store/modules/user";
import auth from "./store/modules/auth";
import sidebar from "./store/modules/sidebar";
Vue.use(Vuex)

/*
const state = {
  sidebarShow: 'responsive',
  sidebarMinimize: false
}

const mutations = {
  toggleSidebarDesktop (state) {
    const sidebarOpened = [true, 'responsive'].includes(state.sidebarShow)
    state.sidebarShow = sidebarOpened ? false : 'responsive'
  },
  toggleSidebarMobile (state) {
    const sidebarClosed = [false, 'responsive'].includes(state.sidebarShow)
    state.sidebarShow = sidebarClosed ? true : 'responsive'
  },
  set (state, [variable, value]) {
    state[variable] = value
  }
}
*/

export default new Vuex.Store({
  modules: {
    user,
    auth,
    sidebar
  }
})