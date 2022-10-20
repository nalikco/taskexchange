<template>
  <div class="grid grid-cols-1 md:grid-cols-4 lg:grid-cols-6 mb-7 gap-2 mx-3 md:mx-0">
    <div class="flex text-center">
      <RouterLink :to="{ name: 'ap' }" class="w-full font-medium py-3 px-2 rounded-lg" :class="{'text-slate-700 bg-slate-200': $route.name === 'ap', 'shadow bg-blue-200 text-blue-700 hover:bg-blue-300 hover:shadow-lg transition duration-300': $route.name !== 'ap'}">
        Статистика
      </RouterLink>
    </div>
    <div class="flex text-center">
      <RouterLink :to="{ name: 'ap-users' }" class="w-full font-medium py-3 px-2 rounded-lg" :class="{'text-slate-700 bg-slate-200': $route.name === 'ap-users', 'shadow bg-blue-200 text-blue-700 hover:bg-blue-300 hover:shadow-lg transition duration-300': $route.name !== 'ap-users'}">
        Пользователи
        <p class="text-xs text-blue-500 -mt-1">управление</p>
      </RouterLink>
    </div>
    <div class="flex text-center">
      <RouterLink :to="{ name: 'ap-options' }" class="w-full font-medium py-3 px-2 rounded-lg" :class="{'text-slate-700 bg-slate-200': $route.name === 'ap-options', 'shadow bg-blue-200 text-blue-700 hover:bg-blue-300 hover:shadow-lg transition duration-300': $route.name !== 'ap-options'}">
        Категории и опции
        <p class="text-xs text-blue-500 -mt-1">управление</p>
      </RouterLink>
    </div>
    <div class="flex text-center">
      <RouterLink :to="{ name: 'ap-tasks' }" class="w-full font-medium py-3 px-2 rounded-lg" :class="{'text-slate-700 bg-slate-200': $route.name === 'ap-tasks', 'shadow bg-blue-200 text-blue-700 hover:bg-blue-300 hover:shadow-lg transition duration-300': $route.name !== 'ap-tasks'}">
        Задачи
        <p class="text-xs text-blue-500 -mt-1">управление</p>
      </RouterLink>
    </div>
    <div class="flex text-center">
      <RouterLink :to="{ name: 'ap-orders' }" class="w-full font-medium py-3 px-2 rounded-lg" :class="{'text-slate-700 bg-slate-200': $route.name === 'ap-orders', 'shadow bg-blue-200 text-blue-700 hover:bg-blue-300 hover:shadow-lg transition duration-300': $route.name !== 'ap-orders'}">
        Заказы
        <p class="text-xs text-blue-500 -mt-1">управление</p>
      </RouterLink>
    </div>
    <div class="flex text-center">
      <RouterLink :to="{ name: 'ap-posts' }" class="w-full font-medium py-3 px-2 rounded-lg" :class="{'text-slate-700 bg-slate-200': $route.name === 'ap-posts', 'shadow bg-blue-200 text-blue-700 hover:bg-blue-300 hover:shadow-lg transition duration-300': $route.name !== 'ap-posts'}">
        Публикации
        <p class="text-xs text-blue-500 -mt-1">управление</p>
      </RouterLink>
    </div>
  </div>
</template>