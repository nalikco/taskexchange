<template>
  <div class="mx-3 md:mx-0 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
    <div class="bg-white shadow rounded-lg text-sm py-2 px-2 hover:shadow-xl transition duration-300">
      <div class="float-left ml-2 mt-1">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-14 h-14 text-gray-800">
          <path fill-rule="evenodd" d="M18.685 19.097A9.723 9.723 0 0021.75 12c0-5.385-4.365-9.75-9.75-9.75S2.25 6.615 2.25 12a9.723 9.723 0 003.065 7.097A9.716 9.716 0 0012 21.75a9.716 9.716 0 006.685-2.653zm-12.54-1.285A7.486 7.486 0 0112 15a7.486 7.486 0 015.855 2.812A8.224 8.224 0 0112 20.25a8.224 8.224 0 01-5.855-2.438zM15.75 9a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z" clip-rule="evenodd" />
        </svg>
      </div>
      <div class="ml-20 py-3">
        {{ user.username }}
        <br>
        <span class="text-gray-500 font-medium">id:{{ user.id }}</span>
      </div>
    </div>
    <div class="bg-white shadow rounded-lg text-sm py-2 px-2 hover:shadow-xl transition duration-300">
      <div class="float-left ml-2 mt-1">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-14 h-14 text-lime-500">
          <path d="M12 7.5a2.25 2.25 0 100 4.5 2.25 2.25 0 000-4.5z" />
          <path fill-rule="evenodd" d="M1.5 4.875C1.5 3.839 2.34 3 3.375 3h17.25c1.035 0 1.875.84 1.875 1.875v9.75c0 1.036-.84 1.875-1.875 1.875H3.375A1.875 1.875 0 011.5 14.625v-9.75zM8.25 9.75a3.75 3.75 0 117.5 0 3.75 3.75 0 01-7.5 0zM18.75 9a.75.75 0 00-.75.75v.008c0 .414.336.75.75.75h.008a.75.75 0 00.75-.75V9.75a.75.75 0 00-.75-.75h-.008zM4.5 9.75A.75.75 0 015.25 9h.008a.75.75 0 01.75.75v.008a.75.75 0 01-.75.75H5.25a.75.75 0 01-.75-.75V9.75z" clip-rule="evenodd" />
          <path d="M2.25 18a.75.75 0 000 1.5c5.4 0 10.63.722 15.6 2.075 1.19.324 2.4-.558 2.4-1.82V18.75a.75.75 0 00-.75-.75H2.25z" />
        </svg>
      </div>
      <div class="ml-20 py-3">
        Баланс
        <br>
        <span class="text-gray-500 font-medium">{{ $filters.currencyFormat(user.balance) }}</span>
      </div>
    </div>
    <div class="bg-white shadow rounded-lg text-sm py-2 px-2 hover:shadow-xl transition duration-300">
      <div class="float-left ml-2 mt-1">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-14 h-14 text-blue-500">
          <path fill-rule="evenodd" d="M9.315 7.584C12.195 3.883 16.695 1.5 21.75 1.5a.75.75 0 01.75.75c0 5.056-2.383 9.555-6.084 12.436A6.75 6.75 0 019.75 22.5a.75.75 0 01-.75-.75v-4.131A15.838 15.838 0 016.382 15H2.25a.75.75 0 01-.75-.75 6.75 6.75 0 017.815-6.666zM15 6.75a2.25 2.25 0 100 4.5 2.25 2.25 0 000-4.5z" clip-rule="evenodd" />
          <path d="M5.26 17.242a.75.75 0 10-.897-1.203 5.243 5.243 0 00-2.05 5.022.75.75 0 00.625.627 5.243 5.243 0 005.022-2.051.75.75 0 10-1.202-.897 3.744 3.744 0 01-3.008 1.51c0-1.23.592-2.323 1.51-3.008z" />
        </svg>
      </div>
      <div class="ml-20 py-3">
        Задачи
        <br>
        <span class="text-gray-500 font-medium">{{ user.activeTasksCount }} {{ $filters.declOfNum(user.activeTasksCount, ['активная', 'активных', 'активных']) }}</span>
      </div>
    </div>
    <div class="bg-white shadow rounded-lg text-sm py-2 px-2 hover:shadow-xl transition duration-300">
      <div class="float-left ml-2 mt-1">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-14 h-14 text-red-500">
          <path fill-rule="evenodd" d="M12.963 2.286a.75.75 0 00-1.071-.136 9.742 9.742 0 00-3.539 6.177A7.547 7.547 0 016.648 6.61a.75.75 0 00-1.152-.082A9 9 0 1015.68 4.534a7.46 7.46 0 01-2.717-2.248zM15.75 14.25a3.75 3.75 0 11-7.313-1.172c.628.465 1.35.81 2.133 1a5.99 5.99 0 011.925-3.545 3.75 3.75 0 013.255 3.717z" clip-rule="evenodd" />
        </svg>

      </div>
      <div class="ml-20 py-3">
        Активность на сайте
        <br>
        <span class="text-gray-500 font-medium">{{ activity }} {{ $filters.declOfNum(activity, ['день', 'дня', 'дней']) }}</span>
      </div>
    </div>
  </div>
</template>

<script>
import {mapState} from "pinia";
import {useUserStore} from "@/stores/user";

export default {
  computed: {
    ...mapState(useUserStore, ['user']),
    activity() {
      let secondsFromUserCreating = (new Date().getTime() - new Date(new Date(this.user.created_at).toDateString()).getTime()) / 1000

      return Math.round(secondsFromUserCreating / 60 / 60 / 24)
    }
  }
}
</script>