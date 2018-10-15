import { mount, createLocalVue } from '@vue/test-utils'
import ModalSignup from '@/components/ModalSignup.vue'
import BootstrapVue from 'bootstrap-vue'

const localVue = createLocalVue()
localVue.use(BootstrapVue)

describe('NavBar.vue', () => {
  const wrapper = mount(ModalSignup, {
    localVue,
    propsData: {  },
    mocks: {
      // $http: {get:
      //   jest.fn(() => Promise.resolve({data: {}}))
      // },
      // $store: {
      //   getters: {
      //     data: "hi"
      //   }
      // }
    }
  })
  it('is shown when clicked', (done) => {
    expect(wrapper.vm.showModal).toBe(false)
    wrapper.find(".nav-link").trigger("click")
    wrapper.vm.$nextTick(() => {
      expect(wrapper.vm.showModal).toBe(true)
      done()
    })
    
  })
})
