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
      meta: { requiredAuth: true }
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
    }, {
      path: '/messages',
      name: 'messages',
      component: () => import('../views/Messages/ListView.vue'),
      meta: { requiredAuth: true }
    }, {
      path: '/payments',
      name: 'payments',
      component: () => import('../views/Payments.vue'),
      meta: { requiredAuth: true }
    }, {
      path: '/@:username',
      name: 'profile',
      component: () => import('../views/Profile.vue'),
      meta: { requiredAuth: true }
    }, {
      path: '/posts',
      name: 'posts',
      component: () => import('../views/Posts/ListView.vue'),
    }, {
      path: '/posts/:id',
      name: 'posts-full',
      component: () => import('../views/Posts/FullView.vue'),
    },

    {
      path: '/ap',
      name: 'ap',
      component: () => import('../views/Admin/Index.vue'),
      meta: { requiredAdmin: true }
    }, {
      path: '/ap/users',
      name: 'ap-users',
      component: () => import('../views/Admin/Users.vue'),
      meta: { requiredAdmin: true }
    }, {
      path: '/ap/options',
      name: 'ap-options',
      component: () => import('../views/Admin/Options.vue'),
      meta: { requiredAdmin: true }
    }, {
      path: '/ap/tasks',
      name: 'ap-tasks',
      component: () => import('../views/Admin/Tasks.vue'),
      meta: { requiredAdmin: true }
    }, {
      path: '/ap/orders',
      name: 'ap-orders',
      component: () => import('../views/Admin/Orders.vue'),
      meta: { requiredAdmin: true }
    }, {
      path: '/ap/posts',
      name: 'ap-posts',
      component: () => import('../views/Admin/Posts/ListView.vue'),
      meta: { requiredAdmin: true }
    }, {
      path: '/ap/posts/create',
      name: 'ap-posts-create',
      component: () => import('../views/Admin/Posts/CreateView.vue'),
      meta: { requiredAdmin: true }
    }, {
      path: '/ap/posts/categories',
      name: 'ap-posts-categories',
      component: () => import('../views/Admin/Posts/CategoriesView.vue'),
      meta: { requiredAdmin: true }
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
  } else if(to.matched.some(record => record.meta.requiredAdmin)) {
    const store = useUserStore()

    if(store.user.type !== 3) {
      next({
        name: 'home',
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
