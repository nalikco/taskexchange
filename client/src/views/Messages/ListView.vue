<script setup>
import {moment} from "@/moment";
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Сообщения</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <div class="flex mx-3 md:mx-0 antialiased text-gray-800">
          <div class="md:flex md:flex-row w-full overflow-x-hidden">
            <div class="flex flex-col bg-white flex-shrink-0 rounded-2xl" :class="{'w-full': selectedConversationIndex === -1 && !newRecipient, 'w-full md:w-64': selectedConversationIndex !== -1 || newRecipient}">
              <div class="flex flex-col pb-2">
                <div class="flex flex-row items-center justify-between rounded-2xl px-3 py-2 text-sm bg-slate-200 font-medium text-slate-900">
                  Беседы
                  <span class="flex items-center justify-center bg-slate-400 text-white h-5 w-5 rounded-full">
                    {{ conversations.length }}
                  </span>
                </div>
                <div class="flex flex-col space-y-1 mt-2 -mx-2 px-3 overflow-y-auto max-h-64">
                  <button v-for="(conversation, i) in conversations" @click="goToDialogWithUser(getConversationRecipient(conversation).id)" class="relative flex flex-row items-center hover:bg-gray-100 rounded-xl p-2">
                    <div v-if="getConversationUnViewedMessagesCount(conversation) > 0" class="absolute right-0 text-sm mr-5 bg-blue-500 rounded-full h-5 w-5 text-white font-semibold">
                      {{ getConversationUnViewedMessagesCount(conversation) }}
                    </div>
                    <div class="flex items-center justify-center h-10 w-10 bg-indigo-200 rounded-full">
                      <img src="@/assets/img/user.png" class="rounded-full">
                    </div>
                    <div class="ml-2 text-sm">
                      <div class="font-medium flex items-center">
                        {{ getConversationRecipient(conversation).username }}
                        <div v-if="checkIsOnline(getConversationRecipient(conversation).last_online)" class="ml-1 h-2.5 w-2.5 rounded-full bg-green-400 mr-2"></div>
                      </div>
                      <div v-if="conversation.messages[0].sender.id === user.id" class="text-xs -mt-1 text-left text-slate-600">
                        <span class="text-slate-400">Вы:</span> {{ conversation.messages[0].text }}
                      </div>
                      <div v-else class="text-xs -mt-1 text-left text-slate-600">
                        {{ conversation.messages[0].text }}
                      </div>
                    </div>
                  </button>
                </div>
              </div>
            </div>
            <div class="flex flex-col flex-auto h-full mt-10 md:mt-0 md:px-4" v-if="selectedConversationIndex !== -1" @timeupdate="scrollToMessagesBottom">
              <div class="flex flex-col flex-auto flex-shrink-0 rounded-2xl bg-gray-100 h-full">
                <div class="bg-white w-full font-medium text-sm shadow rounded-t-2xl py-3 px-5">
                  <div class="flex flex-row items-center">
                    <img src="@/assets/img/user.png" class="h-10 w-10 rounded-full">
                    <div class="ml-4">
                      {{ getConversationRecipient(conversations[selectedConversationIndex]).username }}<br>
                      <div v-if="checkIsOnline(getConversationRecipient(conversations[selectedConversationIndex]).last_online)" class="flex items-center text-slate-500 -mt-1 text-xs">
                        <div class="h-2.5 w-2.5 rounded-full bg-green-400 mr-2"></div> онлайн
                      </div>
                      <div v-else class="flex items-center text-slate-500 -mt-1 text-xs">
                        онлайн {{ moment(getConversationRecipient(conversations[selectedConversationIndex]).last_online).utcOffset(+6, true).fromNow() }}
                      </div>
                    </div>
                  </div>
                </div>
                <div ref="messages" class="flex flex-col overflow-x-auto scrollbar-thin scrollbar-thumb-rounded-xl scrollbar-thumb-gray-400 scrollbar-track-gray-200" style="max-height: 60vh;">
                  <div class="flex flex-col">
                    <div class="flex flex-col-reverse gap-3 bg-gray-200 pb-7 pt-5 px-5 rounded-b-2xl">
                      <div v-for="message in conversations[selectedConversationIndex].messages" class="rounded-lg" :class="{
                        'col-start-1 col-end-8': message.sender.id !== user.id,
                        'col-start-6 col-end-13': message.sender.id === user.id
                      }">
                        <div v-if="message.sender.id !== user.id" class="flex flex-row items-center">
                          <div class="flex items-center justify-center h-10 w-10 text-white rounded-full bg-indigo-500 flex-shrink-0">
                            <img src="@/assets/img/user.png" class="rounded-full">
                          </div>
                          <div class="relative ml-3 text-sm bg-white py-2 px-4 shadow rounded-xl">
                            <div>
                              {{ message.text }}
                              <div class="text-slate-400 text-xs font-medium">{{ dateView(message.created_at) }}</div>
                            </div>
                          </div>
                        </div>
                        <div v-else class="flex items-center justify-start flex-row-reverse" :class="{'mb-5': getLastUserViewedMessage(conversations[selectedConversationIndex]) && getLastUserViewedMessage(conversations[selectedConversationIndex]).id === message.id}">
                          <div class="flex items-center justify-center h-10 w-10 text-white rounded-full bg-indigo-500 flex-shrink-0">
                            <img src="@/assets/img/user.png" class="rounded-full">
                          </div>
                          <div class="relative mr-3 text-sm bg-indigo-100 py-2 px-4 shadow rounded-xl">
                            <div>
                              {{ message.text }}
                              <div class="text-slate-400 text-xs font-medium">{{ dateView(message.created_at) }}</div>
                            </div>
                            <div v-if="getLastUserViewedMessage(conversations[selectedConversationIndex]) && getLastUserViewedMessage(conversations[selectedConversationIndex]).id === message.id" class="absolute text-xs bottom-0 right-0 -mb-5 mr-2 text-gray-500">
                              просмотрено
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="flex flex-row items-center h-16 rounded-xl bg-white w-full px-2">
                  <div class="flex-grow">
                    <div class="relative w-full">
                      <input type="text" v-on:keyup.enter="sendMessage" @input="changeNewMessageText" class="flex w-full border rounded-xl focus:outline-none focus:border-indigo-300 pl-4 h-9 text-sm" :value="newMessageText" placeholder="Введите сообщение...">
                    </div>
                  </div>
                  <div class="ml-2">
                    <button @click="sendMessage" :disabled="newMessageText === ''" class="flex items-center text-sm justify-center bg-indigo-500 rounded-xl text-white px-4 py-2 font-medium flex-shrink-0" :class="{
                      'opacity-50': newMessageText === '',
                      'hover:bg-indigo-600': newMessageText !== ''
                    }">
                      <span>Отправить</span>
                      <span class="ml-2">
                  <svg
                      class="w-4 h-4 transform rotate-45 -mt-px"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                      xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"
                    ></path>
                  </svg>
                </span>
                    </button>
                  </div>
                </div>
              </div>
            </div>
            <div class="flex flex-col flex-auto h-full px-4" v-if="newRecipient !== null">
              <div class="flex flex-col flex-auto flex-shrink-0 rounded-2xl bg-gray-100 h-full">
                <div class="bg-white w-full font-medium text-sm shadow rounded-t-2xl py-3 px-5">
                  {{ newRecipient.username }}<br>
                  <div v-if="checkIsOnline(newRecipient.last_online)" class="flex items-center text-slate-500 -mt-1 text-xs">
                    <div class="h-2.5 w-2.5 rounded-full bg-green-400 mr-2"></div> онлайн
                  </div>
                  <div v-else class="flex items-center text-slate-500 -mt-1 text-xs">
                    онлайн {{ moment(newRecipient.last_online).utcOffset(+6, true).fromNow() }}
                  </div>
                </div>
                <div ref="messages" class="flex flex-col overflow-x-auto" style="max-height: 60vh;">
                  <div class="flex flex-col">
                    <div class="flex flex-col gap-3 bg-gray-200 h-60 items-center text-gray-500 pb-7 pt-5 px-5 rounded-b-2xl">
                      У Вас ещё нет диалога с данным пользователем.
                    </div>
                  </div>
                </div>
                <div class="flex flex-row items-center h-16 rounded-xl bg-white w-full px-2">
                  <div class="flex-grow">
                    <div class="relative w-full">
                      <input type="text" v-on:keyup.enter="sendMessage" @input="changeNewMessageText" class="flex w-full border rounded-xl focus:outline-none focus:border-indigo-300 pl-4 h-9 text-sm" :value="newMessageText" placeholder="Введите сообщение...">
                    </div>
                  </div>
                  <div class="ml-2">
                    <button @click="sendMessage" :disabled="newMessageText === ''" class="flex items-center text-sm justify-center bg-indigo-500 rounded-xl text-white px-4 py-2 font-medium flex-shrink-0" :class="{
                      'opacity-50': newMessageText === '',
                      'hover:bg-indigo-600': newMessageText !== ''
                    }">
                      <span>Отправить</span>
                      <span class="ml-2">
                        <svg
                            class="w-4 h-4 transform rotate-45 -mt-px"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                            xmlns="http://www.w3.org/2000/svg"
                        >
                          <path
                              stroke-linecap="round"
                              stroke-linejoin="round"
                              stroke-width="2"
                              d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"
                          ></path>
                        </svg>
                      </span>
                    </button>
                  </div>
                </div>
              </div>
            </div>
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
import NProgress from "nprogress";
import {emitter} from "@/emitter";

