import { shallowMount } from '@vue/test-utils'
import HelloWorld from '@/components/HelloWorld.vue'

describe('HelloWorld.vue', () => {
  it('renders props.msg when passed', () => {
    const msg = 'new message'
    const wrapper = shallowMount(HelloWorld, {
      propsData: { msg },
      mocks: {
        $http: {get:
          jest.fn(() => Promise.resolve({data: {}}))
        }
      }
    })
    expect(wrapper.text()).toMatch(msg)
  })
})
