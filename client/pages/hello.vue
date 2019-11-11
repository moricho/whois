<template>
  <div>
    <b-field position="is-centered">
      <b-input v-model="name" placeholder="Input Name" type="search" icon="magnify" />
      <button class="button is-hello" @click="sayHello">
        Say Hello!
      </button>
    </b-field>
    <p v-if="message">
      {{ message }}
    </p>
  </div>
</template>

<script>
import { HelloRequest, HelloReply } from '~/proto/hello_pb'
import { GreeterClient } from '~/proto/hello_grpc_web_pb'

export default {
  data () {
    return {
      name: '',
      message: ''
    }
  },
  created: function () {
    this.greeterClient = new GreeterClient('http://localhost:50051', null, null)
  },
  methods: {
    sayHello: function () {
      var req = new HelloRequest()
      req.setName(this.name)
      this.greeterClient.sayHello(req, {}, (err, response) => {
        this.message = response.getMessage()
      })
    }
  }
}
</script>
