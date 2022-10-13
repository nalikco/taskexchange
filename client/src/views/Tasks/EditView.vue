<script setup>
import ProfileInfo from "@/components/ProfileInfo.vue";
import {moment} from "@/moment";
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Редактирование задачи</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <ProfileInfo />
        <div>
          <h2 class="text-2xl font-medium mt-10">Укажите задачу</h2>
          <div class="bg-white mt-3 py-4 px-5 grid grid-cols-1 gap-4 md:grid-cols-4 rounded-xl shadow hover:shadow-lg transition duration-300">
            <div v-for="parent in optionsToShow" @mouseenter="showOptions = parent.id" @mouseleave="showOptions = 0" class="relative">
              <div @click="selectParent(parent)" class="bg-slate-100 cursor-pointer shadow hover:shadow-lg hover:ring ring-slate-100 rounded-lg py-1 px-1 transition duration-300" :class="{'text-white bg-blue-500': selectedParent === parent.id, 'rounded-t-lg': showOptions === parent.id}">
                <div class="py-3 px-4 font-medium text-center">
                  {{ parent.title }}
                  <p class="text-xs -mt-1" :class="{'text-white': selectedParent === parent.id, 'text-gray-500': selectedParent !== parent.id}">{{ $filters.currencyFormat(parent.price) }}</p>
                  <button v-if="selectedParent === parent.id" class="mt-2 text-white w-full text-sm">Выбрано</button>
                  <button v-else class="mt-2 text-gray-500 w-full text-sm">Выбрать</button>
                </div>
              </div>
              <transition name="slide-fade">
                <div v-if="showOptions === parent.id" class="absolute w-full z-50 ring ring-slate-200 bg-slate-200 shadow-lg rounded-b-lg pb-4 pt-2 px-3 text-sm flex flex-col">
                  <div v-if="parent.options.length === 0" class="mt-2 text-center text-slate-500">
                    Нет опций для данной категории
                  </div>
                  <div v-for="option in parent.options" @click="selectOption(option)" class="py-2 px-4 mt-2 cursor-pointer shadow hover:shadow-lg transition duration-300 bg-slate-100 rounded-full" :class="{'bg-blue-500 text-white shadow-lg': selectedOptions.indexOf(option.id) !== -1}">
                    {{ option.title }} <span :class="{'text-white': selectedOptions.indexOf(option.id) !== -1, 'text-slate-500': selectedOptions.indexOf(option.id) === -1}">{{ $filters.currencyFormat(option.price) }}</span>
                  </div>
                </div>
              </transition>
            </div>
          </div>
          <div class="mt-7 bg-white rounded-lg py-3 px-7 shadow">
            <form @submit="onFormSubmit">
              <div class="mt-4">
                <label for="link" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Ссылка на задачу</label>
                <input type="text" @input="onInputChange($event, 'link')" id="link" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Например, ссылка на пост ВКонтакте" :value="link" required>
              </div>
              <div class="mt-6">
                <label for="description" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Описание задачи</label>
                <textarea id="description" @input="onInputChange($event, 'description')" rows="6" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Описание задачи нужно для того, чтобы исполнители поняли суть задачи" :value="description" required></textarea>
              </div>
              <div class="mt-6">
                <label for="amount" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Количество</label>
                <input type="number" min="1" @input="onInputChange($event, 'amount')" id="amount" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Например, количество необходимых комментариев под пост ВКонтакте" :value="amount" required>
              </div>
              <div class="mt-6">
                <label for="amount" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Дата сдачи</label>
                <input type="date" :min="new Date().toLocaleDateString('en-ca')" @input="onInputChange($event, 'deliveryDate')" :value="deliveryDate" id="amount" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Например, количество необходимых комментариев под пост ВКонтакте" required>
              </div>
              <div class="mt-6">
                <transition name="slide-fade">
                  <div v-if="userBalance < priceForCurrentTask" class="mb-3 text-center font-medium text-sm text-red-500">
                    Недостаточно средств на балансе
                  </div>
                </transition>
                <button type="submit" :disabled="!showSubmitButton" class="text-white w-full bg-blue-600 transition duration-300 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2" :class="{'hover:bg-blue-800': showSubmitButton,'opacity-50': !showSubmitButton }">
                  Сохранить задачу
                  <br>
                  <p class="text-xs -mt-1 text-blue-200">
                    Итоговая стоимость данной задачи: {{ $filters.currencyFormat(priceForCurrentTask) }}
                  </p>
                </button>
                <RouterLink :to="{ name: 'tasks-list' }" class="flex flex-col w-full text-center px-auto bg-slate-200 transition duration-300 focus:ring-4 focus:outline-none focus:ring-slate-300 dark:focus:ring-slate-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2 hover:bg-slate-300">
                  К списку задач
                </RouterLink>
              </div>
            </form>
          </div>
        </div>
      </div>
    </main>
  </main>
