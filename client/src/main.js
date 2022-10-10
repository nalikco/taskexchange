import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

import './assets/main.css'

const app = createApp(App)

app.config.globalProperties.$filters = {
    declOfNum(value, text_forms) {
        value = Math.abs(value) % 100;
        let n1 = value % 10;
        if (value > 10 && value < 20) { return text_forms[2]; }
        if (n1 > 1 && n1 < 5) { return text_forms[1]; }
        if (n1 === 1) { return text_forms[0]; }

        return text_forms[2];
    },
    currencyFormat(value) {
        return new Intl.NumberFormat().format(value) + ' руб.'
    },
    dateToHuman(date) {
        // date = new Date(date.toLocaleString() + " GMT+06:00")
        // let time = date.toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'})
        // if (date.toLocaleDateString() === new Date().toLocaleDateString()) return 'Сегодня в ' + time;
        //
        // let yesterday = new Date()
        // yesterday.setDate(yesterday.getDate() - 1)
        // if (date.toLocaleDateString() === yesterday.toLocaleDateString()) return 'Вчера в ' + time;

        let dateString = ""

        return dateString
    }
}

app.use(createPinia())
app.use(router)

app.mount('#app')