export default {
  data() {
    return {
      conversations: [],
      dialogWithUser: 0,
      selectedConversationIndex: -1,
      recipientId: null,
      newRecipient: null,
      newMessageText: '',
      e: emitter
    }
  },
  computed: {
    ...mapState(useUserStore, ['token', 'user'])
  },
  mounted() {
    document.title = 'Сообщения'

    if(this.$route.query.recipient_id) {
      this.recipientId = this.$route.query.recipient_id
    } else this.recipientId = 0

    this.e.on('newMessage', (res, e) => {
      this.newMessageUpdate(res)
    })
  },
  beforeUpdate() {
    if(this.$route.query.recipient_id) {
      this.recipientId = this.$route.query.recipient_id
    } else this.recipientId = 0
  },
  updated() {
    this.scrollToMessagesBottom()
  },
  methods: {
    changeNewMessageText(e) {
      this.newMessageText = e.target.value
    },
    goToDialogWithUser(recipientId) {
      if(this.recipientId == recipientId) {
        this.selectedConversationIndex = -1
        this.$router.push({ name: 'messages' })
        return
      }
      this.recipientId = recipientId
      this.$router.push({ name: 'messages', query: { recipient_id: recipientId }})
    },
    setDialogWithUser(recipientId) {
      this.newRecipient = null

      for(let c = 0; c < this.conversations.length; c++) {
        for(let m = 0; m < this.conversations[c].members.length; m++) {
          if(this.conversations[c].members[m].id == recipientId) {
            this.selectedConversationIndex = c
            document.title = 'Диалог с ' + this.getConversationRecipient(this.conversations[c]).username

            this.viewConversation(this.conversations[c].id)
            this.e.emit('updateMessageCount')
            return
          }
        }
      }

      console.log(this.selectedConversationIndex)

      if (this.selectedConversationIndex === -1) this.getNewRecipient(recipientId)
    },
    getConversationUnViewedMessagesCount(conversation) {
      let count = 0

      for(let m = 0; m < conversation.messages.length; m++) {
        if(!conversation.messages[m].viewed_at && conversation.messages[m].sender.id !== this.user.id) {
          count++
        }
      }

      return count
    },
    dateView(date) {
      let dateObj = moment(date).utc(0)
      let currentDateObj = moment()

      if (dateObj.isSame(currentDateObj, 'days')) {
        return "сегодня в " + dateObj.format('HH:mm')
      }
      else if (dateObj.subtract(1, 'days').isSame(currentDateObj, 'days')) {
        return "вчера в " + dateObj.format('HH:mm')
      } else {
        return dateObj.format('DD MMM в HH:mm')
      }
    },
    getLastUserMessage(conversation) {
      let userMessages = []

      for(let i = 0; i < conversation.messages.length; i++) {
        if (conversation.messages[i].sender.id === this.user.id) {
          userMessages.push(conversation.messages[i])
        }
      }

      return userMessages[0]
    },
    getLastUserViewedMessage(conversation) {
      let userMessages = []

      for(let i = 0; i < conversation.messages.length; i++) {
        if (conversation.messages[i].sender.id === this.user.id && conversation.messages[i].viewed_at) {
          userMessages.push(conversation.messages[i])
        }
      }

      return userMessages[0]
    },
    viewConversation(conversationId) {
      axios.put(import.meta.env.VITE_API_URL + 'messages/view', {
        conversation_id: conversationId
      }, {
        headers: { Authorization: `Bearer ${this.token}` }
      }).then(res => {
        this.getConversations()
      })
    },
    sendMessage() {
      axios.post(import.meta.env.VITE_API_URL + 'messages/send', {
        recipient_id: parseInt(this.recipientId),
        text: this.newMessageText
      }, {
        headers: { Authorization: `Bearer ${this.token}` }
      }).then(res => {
        this.newMessageText = ''
        this.getConversations(this.recipientId)
      })
    },
    scrollToMessagesBottom() {
      let el = this.$refs.messages

      if(el) {
        el.scrollTop = el.scrollHeight
      }
    },
    getConversationRecipient(conversation) {
      let recipient = {}

      for(let i = 0; i < conversation.members.length; i++) {
        if (conversation.members[i].id !== this.user.id) {
          recipient = conversation.members[i]
          break
        }
      }

      return recipient
    },
    newMessageUpdate(res) {
      if(this.selectedConversationIndex !== -1) {
        if(res.data.data.conversation_id === this.conversations[this.selectedConversationIndex].id) {
          this.viewConversation(this.conversations[this.selectedConversationIndex].id)
          return
        }

        this.getConversations()
      }
    },
    getNewRecipient(id){
      NProgress.start();

      axios.get(import.meta.env.VITE_API_URL + 'users/' + id, {
        headers: { Authorization: `Bearer ${this.token}` }
      }).then(res => {
        this.newRecipient = res.data.data
        document.title = 'Диалог с ' + this.newRecipient.username

        NProgress.done();
      }).catch(err => {
        this.recipientId = 0
        this.$router.push({ name: 'messages' })
      })
    },
    getConversations(afterNewRecipient = 0) {
      NProgress.start();

      axios.get(import.meta.env.VITE_API_URL + 'messages/conversations', {
        headers: { Authorization: `Bearer ${this.token}` }
      }).then(res => {
        if(res.data.data) {
          this.conversations = res.data.data
        }

        if(afterNewRecipient !== 0){
          this.setDialogWithUser(afterNewRecipient)
        }

        NProgress.done();
      })
    },
    checkIsOnline(onlineDate) {
      let onlineDateObj = moment(onlineDate).utcOffset(+6, true)
      let currentDateObj = moment()

      if (onlineDateObj.diff(currentDateObj, 'minutes') > -15) {
        return true
      }

      return false
    },
  },
  watch: {
    recipientId() {
      if(this.recipientId !== null) {
        this.getConversations(this.recipientId)
      } else {
        this.selectedConversationIndex = -1
        document.title = 'Сообщения'

        this.getConversations()
      }
    }
  }
}
</script>