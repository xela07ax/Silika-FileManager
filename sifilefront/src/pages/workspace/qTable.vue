<template>
  <div class="q-pa-md">
    <!--    <q-btn id="finder"-->
    <!--           :loading="loading3"-->
    <!--           :percentage="percentage3"-->
    <!--           dark-percentage-->
    <!--           unelevated-->
    <!--           color="orange"-->
    <!--           text-color="grey-9"-->
    <!--           @click="startComputing(3)"-->
    <!--           icon="youtube_searched_for"-->
    <!--           style="width: 200px; color: white"-->
    <!--    >Обновить</q-btn>-->
<!--{{load}}-->
<!--    <vue-friendly-iframe src="http://172.31.50.74:8080/#/" @load="loadCheck()"></vue-friendly-iframe>-->
<!--    <q-btn @click="getList()" style="width: 200px;margin-left: 10px">Показать сообщения</q-btn>-->
<!--    <q-btn @click="addTest('Test1')" style="width: 200px;margin-left: 10px">Запустить</q-btn>-->
    <!--    <q-btn @click="updateData()" style="width: 200px;margin-left: 10px">Dep-Обновление</q-btn>-->
    <!--    {{this.$store.state.myStore.list}}-->
    <br><br>
    <q-table
      title="Тесты"
      :data="this.$store.state.myStore.list"
      :columns="columns"
      row-key="Id"
    >
      <template v-slot:header="props">
        <q-tr :props="props">
          <q-th auto-width></q-th>
          <q-th
            v-for="col in props.cols"
            :key="col.name"
            :props="props"
          >
            {{ col.label }}
          </q-th>
        </q-tr>
      </template>

      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td auto-width>
            <q-btn size="sm" round @click="addTest(props.row)" color="deep-orange" icon="play_arrow"  style="margin-left: 10px;margin-right: 20px"/>
<!--            <q-btn @click="addTest(props.row)" style="width: 200px;margin-left: 10px">Запустить</q-btn>-->
            <q-btn size="sm" color="accent" round dense @click="addCom(props)" :icon="props.expand ? 'remove' : 'add'"></q-btn>
          </q-td>
          <q-td
            v-for="col in props.cols"
            :key="col.name"
            :props="props"
          >
            <div v-if="col.name === 'Status'">
              <!--0=ok,1=error,4=runing-->
              <q-linear-progress v-show="props.row.Status === 4" dark query color="cyan" class="q-mt-sm" />
              <q-icon v-show="props.row.Status === 0" name="done_outline" class="text-positive" style="font-size: 1.4rem;" />
              <q-icon v-show="props.row.Status === 1" name="warning" class="text-red" style="font-size: 1.4rem;" />
            </div>
            <div v-else>{{ col.value}}</div>
          </q-td>
        </q-tr>
        <q-tr v-if="props.expand" :props="props">
          <q-td colspan="100%">
            <!--            <q-input-->
            <!--              v-model="!disabledStore[props.row.Id]?props.row.code:dised[props.row.Id].code"-->
            <!--              filled-->
            <!--              label="Вы можете ввести код теста"-->
            <!--              stack-label-->
            <!--              hint="не обязательно*"-->
            <!--              type="textarea"-->
            <!--              :disable="!disabledStore[props.row.Id]"-->
            <!--            />-->
            <q-card class="my-card" flat bordered>
              <q-card-section>
                <q-input  class="text-h5 q-mt-sm q-mb-xs" filled v-model="!disabledStore[props.row.Id]?props.row.Name:dised[props.row.Id].Name"
                         :disable="!disabledStore[props.row.Id]" label="Название"/>
              </q-card-section>

              <q-card-actions>
