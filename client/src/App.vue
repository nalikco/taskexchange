<script setup>
import { RouterLink, RouterView } from 'vue-router'
import NotificationsPopup from "@/components/NotificationsPopup.vue";
</script>

<template>
  <div class="min-h-full">
    <div class="absolute right-5 top-5 z-50">
      <transition-group name="slide-fade">
          <div v-for="(alert, i) in alerts" v-bind:key="i" @click="alerts.splice(i, 1)" class="p-4 mb-4 cursor-pointer shadow-lg text-sm rounded-lg bg-opacity-95" :class="[alert.classes]">
            <span class="font-medium">{{ alert.title }}</span>
            {{ alert.message }}
          </div>
      </transition-group>
    </div>

    <nav class="bg-gray-800">
      <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div class="flex h-16 items-center justify-between">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <img class="h-8 w-8" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500" alt="Tasks Exchange">
            </div>
            <div class="hidden md:block">
              <div class="ml-10 flex items-baseline space-x-4">
                <RouterLink v-for="elem in topMenu" :to="{'name': elem.to}" class="rounded-md text-sm font-medium px-3 py-2 transition duration-300" :class="{'bg-gray-900 text-white': $route.name === elem.to, 'text-gray-300 hover:bg-gray-700 hover:text-white': $route.name !== elem.to}">{{ elem.title }}</RouterLink>
              </div>
            </div>
          </div>
          <div class="hidden md:block">
            <div class="ml-4 flex items-center md:ml-6">
              <button v-if="user.type !== 0" @click="toggleNotifications" type="button" class="rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800" id="user-notifications-button" aria-expanded="false" aria-haspopup="true">
                <span class="sr-only">View notifications</span>
                <!-- Heroicon name: outline/bell -->
                <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" id="user-notifications-svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0" />
                </svg>
              </button>

              <NotificationsPopup
                  v-if="user.type !== 0"
                  :show-notifications="showNotifications"
                  :events="events" is-mobile="false"
                  @closePopup="showNotifications = false"
                  @updateEvents="getNewEvents" />

              <div class="relative ml-3">
                <div>
                  <button @click="toggleProfileMenu" type="button" class="flex max-w-xs items-center rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800" id="user-menu-button" aria-expanded="false" aria-haspopup="true">
                    <span class="sr-only">Open user menu</span>
                    <img class="h-8 w-8 rounded-full" src="@/assets/img/user.png" alt="User">
                  </button>
                </div>

                <transition
                    enter-active-class="transition ease-out duration-100"
                    enter-from-class="transform opacity-0 scale-95"
                    enter-to-class="transform opacity-100 scale-100"
                    leave-active-class="transition ease-in duration-75"
                    leave-from-class="transform opacity-100 scale-100"
                    leave-to-class="transform opacity-0 scale-95"
                >
                <div v-if="showProfileMenu" class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none" role="menu" aria-orientation="vertical" aria-labelledby="user-menu-button" tabindex="-1">
                  <div v-if="user.type !== 0" class="text-sm text-gray-700 px-4 py-2 border-b-2 border-gray-200">
                    {{ user.username }}
                    <br>
                    <span class="text-gray-500">{{ $filters.currencyFormat(user.balance) }}</span>
                  </div>
                  <RouterLink v-for="elem in profileMenu" :to="{'name': elem.to}" @click="showProfileMenu = false" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-200 transition duration-300" role="menuitem" tabindex="-1" id="user-menu-item-0">{{ elem.title }}</RouterLink>
                  <a href="#" v-if="user.type !== 0" @click="logout" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-200 transition duration-300" role="menuitem" tabindex="-1" id="user-menu-item-0">Выйти</a>
                </div>
                </transition>
              </div>
            </div>
          </div>
          <div class="-mr-2 flex md:hidden">
            <!-- Mobile menu button -->
            <button type="button" @click="showMobileMenu = !showMobileMenu" class="inline-flex items-center justify-center rounded-md bg-gray-800 p-2 text-gray-400 hover:bg-gray-700 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800" aria-controls="mobile-menu" aria-expanded="false">
              <span class="sr-only">Open main menu</span>
              <svg class="block h-6 w-6" :class="{'hidden': showMobileMenu, 'block': !showMobileMenu}" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
              </svg>
              <svg class="h-6 w-6" :class="{'hidden': !showMobileMenu, 'block': showMobileMenu}" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>
      </div>

      <transition
          enter-active-class="transition ease-out duration-100"
          enter-from-class="transform opacity-0 scale-95"
          enter-to-class="transform opacity-100 scale-100"
          leave-active-class="transition ease-in duration-75"
          leave-from-class="transform opacity-100 scale-100"
          leave-to-class="transform opacity-0 scale-95"
      >
        <div v-if="showMobileMenu" class="md:hidden" id="mobile-menu">
          <div class="space-y-1 px-2 pt-2 pb-3 sm:px-3">
            <RouterLink v-for="elem in topMenu" :to="{'name': elem.to}" @click="showMobileMenu = false" class="block px-3 py-2 rounded-md text-base font-medium" :class="{'bg-gray-900 text-white': $route.name === elem.to, 'text-gray-300 hover:bg-gray-700 hover:text-white': $route.name !== elem.to}">{{ elem.title }}</RouterLink>
          </div>
          <div class="border-t border-gray-700 pt-4 pb-3">
            <div class="flex items-center px-5">
              <div class="flex-shrink-0">
                <img class="h-10 w-10 rounded-full" src="@/assets/img/user.png" alt="User">
              </div>
              <div v-if="user.type !== 0" class="ml-3">
                <div class="text-base font-medium leading-none text-white">{{ user.username }}</div>
                <div class="text-sm mt-1 font-medium leading-none text-gray-400">{{ $filters.currencyFormat(user.balance) }}</div>
              </div>
              <button v-if="user.type !== 0" @click="toggleNotifications" type="button" class="ml-auto flex-shrink-0 rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800">
                <span class="sr-only">View notifications</span>
                <!-- Heroicon name: outline/bell -->
                <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0" />
                </svg>
              </button>
              <NotificationsPopup
                  v-if="user.type !== 0"
                  :show-notifications="showNotifications"
                  :events="events"
                  is-mobile="true"
                  @closePopup="showNotifications = false"
                  @updateEvents="getNewEvents"
              />
            </div>
            <div class="mt-3 space-y-1 px-2">
              <RouterLink v-for="elem in profileMenu" :to="{'name': elem.to}" @click="showMobileMenu = false" class="block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white">{{ elem.title }}</RouterLink>
              <a href="#" v-if="user.type !== 0" @click="logout" class="block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white">Выйти</a>
            </div>
          </div>
        </div>
      </transition>
    </nav>

    <router-view v-slot="{ Component }">
      <transition name="slide-fade" mode="out-in">
        <component :is="Component" />
      </transition>
    </router-view>

    <div class="mx-auto max-w-7xl text-center text-sm text-gray-500 border-t-2 py-7">
      &copy; 2022, Биржа работ. Все права защищены.
    </div>
  </div>
