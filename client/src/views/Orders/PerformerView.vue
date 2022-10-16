<script setup>
import ProfileInfo from "@/components/ProfileInfo.vue";
import {moment} from "@/moment";
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Мои задачи</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <ProfileInfo />
        <div v-for="order in orders" class="bg-white rounded-xl shadow">
          <div class="grid grid-cols-1 md:text-sm mt-4 md:grid-cols-6 py-3 px-4 cursor-pointer" @click="openFull(order.id)">
            <div class="text-center md:text-left py-1 md:py-0">
              <div class="text-slate-500">ID:</div>
              <div class="font-medium">#{{ order.id }}</div>
            </div>
            <div class="text-center md:text-left border-t-2 md:border-none py-1 md:py-0">
              <div class="text-slate-500">Категория:</div>
              <div class="font-medium">{{ order.task.structed.main.title }}</div>
            </div>
            <div class="text-center md:text-left border-t-2 md:border-none py-1 md:py-0">
              <div class="text-slate-500">Стоимость:</div>
              <div class="font-medium">{{ $filters.currencyFormat(order.task.structed.price)}}</div>
            </div>
            <div class="text-center md:text-left border-t-2 md:border-none py-1 md:py-0">
              <div class="text-slate-500">Заказчик:</div>
              <div>
                <span class="font-medium">{{ order.task.customer.username }}</span><br>
                <span class="text-slate-500">онлайн {{ moment(order.task.customer.last_online).utcOffset(+6, true).fromNow() }}</span>
              </div>
            </div>
            <div class="text-center md:text-left border-t-2 md:border-none py-1 md:py-0">
              <div class="text-slate-500">Создано:</div>
              <div class="font-medium">{{ moment(order.created_at).utc(0).format('lll') }}</div>
            </div>
            <div class="border-t-2 md:border-none py-1 md:py-0">
              <button v-if="order.status === 0" disabled class="bg-blue-200 text-blue-700 w-full font-medium py-3 px-2 rounded-lg shadow">Выполняется</button>
              <button v-if="order.status === 1" disabled class="bg-yellow-200 text-yellow-700 w-full font-medium py-3 px-2 rounded-lg shadow">Сдан на проверку</button>
              <button v-if="order.status === 2" disabled class="bg-green-200 text-green-700 w-full font-medium py-3 px-2 rounded-lg shadow">Выполнено</button>
              <button v-if="order.status === 3 && order.canceled_user_id === order.task.customer_id" disabled class="bg-red-200 text-red-700 w-full font-medium py-3 px-2 rounded-lg shadow">Отменено заказчиком</button>
              <button v-if="order.status === 3 && order.canceled_user_id === order.offer.performer_id" disabled class="bg-red-200 text-red-700 w-full font-medium py-3 px-2 rounded-lg shadow">Отменено Вами</button>
            </div>
          </div>
          <transition name="slide-fade">
            <div v-if="fullShow === order.id" class="border-t-2">
              <div class="py-3 px-4">
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-2">
                  <div class="text-sm py-2 px-4 rounded-full bg-slate-500 text-white shadow-lg">
                    Категория: <span class="font-medium">{{ order.task.structed.main.title }}</span> <span class="text-white">{{ $filters.currencyFormat(order.task.structed.main.price) }}</span>
                  </div>
                  <div v-for="option in order.task.structed.options" class="text-sm py-2 px-4 rounded-full bg-blue-500 text-white shadow-lg">
                    <span class="font-medium">{{ option.title }}</span> <span class="text-white">{{ $filters.currencyFormat(option.price) }}</span>
                  </div>
                </div>
                <div v-if="order.cancel_comment" class="mt-2 text-sm bg-red-200 text-red-700 py-2 px-3 rounded-full">
                  Комментарий отмены: <span class="font-medium">{{ order.cancel_comment }}</span>
                </div>
                <div v-if="order.status === 1 && order.surrender_comment" class="mt-2 text-sm bg-yellow-200 text-yellow-700 py-2 px-3 rounded-full">
                  Ваш комментарий сдачи работы: <span class="font-medium">{{ order.surrender_comment }}</span>
                </div>
                <div v-if="order.status === 0 && order.return_comment" class="mt-2 text-sm bg-yellow-200 text-yellow-700 py-2 px-3 rounded-full">
                  Комментарий заказчика о возврате в работу: <span class="font-medium">{{ order.return_comment }}</span>
                </div>
              </div>
              <div v-if="order.status < 2" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 text-sm border-t-2 py-2 px-2 gap-2">
                <button v-if="order.status === 0 && cancelShow !== order.id" @click="openCancel(order.id)" class="bg-red-200 text-red-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:bg-red-300 transition duration-300">
                  Отменить
                </button>
                <button v-if="order.status === 0 && cancelShow === order.id" @click="openCancel(order.id)" class="bg-slate-200 text-slate-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:bg-slate-300 transition duration-300">
                  Закрыть
                </button>
                <button v-if="order.status === 0 && surrenderShow !== order.id" @click="openSurrender(order.id)" class="bg-yellow-200 text-yellow-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:bg-yellow-300 transition duration-300">
                  Сдать работу
                </button>
                <button v-if="order.status === 0 && surrenderShow === order.id" @click="openSurrender(order.id)" class="bg-slate-200 text-slate-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:bg-slate-300 transition duration-300">
                  Закрыть
                </button>
              </div>
              <transition name="slide-fade">
                <div v-if="cancelShow === order.id" class="border-t-2 text-sm py-3 px-4">
                  <div>
                    <label for="cancel_comment" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Комментарий отмены</label>
                    <input type="text" id="cancel_comment" @input="updateInput($event, 'cancelComment')" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Введите комментарий..." required>
                  </div>
                  <button @click="cancelOrder(order.id)" :disabled="cancelComment === ''" class="bg-red-200 text-red-700 mt-3 w-full font-medium py-3 px-2 rounded-lg shadow" :class="{'hover:bg-red-300 transition duration-300': cancelComment !== '', 'opacity-50': cancelComment === ''}">Отменить</button>
                </div>
              </transition>
              <transition name="slide-fade">
                <div v-if="surrenderShow === order.id" class="border-t-2 text-sm py-3 px-4">
                  <div class="mt-2 text-sm bg-yellow-200 text-yellow-600 py-2 px-3 rounded-full mb-4">
                    Вы можете указать, например, имя в социальной сети. Это поможет заказчику проверить выполненную Вами работу.
                  </div>
                  <div>
                    <label for="surrender_comment" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Комментарий для сдачи</label>
                    <input type="text" id="surrender_comment" @input="updateInput($event, 'surrenderComment')" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Введите комментарий..." required>
                  </div>
                  <button @click="surrenderOrder(order.id)" :disabled="surrenderComment === ''" class="bg-yellow-200 text-yellow-700 mt-3 w-full font-medium py-3 px-2 rounded-lg shadow" :class="{'hover:bg-red-300 transition duration-300': surrenderComment !== '', 'opacity-50': surrenderComment === ''}">Сдать работу</button>
                </div>
              </transition>
            </div>
          </transition>
        </div>
      </div>
    </main>
  </main>
