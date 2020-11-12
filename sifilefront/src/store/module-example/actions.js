export function someAction (/* context */) {
}
import moment from 'moment'
const Notify = {FuncName: "",List: [],Show: false,Status: 0,Text: "",UpdNum: 0}
const Test = {Code: null,CreateDate: "2020-08-02T21:54:05.6641194+03:00",EndDate: "", Id: 6800001,Lang: "Java",Name: "Test1",Status: 0,UpdateDate: ""}
let name = 'actions.js'
let pool = {
  name: 'nilName',
  UID: 0,
  UpdateNumber: 0, // При первом запросе должны сразу получить данные
  getters: {},
  commit: {},
  _vm: {},
  store: {},
  loading: {}, // Статус обновления системы
  winMode: {},
  watchData: {},
  ChatDID: 0,
  ChatMessages: [],
  ChatDisable: true,
  HistoryActive: {},
  OperListActive: {},
  OperListWaiters: {},
  OperListTickets: {},
  pollObjCreated: false, // (true) Объект пулла инициализирован, второй раз это делать нечего
  poolActive: false, // (true) Пулл активирован и опрашивает сервер
  poolSuccess: true, // (false) Пулл должен прекратить работу
  printErr: function (error) {
    if (error.response) {
      if (error.response.status === 400) {
        let myRegexp = /COM:(.*?\|)/g
        let match = myRegexp.exec(error.response.data)
        this.$q.notify({ message: match[1], color: 'warning' })
      }
    } else if (error.request) {
      this.$q.notify({ message: error.request, color: 'warning' })
    } else {
      // Something happened in setting up the request that triggered an Error
      this.$q.notify({ message: error.message, color: 'negative' })
    }
  },
  renderNotify(notify){
    //console.log("Status")
    //console.log(notify.Show)
    if (notify.Show) {
      //console.log(notify.Status)
      switch (notify.Status) {
        case 1:
          this.showNotif('top', 'negative', notify.Text, 'report_problem')
          break
        case 2:
          this.showNotif('top', 'warning', notify.Text, 'remove_shopping_cart')
          break
        default:
          this.showNotif('top', 'teal', notify.Text, 'thumb_up')
      }
    }
  },
  showNotif (position, color, message, icon) {
    //console.log(position + color + message + icon)
    const buttonColor = color ? 'white' : void 0
    this._vm.$q.notify({
      color,
      icon: icon,
      message,
      position,
      actions: [{ color: buttonColor, handler: () => { /* console.log('wooow') */ } }],
      timeout: Math.random() * 5000 + 3000
    })
  },

  GetTestExample: function(){
    return JSON.parse(JSON.stringify(Test));
  },
  GetNotifyExample: function(){
    return JSON.parse(JSON.stringify(Notify));
  },
  LongPool: function (initiator) {
    let curFName = 'LongPool'
    console.log(`Pooling: ${this.UpdateNumber}`)
    this._vm.$axios.post('/api/pooling', this.UpdateNumber)
      .then(response => {
        // console.log(`=>${name}=>${curFName}|COM:Запрос удачен`)
        // console.log(response.data)
        if (!response.data) {
          //console.log("Пусто")
          return
        }
        let notify = response.data
        console.log(`=>${name}=>${curFName}|COM:Пришло обновление| ${notify.Text}`)
        pool.renderNotify(notify)
        console.log(notify)
        pool.UpdateNumber = notify.UpdNum
        // Отсортируем список по id
        notify.List.sort(function (a, b) {
            return a.Id - b.Id
        })
        // Установка списка
        // console.log('В таблицу')
        this.commit('setList', notify.List.reverse())
      })
      .catch(error => {
        console.log(`=>${name}=>${curFName}|COM:Запрос не удачен :(`)
        let notify = pool.GetNotifyExample()
        notify.Text = `${error.message}`
        notify.Show = true
        notify.Status = 1
        this.renderNotify(notify)
      })
      .finally(() => {
        setTimeout(function () {
          pool.LongPool()
        }, 500);
      })
  },
  CreateTest: function (testy) {
    let curFName = 'CreateTest'
    console.log(`=>${name}=>${curFName}|COM:Создаем новый тест`)
    console.log(testy)
    this._vm.$axios.post('/qapi/createTest', testy)
      .then(response => {
        console.log(`=>${name}=>${curFName}|COM:Запрос удачен`)
        //console.log(response.data)
        //console.log("-----")
        //{"FuncName":"CreateTest","Text":"Имя теста не может быть пустым","Status":1,"List":null,"Show":true,"UpdNum":0}{"FuncName":"CreateTest","Text":"Тест создан","Status":0,"List":null,"Show":true,"UpdNum":0}
        if (!response.data) {
          //console.log("Пусто")
          return
        }
        let notify = response.data
        //console.log("notify")
        // console.log(response.data.Show)
        console.log(notify)
        pool.renderNotify(notify)
      })
      .catch(error => {
        console.log(`=>${name}=>${curFName}|COM:Запрос не удачен :(`)
        this.showNotif('top', 'warning', `Ошибка сервера: ${error.message}`, 'remove_shopping_cart')
      })

  }
}
export function PoolInit ({ commit, getters }, payload) {
  // Инициализируем объект Пула
  console.log("Инициализируем объект Пула")
  let fName = 'PoolInit'
  if (pool.pollObjCreated) {
    console.log(`=>${name}=>${fName}|COM:Объект управления пуллом уже инициализирован`)
    return
  }
  this.commit('setLoading', true)
  // Создадим объект управления пуллом
  console.log(`=>${name}=>${fName}|COM:Создадим объект управления пуллом`)
  Object.defineProperties(pool, {
    name: {
      value: fName
    },
    _vm: {
      //value: payload//this._vm
      value: this._vm
    },
    store: {
      value: this
    },
    UID: {
      get: function () {
        return this.getters.getUID
      },
      set: function (value) {
        console.log(`=>${this.name}=>${name}|COM:Установить значение UID нельзя ${value}`)
      }
    },
    getters: {
      value: getters
    },
    commit: {
      value: commit
    },
    loading: {
      get: function () {
        return this.getters.getLoading
      },
      set: function (value) {
        //console.log(`Система обновилась ${value}`)
        this.commit('setLoading', value)
      }
    }
  })
  // Отправим серверу запрос на выдачу первой партии данных
  // console.log('insertMeData')
  // this._vm.$axios.post('/api/insertMeData')
  //   .then(response => {
  //     console.log('Отправил серверу запрос на выдачу первой партии данных')
  //   })
  //   .catch(error => {
  //     console.log(error)
  //     this.showNotif('top', 'negative', 'Ошибка сервера', 'report_problem')
  //   })
  console.log(`Инициализация пула закончена ${!pool.loading}`)
  pool.loading = false
}

