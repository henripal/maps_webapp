import { shallowMount } from '@vue/test-utils'
import HelloWorld from '@/components/HelloWorld.vue'

describe('HelloWorld.vue', () => {
  it('renders props.msg when passed', () => {
    const msg = 'newmessage'
    const wrapper = shallowMount(HelloWorld, {
      propsData: { "wowi" },
      mocks: {
        $http: {get:
          jest.fn(() => Promise.resolve({data: {}}))
        },
        $store: {
          getters: {
            data: "hi"
          }
        }
      }
    })
    expect(wrapper.text()).toMatch(msg)
  })
})
