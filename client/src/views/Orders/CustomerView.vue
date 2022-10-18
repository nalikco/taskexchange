<script setup>
import ProfileInfo from "@/components/ProfileInfo.vue";
import {moment} from "@/moment";
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Мои заказы</h1>
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
              <div class="text-slate-500">Исполнитель:</div>
              <div>
                <span class="font-medium">{{ order.offer.performer.username }}</span><br>
                <span class="text-slate-500">онлайн {{ moment(order.offer.performer.last_online).utcOffset(+6, true).fromNow() }}</span>
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
              <button v-if="order.status === 3 && order.canceled_user_id === order.task.customer_id" disabled class="bg-red-200 text-red-700 w-full font-medium py-3 px-2 rounded-lg shadow">Отменено Вами</button>
              <button v-if="order.status === 3 && order.canceled_user_id === order.offer.performer_id" disabled class="bg-red-200 text-red-700 w-full font-medium py-3 px-2 rounded-lg shadow">Отменено исполнителем</button>
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
                  Комментарий исполнителя: <span class="font-medium">{{ order.surrender_comment }}</span>
                </div>
                <div v-if="order.status === 0 && order.return_comment" class="mt-2 text-sm bg-yellow-200 text-yellow-700 py-2 px-3 rounded-full">
                  Ваш комментарий о возврате в работу: <span class="font-medium">{{ order.return_comment }}</span>
                </div>
              </div>
              <div v-if="order.status < 2" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 text-sm border-t-2 py-2 px-2 gap-2">
                <button v-if="order.status === 0 && cancelShow !== order.id" @click="openCancel(order.id)" class="bg-red-200 text-red-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:bg-red-300 transition duration-300">
                  Отменить
                </button>
                <button v-if="order.status === 0 && cancelShow === order.id" @click="openCancel(order.id)" class="bg-slate-200 text-slate-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:bg-slate-300 transition duration-300">
                  Закрыть
                </button>
                <button v-if="order.status === 1" @click="acceptOrder(order.id)" class="bg-green-200 text-green-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:bg-green-300 transition duration-300">
                  Подтвердить
                  <p class="text-xs -mt-1">выполнение</p>
                </button>
                <button v-if="order.status === 1 && returnShow !== order.id" @click="openReturn(order.id)" class="bg-yellow-200 text-yellow-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:bg-yellow-300 transition duration-300">
                  Вернуть
                  <p class="text-xs -mt-1">в работу</p>
                </button>
                <button v-if="order.status === 1 && returnShow === order.id" @click="openReturn(order.id)" class="bg-slate-200 text-slate-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:bg-slate-300 transition duration-300">
                  Закрыть
                </button>
              </div>
              <div class="flex flex-col items-center px-2 pb-2">
                <RouterLink :to="{ name:'messages', query: { recipient_id: order.offer.performer.id } }" class="bg-green-200 text-sm flex place-content-center font-medium text-green-700 w-full font-medium py-3 px-2 rounded-lg shadow hover:bg-green-300 transition duration-300">
                  Написать исполнителю
                </RouterLink>
              </div>
              <transition name="slide-fade">
                <div v-if="cancelShow === order.id" class="border-t-2 text-sm py-3 px-4">
                  <div class="mt-2 text-sm bg-red-200 text-red-600 py-2 px-3 rounded-full mb-4">
                    После отмены количество у задачи, по которой был создан заказ, увеличится на 1, т.е. вернется на прежнее значение и средства не будут утеряны.
                  </div>
                  <div>
                    <label for="cancel_comment" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Комментарий отмены</label>
                    <input type="text" id="cancel_comment" @input="updateInput($event, 'cancelComment')" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Введите комментарий..." required>
                  </div>
                  <button @click="cancelOrder(order.id)" :disabled="cancelComment === ''" class="bg-red-200 text-red-700 mt-3 w-full font-medium py-3 px-2 rounded-lg shadow" :class="{'hover:bg-red-300 transition duration-300': cancelComment !== '', 'opacity-50': cancelComment === ''}">Отменить</button>
                </div>
              </transition>
              <transition name="slide-fade">
                <div v-if="returnShow === order.id" class="border-t-2 text-sm py-3 px-4">
                  <div class="mt-2 text-sm bg-yellow-200 text-yellow-600 py-2 px-3 rounded-full mb-4">
                    Комментарий необходим для того, чтобы исполнителю было легче понять, какие ошибки в процессе работы были допущены.
                  </div>
                  <div>
                    <label for="return_comment" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Комментарий возврата в работу</label>
                    <input type="text" id="return_comment" @input="updateInput($event, 'returnComment')" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Введите комментарий..." required>
                  </div>
                  <button @click="returnOrder(order.id)" :disabled="returnComment === ''" class="bg-yellow-200 text-yellow-700 mt-3 w-full font-medium py-3 px-2 rounded-lg shadow" :class="{'hover:bg-yellow-300 transition duration-300': returnComment !== '', 'opacity-50': returnComment === ''}">Отменить</button>
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
      cancelShow: 0,
      cancelComment: '',
      returnShow: 0,
      returnComment: ''
    }
  },
  computed: {
    ...mapState(useUserStore, ['user', 'token'])
  },
  mounted() {
    document.title = 'Мои заказы'

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
      if (this.cancelShow === orderId)this.cancelShow = 0
      else {
        this.cancelShow = orderId
        this.returnShow = 0
        this.returnComment = ''
      }
    },
    openReturn(orderId) {
      if (this.returnShow === orderId)this.returnShow = 0
      else {
        this.returnShow = orderId
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

      axios.get(import.meta.env.VITE_API_URL + 'orders/customer', {
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
    returnOrder(orderId) {
      NProgress.start();

      axios.put(import.meta.env.VITE_API_URL + 'orders/' + orderId, {
        status: 0,
        return_comment: this.returnComment
      }, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.getOrders()
        this.returnComment = ''
        this.returnShow = 0
      })
    },
    acceptOrder(orderId) {
      NProgress.start();

      axios.put(import.meta.env.VITE_API_URL + 'orders/' + orderId, {
        status: 2
      }, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.getOrders()
        this.fullShow = 0
        this.cancelShow = 0
        this.cancelComment = ''
        this.returnShow = 0
        this.returnComment = ''
      })
    }
  }
}
</script>