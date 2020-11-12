<template>
    <q-form
      @submit="onSubmit"
      @reset="onReset"
      class="q-gutter-md"
    > <div class="q-pa-md row items-start q-gutter-md">
      <q-card class="bg-grey-3 relative-position card-example" style="width:500px">
        <q-card-section>
          <q-input
            filled
            v-model="stpData.NerFinal.EMAIL.Value"
            label="Почтовый адрес"
          />
        </q-card-section>
        <q-card-section>
          <q-input
            filled
            v-model="stpData.NerFinal.PERSON.Value"
            label="Имя заявителя"
          />
        </q-card-section>
        <q-card-section>
          <q-input
            filled
            v-model="stpData.NerFinal.PHONE.Value"
          />
        </q-card-section>
        <q-card-section>
          <q-input
            filled
            v-model="stpData.NerFinal.GPE.Value"
            label="Регион"
            :hint="stpData.NerFinal.GPE.RawVal"
          />
        </q-card-section>
      </q-card>
<!--      <q-card class="bg-grey-3 relative-position card-example" style="width:500px">-->
<!--        <div class="q-pa-md" style="max-width: 500px">-->
<!--          <Maximized :Vnc="stpData.Vnc"></Maximized>-->
<!--        </div>-->
<!--      </q-card>-->
    </div>
      <div>
        <q-btn label="Заполнить" color="primary" @click="sentRegister(stpData.LetterID, 'prod', 'fill')"/>
        <q-btn label="Отменить" type="reset" color="primary" flat class="q-ml-sm" />
        <q-btn label="Тест" color="primary" @click="sentRegister(stpData.LetterID, 'dev', 'fill')"/>
      </div>
    </q-form>
</template>

120
<script>
import vncScreen from './vncScreen'
import Maximized from './Maximized'

export default {
  name: 'regTicket',
  props: ['stpData'],
  data () {
    return {
      maximum: false,
      longerNerFinal: {},
      longerSysRating: {},
      longerSysComment: {}
    }
  },
  components: {
    vncScreen,
    Maximized
  },
  methods: {
    onSubmit () {
      if (this.accept !== true) {
        this.$q.notify({
          color: 'red-5',
          textColor: 'white',
          icon: 'warning',
          message: 'You need to accept the license and terms first'
        })
      } else {
        this.$q.notify({
          color: 'green-4',
          textColor: 'white',
          icon: 'cloud_done',
          message: 'Submitted'
        })
      }
    },
    onReset () {
      this.name = null
      this.age = null
      this.accept = false
    },
    sentRegister: function (id, server, steps, vnc, rec) {
      console.log('На регистрацию')
      let ticket = {
        LetterID: id,
        Server: server,
        Step: steps,
        Vnc: vnc,
        Rec: rec
      }
      console.log(this.stpData.NerFinal)
      let check = function (foo) {
        let bad = true
        if (foo) {
          if (foo !== '') {
            bad = false
          }
        }
        if (bad) {
          console.log('Не введены все поля для регистрации')
          return false
        }
        return true
      }
      if (check(this.stpData.NerFinal.EMAIL.Value)) {
        ticket.Email = this.stpData.NerFinal.EMAIL.Value
      } else {
        this.showNotif('top', 'negative', 'Не введены все поля для регистрации: Почта', 'report_problem')
        return
      }
      if (check(this.stpData.NerFinal.PERSON.Value)) {
        ticket.Fio = this.stpData.NerFinal.PERSON.Value
      } else {
        this.showNotif('top', 'negative', 'Не введены все поля для регистрации: Имя заявителя', 'report_problem')
        return
      }
      if (check(this.stpData.NerFinal.PHONE.Value)) {
        ticket.Phone = this.stpData.NerFinal.PHONE.Value
      } else {
        this.showNotif('top', 'negative', 'Не введены все поля для регистрации: Телефон', 'report_problem')
        return
      }
      if (check(this.stpData.NerFinal.GPE.Value)) {
        ticket.Gpe = this.stpData.NerFinal.GPE.Value
      } else {
        this.showNotif('top', 'negative', 'Не введены все поля для регистрации: Регион', 'report_problem')
        return
      }
      ticket.Rating = this.stpData.SysRating
      ticket.CommentRating = this.stpData.SysComment
      console.log(ticket)
      this.$axios.post('/api/sentRegisterTicket', ticket)
        .then(response => {
          console.log('Отправлена на регистрацию')
          console.log(response)
          this.showNotif('top', 'teal', 'Заявка отправлеа на регистрацию', 'thumb_up')
        })
        .catch(error => {
          console.log(error)
          this.showNotif('top', 'negative', 'Непредвиденная ошибка', 'report_problem')
        })
      setTimeout(function () {
        document.getElementById('finder').click() // Нажатие кнопки обновдения данных
      }, 1000)
    }
  }
}
</script>

<style scoped>

</style>
