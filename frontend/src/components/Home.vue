<template>

  <b-jumbotron header="Your City From Space" lead="Enter your address or the name of your city (or any other place or landmark); we'll find a satellite image of the neighborhood." >
    <b-form @submit="onSubmit" inline>
        <!-- <label class="sr-only" label-for="address">Your Address:</label> -->
        <b-input class="mb-2 mr-sm-2 mb-sm-0" id="address" v-model="address" placeholder="Your Address" />
        <b-button type="submit" variant="primary">Submit</b-button>
      </b-form>
    <div v-if="firstsubmit==false">
      <p>Using only this image, we'll tell you which country we think your city is in!</p>
      <p>More fun stuff to come.</p>
    </div>
    <div v-else>
      <br>
      <b-container>
            <b-card-group deck fluid>
              <b-card style="max-width: 20rem;" title="Your City:">
                <b-alert v-if="status0 != ''" show variant="primary">{{ status1 }}</b-alert>
                <b-img v-if="image!=null" center thumbnail v-bind:src="'data:image/png;base64,'+  this.image" />
              </b-card>
              <b-card style="max-width: 20rem;" title="Looks like other cities in:">
                <b-alert v-if="status1 != ''" show variant="primary">{{ status1 }}</b-alert>
                <b-alert v-if="prediction != ''" show variant="primary">
                  {{ prediction }}
                </b-alert>
                <p>This was determined only by looking at the satellite picture on the left.</p>
              </b-card>
              <b-card style="max-width: 20rem;" title="Feature Heatmap">
                <b-alert v-if="status2 != ''" show variant="primary">{{ status2 }}</b-alert>
                 <b-img v-if="heatmap!=null" center thumbnail width="234" height="234" v-bind:src="'data:image/png;base64,'+  this.heatmap" />
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
        prediction: "",
        heatmap: null,
        status0: "",
        status1: "",
        status2: "",
        firstsubmit: false
    }
  },
  methods: {
     onSubmit (evt) {
       this.firstsubmit = true
       this.prediction = ""
       this.image = null
       this.heatmap = null
       this.status0= "Loading . . ."
       this.status1= "Loading . . ."
       this.status2= "Loading . . ."
      evt.preventDefault();
      this.$http.get(process.env.VUE_APP_BACKEND_URL + "maps", {
         responseType: 'arraybuffer',
        params : {location: this.address}
      }).then( (res) => {
        this.image = new Buffer(res.data, 'binary').toString('base64')
        this.status0 = ""
        let fd= new FormData()
        fd.append('file', this.image)
        this.$http.post(process.env.VUE_APP_MODEL_BACKEND_URL + "upload", fd
        ).then( (res) =>{
          this.status1 = ""
          this.prediction = res.data.predictions
          this.$http.post(process.env.VUE_APP_MODEL_BACKEND_URL + "heatmap", fd)
          .then((res) => {
            this.status2 = ""
            this.heatmap = res.data.predictions
            console.log(res)
          })
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