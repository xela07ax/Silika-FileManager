<template>
  <form @submit.prevent="simulateSubmit" class="q-pa-md">
    <div class="q-gutter-md" style="max-width: 300px">
      <q-input filled v-model="name" label="Название"/>
      <q-input
        v-model="code"
        filled
        label="Вы можете ввести код теста"
        stack-label
        hint="не обязательно*"
        type="textarea"
      />
      <div class="row justify-end">
        <q-btn
          type="submit"
          :loading="submitting"
          label="Создать"
          class="q-mt-md"
          color="teal"
        >
          <template v-slot:loading>
            <q-spinner-facebook/>
          </template>
        </q-btn>
      </div>
    </div>
  </form>
</template>

<script>
  export default {
    data() {
      return {
        name: '',
        code: '',
        lang: 'Java',
        submitting: false
      }
    },
    methods: {
      simulateSubmit() {
        this.submitting = true
        let test = {
          Code: this.code,
          Lang: this.lang,
          Name: this.name,
        }
        this.$store.dispatch('myStore/ExecFunc', {'Name': 'CreateTest', 'Data': test})
        setTimeout(() => {
          // delay simulated, we are done,
          // now restoring submit to its initial state
          //this.submitting = false
          this.$router.push({ path: `/` })
        }, 2000)
      },
      onItemClick(lang) {
        this.lang = lang
      }
    },
  }
</script>
