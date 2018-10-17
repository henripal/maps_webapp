import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    data: null
  },
  mutations: {
    setData(currentState, data) {
      currentState.data = data
    }
  },
  actions: {
    getData(context) {
      axios.get(process.env.VUE_APP_DATA_URL).then((response) =>{
        context.commit("setData", response.data)
      })
    },
    postSignup(context, payload) {
      return axios.post(process.env.VUE_APP_DATA_URL + "signup", payload)      
    }
  },
  getters: {
    data (state) {
      return state.data
    }
  }
})