</template>

<script>
import {emitter} from "@/emitter";
import {mapActions, mapState} from "pinia";
import {useUserStore} from "@/stores/user";
import axios from "axios";

export default {
  data() {
    return {
      topMenu: [
        {'to': 'home', 'title': 'Главная'},
      ],
      alerts: [],
      alertTimers: [],
      showMobileMenu: false,
      showProfileMenu: false,
      showNotifications: false,
      e: emitter,
      events: []
    }
  },
  computed: {
    profileMenu() {
      if (this.user.type === 0) {
        return [
          {'to': 'sign-in', 'title': 'Войти'},
          {'to': 'sign-up', 'title': 'Зарегистрироваться'},
        ]
      } else if (this.user.type === 1) {
        return [
          {'to': 'cabinet', 'title': 'Кабинет исполнителя'},
          {'to': 'home', 'title': 'Мои задачи'},
        ]
      } else if (this.user.type === 2) {
        return [
          {'to': 'cabinet', 'title': 'Кабинет заказчика'},
          {'to': 'create-task', 'title': 'Добавить задачу'},
          {'to': 'tasks-list', 'title': 'Мои задачи'},
          {'to': 'home', 'title': 'Мои заказы'},
        ]
      } else if (this.user.type === 3) {
        return [
          {'to': 'home', 'title': 'Панель администратора'},
        ]
      }
    },
    ...mapState(useUserStore, ['token', 'user', 'eventsAfter']),
  },
  mounted() {
    this.e.on('alert', (alert, e) => {
      this.addAlert(alert.title, alert.message, alert.alertType)
    })

    this.e.on('updateUser', (redirect = false, e) => {
      this.updateUser(redirect)
    })

    this.updateUser(false)

    this.eventsPolling()
  },
  methods: {
    ...mapActions(useUserStore, ['setToken', 'setUser', 'setEventsAfter']),
    toggleNotifications() {
      this.showNotifications = !this.showNotifications
      this.showProfileMenu = false
    },
    toggleProfileMenu() {
      this.showProfileMenu = !this.showProfileMenu
      this.showNotifications = false
    },
    addAlert(title, message, alertType) {
      let alertClasses = ''

      switch(alertType){
        case 1:
          alertClasses = 'text-green-700 bg-green-100 dark:bg-green-200 dark:text-green-800'
          break;
        case 3:
          alertClasses = 'text-blue-700 bg-blue-100 dark:bg-blue-200 dark:text-blue-800'
          break;
        default:
          alertClasses = 'text-red-700 bg-red-100 dark:bg-red-200 dark:text-red-800'
          break;
      }

      let alert = {
        title: title,
        message: message,
        classes: alertClasses
      }

      this.alerts.splice(this.alerts.length - 1, 0, alert)

      this.alertTimers.push(setTimeout(() => {
        this.alerts.splice(this.alerts.length - 1, 1)
      }, 3000))
    },
    updateUser(redirect = false) {
      if(this.token !== '') {
        axios.get(import.meta.env.VITE_API_URL + 'users/me', {
          headers: { Authorization: `Bearer ${this.token}` }
        }).then(res => {
          this.setUser(res.data.data)
          this.getNewEvents()

          if(redirect) this.e.emit('redirectAfterLogin')
        })
      }
    },
    logout(e) {
      e.preventDefault()

      this.setUser({type: 0})
      this.setToken('')

      this.showProfileMenu = false
      this.showMobileMenu = false
      this.showNotifications = false

      if(this.$route.meta.requiredAuth) {
        return this.$router.push({ name: 'sign-in' })
      }

      if(this.$route.meta.requiredCustomer) {
        return this.$router.push({ name: 'sign-in' })
      }
    },
    getNewEvents() {
      axios.get(import.meta.env.VITE_API_URL + 'events/new', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if (res.data.data && res.data.data.length > 0) {
          this.events = []

          for (let i = 0; i < res.data.data.length; i++){
            this.events.push({
              id: res.data.data[i].id,
              message: res.data.data[i].message,
              link: res.data.data[i].link,
              created_at: res.data.data[i].created_at
            })
          }
        }
      })
    },
    eventsPolling() {
      axios.get(import.meta.env.VITE_API_URL + 'events/polling?after=' + this.eventsAfter, {
        headers: { Authorization: `Bearer ${this.token}` },
        timeout: 25000,
      }).then(res => {
        if(res.data.data.length > 0) {
          this.setEventsAfter(res.data.data[0].id)
          this.getNewEvents()

          for(let i = 0; i < res.data.data.length; i++) {
            this.addAlert('Новое уведомление:', res.data.data[i].message, 3)
          }
        }

        this.eventsPolling()
      }).catch(err => {
        setTimeout(this.eventsPolling, 10000)
      })
    }
  }
}
</script>

<style>
.slide-fade-enter-active {
  transition: all 0.5s ease-out;
}
.slide-fade-leave-active {
  transition: all 0.5s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(50px);
  opacity: 0;
}
</style>