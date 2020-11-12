import Vue from 'vue'
import axios from 'axios'
// const BudgetManagerAPI = `http:///172.31.11.240:3022/` // ${window.location.hostname}
// const BudgetManagerAPI = `http://172.31.50.56:30339/` // ${window.location.hostname}
const BudgetManagerAPI = `http://127.0.0.1:30339/` // ${window.location.hostname}

//Vue.prototype.$axios = axios
axios.defaults.baseURL = BudgetManagerAPI
axios.defaults.credentials = true
axios.defaults.withCredentials = false
Vue.prototype.$axios = axios