<!--                <q-btn class="q-mt-md" @click="addCom(props.row)" style="margin-left: 10px">Редактировать</q-btn>-->
                <q-btn
                  :loading="submitting"
                  label="Сохранить"
                  class="q-mt-md"
                  color="teal"
                  style="margin-left: 10px"
                  @click="editPush(props.row.Id)"
                >
                  <template v-slot:loading>
                    <q-spinner-facebook/>
                  </template>
                </q-btn>
                <deleter :row="props.row"></deleter>
                <q-btn round @click="addTest(props.row)" color="deep-orange" icon="play_arrow" class="q-mt-md"  style="margin-left: 10px"/>
                <q-toggle style="margin-left: 10px;margin-top: 20px" v-model="record" label="Запись" />
                <q-toggle style="margin-left: 10px;margin-top: 20px" v-model="vnc" label="Vnc" />
                <div class="bg-grey-2 q-pa-sm rounded-borders" style="margin-left: 10px">
                  <q-checkbox
                    name="accept_agreement"
                    v-model="chrome"
                    label="Chrome"
                  />

                  <q-checkbox
                    name="subscribe_newsletter"
                    v-model="ie"
                    label="Internet Explorer"
                  />
                  <q-checkbox
                    name="subscribe_newsletter"
                    v-model="mozilla"
                    label="Mozilla"
                  />
                </div>
                <q-space />

                <q-btn
                  color="grey"
                  round
                  flat
                  dense
                  :icon="expanded ? 'keyboard_arrow_up' : 'keyboard_arrow_down'"
                  @click="expanded = !expanded"
                />
              </q-card-actions>

              <q-slide-transition>
                <div v-show="expanded">
                  <q-separator />
                  <q-card-section class="text-subitle2">
                    <div v-if="dised[props.row.Id].Code">
                      <prism-editor v-model="dised[props.row.Id].Code" language="java" :line-numbers="lineNumbers" class="my-editor" style="max-height: 32em;"/>
                    </div>
                  </q-card-section>
                </div>
              </q-slide-transition>
            </q-card>
          <ReportTable style="margin-top: 10px" v-if="props.row.ReportArray" :repoList="props.row.ReportArray" :TestId="props.row.Id"></ReportTable>
          </q-td>
        </q-tr>
      </template>

    </q-table>
  </div>
</template>