</template>

<script>
import NProgress from "nprogress";
import axios from "axios"
import {mapState} from "pinia";
import {useUserStore} from "@/stores/user";

export default {
  data() {
    return {
      orders: [],
      fullShow: 0,
      surrenderShow: 0,
      surrenderComment: '',
      cancelShow: 0,
      cancelComment: '',
    }
  },
  computed: {
    ...mapState(useUserStore, ['user', 'token'])
  },
  mounted() {
    document.title = 'Мои задачи'

    this.getOrders()
  },
  methods: {
    updateInput(e, field) {
      this[field] = e.target.value
    },
    openFull(orderId) {
      if (this.fullShow === orderId) this.fullShow = 0
      else this.fullShow = orderId
    },
    openCancel(orderId) {
      if (this.cancelShow === orderId) this.cancelShow = 0
      else {
        this.cancelShow = orderId
        this.surrenderShow = 0
        this.surrenderComment = ''
      }
    },
    openSurrender(orderId) {
      if (this.surrenderShow === orderId) this.surrenderShow = 0
      else {
        this.surrenderShow = orderId
        this.cancelShow = 0
        this.cancelComment = ''
      }
    },
    getTaskOptionStructured(task) {
      let mainOption = {}
      let options = []
      let price = 0

      for(let i = 0; i < task.options.length; i++) {
        if (!task.options[i].parent_id) {
          mainOption = task.options[i]
        } else {
          options.push(task.options[i])
        }
        price += task.options[i].price
      }

      return {
        'main': mainOption,
        'price': price,
        'options': options
      }
    },
    getOrders() {
      NProgress.start();

      axios.get(import.meta.env.VITE_API_URL + 'orders/performer', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.orders = res.data.data

          for(let i = 0; i < this.orders.length; i++) {
            this.orders[i].task.structed = this.getTaskOptionStructured(this.orders[i].task)
          }
        }

        NProgress.done();
      })
    },
    cancelOrder(orderId) {
      NProgress.start();

      axios.put(import.meta.env.VITE_API_URL + 'orders/' + orderId, {
        status: 3,
        cancel_comment: this.cancelComment
      }, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.getOrders()
        this.cancelComment = ''
        this.cancelShow = 0
      })
    },
    surrenderOrder(orderId) {
      NProgress.start();

      axios.put(import.meta.env.VITE_API_URL + 'orders/' + orderId, {
        status: 1,
        surrender_comment: this.surrenderComment
      }, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.getOrders()
        this.surrenderComment = ''
        this.surrenderShow = 0
      })
    }
  }
}
</script>