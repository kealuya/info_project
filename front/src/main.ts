import { createApp } from 'vue'
// @ts-ignore
import { createPinia} from 'pinia'
import piniaPluginPersist from 'pinia-plugin-persist';
import {useInvoiceData, userInfoData} from "./store";
import App from './App.vue'
import Vant from 'vant';
import router from './router'
import 'vant/es/notify/style';
import 'vant/es/dialog/style';
import './assets/main.css'
// @ts-ignore
import VueHashCalendar from 'vue3-hash-calendar';
import 'vue3-hash-calendar/es/index.css';
const app = createApp(App)
app.use(createPinia().use(piniaPluginPersist))
app.use(router)
app.use(Vant)
app.use(VueHashCalendar);
app.provide('userInfo', userInfoData());
app.provide('invoiceData', useInvoiceData());
app.mount('#app')
