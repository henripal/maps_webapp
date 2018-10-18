import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    loggedIn: false,
    user: null
  },
  mutations: {
    logUserIn(currentState, user) {
      currentState.user = user
      currentState.loggedIn = true
    },
    logUserOut(currentState) {
      currentState.user = null
      currentState.loggedIn = false
    }
  },
  actions: {
    logOut(context) {
      axios({
        method: "get",
        url: process.env.VUE_APP_DATA_URL + "logout",
        withCredentials: true
      }).then(() => {
        context.commit("logUserOut")
      }).catch((error) => {
        console.log(error.response)
      })
    },
    postSignup(context, payload) {
      return axios({
        method: "post",
        url: process.env.VUE_APP_DATA_URL + "signup",
        data: payload,
        withCredentials: true
      })      
    },
    postSignin(context, payload) {
      return axios({
        method: "post",
        url: process.env.VUE_APP_DATA_URL + "signin",
        data: payload,
        withCredentials: true
      })      
    },
    getUser(context) {
      axios({
        method: "get",
        url: process.env.VUE_APP_DATA_URL + "user",
        withCredentials: true
      }).then( (response) => {
        context.commit("logUserIn", response.data)
      })      
    }
  },
  getters: {
    firstName (state) {
      if (state.user) {
        return state.user.firstName
      }
      return ''
    },
    loggedIn (state) {
      return state.loggedIn
    }
  }
})
