<template>
  <div>
    <h1 class="title">
      {{ $route.params.domain }}
    </h1>
    <h2 class="subtitle">
      Related information
    </h2>
    <div class="card">
      <header class="card-header">
        <p class="card-header-title">
          Whois Result
        </p>
        <a href="#" class="card-header-icon" aria-label="more options">
          <span class="icon">
            <i class="fas fa-angle-down" aria-hidden="true" />
          </span>
        </a>
      </header>
      <div class="card-content">
        <div class="content">
          <template>
            <pre>{{ whoisResp.info }}</pre>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  async asyncData (context) {
    const DOMAIN_NAME = context.params.domain
    const API_SERVER = process.env.apiServer
    const { data } = await axios.get(
      API_SERVER + '/api/v1/whois/' + DOMAIN_NAME
    )
    return {
      whoisResp: data
    }
  }
}
</script>
