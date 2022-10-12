<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Вход</h1>
      </div>
    </header>
    <main>
      <div class="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
        <div class="w-full max-w-md space-y-8">
          <div>
            <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">Войдите в свой аккаунт</h2>
            <p class="mt-2 text-center text-sm text-gray-600">
              Или
              <RouterLink :to="{'name': 'sign-up'}" class="font-medium text-blue-600 hover:text-blue-500">зарегистрируйте новый</RouterLink>
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
                <label for="password" class="sr-only">Password</label>
                <input id="password" name="password" v-model="password" type="password" required class="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-blue-500 focus:outline-none focus:ring-blue-500 sm:text-sm" placeholder="Пароль (qwerty)">
              </div>
            </div>

            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <input id="remember-me" v-model="rememberMe" name="remember-me" type="checkbox" class="h-4 w-4 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <label for="remember-me" class="ml-2 block font-medium text-sm text-gray-900">Запомнить меня</label>
              </div>

              <div class="text-sm">
                <RouterLink :to="{'name': 'sign-in'}" class="font-medium text-blue-600 hover:text-blue-500">Забыли пароль?</RouterLink>
              </div>
            </div>

            <div>
              <button type="submit" class="group relative flex w-full justify-center rounded-md border border-transparent bg-blue-600 py-2 px-4 text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2">
                <span class="absolute inset-y-0 left-0 flex items-center pl-3">
                  <svg class="h-5 w-5 text-blue-500 group-hover:text-blue-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                    <path fill-rule="evenodd" d="M10 1a4.5 4.5 0 00-4.5 4.5V9H5a2 2 0 00-2 2v6a2 2 0 002 2h10a2 2 0 002-2v-6a2 2 0 00-2-2h-.5V5.5A4.5 4.5 0 0010 1zm3 8V5.5a3 3 0 10-6 0V9h6z" clip-rule="evenodd" />
                  </svg>
                </span>
                Войти
              </button>
            </div>
          </form>
        </div>
      </div>
    </main>
  </main>
</template>

<script>
import {mapActions} from "pinia";
import {useUserStore} from "@/stores/user";
import axios from "axios";
import {emitter} from "@/emitter";
import NProgress from "nprogress";

export default {
  data() {
    return {
      email: '',
      password: '',
      rememberMe: false,
      e: emitter
    }
  },
  mounted() {
    NProgress.done()
    document.title = 'Вход'

    this.e.on('redirectAfterLogin', e => {
      if(this.$route.query.redirect) this.$router.push(this.$route.query.redirect)
      else this.$router.push({ name: 'cabinet' })
    })
  },
  methods: {
    ...mapActions(useUserStore, ['setToken']),
    onSubmit(e) {
      e.preventDefault()

      NProgress.start()
      axios.post(import.meta.env.VITE_API_URL + "auth/sign-in", {
        email: this.email,
        password: this.password,
      }).then(res => {
        if(res.data.token) {
          this.setToken(res.data.token)

          this.e.emit('alert', {
            title: 'Успешно!',
            message: 'Вы вошли.',
            alertType: 1
          })
          this.e.emit('updateUser', true)

        } else this.e.emit('alert', {
          title: 'Ошибка!',
          message: 'Произошла ошибка сервера. Повторите позже.',
          alertType: 0
        })
      }).catch(err => {
        if(err.response.data.message && err.response.data.message === 'sql: no rows in result set') {
          this.e.emit('alert', {
            title: 'Ошибка!',
            message: 'Неверный адрес электронной почты или пароль.',
            alertType: 0
          })
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