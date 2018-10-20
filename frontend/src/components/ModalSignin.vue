<template>
  <div>
    <b-nav-item v-b-modal.signin-form>Signin</b-nav-item>
    <b-modal id="signin-form" title="Sign In" v-model="showModal"
      @shown="focusEmail"
      >
      <b-form v-if="showForm">

        <b-form-group horizontal
          label="Email address:"
          label-text-align="left"
        >
          <b-form-input type="email"
            ref="focusThis"
            v-model="form.email"
            required
            placeholder="Enter email"
          ></b-form-input>
        </b-form-group>

        <b-form-group horizontal
          label="Password:"
          label-text-align="left"
        >
          <b-form-input type="password"
            v-model="form.password"
            required
            placeholder="Enter password"
            v-on:keyup.enter.native="onSubmit"
          ></b-form-input>
        </b-form-group>
        <b-alert :show="showAlert" variant="danger">{{ alertText }}</b-alert>
      </b-form>
        <div slot="modal-footer">
          <b-button variant="primary" @click="onSubmit">Submit</b-button>

          <b-button variant="danger" @click="onReset">Reset</b-button>
        </div>
    </b-modal>
  </div>
</template>

<script>

export default {
  name: 'ModalSignin',
  data() {
    return {
      form: {
        email: '',
        password: '',
      },
      showAlert: false,
      alertText: "",
      showModal: false,
      showForm: true,
    }
  },
  watch: {
    showModal: function (old, n) {
      if (old === true && n === false ) {
        this.showAlert = false
        this.alertText = ""
      }

    }
  },
  methods: {
    focusEmail () {
      this.$refs.focusThis.focus()
    },
    onSubmit (evt) {
      evt.preventDefault();
      this.$store.dispatch("postSignin", JSON.stringify(this.form))
      .then((response) => {
        console.log(response);
        this.showModal = false
        this.$store.dispatch('getUser')
      })
      .catch((error) => {
        this.alertText = error.response.data
        this.showAlert = true
      })
    },
    onReset (evt) {
      evt.preventDefault();
      /* Reset our form values */
      this.form.email = '';
      this.form.password = '';
      /* Trick to reset/clear native browser form validation state */
      this.showForm = false;
      this.$nextTick(() => { this.showForm = true });
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.btn {
  margin-right: 10px;
}
</style>