</template>

<script>
import axios from 'axios';
import NProgress from "nprogress";
import {useUserStore} from "@/stores/user";
import {mapState} from "pinia";
import {emitter} from "@/emitter";

export default {
  data() {
    return {
      options: [],
      selectedParent: 0,
      showOptions: 0,
      link: "",
      description: "",
      amount: 1,
      selectedOptions: [],
      deliveryDate: "",
      loading: false,
      userBalance: 0,
      task: {},
      e: emitter
    }
  },
  computed: {
    showSubmitButton() {
      if(this.selectedParent === 0) return false
      if (this.link === '') return false
      if (this.description === '') return false
      if (this.deliveryDate === '') return false
      if (parseInt(this.amount) <= 0) return false
      if (this.userBalance < this.priceForCurrentTask) return false
      if (this.loading) return false

      return true
    },
    priceForCurrentTask() {
      let price = 0

      if (this.selectedParent !== 0){
        for (let i = 0; i < this.options.length; i++) {
          if (this.options[i].id === this.selectedParent) {
            price += this.options[i].price
            break
          }
        }

        for (let s = 0; s < this.selectedOptions.length; s++) {
          for (let o = 0; o < this.options.length; o++) {
            if (this.options[o].id === this.selectedOptions[s]) {
              price += this.options[o].price
            }
          }
        }
      }

      return price * parseInt(this.amount)
    },
    optionsToShow() {
      let parents = []

      for (let i = 0; i < this.options.length; i++) {
        if(!this.options[i].parent_id) {
          let parent = this.options[i]
          parent.options = []

          for (let i = 0; i < this.options.length; i++) {
            if(this.options[i].parent_id && this.options[i].parent_id === parent.id) {
              parent.options.push(this.options[i])
            }
          }

          parents.push(parent)
        }
      }

      return parents
    },
    ...mapState(useUserStore, ['user', 'token']),
  },
  mounted() {
    document.title = 'Редактирование задачи'

    this.getOptions()
    this.getTask(this.$route.params.id)

    this.userBalance = this.user.balance
  },
  methods: {
    onInputChange(e, field) {
      this[field] = e.target.value
    },
    getOptions() {
      axios.get(import.meta.env.VITE_API_URL + 'options/').then(res => {
        if(res.data.data) {
          this.options = res.data.data
        }
      })

      NProgress.done()
    },
    getTask(taskId) {
      axios.get(import.meta.env.VITE_API_URL + 'tasks/' + taskId).then(res => {
        if(res.data.data) {
          this.task = res.data.data
          document.title = document.title + ' #' + this.task.id

          this.link = this.task.link
          this.description = this.task.description
          this.amount = this.task.amount
          this.deliveryDate = moment(this.task.delivery_date).format('YYYY-MM-DD')

          let options = []

          for(let i = 0; i < this.task.options.length; i++) {
            if (!this.task.options[i].parent_id) {
              this.selectedParent = this.task.options[i].id
            } else {
              this.selectedOptions.push(this.task.options[i].id)
            }
          }
        }
      })

      NProgress.done()
    },
    onFormSubmit(e) {
      e.preventDefault()
      NProgress.start()

      let requestOptions = []
      requestOptions.push(this.selectedParent)
      for(let i = 0; i < this.selectedOptions.length; i++) {
        requestOptions.push(this.selectedOptions[i])
      }

      axios.put(import.meta.env.VITE_API_URL + 'tasks/' + this.task.id, {
        amount: parseInt(this.amount),
        delivery_date: this.deliveryDate,
        link: this.link,
        description: this.description,
        options: requestOptions,
      }, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.e.emit('updateUser', false)

        this.e.emit('alert', {
          title: 'Успешно',
          message: 'сохранено.',
          alertType: 1
        })

        NProgress.done()
      }).catch(err => {
        if (err.response.data.message) {
          this.e.emit('alert', {
            title: 'Ошибка:',
            message: err.response.data.message,
            alertType: 2
          })
        }

        NProgress.done()
      })
    },
    selectParent(parent) {
      if(this.selectedParent === 0 || this.selectedParent !== parent.id) this.selectedParent = parent.id
      else this.selectedParent = 0

      this.selectedOptions = []
    },
    selectOption(option) {
      if(this.selectedParent !== option.parent_id) {
        this.selectedParent = option.parent_id
        this.selectedOptions = []
      }

      let optionIndex = this.selectedOptions.indexOf(option.id)
      if (optionIndex === -1){
        this.selectedOptions.push(option.id)
      } else {
        this.selectedOptions.splice(optionIndex, 1)
      }
    },
  }
}
</script>