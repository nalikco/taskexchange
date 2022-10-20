<script setup>
import AdminMenu from '@/components/AdminMenu.vue'
</script>
<template>
  <main>
    <header class="bg-white shadow">
      <div class="mx-auto max-w-7xl py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold tracking-tight text-gray-900">Создание публикации</h1>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <AdminMenu />
        <div class="bg-white rounded-md shadow-md py-5 px-7">
          <div class="mb-6">
            <label for="category" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-400">Выберите категорию</label>
            <select id="category" @input="onEditFormChange($event, 'category')" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
              <option v-for="categoryDb in categories" :value="categoryDb.id" :selected="categoryDb.id === category">{{ categoryDb.title }}</option>
            </select>
          </div>
          <div class="mt-6 mb-6">
            <label for="username" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Заголовок</label>
            <input type="text" @input="onEditFormChange($event, 'title')" name="username" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" :value="title" placeholder="Введите заголовок" required>
          </div>
          <div class="mt-6 mb-6">
            <label for="username" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Краткий текст</label>
            <QuillEditor ref="shortEditor" theme="snow" contentType="html" placeholder="Введите краткий текст" v-model:content="postContent.short" toolbar="essential" />
          </div>
          <div class="mt-6 mb-6">
            <label for="username" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Текст</label>
            <QuillEditor ref="textEditor" theme="snow" contentType="html" placeholder="Введите текст" v-model:content="postContent.text" toolbar="essential" />
          </div>
          <div class="mt-6 mb-6">
            <input @change="changeImage" accept="image/png, image/jpg, image/jpeg" class="block w-full text-sm text-slate-900 bg-gray-100 rounded-lg border border-gray-100 cursor-pointer dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400" aria-describedby="file_input_help" id="file_input" type="file">
            <p class="mt-1 text-sm text-gray-500 dark:text-gray-300" id="file_input_help">только .png, .jpeg, .jpg</p>
          </div>
          <div class="mt-6 mb-6">
            <button v-if="editPost === null" @click="createPost" class="w-full bg-indigo-500 text-white font-semibold py-2 rounded-md">
              Создать
            </button>
            <button v-else @click="savePost" class="w-full bg-indigo-500 text-white font-semibold py-2 rounded-md">
              Сохранить
            </button>
          </div>
        </div>
      </div>
    </main>
  </main>
</template>

<script>
import NProgress from "nprogress";
import axios from "axios"
import {mapState} from "pinia";
import {useUserStore} from "@/stores/user";
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';

export default {
  components: {
    QuillEditor
  },
  data() {
    return {
      categories: [],
      editPost: null,
      category: 0,
      title: '',
      postContent: {
        short: '',
        text: ''
      },
      image: null,
    }
  },
  mounted() {
    document.title = 'Управление пользователями'

    this.getCategories()

    if(this.$route.query.edit_id) {
      this.getEditPost(this.$route.query.edit_id)
    }
  },
  computed: {
    ...mapState(useUserStore, ['user', 'token'])
  },
  methods: {
    onEditFormChange(e, field) {
      this[field] = e.target.value
    },
    changeImage(e) {
      this.image = e.target.files[0]
    },
    getCategories() {
      NProgress.start()

      axios.get(import.meta.env.VITE_API_URL + 'posts/categories/', {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.categories = res.data.data
          this.category = this.categories[0].id

          NProgress.done()
        }
      })
    },
    getEditPost(postId) {
      NProgress.start()

      axios.get(import.meta.env.VITE_API_URL + 'posts/' + postId, {
        headers: { Authorization: `Bearer ${this.token}` },
      }).then(res => {
        if(res.data.data) {
          this.editPost = res.data.data
          this.category = this.editPost.categories[0].id
          this.title = this.editPost.title

          this.$refs.shortEditor.setHTML(this.editPost.short)
          this.$refs.textEditor.setHTML(this.editPost.text)

          NProgress.done()
        }
      }).catch(err => {
        NProgress.done()
      })
    },
    createPost() {
      NProgress.start()

      let data = {
        title: this.title,
        status: 1,
        short: this.postContent.short,
        text: this.postContent.text,
      }

      if(this.category !== 0) data.categories = [parseInt(this.category)]

      axios.post(import.meta.env.VITE_API_URL + 'posts/', data, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        }
      }).then(res => {
        this.updateImage(res.data.id)
      })
    },
    savePost() {
      NProgress.start()

      let data = {
        title: this.title,
        short: this.postContent.short,
        text: this.postContent.text,
      }

      if (this.title !== this.editPost.title) data.title = this.title
      if (this.postContent.short !== this.editPost.short) data.short = this.postContent.short
      if (this.postContent.text !== this.editPost.text) data.text = this.postContent.text
      if(parseInt(this.category) !== this.editPost.categories[0].id) data.categories = [parseInt(this.category)]

      axios.put(import.meta.env.VITE_API_URL + 'posts/' + this.editPost.id, data, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        }
      }).then(res => {
        if(this.image !== null) this.updateImage(this.editPost.id)
        else this.$router.push({ name: 'ap-posts' })
      })
    },
    updateImage(postId) {
      let formData = new FormData();
      formData.append('main_image', this.image);

      axios.put(import.meta.env.VITE_API_URL + 'posts/img/' + postId, formData, {
        headers: {
          Authorization: `Bearer ${this.token}`,
          'Content-Type': 'multipart/form-data'
        }
      }).then(res => {
        this.$router.push({ name: 'ap-posts' })
      })
    }
  },
}
</script>