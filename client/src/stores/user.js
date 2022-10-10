import { defineStore } from 'pinia'
import { useStorage } from '@vueuse/core'

export const useUserStore = defineStore('user', {
    state: () => {
        return {
            token: useStorage('token', ''),
            user: useStorage('user', {type: 0}),
            eventsAfter: useStorage('eventsAfter', ''),
        }
    },
    actions: {
        setToken(token) {
            this.token = token
        },
        setUser(user) {
            this.user = user
        },
        setEventsAfter(eventsAfter) {
            this.eventsAfter = eventsAfter
        }
    },
})
