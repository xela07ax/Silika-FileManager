<template>
  <i>
    <q-btn round @click="confirm = true" color="deep-orange" icon="delete" size="sm" style="margin-left: 10px"/>
    <q-dialog v-model="confirm" persistent>
      <q-card>
        <q-card-section class="row items-center">
          <q-avatar icon="remove_shopping_cart" color="primary" text-color="white" />
          <span class="q-ml-sm">Это действие нельзя будет отменить</span>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="Отменить" color="primary" v-close-popup />
          <q-btn flat label="Удалить" color="negative" v-close-popup @click="delCom(row.Sid)" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </i>
</template>

<script>
  export default {
    props: ['row','TestId'],
    data () {
      return {
        confirm: false
      }
    },
    methods: {
      delCom(id) {
        let curFName = 'DeleteRepo'
        console.log(this.row)
        console.log(`=>${name}=>${curFName}|COM:Удаляем [id:${id}]`)
        this.$axios.get(`/qapi/setStatus?TestId=${this.TestId}&ReportId=${id}&Data=&Step=DelRepo`)
          .then(response => {
            console.log(`=>${name}=>${curFName}|COM:Запрос удачен`)
            if (!response.data) {
              return
            }
            console.log(response.data)
            let notify = response.data
            this.$store.dispatch('myStore/ExecFunc', {'Name': 'ToNotify', 'Data': notify})
          })
          .catch(error => {
            console.log(`=>${name}=>${curFName}|COM:Запрос не удачен :(`)
            const Notify = {FuncName: "",List: [],Show: true,Status: 1,Text: `Ошибка сервера: ${error.message}`,UpdNum: 0}
            this.$store.commit('myStore/ToNotify', Notify)
            //this.showNotif('top', 'warning', `Ошибка сервера: ${error.message}`, 'remove_shopping_cart')
          })
        .finally(

        )
      },
    }
  }
</script>

<style scoped>

</style>
