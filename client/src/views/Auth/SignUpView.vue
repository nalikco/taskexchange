<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Регистрация</h1>
      </div>
    </header>
    <main>
      <div class="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
        <div class="w-full max-w-lg space-y-8">
          <div>
            <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">Зарегистрируйте аккаунт</h2>
            <p class="mt-2 text-center text-sm text-gray-600">
              Или
              <RouterLink :to="{'name': 'sign-in'}" class="font-medium text-blue-600 hover:text-blue-500">войдите, если он уже есть</RouterLink>
            </p>
          </div>
          <form class="mt-8 space-y-6" @submit="onSubmit" method="POST">
            <input type="hidden" name="remember" value="true">
            <div class="-space-y-px rounded-md shadow-sm">
              <div>
                <label for="email-address" class="sr-only">Email address</label>
                <input id="email-address" name="email" v-model="email" type="email" required class="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-blue-500 focus:outline-none focus:ring-blue-500 sm:text-sm" placeholder="Адрес электронной почты (example@gmail.com)">
              </div>
              <div>
                <label for="username" class="sr-only">Email address</label>
                <input id="username" name="username" v-model="username" type="text" required class="relative block w-full appearance-none rounded-none border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-blue-500 focus:outline-none focus:ring-blue-500 sm:text-sm" placeholder="Имя (Илья Васильев)">
              </div>
              <div>
                <label for="password" class="sr-only">Password</label>
                <input id="password" name="password" v-model="password" type="password" required class="relative block w-full appearance-none rounded-none border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-blue-500 focus:outline-none focus:ring-blue-500 sm:text-sm" placeholder="Пароль (qwerty)">
              </div>
              <div>
                <label for="r_password" class="sr-only">Password</label>
                <input id="r_password" name="r_password" v-model="r_password" type="password" required class="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-blue-500 focus:outline-none focus:ring-blue-500 sm:text-sm" placeholder="Повторите пароль (qwerty)">
              </div>
            </div>

            <div class="flex">
              <div class="flex items-center mr-4">
                <input checked id="inline-checked-radio" type="radio" v-model="accountType" value="1" name="type" class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
                <label for="inline-checked-radio" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Я исполнитель</label>
              </div>
              <div class="flex items-center mr-4">
                <input id="inline-2-radio" type="radio" v-model="accountType" value="2" name="type" class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
                <label for="inline-2-radio" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Я заказчик</label>
              </div>
            </div>

            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <input id="rules" v-model="isRulesChecked" name="rules" type="checkbox" class="h-4 w-4 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <label for="rules" class="ml-2 block text-sm font-medium text-gray-900">
                  Согласен(-а) на обработку персональных данных и правилами сайта
                </label>
              </div>
            </div>

            <div>
              <button :disabled="!isRulesChecked" type="submit" class="group relative flex w-full justify-center rounded-md border border-transparent bg-blue-600 py-2 px-4 text-sm font-medium text-white" :class="{'opacity-50': !isRulesChecked, 'hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2': isRulesChecked}">
                <span class="absolute inset-y-0 left-0 flex items-center pl-3">
                  <svg class="h-5 w-5 text-blue-500" :class="{'group-hover:text-blue-400': isRulesChecked}" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                    <path fill-rule="evenodd" d="M10 1a4.5 4.5 0 00-4.5 4.5V9H5a2 2 0 00-2 2v6a2 2 0 002 2h10a2 2 0 002-2v-6a2 2 0 00-2-2h-.5V5.5A4.5 4.5 0 0010 1zm3 8V5.5a3 3 0 10-6 0V9h6z" clip-rule="evenodd" />
                  </svg>
                </span>
                Зарегистрироваться
              </button>
            </div>
          </form>
        </div>
      </div>
    </main>
  </main>
</template>

<script>
import axios from "axios";
import {emitter} from "@/emitter";
import {mapActions} from "pinia";
import {useUserStore} from "@/stores/user";
import NProgress from "nprogress";

export default {
  data() {
    return {
      email: '',
      username: '',
      password: '',
      r_password: '',
      accountType: 1,
      isRulesChecked: false,
      e: emitter,
    }
  },
  mounted() {
    document.title = 'Регистрация'

    this.e.on('redirectAfterLogin', e => {
      this.$router.push({ name: 'cabinet' })
    })

    NProgress.done()
  },
  methods: {
    ...mapActions(useUserStore, ['setToken']),
    onSubmit(e) {
      e.preventDefault()

      if (this.password !== this.r_password) {
        this.e.emit('alert', {
          title: 'Ошибка!',
          message: 'Пароль повторен неверно.',
          alertType: 0
        })

        return
      }

      axios.post(import.meta.env.VITE_API_URL + "auth/sign-up", {
        username: this.username,
        email: this.email,
        password: this.password,
        type: parseInt(this.accountType)
      }).then(res => {
        axios.post(import.meta.env.VITE_API_URL + "auth/sign-in", {
          email: this.email,
          password: this.password,
        }).then(res => {
          if(res.data.token) {
            NProgress.start()

            this.setToken(res.data.token)

            this.e.emit('alert', {
              title: 'Успешно!',
              message: 'Вы зарегистрировались.',
              alertType: 1
            })
            this.e.emit('updateUser', true)

          } else this.e.emit('alert', {
            title: 'Ошибка!',
            message: 'Произошла ошибка сервера. Повторите позже.',
            alertType: 0
          })
        }).catch(err => {
          this.e.emit('alert', {
            title: 'Ошибка!',
            message: 'Произошла ошибка сервера. Повторите позже.',
            alertType: 0
          })
        })
      }).catch(err => {
        if(err.response.data.message) {
          switch(err.response.data.message) {
            case "not valid username":
              this.e.emit('alert', {
                title: 'Ошибка!',
                message: 'Укажите верное имя.',
                alertType: 0
              })
              break;
            case "email is already taken":
              this.e.emit('alert', {
                title: 'Ошибка!',
                message: 'Адрес электронной почты занят.',
                alertType: 0
              })
              break;
            default:
              this.e.emit('alert', {
                title: 'Ошибка!',
                message: 'Произошла ошибка сервера. Повторите позже.',
                alertType: 0
              })
              break;
          }
        } else this.e.emit('alert', {
          title: 'Ошибка!',
          message: 'Произошла ошибка сервера. Повторите позже.',
          alertType: 0
        })
      })
    }
  }
}
</script>