export function ExecFunc ({ commit, getters }, payload) { // payload = {"Name":"PoolStart","Data":{}}
  let fName = 'ExecFunc'
  if (payload.Name) {
    // Проверяем состояние загрузки системы, не запускаем функцию пока она не закончится
    let intervalID = setInterval(function () {
      if (pool.loading) {
          console.log(`=>${name}=>${fName}=>setInterval|COM:Система обновляется|DT:loading:${pool.loading}`)
      } else {
        clearInterval(intervalID)
        console.log(`=>${name}=>${fName}=>setInterval|COM:Выполняем функцию|DT:funcName:${payload.Name}`)
        switch (payload.Name) {
          case 'PoolStart':
            // Это первый запрос, в нем получим первичные данные для отрисовки


            if (!pool.poolActive) {
              // Проставим на всякий случий разрешающий флаг работы пулла
              if (!pool.poolSuccess) {
                pool.poolSuccess = true
              }
              // Проставим флаг активной работы пулла, что-бы не запустить второй раз
              // Запускаем пулл
              console.log('Запускаем пулл')
              pool.LongPool(`actons.js=>${fName}`)
            } else {
              console.log(`=>${name}=>${fName}=>setInterval|COM:Пулл уже работает|DT:funcName:${payload.Name}`)
            }
            break
          case 'PoolEnd': // if (x === 'value1')
            // Проставим флаг запрета работы пулла
            if (pool.poolSuccess) {
              pool.poolSuccess = false
            }
            break
          case 'CloseUser':
            console.log(`=>${name}=>${fName}=>CloseUser|COM:Команда принята|DT:funcName:${payload.Name}`)
            pool.CloseUser()
            if (pool.poolSuccess) {
              pool.poolSuccess = false
            }
            break
          case 'CreateTest': // if (x === 'value1')
            pool.CreateTest(payload.Data)
            break
          case 'ToNotify':
            console.log(`=>${name}=>${fName}=>setInterval|COM:Отображаем сообщение|DT:funcName:${payload.Data}`)
            pool.renderNotify(payload.Data)
            break
          case 'SetDialog': // if (x === 'value1')
            // Принимаем на вход UID, ожидам на выход его свалку
            pool.SetDialog(payload.Data)
            break
          case 'SetClientInActive': // if (x === 'value1')
            // Принимаем на вход UID, ожидам на выход его свалку
            pool.SetClientInActive(payload.Data)
            break
          default:
            console.log(`=>${name}=>${fName}|COM:Ошибка, имя вызываемой функции не найдено в ${fName}|DT:payload:${JSON.stringify(payload)}`)
        }
      }
    }, 500)
  } else {
    console.log(`=>${name}=>${fName}|COM:Ошибка, имя вызываемой функции не задано|DT:payload:${JSON.stringify(payload)}`)
  }
}

