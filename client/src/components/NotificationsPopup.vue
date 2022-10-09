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
        <RouterLink :to="{'name': 'home'}" @click="showNotifications = false" class="block px-4 py-2 text-sm text-gray-700 transition duration-300 hover:bg-gray-200" role="menuitem" tabindex="-1" id="user-menu-item-0">
          Новое предложение по заказу
          <br>
          <span class="text-gray-500 text-xs">Сегодня в 15:56</span>
        </RouterLink>
        <RouterLink :to="{'name': 'home'}" @click="showNotifications = false" class="block px-4 py-2 text-sm text-gray-700 transition duration-300 hover:bg-gray-200" role="menuitem" tabindex="-1" id="user-menu-item-0">
          Заказ сдан на проверку
          <br>
          <span class="text-gray-500 text-xs">Сегодня в 15:02</span>
        </RouterLink>
        <div class="flex items-center">
          <RouterLink :to="{'name': 'home'}" @click="showNotifications = false" class="text-sm mx-4 w-full my-2 bg-green-500 hover:bg-green-700 shadow-md hover:shadow-none transition duration-300 text-center text-white font-medium py-1 rounded-lg">Все уведомления</RouterLink>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  props: ['showNotifications', 'isMobile']
}
</script>