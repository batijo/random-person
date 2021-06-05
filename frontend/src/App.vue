<template>
  <div>
    <Header title="Random Person" />
    <Person :person="person" />
    <Button @click="generateNew" text="random" color="#424963" />
  </div>
</template>

<script>
import Header from './components/Header'
import Button from './components/Button'
import Person from './components/Person'

const version = "v0"

export default {
  name: 'App',
  components: {
    Header,
    Button,
    Person,
  },
  data() {
    return {
      person: Object
    }
  },
  methods: {
    async generateNew() {
      this.person = await this.fetchPerson()
    },

    async fetchPerson() {
      const res = await fetch(`api/${version}/person`)
      // TODO: this is temporary and absolutely not ok
      if (!res.ok) {
        const error = await res.json()
        return { name: error.message }
      }
      const data = await res.json()
      return data
    },
  },
  async created() {
    this.person = await this.fetchPerson()
  },
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400&display=swap');
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}
body {
  font-family: 'Poppins', sans-serif;
}
.container {
  max-width: 500px;
  margin: 30px auto;
  overflow: auto;
  min-height: 300px;
  border: 1px solid steelblue;
  padding: 30px;
  border-radius: 5px;
}
.btn {
  display: inline-block;
  color: #fff;
  border: none;
  padding: 10px 20px;
  margin: 5px;
  border-radius: 5px;
  cursor: pointer;
  text-decoration: none;
  font-size: 15px;
  font-family: inherit;
}
.btn:focus {
  outline: none;
}
.btn:active {
  transform: scale(0.98);
}
.btn-block {
  display: block;
  width: 100%;
}
</style>
