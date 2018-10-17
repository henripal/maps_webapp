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
      mocks: {
        // $http: {get:
        //   jest.fn(() => Promise.resolve({data: {}}))
        // },
        $store: {
          getters: {
            loggedIn: false
          }
        }
      }
    })
    expect(wrapper.find(".navbar-brand").text()).toMatch("NewCo")
  })
})
