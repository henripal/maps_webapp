import { mount, createLocalVue } from '@vue/test-utils'
import ModalSignup from '@/components/ModalSignup.vue'
import BootstrapVue from 'bootstrap-vue'

const localVue = createLocalVue()
localVue.use(BootstrapVue)

describe('ModalSignup.vue', () => {
  const wrapper = mount(ModalSignup, {
    localVue,
    propsData: {  },
    mocks: {
      $store: {
        dispatch() {
          return new Promise(function(resolve) {
            resolve('foo')
          })
        },
        actions : {
          postSignup() {
            return new Promise(function(resolve) {
              resolve('foo')
            })
          }
        }
      }
    }
  })
  it('is shown when clicked', (done) => {
    expect(wrapper.vm.showModal).toBe(false)
    wrapper.find("button.btn").trigger("click")
    wrapper.vm.$nextTick(() => {
      expect(wrapper.vm.showModal).toBe(true)
      done()
    })
  })
  it('it closes after data is given to it', (done) => {
    wrapper.find("button.btn").trigger("click")
    wrapper.vm.$nextTick(() => {
      const input = wrapper.find("input")
      input.element.value = "name"
      input.trigger('input') 
      wrapper.find("button.btn.btn-primary").trigger("click")
      wrapper.vm.$nextTick(() => {
        expect(wrapper.vm.showModal).toBe(false)
        done()
      })
    })
  })
  it('sends the right data to vuex', (done) => {
    wrapper.find("button.btn").trigger("click")
    wrapper.vm.$nextTick(() => {
      const input = wrapper.find("input")
      input.element.value = "name"
      input.trigger('input') 
      wrapper.vm.$nextTick(() => {
        wrapper.find("button.btn.btn-primary").trigger("click")
        wrapper.vm.$nextTick(() => {
          expect(wrapper.vm.form.email).toEqual("name")
          done()
        })
      })
    })
  })
})
