<script setup>
import {moment} from "@/moment";
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Уведомления</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <div v-if="events.length > 0" class="text-gray-500 text-sm px-3 pb-4 text-center">
          При нажатии на уведомление Вы перейдете на источник уведомления.
        </div>
        <transition name="slide-fade" mode="out-in">
          <div v-if="loading" class="text-center mb-5">
              <div role="status">
                  <svg class="inline mr-2 w-8 h-8 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
                      <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
                  </svg>
                  <span class="sr-only">Loading...</span>
              </div>
          </div>
        </transition>
        <RouterLink v-for="event in events" v-bind:key="event.id" :to="event.link">
          <div class="bg-white mt-2 text-sm rounded-lg shadow py-3 px-3 hover:shadow-lg transition duration-300">
            <div class="text-gray-500 text-xs">
              {{ moment(event.created_at).utc(0).format('lll') }}
              <span v-if="!event.viewed_at" class="float-right text-blue-500 font-medium">новое</span>
              <span v-else class="float-right font-medium">просмотрено {{ moment(event.viewed_at).utc(0).format('lll') }}</span>
            </div>
            <div class="mt-2">
              {{ event.message }}
            </div>
          </div>
        </RouterLink>
        <div v-if="events.length === 0" class="text-gray-500 text-sm text-center">
          У Вас пока нет уведомлений
        </div>
        <div v-if="events.length > 0" class="mt-7 text-sm grid justify-items-center">
          <div class="text-sm text-gray-500 dark:text-gray-400">
              Показано <span class="font-semibold text-gray-700 dark:text-white">{{ offset + 1 }}-{{ offsetEnd }}</span> из <span class="font-semibold text-gray-700 dark:text-white">{{ count }}</span> {{ $filters.declOfNum(count, ['уведомления', 'уведомлений', 'уведомлений']) }}
          </div>
          <div>
            <nav class="mt-8">
              <ul class="inline-flex -space-x-px">
                <li v-if="currentPage !== 1">
                  <RouterLink :to="'?page=' + (currentPage - 1)" reload class="py-2 px-3 ml-0 leading-tight rounded-l-lg text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">Назад</RouterLink>
                </li>
                <li v-for="page in pages">
                  <RouterLink :to="'?page=' + page" @click="currentPage = page" :class="{'py-2 px-3 text-blue-600 bg-blue-50 border border-gray-300 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white': page === currentPage, 'py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white': page !== currentPage }">{{ page }}</RouterLink>
                </li>
                <li v-if="currentPage < pages">
                  <RouterLink :to="'#'" class="py-2 px-3 leading-tight text-gray-500 bg-white rounded-r-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">Вперед</RouterLink>
                </li>
              </ul>
            </nav>
          </div>
        </div>
      </div>
    </main>
  </main>
</template>

<script>
import axios from "axios"
import {mapState} from "pinia";
import {useUserStore} from "@/stores/user";

export default {
  data() {
    return {
      events: [],
      currentPage: 1,
      perPage: 25,
      offset: 0,
      pages: 0,
      count: 0,
      loading: false,
    }
  },
  computed: {
    offsetEnd() {
      let offsetEnd = this.offset + this.perPage
      if (offsetEnd > this.count) offsetEnd = this.count
      return offsetEnd
    },
    ...mapState(useUserStore, ['token'])
  },
  mounted() {
    document.title = 'Уведомления'

    if (this.$route.query.page) {
      this.currentPage = parseInt(this.$route.query.page)
    }

    this.getEvents()

    this.viewAllEvents()
  },
  methods: {
    getEvents() {
      this.loading = true

      axios.get(import.meta.env.VITE_API_URL + 'events/?page=' + this.currentPage + '&per_page=' + this.perPage, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.events = res.data.data
          this.pages = res.data.pagination.pages
          this.count = res.data.pagination.count
          this.offset = res.data.pagination.offset

        }

        this.loading = false
      })
    },
    viewAllEvents() {
      axios.put(import.meta.env.VITE_API_URL + 'events/view-all', {}, {
        headers: { Authorization: `Bearer ${this.token}` },
      })
    }
  },
  watch: {
    currentPage() {
      this.getEvents()
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