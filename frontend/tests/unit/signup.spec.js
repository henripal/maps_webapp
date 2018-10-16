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
      $store: {
        dispatch(a, b) {
          return new Promise(function(resolve, reject) {
            resolve('foo')
          })
        },
        actions : {
          postSignup(ctx, payload) {
            return new Promise(function(resolve, reject) {
              resolve('foo')
            })
          }
        }
      }
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
  it('it closes after data is given to it', (done) => {
    wrapper.find(".nav-link").trigger("click")
    wrapper.vm.$nextTick(() => {
      wrapper.find("input").value = "name"
      wrapper.find("input").trigger('input') 
      wrapper.find("button").trigger("click")
      wrapper.vm.$nextTick(() => {
        expect(wrapper.vm.showModal).toBe(false)
        done()
      })
    })
  })
  it('sends the right data to vuex', (done) => {
    wrapper.find(".nav-link").trigger("click")
    wrapper.vm.$nextTick(() => {
      wrapper.find("input").element.value = "name"
      wrapper.find("input").trigger('input') 
      console.log(wrapper.find("input").html())
      wrapper.vm.$nextTick(() => {
        wrapper.find(".btn-primary").trigger("click")
        wrapper.vm.$nextTick(() => {
          console.log(wrapper.vm.form)
          expect(wrapper.vm.form.email).toEqual("name")
          done()
        })
      })
    })
  })
})
