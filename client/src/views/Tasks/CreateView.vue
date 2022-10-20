<script setup>
import ProfileInfo from "@/components/ProfileInfo.vue";
import {moment} from "@/moment";
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Добавление задачи</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <ProfileInfo />
        <div class="mt-10 mx-3 md:mx-0">
          <transition name="slide-fade">
            <div v-if="selectedType === 1 || selectedType === 0">
              <h2 class="text-2xl font-medium">Что хотите заказать? - <span class="underline">для разовой задачи</span></h2>
              <p class="text-gray-700 mt-3 text-sm">
                Краткое описание для разовой задачи
              </p>
            </div>
          </transition>
          <transition name="slide-fade">
            <div v-if="selectedType === 2">
              <h2 class="text-2xl font-medium">Что хотите заказать? - <span class="underline">для большой задачи</span></h2>
              <p class="text-gray-700 mt-3 text-sm">
                Краткое описание для большой задачи
              </p>
            </div>
          </transition>
          <transition name="slide-fade">
            <div v-if="selectedType === 3">
              <h2 class="text-2xl font-medium">Что хотите заказать? - <span class="underline">для Excel-файла</span></h2>
              <p class="text-gray-700 mt-3 text-sm">
                Допускаются только файлы формата .xlsx. Заполните файл необходимыми данными (ссылка, описание, дата сдачи, категория и опции), каждая задача в отдельном листе.
                <br><br>
                1-A — ссылка на задачу, 1-B — описание задачи, 1-C — дата сдачи в формате xx.xx.xxxx, 1-D - количество выполнений
                <br>
                2-A — название категории, 3-... — опции (могут быть пустыми)
                <br><br>
                Категории и опции можете взять с данной страницы, выбрав разовую или большую задачу. Название категории и опции должны быть точные, но без учёта регистра.
                <br>
                Пример заполнения:
                <img src="@/assets/img/create-tasks-excel-example.png" class="m-auto mt-5 mb-10 border border-2 rounded-lg w-full md:w-3/4 lg:w-1/2">
              </p>
            </div>
          </transition>
          <div class="bg-white mt-7 py-4 px-5 grid grid-cols-1 gap-4 md:grid-cols-4 rounded-xl shadow hover:shadow-lg transition duration-300">
            <div @click="selectType(1)" class="shadow hover:shadow-lg hover:ring ring-slate-100 rounded-lg py-1 px-1 cursor-pointer transition duration-300" :class="{'bg-blue-500 text-white ring-blue-300': selectedType === 1, 'bg-slate-100': selectedType !== 1}">
              <div class="float-left ml-2 mt-2">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-14 h-14">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <div class="ml-20 py-3 font-medium">
                Разовая задача
                <br>
                <span v-if="selectedType === 1" class="text-white text-sm">Выбрано</span>
                <span v-else class="text-gray-500 text-sm">Выбрать</span>
              </div>
            </div>
            <div @click="selectedType = 2" class=" shadow hover:shadow-lg hover:ring ring-slate-100 rounded-lg py-1 px-1 cursor-pointer transition duration-300" :class="{'bg-blue-500 text-white ring-blue-300': selectedType === 2, 'bg-slate-100': selectedType !== 2}">
              <div class="float-left ml-2 mt-2">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-14 h-14">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M11.35 3.836c-.065.21-.1.433-.1.664 0 .414.336.75.75.75h4.5a.75.75 0 00.75-.75 2.25 2.25 0 00-.1-.664m-5.8 0A2.251 2.251 0 0113.5 2.25H15c1.012 0 1.867.668 2.15 1.586m-5.8 0c-.376.023-.75.05-1.124.08C9.095 4.01 8.25 4.973 8.25 6.108V8.25m8.9-4.414c.376.023.75.05 1.124.08 1.131.094 1.976 1.057 1.976 2.192V16.5A2.25 2.25 0 0118 18.75h-2.25m-7.5-10.5H4.875c-.621 0-1.125.504-1.125 1.125v11.25c0 .621.504 1.125 1.125 1.125h9.75c.621 0 1.125-.504 1.125-1.125V18.75m-7.5-10.5h6.375c.621 0 1.125.504 1.125 1.125v9.375m-8.25-3l1.5 1.5 3-3.75" />
                </svg>
              </div>
              <div class="ml-20 py-3 font-medium">
                Большая задача
                <br>
                <span v-if="selectedType === 2" class="text-white text-sm">Выбрано</span>
                <span v-else class="text-gray-500 text-sm">Выбрать</span>
              </div>
            </div>
            <div @click="selectedType = 3" class=" shadow hover:shadow-lg hover:ring ring-slate-100 rounded-lg py-1 px-1 cursor-pointer transition duration-300" :class="{'bg-blue-500 text-white ring-blue-300': selectedType === 3, 'bg-slate-100': selectedType !== 3}">
              <div class="float-left ml-2 mt-2">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-14 h-14">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M10.125 2.25h-4.5c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125v-9M10.125 2.25h.375a9 9 0 019 9v.375M10.125 2.25A3.375 3.375 0 0113.5 5.625v1.5c0 .621.504 1.125 1.125 1.125h1.5a3.375 3.375 0 013.375 3.375M9 15l2.25 2.25L15 12" />
                </svg>
              </div>
              <div class="ml-20 py-3 font-medium">
                Excel-файл
                <br>
                <span v-if="selectedType === 3" class="text-white text-sm">Выбрано</span>
                <span v-else class="text-gray-500 text-sm">Выбрать</span>
              </div>
            </div>
          </div>
          <transition name="slide-fade">
            <div v-if="selectedType === 2" class="mt-10">
              <h2 class="text-2xl font-medium mt-10">Добавленные задачи</h2>
              <div v-if="tasks.length === 0" class="text-slate-500 mt-2">
                Вы ещё не добавили ни одной задачи
              </div>
              <transition-group name="slide-fade">
                <div @click="editTask(i)" v-for="(task, i) in tasks" v-bind:key="(i + 1)" @mouseenter="showTask = (i + 1)" @mouseleave="showTask = 0" class="cursor-pointer relative bg-white shadow hover:shadow-xl mt-3 transition duration-300" :class="{'rounded-none rounded-t-xl': showTask === (i + 1), 'rounded-xl': showTask !== (i + 1)}">
                  <div class="py-3 px-4 font-medium">
                    {{ task.taskOptions.title }},
                    <span class="text-slate-500">итого: {{ $filters.currencyFormat(task.taskOptions.amountPrice) }}</span>
                    <span class="float-right text-sm text-slate-500">Дата сдачи: {{ moment(task.delivery_date).calendar() }}</span>
                  </div>
                  <transition name="slide-fade">
                    <div v-if="showTask === (i + 1)" class="absolute w-full z-20 bg-slate-100 shadow-lg rounded-b-lg pb-4 pt-2 px-3 flex flex-col">
                      <div>
                        Ссылка: <a :href="task.link" class="text-blue-500 font-medium" target="_blank">{{ task.link }}</a>
                      </div>
                      <div>
                        Количество: <span class="font-medium">{{ task.amount }} шт.</span>
                      </div>
                      <div>
                        Описание: {{ task.description }}
                      </div>
                      <div>
                        <h3 class="text-lg text-center">Опции:</h3>
                        <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
                          <div class="text-sm py-2 px-4 mt-2 rounded-full bg-slate-500 text-white shadow-lg">
                            Категория: <span class="font-medium">{{ task.taskOptions.title }}</span> <span class="text-white">{{ $filters.currencyFormat(task.taskOptions.price) }}</span>
                          </div>
                          <div v-for="option in task.taskOptions.options" class="text-sm py-2 px-4 mt-2 rounded-full bg-blue-500 text-white shadow-lg">
                            <span class="font-medium">{{ option.title }}</span> <span class="text-white">{{ $filters.currencyFormat(option.price) }}</span>
                          </div>
                        </div>
                      </div>
                    </div>
                  </transition>
                </div>
              </transition-group>
              <div class="text-center text-slate-500 mt-5">
                Итого: <span class="font-medium">{{ $filters.currencyFormat(calculateAllTasksPrice) }}</span>
              </div>
              <div v-if="closeSubmitForm">
                <button @click="closeSubmitForm = false" class="mt-5 text-white w-full bg-blue-600 transition duration-300 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center hover:bg-blue-800 mr-2 mb-2">
                  Добавить ещё одну задачу
                </button>
                <div class="text-center text-slate-500">
                  или
                </div>
                <button @click="onFormAllTasksSubmit" :disabled="!showSubmitAllTasksButton" class="mt-3 text-white w-full bg-blue-600 transition duration-300 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2" :class="{'hover:bg-blue-800': showSubmitAllTasksButton,'opacity-50': !showSubmitAllTasksButton }">
                  Сохранить добавленные задачи
                  <br>
                  <p class="text-xs -mt-1 text-blue-200">
                    Итоговая стоимость всех задач: {{ $filters.currencyFormat(calculateAllTasksPrice) }}
                  </p>
                </button>
              </div>
            </div>
          </transition>
          <transition name="slide-fade">
            <div v-if="selectedType === 1 && !closeSubmitForm || selectedType === 2 && !closeSubmitForm">
              <h2 class="text-2xl font-medium mt-10">Укажите задачу</h2>
              <div class="bg-white mt-3 py-4 px-5 grid grid-cols-1 gap-4 md:grid-cols-4 rounded-xl shadow hover:shadow-lg transition duration-300">
                <div v-for="parent in optionsToShow" @mouseenter="showOptions = parent.id" @mouseleave="showOptions = 0" class="relative">
                  <div @click="selectParent(parent)" class=" cursor-pointer shadow hover:shadow-lg hover:ring ring-slate-100 rounded-lg py-1 px-1 transition duration-300" :class="{'text-white bg-sky-800': selectedParent === parent.id, 'bg-slate-100': selectedParent !== parent.id, 'rounded-t-lg': showOptions === parent.id}">
                    <div class="py-3 px-4 font-medium flex items-center flex-row gap-4">
                      <div class="basis-1/5">
                        <div class="uppercase h-16 w-16 flex justify-center items-center text-xl rounded-md font-semibold" :class="{
                          'bg-white text-sky-800': selectedParent === parent.id,
                          'bg-sky-800 text-white': selectedParent !== parent.id
                        }">
                          {{ parent.short }}
                        </div>
                      </div>
                      <div class="basis-3/5 text-sm font-semibold">
                        <p class="leading-none">{{ parent.title }}</p>
                        <p class="text-xs mt-1" :class="{'text-white': selectedParent === parent.id, 'text-gray-500': selectedParent !== parent.id}">от {{ $filters.currencyFormat(parent.price) }}</p>
                      </div>
                      <div class="basis-1/5 text-sm font-semibold">
                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                          <path fill-rule="evenodd" clip-rule="evenodd" d="M10 5C10 6.10457 10.8954 7 12 7C13.1046 7 14 6.10457 14 5C14 3.89543 13.1046 3 12 3C10.8954 3 10 3.89543 10 5ZM12 14C10.8954 14 10 13.1046 10 12C10 10.8954 10.8954 10 12 10C13.1046 10 14 10.8954 14 12C14 13.1046 13.1046 14 12 14ZM12 21C10.8954 21 10 20.1046 10 19C10 17.8954 10.8954 17 12 17C13.1046 17 14 17.8954 14 19C14 20.1046 13.1046 21 12 21Z" fill="#2E363E"/>
                        </svg>
                      </div>
                    </div>
                  </div>
                  <transition name="slide-fade">
                    <div v-if="showOptions === parent.id" class="absolute w-full z-50 ring ring-slate-200 bg-slate-200 shadow-lg rounded-b-lg pb-4 pt-2 px-3 text-sm flex flex-col">
                      <div v-if="parent.options.length === 0" class="mt-2 text-center text-slate-500">
                        Нет опций для данной категории
                      </div>
                      <div v-for="option in parent.options" @click="selectOption(option)" class="py-2 px-4 mt-2 cursor-pointer shadow hover:shadow-lg transition duration-300 rounded-full" :class="{'bg-blue-500 text-white shadow-lg': selectedOptions.indexOf(option.id) !== -1, 'bg-slate-100': selectedOptions.indexOf(option.id) === -1}">
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
                      <span v-if="editTaskIndex !== null">Сохранить задачу</span>
                      <span v-else>Добавить</span>
                      <br>
                      <p class="text-xs -mt-1 text-blue-200">
                        Итоговая стоимость данной задачи: {{ $filters.currencyFormat(priceForCurrentTask) }}
                      </p>
                    </button>
                  </div>
                </form>
              </div>
            </div>
          </transition>
          <transition name="slide-fade">
            <div v-if="selectedType === 3" class="bg-white rounded-xl shadow hover:shadow-xl transition duration-300 py-7 px-5 mt-5">
              <input @change="changeExcelFile" accept="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" class="block w-full text-sm text-slate-900 bg-gray-100 rounded-lg border border-gray-100 cursor-pointer dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400" aria-describedby="file_input_help" id="file_input" type="file">
              <p class="mt-1 text-sm text-gray-500 dark:text-gray-300" id="file_input_help">только .xlsx</p>

              <button @click="submitExcelFileForm" :disabled="excelFile === null" class="mt-5 text-white w-full bg-blue-600 transition duration-300 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2" :class="{'hover:bg-blue-800': excelFile !== null,'opacity-50': excelFile === null }">
                Добавить
              </button>
            </div>
          </transition>
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
      selectedType: 0,
      options: [],
      selectedParent: 0,
      showOptions: 0,
      link: "",
      description: "",
      amount: 1,
      selectedOptions: [],
      deliveryDate: "",
      tasks: [],
      showTask: 0,
      editTaskIndex: null,
      closeSubmitForm: false,
      userBalance: 0,
      loading: false,
      excelFile: null,
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
    showSubmitAllTasksButton() {
      if (this.userBalance < this.calculateAllTasksPrice) return false
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
    calculateAllTasksPrice() {
      let price = 0

      for (let i = 0; i < this.tasks.length; i++) {
        price += this.tasks[i].taskOptions.amountPrice
      }

      return price
    },
    ...mapState(useUserStore, ['user', 'token']),
  },
  mounted() {
    document.title = 'Добавление задачи'

    this.getOptions()

    this.userBalance = this.user.balance
  },
  methods: {
    onInputChange(e, field) {
      this[field] = e.target.value
    },
    changeExcelFile(e) {
      this.excelFile = e.target.files[0]
    },
    editTask(taskIndex) {
      this.editTaskIndex = taskIndex
      this.selectedParent = this.tasks[taskIndex].taskOptions.id

      for (let i = 0; i < this.tasks[taskIndex].taskOptions.options.length; i++) {
        this.selectedOptions.push(this.tasks[taskIndex].taskOptions.options[i].id)
      }
      this.amount = parseInt(this.tasks[taskIndex].amount)
      this.link = this.tasks[taskIndex].link
      this.description = this.tasks[taskIndex].description
      this.deliveryDate = this.tasks[taskIndex].delivery_date

      this.closeSubmitForm = false
    },
    selectType(type) {
      this.selectedType = type
      if (type === 1) {
        this.closeSubmitForm = false
      }
    },
    taskOptionsToShow(taskOptions) {
      let parent = {}
      let options = []
      let amountPrice = 0

      for (let t = 0; t < taskOptions.length; t++) {
        for (let o = 0; o < this.options.length; o++) {
          if(taskOptions[t] === this.options[o].id) {
            if(!this.options[o].parent_id) {
              parent = this.options[o]
            } else options.push(this.options[o])
            amountPrice += this.options[o].price
          }
        }
      }

      parent.options = options
      parent.amountPrice = amountPrice * parseInt(this.amount)

      return parent
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
    collectSelectedOptions() {
      let options = []
      for (let i = 0; i < this.selectedOptions.length; i++) {
        options.push(this.selectedOptions[i])
      }
      options.push(this.selectedParent)

      return options
    },
    getOptions() {
      axios.get(import.meta.env.VITE_API_URL + 'options/').then(res => {
        if(res.data.data) {
          this.options = res.data.data
        }
      })

      NProgress.done()
    },
    onFormAllTasksSubmit() {
      this.loading = true
      NProgress.start()
      let tasksForRequest = []

      for (let i = 0; i < this.tasks.length; i++) {
        tasksForRequest.push({
          status: 1,
          amount: this.tasks[i].amount,
          delivery_date: this.tasks[i].delivery_date,
          link: this.tasks[i].link,
          description: this.tasks[i].description,
          options: this.tasks[i].options,
        })
      }

      console.log(tasksForRequest)

      axios.post(import.meta.env.VITE_API_URL + 'tasks/', {
        tasks: tasksForRequest
      }, {
        headers: { Authorization: `Bearer ${this.token}` }
      }).then(res => {
        if(res.data.status) {
          this.e.emit('alert', {
            title: 'Успешно!',
            message: 'Задачи созданы.',
            alertType: 1
          })
          this.e.emit('updateUser', false)

          this.$router.push({'name': 'tasks-my'})
        }
      }).catch(err => {
        if(err.response.data.message) this.e.emit('alert', {
          title: 'Ошибка.',
          message: err.response.data.message,
          alertType: 2
        })
        else this.e.emit('alert', {
          title: 'Ошибка.',
          message: 'Произошла внутренняя ошибка сервера.',
          alertType: 2
        })

        NProgress.done()
      })
    },
    onFormSubmit(e) {
      e.preventDefault()

      if(this.editTaskIndex !== null) {
        this.tasks[this.editTaskIndex] = {
          status: 1,
          amount: parseInt(this.amount),
          delivery_date: this.deliveryDate,
          link: this.link,
          description: this.description,
          options: this.collectSelectedOptions(),
          taskOptions: this.taskOptionsToShow(this.collectSelectedOptions())
        }

        this.closeSubmitForm = true
        this.editTaskIndex = null

        return
      }

      if(this.selectedType === 1) {
        this.loading = true
        NProgress.start()

        axios.post(import.meta.env.VITE_API_URL + 'tasks/', {
          tasks: [
            {
              status: 1,
              amount: parseInt(this.amount),
              delivery_date: this.deliveryDate,
              link: this.link,
              description: this.description,
              options: this.collectSelectedOptions(),
            }
          ]
        }, {
          headers: { Authorization: `Bearer ${this.token}` }
        }).then(res => {
          if(res.data.status) {
            this.e.emit('alert', {
              title: 'Успешно!',
              message: 'Задача создана.',
              alertType: 1
            })
            this.e.emit('updateUser', false)

            this.$router.push({'name': 'tasks-my'})
          }
        }).catch(err => {
          if(err.response.data.message) this.e.emit('alert', {
            title: 'Ошибка.',
            message: err.response.data.message,
            alertType: 2
          })
          else this.e.emit('alert', {
            title: 'Ошибка.',
            message: 'Произошла внутренняя ошибка сервера.',
            alertType: 2
          })

          NProgress.done()
        })
      } else if (this.selectedType === 2) {
        let task = {
          status: 1,
          amount: parseInt(this.amount),
          delivery_date: this.deliveryDate,
          link: this.link,
          description: this.description,
          options: this.collectSelectedOptions(),
          taskOptions: this.taskOptionsToShow(this.collectSelectedOptions())
        }

        this.tasks.push(task)
        this.closeSubmitForm = true

        this.clearForm()
      }
    },
    submitExcelFileForm() {
      if(!this.excelFile) return
      if(this.loading) return

      let formData = new FormData();
      formData.append('file', this.excelFile);

      axios.post(import.meta.env.VITE_API_URL + 'tasks/excel', formData, {
        headers: {
          Authorization: `Bearer ${this.token}`,
          'Content-Type': 'multipart/form-data'
        }
      }).then(res => {
        if(res.data.status) {
          this.e.emit('alert', {
            title: 'Успешно!',
            message: 'Задачи созданы.',
            alertType: 1
          })
          this.e.emit('updateUser', false)

          this.$router.push({'name': 'tasks-my'})
        }
      }).catch(err => {
        if(err.response.data.message) this.e.emit('alert', {
          title: 'Ошибка.',
          message: err.response.data.message,
          alertType: 2
        })
        else this.e.emit('alert', {
          title: 'Ошибка.',
          message: 'Произошла внутренняя ошибка сервера.',
          alertType: 2
        })

        NProgress.done()
      })
    },
    clearForm() {
      this.amount = 1
      this.deliveryDate = ''
      this.link = ''
      this.description = ''
      this.selectedParent = 0
      this.selectedOptions = []
    }
  }
}
</script>

<style>
.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}
.slide-fade-leave-active {
  transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(50px);
  opacity: 0;
}
</style>