import { mount, createLocalVue } from '@vue/test-utils'
import NavBar from '@/components/NavBar.vue'
import BootstrapVue from 'bootstrap-vue'

const localVue = createLocalVue()
localVue.use(BootstrapVue)

describe('NavBar.vue', () => {
  it('renders the title properly', () => {
    const wrapper = mount(NavBar, {
      localVue,
      propsData: {  },
      mocks: { $store: { getters: { loggedIn: false }}}
    })
    expect(wrapper.find(".navbar-brand").text()).toMatch("NewCo")
  }),
  it("shows signup and signin when user isn't logged in", () => {
    const wrapper = mount(NavBar, {
      localVue,
      propsData: {  },
      mocks: { $store: { getters: { loggedIn: false }}}
    })
    expect(wrapper.find(".nav-link").text()).toMatch("Signin")
  })
  it("has two elements when user is not logged in", () => {
    const wrapper = mount(NavBar, {
      localVue,
      propsData: {  },
      mocks: { $store: { getters: { loggedIn: true }}}
    })
    expect(wrapper.findAll(".nav-link").length).toBe(2)
  })
})