<script>

  import regTicket from './regTicket'
  import Maximized from './Maximized'
  import ReportTable from "pages/workspace/ReportTable";
  import deleter from "pages/workspace/deleter";
  import {date} from 'quasar'
  import PrismEditor from 'vue-prism-editor'
  import moment from 'moment'

  export default {
    data() {
      return {
        mozilla: false,
        ie: false,
        chrome: true,
        expanded: false,
        load: 1,
        record : true,
        vnc: true,
        code: `By user = By.xpath("//input[@id='user']");
        By password = By.xpath("//input[@id='psw']");
        By buttonOk = By.xpath("//input[@id='okButton']");
        String name = "Sazonov.OI";
        String psw =  "Oracle33";
        driver.findElement(user).sendKeys(name);
        driver.findElement(password).sendKeys(psw);
        driver.findElement(buttonOk).click();`,
        lineNumbers: true,
        dised: {},
        submitting: false,
        tab: 'aboutLetter',
        splitterModel: 20,
        step: 2,
        folderId: null,
        selected: [],
        loader: null,
        loading: false,
        loading3: false,
        percentage3: 0,
        pagination: {
          page: 1,
          sortBy: 'Delivered',
          descending: true,
          rowsPerPage: 10
        },
        columns: [
          {
            name: 'Id',
            required: true,
            label: 'ID',
            align: 'left',
            field: row => row.Id,
            format: val => `${val}`,
            sortable: true
          },
          {name: 'Name', align: 'center', label: 'Название', field: 'Name'},
          {
            name: 'Created',
            align: 'center',
            label: 'Создан',
            field: row => row.CreateDate,
            format: val => date.formatDate(val, 'DD-MM-YYYY'),
            sortable: true
          },
          {name: 'Status', align: 'center', label: 'Статус', field: 'Status', sortable: true, style: 'width: 10px'}
        ],
        data2: [
          {
            IdLetter: 'Frozen Yogurt',
            LongName: 159,
            StatusProgress: 6.0,
            Warnings: 24,
            StpTicket: 4.0
          },
          {
            IdLetter: 'Ice cream sandwich',
            LongName: 237,
            StatusProgress: 9.0,
            Warnings: 37,
            StpTicket: 4.3
          }
        ],
        ratingModel: 3,
        icons: [
          'sentiment_very_dissatisfied',
          'sentiment_dissatisfied',
          'sentiment_satisfied',
          'sentiment_very_satisfied'
        ],
        data: [],
        currentProps: null,
        longerKurrentKey: null,
        longerFolderId: {},
        longerFolderName: {},
        longerDelivereddate: {},
        longerDescription: {},
        longerFrom: {},
        longerSendTo: {},
        longerSubject: {},
        longerNer_TextBody: {},
        longerattachments: {},
        longerNerFinal: {},
        longerSysRating: {},
        longerSysComment: {}
      }
    },
    components: {
      regTicket,
      Maximized,
      ReportTable,
      PrismEditor,
      deleter
    },
    computed: {
      disabledStore: {
        get: function (id) {
          console.log(`Ого`)
          let obj = this.$store.state.myStore.dised
          console.log(obj)
          return obj
        },
        set: function (value) {
          console.log(`=>|COM:Установить значение ${value} надо через addCom()`)
        }
      }
    },
    methods: {
      loadCheck() {
        this.load = 2
      },
      editPush(Id) {
        let curFName = 'UpdateTest'
        this.submitting = true
        let obj = this
        console.log(`=>${name}=>${curFName}|COM:Обновляем тест:${Id}`)
        console.log(this.dised[Id])
        let test = {
          Id: Id,
          Code: this.dised[Id].Code,
          Lang: this.dised[Id].Lang,
          Name: this.dised[Id].Name,
        }
        this.$axios.post('/qapi/updateTest', test)
          .then(response => {
            console.log(`=>${name}=>${curFName}|COM:Запрос удачен`)
            if (!response.data) {
              return
            }
            let notify = response.data
            this.$store.dispatch('myStore/ExecFunc', {'Name': 'ToNotify', 'Data': notify})
          })
          .catch(error => {
            console.log(`=>${name}=>${curFName}|COM:Запрос не удачен :(`)
            const Notify = {FuncName: "",List: [],Show: true,Status: 1,Text: `Ошибка сервера: ${error.message}`,UpdNum: 0}
            this.$store.dispatch('myStore/ExecFunc', {'Name': 'ToNotify', 'Data': Notify})
            //this.$store.commit('myStore/ToNotify', Notify)
            //this.showNotif('top', 'warning', `Ошибка сервера: ${error.message}`, 'remove_shopping_cart')
          })
          .finally(() => {
            setTimeout(function () {
              obj.submitting = false
            }, 500);
          })
      },
      addCom(props) {
        console.log(props.row)
        props.expand = !props.expand
        this.$store.commit('myStore/setDised', {id: props.row.Id, val: true})
        const clone = JSON.parse(JSON.stringify(this.dised))
        clone[props.row.Id] = JSON.parse(JSON.stringify(props.row))
        if (!clone[props.row.Id].Code) {
          clone[props.row.Id].Code = " " // Не можен Призм работать с null
        }
        this.dised = clone
      },
      disabledGet(id) {
        if (this.disabledAll.get(id)) {
          console.log("gec")
          return this.disabledAll.get(id)
        }
        console.log("not")
        return true
      },
      disabledSet(id, value) {
        console.log(this.disabledAll.get(id))
        this.disabledAll.set(id, value)
        console.log(this.disabledAll.get(id))
      },
      simulateSubmit(test) {
        this.submitting = true
        this.$store.dispatch('myStore/ExecFunc', {'Name': 'UpdateTest', 'Data': test})
        setTimeout(() => {
          // delay simulated, we are done,
          // now restoring submit to its initial state
          //this.submitting = false
          this.$router.push({path: `/`})
        }, 2000)
      },
      tstDate(dat) {
        return date.formatDate(dat, 'DD.MM.YYYY')
      },
      toStpData(row) {
        return {
          NerFinal: this.longerNerFinal[row.IdLetter],
          SysRating: this.longerSysRating[row.IdLetter],
          SysComment: this.longerSysComment[row.IdLetter],
          Vnc: row.Vnc,
          LetterID: row.IdLetter
        }
      },
      circleUpdate() {
        setTimeout(function () {
          document.getElementById('finder').click() // Нажатие кнопки обновдения данных
          this.circleUpdate()
        }, 10000)
      },
      replaceRegisterTicket: function (props) {
        this.$axios.post('/api/returnRegisterTicket', props.key)
          .then(response => {
            console.log('Отправлена на перерегистрацию')
            console.log(response)
            this.showNotif('top', 'teal', 'Заявка отправлена на перерегистрацию', 'thumb_up')
          })
          .catch(error => {
            console.log(error)
            this.showNotif('top', 'negative', 'Ошибка соединения', 'report_problem')
          })
        setTimeout(function () {
          document.getElementById('finder').click() // Нажатие кнопки обновдения данных
        }, 1000)
      },
      moment(dateForm) {
        return moment(dateForm).format()
      },
      RunWatcher() {
        setTimeout(function () {
          console.log('Обновление')
          document.getElementById('finder').click() // Нажатие кнопки обновдения данных
          if (this.longerKurrentKey !== undefined && this.longerKurrentKey !== null) {
            this.getLong(this.longerKurrentKey)
          }
        }, 1000)
      },
      onSubmit() {
        // Запрос к базе
        ServiceTable.execBaseData()
        this.updateData()
      },
      propActive(propExpand) {
        console.log('propExpand')
        console.log(propExpand)
        propExpand.expand = !propExpand.expand
        return propExpand
        // props.expand = !props.expand
      },
      getLong(key) {
        this.$axios.post('/api/givMeLongLetter', key)
          .then(response => {
            this.longerKurrentKey = key // Проставляем для его обновления
            console.log('Пришло подробнее о письме, из базы')
            console.log(response.data)
            // console.log('отправляем в qtable')
            // this.$emit('execbasedata', response.data)
            // this.folderId = response.data.FolderId
            // this.longerFolderId[key] = response.data.FolderId
            let a = {}
            a[key] = response.data.FolderName
            this.longerFolderName = a
            let b = {}
            b[key] = response.data.FromNode.Delivereddate
            this.longerDelivereddate = b
            let c = {}
            c[key] = response.data.FromNode.Description
            this.longerDescription = c
            let e = {}
            e[key] = response.data.FromNode.From
            this.longerFrom = e
            let f = {}
            f[key] = response.data.FromNode.SendTo
            this.longerSendTo = f
            let g = {}
            g[key] = response.data.FromNode.Subject
            this.longerSubject = g

            // console.log('Replacer')
            // console.log(item)
            // console.log(response.data.FromNode.textBody)
            // console.log(item.replace('\\n', '\n'))

            let h = {}
            h[key] = response.data.FromNode.attachments
            this.longerattachments = h
            let k = {}
            k[key] = response.data.FromNode.Ner_Final
            this.longerNerFinal = k
            if (response.data.FromNode.Ner_TextBody) {
              let count = response.data.FromNode.Ner_TextBody.length
              for (let i = 0; i < count; i++) {
                let item = response.data.FromNode.Ner_TextBody[i]
                response.data.FromNode.Ner_TextBody[i].TextRaw = item.TextRaw.replace('\\n', '\n')
              }
              let l = {}
              l[key] = response.data.FromNode.Ner_TextBody
              this.longerNer_TextBody = l
            }
            this.longerSysRating[key] = 3
          })
          .catch(error => {
            console.log(error)
            this.showNotif('top', 'negative', 'Непредвиденная ошибка', 'report_problem')
          })
      },
      getFolderIdInLongDesk(props) {
        console.log('Запрос на FolderId')
        // cols: Array(4)
        //let Example = {cols: '(...)',colsMap: (...),expand: (...),key: (...),pageIndex: (...),row: (...),rowIndex: (...),selected: (...)
        console.log(this.currentProps)
        console.log(props)
        //this.getLong(props.key)
        if (this.currentProps) {
          if (this.currentProps.row.Id !== props.row.Id) {
            this.currentProps.expand = false
            console.log(props.expand)
          }
        }
        this.currentProps = props
        props.expand = !props.expand

        return props.expand
      },
      updateData2() {
        console.log('Запрос на обновление обрабатывается')
        //return
        // this.$emit('execBaseData', dt)
        if (this.currentProps) {
          this.getLong(this.currentProps.key)
        }
        this.$axios.post('/api/givMeData')
          .then(response => {
            console.log('пришел ответ по короткому, из базы')
            // console.log('отправляем в qtable')
            if (response.data == null) {
              this.showNotif('top', 'warning', 'В целевой папке нет ни одного письма', 'remove_shopping_cart')
              return
            }
            console.log('отправляем в qtable')
            // this.$emit('execbasedata', response.data)

            let count = response.data.length
            for (let i = 0; i < count; i++) {
              response.data[i].Progress = 0 // Проставляем 0, потом поменяем
              response.data[i].Status = 0
              if (response.data[i].Messages) {
                let iMesCoin = response.data[i].Messages.length
                // Проставляем последние актуальные статусы
                response.data[i].Progress = response.data[i].Messages[iMesCoin - 1].Progress
                response.data[i].Status = response.data[i].Messages[iMesCoin - 1].Status
                response.data[i].StepStatus = {}
                for (let iMes = 0; iMes < iMesCoin; iMes++) {
                  response.data[i].StepStatus[response.data[i].Messages[iMes].Progress] = response.data[i].Messages[iMes].Status
                }
              }
              if (response.data[i].Delivered) {
                let rawDate = response.data[i].Delivered
                response.data[i].Delivered = moment(rawDate, 'DD.MM.YYYYTHH:mm:ss').toDate()
                let rawText = response.data[i].LongName
                if (rawText.length > 90) {
                  response.data[i].LongName = rawText.substr(0, 90)
                }
              }
            }
            let array = []
            array = response.data
            array.sort(function (a, b) {
              var c = new Date(a.Delivered)
              var d = new Date(b.Delivered)
              return c - d
            })
            console.log('В таблицу ')
            console.log(array)
            this.data = array.reverse()
          })
          .catch(error => {
            console.log('пришла ошибка ;(')
            console.log(error)
            this.showNotif('top', 'negative', 'Непредвиденная ошибка', 'report_problem')
          })
      },
      setNotify() {
        this.$axios.post('/api/setNotify', "")
      },
      getList() {
        this.$axios.post('/qapi/getAllTests', "")
          .then(response => {
            console.log('Отправил серверу запрос на выдачу первой партии данных')
            console.log(response.data)
          })
          .catch(error => {
            console.log(error)
          })
      },
      addTest(row) {
        let ticket = {
          LetterID: row.Id,
          Server: "",
          Step: "StartTest",
          Vnc: this.vnc,
          Rec: this.record
        }
        this.$axios.post('/qapi/startTests', ticket)
          .then(response => {
            console.log(`Запустить тест: ${row.Id}`)
            console.log(response)
            this.$store.dispatch('myStore/ExecFunc', {'Name': 'ToNotify', 'Data': response.data})
          })
          .catch(error => {
            console.log(error)
          })
      },
      PoolInit() {
        console.clear()
        this.$store.dispatch('myStore/ExecFunc', {'Name': 'PoolStart'})
        // this.$store.dispatch('myStore/PoolInit')
      },
      startComputing(id) {
        this.onSubmit()
        this[`loading${id}`] = true
        this[`percentage${id}`] = 0
        this[`interval${id}`] = setInterval(() => {
          this[`percentage${id}`] += Math.floor(Math.random() * 8 + 10)
          if (this[`percentage${id}`] >= 100) {
            clearInterval(this[`interval${id}`])
            this[`loading${id}`] = false
          }
        }, 700)
        // Асинхронная функция выше будет прорисовывать прогресс независимо от реального представления дел ниже
      },
      showNotif(position, color, message, icon) {
        console.log(position + color + message + icon)
        const buttonColor = color ? 'white' : void 0
        this.$q.notify({
          color,
          icon: icon,
          message,
          position,
          actions: [{
            color: buttonColor, handler: () => { /* console.log('wooow') */
            }
          }],
          timeout: Math.random() * 5000 + 3000
        })
      },
      progressNum: function (LetterID, num) {
        console.log('На изменение прогресса')
        let dat = {}
        dat.Id = LetterID
        dat.Progress = num
        console.log(dat)
        this.$axios.post('/api/progressNum', dat)
          .then(response => {
            console.log('Отправлена на изменение прогресса')
            console.log(response)
            this.showNotif('top', 'teal', 'Заявка отправлеа на изменение прогресса', 'thumb_up')
          })
          .catch(error => {
            console.log(error)
            this.showNotif('top', 'negative', 'Непредвиденная ошибка', 'report_problem')
          })

      },
    },
    beforeDestroy() {
      clearInterval(this.interval1)
      clearInterval(this.interval2)
    },
    created() {
      this.RunWatcher()
    }
  }
</script>

<style scoped>
  .my-table-details {
    font-size: 0.85em;
    font-style: italic;
    max-width: 200px;
    white-space: normal;
    color: #555;
    margin-top: 4px;
  }
</style>
