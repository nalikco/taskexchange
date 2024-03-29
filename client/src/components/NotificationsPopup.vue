<script setup>
import {moment} from "@/moment";
</script>
<template>
  <div>
    <transition
        enter-active-class="transition ease-out duration-100"
        enter-from-class="transform opacity-0 scale-95"
        enter-to-class="transform opacity-100 scale-100"
        leave-active-class="transition ease-in duration-75"
        leave-from-class="transform opacity-100 scale-100"
        leave-to-class="transform opacity-0 scale-95"
    >
      <div v-if="showNotifications" class="absolute z-10 w-80 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none" :class="{'right-5 mt-5': isMobile === 'true', 'right-40 mt-5': isMobile !== 'true'}" role="menu" aria-orientation="vertical" aria-labelledby="user-notifications-button" tabindex="-1">
        <div v-if="events.length > 0">
          <RouterLink v-for="event in showedEvents" v-bind:key="event.id" :to="event.link" @click="onEventClick(event)" class="block px-4 py-2 text-sm text-gray-700 transition duration-300 hover:bg-gray-200" role="menuitem" tabindex="-1" id="user-menu-item-0">
            {{ event.message }}
            <br>
            <span class="text-gray-500 text-xs">{{ moment(event.created_at).utcOffset(+6, true).fromNow() }}</span>
          </RouterLink>
        </div>
        <div v-else class="text-center mt-3 mb-2 text-gray-500 text-sm">
          Нет новых уведомлений
        </div>
        <div v-if="events.length > 5" class="text-center mt-3 mb-2 text-sm text-gray-500">
          и ещё {{ events.length - 5 }} {{ $filters.declOfNum(events.length - 5, ['уведомление', 'уведомления', 'уведомлений'])}}
        </div>
        <div class="flex items-center">
          <RouterLink :to="{'name': 'events'}" @click="showNotifications = false" class="text-sm mx-4 w-full my-2 bg-green-500 hover:bg-green-700 shadow-md hover:shadow-none transition duration-300 text-center text-white font-medium py-1 rounded-lg">Все уведомления</RouterLink>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
import axios from "axios";
import {mapState} from "pinia";
import {useUserStore} from "@/stores/user";

export default {
  props: ['showNotifications', 'isMobile', 'events'],
  computed: {
    showedEvents() {
      if(this.events.length > 5) {
        return this.events.slice(0, 5)
      }
      return this.events
    },
    ...mapState(useUserStore, ['token']),
  },
  methods: {
    onEventClick(event) {
      this.$emit('closePopup')

      axios.put(import.meta.env.VITE_API_URL + 'events/' + event.id, {}, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        this.$emit('updateEvents')
      })
    }
  }
}
</script>