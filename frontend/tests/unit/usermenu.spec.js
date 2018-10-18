import { mount, createLocalVue } from '@vue/test-utils'
import UserMenu from '@/components/UserMenu.vue'
import BootstrapVue from 'bootstrap-vue'

const localVue = createLocalVue()
localVue.use(BootstrapVue)

describe('UserMenu.vue', () => {
  const wrapper = mount(UserMenu, {
    localVue,
    propsData: {  },
    mocks: {
      $store: {
        dispatch() {
          return new Promise(function(resolve) {
            resolve('foo')
          })
        },
        getters : { firstName: "daffy" }
      }
    }
  })
  it('is displays name when store is set', () => {
    console.log(wrapper.html())
    expect(wrapper.find(".nav-link").text()).toBe("daffy")
  })
})
