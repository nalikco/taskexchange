<script setup>
import {moment} from "@/moment";
</script>
<template>
  <main>
    <section v-if="userProfile" class="w-1/3 my-10 mx-auto bg-[#20354b] rounded-2xl px-8 py-6 shadow-lg">
      <div class="flex items-center justify-between">
        <span class="text-gray-300 text-sm font-semibold">
          {{ getOnline(userProfile.last_online) }}
        </span>
      </div>
      <div class="mt-6 w-fit mx-auto">
        <img src="@/assets/img/user.png" class="rounded-full w-28 " alt="profile picture" srcset="">
      </div>

      <div class="mt-8 ">
        <h2 class="text-white font-bold text-2xl tracking-wide">
          {{ userProfile.username }} <br/>
        </h2>
      </div>
      <p v-if="userProfile.type === 1" class="text-emerald-400 font-semibold mt-2.5" >
        Исполнитель
      </p>
      <p v-if="userProfile.type === 2" class="text-emerald-400 font-semibold mt-2.5" >
        Заказчик
      </p>
      <p v-if="userProfile.type === 3" class="text-emerald-400 font-semibold mt-2.5" >
        Администратор
      </p>
      <div class="mt-3 text-white text-sm">
        <span class="text-gray-400 font-semibold">Зарегистрирован</span>
        <span class="ml-1">{{ moment(userProfile.created_at).utc(0).format('DD MMMM YYYY г.')}}</span>
      </div>
      <RouterLink v-if="userProfile.id !== user.id" :to="{'name': 'messages', query: { recipient_id: userProfile.id }}" class="mt-3 text-white bg-green-500 font-semibold shadow rounded-md py-2 text-base text-sm grid place-items-center">
        Написать
      </RouterLink>
    </section>
    <section v-else class="w-1/3 my-10 mx-auto bg-[#20354b] text-white text-center font-semibold rounded-2xl px-8 py-6 shadow-lg">
      Пользователь не найден.
    </section>
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
      userProfile: null,
    }
  },
  computed: {
    ...mapState(useUserStore, ['user']),
  },
  mounted() {
    document.title = 'Профиль'

    this.getUser(this.$route.params.user_id)
  },
  methods: {
    getUser(userId) {
      axios.get(import.meta.env.VITE_API_URL + 'users/' + userId, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.userProfile = res.data.data

          NProgress.done()
        }
      }).catch(err => {
        this.user = null
        NProgress.done()
      })
    },
    getOnline(onlineDate) {
      let onlineDateObj = moment(onlineDate).utcOffset(+6, true)
      let currentDateObj = moment()

      if (onlineDateObj.diff(currentDateObj, 'minutes') > -15) {
        return 'онлайн'
      }

      return 'был(-а) онлайн ' + onlineDateObj.utcOffset(+6, true).fromNow()
    },
  }
}
</script>