import { createRouter, createWebHistory } from 'vue-router'
import {useUserStore} from "@/stores/user";

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
  } else {
    next()
  }
})

export default router
