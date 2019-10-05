<template>
  <section class="section">
    <div class="columns is-mobile" />
    <hero />
    <div id="inputfield">
      <b-field position="is-centered">
        <b-input v-model="domain" placeholder="Input Domain or IP address..." type="search" icon="magnify" />
        <a :href="'/whois/' + domain">
          <p class="control">
            <button class="button is-info">
              Lookup
            </button>
          </p>
        </a>
      </b-field>
    </div>
    <description />
  </section>
</template>

<script>
import Hero from '~/components/Hero.vue'

export default {
  components: {
    Hero
  },
  data () {
    return {
      domain: ''
    }
  }
}
</script>

<style>
#inputfield {
  margin: 40px;
}
</style>
