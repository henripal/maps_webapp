<template>

  <b-jumbotron header="Your City From Space" lead="Enter your address or the name of your city (or any other place or landmark); we'll find a satellite image of the neighborhood." >
    <b-form @submit="onSubmit" inline>
        <!-- <label class="sr-only" label-for="address">Your Address:</label> -->
        <b-input class="mb-2 mr-sm-2 mb-sm-0" id="address" v-model="address" placeholder="Your Address" />
        <b-button type="submit" variant="primary">Submit</b-button>
      </b-form>
    <div v-if="image==null">
      <p>Using only this image, we'll tell you which country we think your city is in!</p>
      <p>More fun stuff to come.</p>
    </div>
    <div v-else>
      <br>
      <b-container>
            <b-card-group deck fluid>
              <b-card style="max-width: 20rem;" title="Your City:">
                <b-img center thumbnail v-bind:src="'data:image/png;base64,'+  this.image" />
              </b-card>
              <b-card style="max-width: 20rem;" title="Looks like other cities in:">
                <b-alert show variant="primary">
                  {{ prediction }}
                </b-alert>
                <p>This was determined only by looking at the satellite picture on the left.</p>
              </b-card>
            </b-card-group>
      </b-container>
    </div>
  </b-jumbotron>
</template>

<script>


export default {
  name: 'home',
  data() {
    return {
        address: "",
        image: null,
        prediction: ""
    }
  },
  methods: {
     onSubmit (evt) {
      evt.preventDefault();
      this.$http.get(process.env.VUE_APP_BACKEND_URL + "maps", {
         responseType: 'arraybuffer',
        params : {location: this.address}
      }).then( (res) => {
        this.image = new Buffer(res.data, 'binary').toString('base64')
        let fd= new FormData()
        fd.append('file', this.image)
        this.$http.post(process.env.VUE_APP_MODEL_BACKEND_URL + "upload", fd
        ).then( (res) =>{
          this.prediction = res.data.predictions
        }).catch((err) => {
          console.log("error")
        })
      })
    },
  },
  components: {
  }, 
  mounted() {
  }
}
</script>

<style>
.jumbotron {
  height: 100vh !important; 
}
</style>