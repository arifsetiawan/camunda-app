import Vue from 'vue'
import Router from 'vue-router'
import store from "../store";

// Containers
const TheContainer = () => import('@/containers/TheContainer')

// Nexus
const Home = () => import('@/views/Home')
const Tasks = () => import('@/views/tasks/Tasks')
const LeaveRequestApproval = () => import('@/views/forms/LeaveRequestApproval')
const Requests = () => import('@/views/tasks/Requests')
const LeaveRequest = () => import('@/views/forms/LeaveRequest')
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
          component: Home,
        },
        {
          path: 'tasks',
          name: 'Tasks',
          component: Tasks,
        },
        {
          path: 'tasks/:taskId',
          name: 'Tasks Detail',
          component: LeaveRequestApproval,
        },
        {
          path: 'requests',
          name: 'Requests',
          component: Requests,
        },
        {
          path: 'forms/leave',
          name: 'Leave Request',
          component: LeaveRequest,
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

