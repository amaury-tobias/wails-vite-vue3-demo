<template>
  <h1>{{ msg }}</h1>
  <h2>Message from Go: {{ result }}</h2>

  <button @click="count++">count is: {{ count }}</button>
  <br>
  <input v-model="message" type="text" placeholder="Send text to Wails (Go)" />
  <br>
  <button @click="callHi">call 'Greet' on wails</button>
</template>

<script lang="ts">
import { ref, defineComponent } from 'vue'
export default defineComponent({
  name: 'HelloWorld',
  props: {
    msg: {
      type: String,
      required: true
    }
  },
  setup: () => {
    const count = ref(0)
    const message = ref('')
    const result = ref('')

    const callHi = async () => {
      // @ts-ignore
      result.value = await window.backend.main.app.Greet(message.value)
    }
    return { count, result, callHi, message }
  }
})
</script>

<style scoped>
a {
  color: #42b983;
}

label {
  margin: 0 0.5em;
  font-weight: bold;
}

code {
  background-color: #eee;
  padding: 2px 4px;
  border-radius: 4px;
  color: #42b983;
}
</style>
