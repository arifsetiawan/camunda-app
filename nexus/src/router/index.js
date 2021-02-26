import Vue from 'vue'
import Router from 'vue-router'
import store from "../store";

// Containers
const TheContainer = () => import('@/containers/TheContainer')

// Nexus
const NexusHome = () => import('@/views/Home')
const NexusTables = () => import('@/views/nexus/Tables')
const NexusLeaveRequest = () => import('@/views/nexus/LeaveRequest')
const Login = () => import('@/views/pages/Login')

Vue.use(Router)

const ifNotAuthenticated = (to, from, next) => {
  if (!store.getters.isAuthenticated) {
    next();
    return;
  }
  next("/");
};

const ifAuthenticated = (to, from, next) => {
  if (store.getters.isAuthenticated) {
    next();
    return;
  }
  next("/login");
};

export default new Router({
  mode: 'hash', // https://router.vuejs.org/api/#mode
  linkActiveClass: 'active',
  scrollBehavior: () => ({ y: 0 }),
  routes: configRoutes()
})

function configRoutes () {
  return [
    {
      path: '/',
      redirect: '/home',
      name: 'Home',
      component: TheContainer,
      children: [
        {
          path: 'home',
          name: 'Home',
          component: NexusHome,
        },
        {
          path: 'tasks',
          name: 'Tasks',
          component: NexusTables,
        },
        {
          path: 'workflow/leave',
          name: 'Leave Request',
          component: NexusLeaveRequest,
        }
      ],
      beforeEnter: ifAuthenticated  
    },
    {
      path: '/login',
      name: 'Login',
      component: Login,
      beforeEnter: ifNotAuthenticated
    }
  ]
}

