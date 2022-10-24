<script setup>
import {moment} from "@/moment";
</script>
<template>
  <main>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <section class="text-gray-600 body-font">
          <div class="container px-5 mx-auto max-w-7x1">
            <div class="flex flex-wrap w-full mb-4 p-4">
              <div class="w-full mb-6 lg:mb-0">
                <h1 class="sm:text-4xl text-5xl font-medium font-bold title-font mb-2 text-gray-900" v-html="post.title"></h1>
                <div class="h-1 w-40 bg-indigo-500 rounded"></div>
              </div>
            </div>
            <div class="">
              <div v-if="post" class="p-4">
                <div class="bg-white p-6 rounded-lg">
                  <img class="lg:h-96 xl:h-96 md:h-64 sm:h-72 xs:h-72 h-72  rounded w-full object-cover object-center mb-6" :src="mainUrl + post.main_image" alt="Image Size 720x400">
                  <h3 class="tracking-widest text-indigo-500 text-xs font-medium title-font">{{ post.categories[0].title }}</h3>
                  <p class="leading-relaxed text-base mt-2" v-html="post.text"></p>
                  <p class="text-sm text-gray-500 mt-3 text-right">{{ moment(post.created_at).utc(0).format('DD MMMM YYYY г.') }}</p>
                </div>
              </div>
            </div>
          </div>
        </section>
      </div>
    </main>
  </main>
</template>

<script>
import axios from 'axios'
import NProgress from "nprogress"

export default {
  data() {
    return {
      post: null
    }
  },
  computed: {
    mainUrl() {
      return import.meta.env.VITE_URL
    },
  },
  mounted() {
    document.title = 'Новости'

    this.getPost(this.$route.params.id)
  },
  methods: {
    getPost(postId) {
      NProgress.start()

      axios.get(import.meta.env.VITE_API_URL + 'posts/' + postId).then(res => {
        this.post = res.data.data

        NProgress.done()
      }).catch(res => {
        this.$router.push({ name: 'posts' })
      })
    }
  }
}
</script>