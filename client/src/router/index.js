import { createRouter, createWebHistory } from 'vue-router'
import {useUserStore} from "@/stores/user";
import NProgress from 'nprogress';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue')
    }, {
      path: '/sign-in',
      name: 'sign-in',
      component: () => import('../views/Auth/SignInView.vue'),
      meta: { requiredGuest: true }
    }, {
      path: '/sign-up',
      name: 'sign-up',
      component: () => import('../views/Auth/SignUpView.vue'),
      meta: { requiredGuest: true }
    }, {
      path: '/cabinet',
      name: 'cabinet',
      component: () => import('../views/CabinetView.vue'),
      meta: { requiredAuth: true }
    }, {
      path: '/events',
      name: 'events',
      component: () => import('../views/EventsView.vue'),
      meta: { requiredAuth: true }
    }, {
      path: '/tasks/create',
      name: 'create-task',
      component: () => import('../views/Tasks/CreateView.vue'),
      meta: { requiredCustomer: true }
    }, {
      path: '/tasks/my',
      name: 'tasks-my',
      component: () => import('../views/Tasks/MyView.vue'),
      meta: { requiredCustomer: true }
    }, {
      path: '/tasks/:id',
      name: 'edit-task',
      component: () => import('../views/Tasks/EditView.vue'),
      meta: { requiredCustomer: true }
    }, {
      path: '/tasks',
      name: 'tasks',
      component: () => import('../views/Tasks/ListView.vue'),
    }, {
      path: '/orders/customer',
      name: 'orders-customer',
      component: () => import('../views/Orders/CustomerView.vue'),
      meta: { requiredCustomer: true }
    }, {
      path: '/orders/performer',
      name: 'orders-performer',
      component: () => import('../views/Orders/PerformerView.vue'),
      meta: { requiredPerformer: true }
    },
  ]
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiredGuest)) {
    const store = useUserStore()

    if(store.user.type !== 0) {
      next({
        name: 'cabinet',
      })
    } else next()
  } else if(to.matched.some(record => record.meta.requiredAuth)) {
    const store = useUserStore()

    if(store.user.type === 0) {
      next({
        name: 'sign-in',
        query: { redirect: to.fullPath }
      })
    } else next()
  } else if(to.matched.some(record => record.meta.requiredCustomer)) {
    const store = useUserStore()

    if(store.user.type !== 2) {
      next({
        name: 'cabinet',
      })
    } else next()
  } else if(to.matched.some(record => record.meta.requiredPerformer)) {
    const store = useUserStore()

    if(store.user.type !== 1) {
      next({
        name: 'cabinet',
      })
    } else next()
  } else {
    next()
  }
})

router.beforeResolve((to, from, next) => {
  if (to.name) {
    NProgress.start()
  }
  next()
})

export default router
