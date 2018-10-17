<template>
  <div>
    <b-nav-item v-b-modal.signup-form>Signup</b-nav-item>
    <b-modal id="signup-form" title="Sign Up" v-model="showModal">
      <b-form v-if="showForm">

        <b-form-group horizontal
          label="Email address:"
          label-text-align="left"
          :invalid-feedback="emailAlreadyInDatabase"
          :state = "validInput"
        >
          <b-form-input type="email"
            v-model="form.email"
            required
            :state = "validInput"
            placeholder="Enter email"
          ></b-form-input>
        </b-form-group>

        <b-form-group horizontal
          label="First Name:"
          label-text-align="left"
        >
          <b-form-input type="text"
            v-model="form.firstName"
            required
            placeholder="Enter first name"
          ></b-form-input>
        </b-form-group>
        <b-form-group horizontal
          label="Last Name:"
          label-text-align="left"
        >
          <b-form-input type="text"
            v-model="form.lastName"
            required
            placeholder="Enter last name"
          ></b-form-input>
        </b-form-group>
        <b-form-group horizontal
          label="Password:"
          label-text-align="left"
        >
          <b-form-input type="text"
            v-model="form.password"
            required
            placeholder="Enter password"
          ></b-form-input>
        </b-form-group>
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
  name: 'ModalSignup',
  data() {
    return {
      form: {
        email: '',
        firstName: '',
        lastName: '',
        password: '',
      },
      showModal: false,
      showForm: true,
      validInput: null
    }
  },
  props: {
  },
  methods: {
    onSubmit (evt) {
      evt.preventDefault();
      this.$store.dispatch("postSignup", JSON.stringify(this.form))
      .then((response) => {
        this.showModal = false
      })
      .catch((error) => {
        this.validInput = false
        console.log(error.response)})
    },
    onReset (evt) {
      evt.preventDefault();
      /* Reset our form values */
      this.form.email = '';
      this.form.firstName = '';
      this.form.lastName = '';
      this.form.password = '';
      /* Trick to reset/clear native browser form validation state */
      this.showForm = false;
      this.$nextTick(() => { this.showForm = true });
    }
  },
  computed: {
    emailAlreadyInDatabase () {
      if (this.validInput) {
        return ""
      } else {
        return "Username already in database"
      }
    }
  },
  mounted () {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.btn {
  margin-right: 10px;
}
</style>
