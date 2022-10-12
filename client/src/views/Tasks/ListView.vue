<script setup>
import ProfileInfo from "@/components/ProfileInfo.vue";
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
        tasks
      </div>
    </main>
  </main>
</template>

<script>
import NProgress from "nprogress";
export default {
  mounted() {
    document.title = 'Мои задачи'
    NProgress.done();
  }
}
</script>