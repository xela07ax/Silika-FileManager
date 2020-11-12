<template>
  <q-table
    :data="repoList"
    :columns="columns"
    row-key="Start"
  >
    <template v-slot:header="qprops">
      <q-tr :qprops="qprops">
        <q-th auto-width></q-th>
        <q-th
          v-for="col in qprops.cols"
          :key="col.name"
          :qprops="qprops"
        >
          {{ col.label }}
        </q-th>
      </q-tr>
    </template>

    <template v-slot:body="qprops">
      <q-tr :qprops="qprops">
        <q-td auto-width>
          <q-btn dense round flat :icon="qprops.expand ? 'arrow_drop_up' : 'arrow_drop_down'" @click="openExpand(qprops)" />
          <deleter style="height: 10px" :row="qprops.row" :TestId="TestId"></deleter>
        </q-td>
        <q-td
          v-for="col in qprops.cols"
          :key="col.name"
          :qprops="qprops"
        >
          <div v-if="col.name === 'Status'">
<!--            {{dateConvert(qprops.row.End)}}-->
            <!--0=ok,1=error,4=runing-->
            <div  v-if="qprops.row.Status === 0">
              {{upd(qprops)}}
            </div>
            <div  v-if="qprops.row.Status === 1">
              {{upd(qprops)}}
            </div>
            <q-linear-progress v-if="dateConvert(qprops.row.End) === '1754'" dark query color="cyan" class="q-mt-sm" />
            <q-icon v-else-if="qprops.row.Status === 0" name="done_outline" class="text-positive" style="font-size: 1.4rem;" />
            <q-icon v-else name="warning" class="text-red" style="font-size: 1.4rem;" />
          </div>
          <div v-else>{{ col.value}}</div>
        </q-td>
      </q-tr>
      <q-tr v-if="qprops.expand" :qprops="qprops">
        <q-td colspan="100%">
          <q-card class="my-card">
            <q-card-section>
              <q-splitter
                v-model="splitterModel"
                style="height: 300px"
              >
                <template v-slot:before>
                  <q-tabs
                    v-model="tab"
                    vertical
                    class="text-teal"
                  >
                    <q-tab name="about" icon="mail" label="О тесте" />
                    <q-tab v-if="qprops.row.SelenoidId && qprops.row.Record" name="vidos" icon="movie" label="Видео" />
                    <q-tab v-if="dateConvert(qprops.row.End) === '1754' && qprops.row.SelenoidId" name="textLetter" icon="video_settings" label="Vnc" />
                    <q-tab name="logCmd" icon="subtitles" label="Лог CMD" />
                    <q-tab name="errCmd" icon="subtitles" label="Ошибки CMD" />
                    <q-tab name="step" icon="subtitles" label="Шаги" />
                  </q-tabs>
                </template>

                <template v-slot:after>
                  <q-tab-panels
                    v-model="tab"
                    animated
                    transition-prev="jump-up"
                    transition-next="jump-up"
                  >
                    <q-tab-panel name="about">
                      <div class="text-h6 q-mb-md">Статус</div>
                      <q-toggle style="margin-left: 10px;margin-top: 20px" v-model="qprops.row.Record" label="Запись" />
                      <q-toggle style="margin-left: 10px;margin-top: 20px" v-model="qprops.row.Vnc" label="Vnc" />
<!--                          {{qprops.row}}-->
                      <p>Текст сообщения: {{qprops.row.StatusText}}</p>
                      <p>Sewlenoid Id: {{qprops.row.SelenoidId}}</p>
<!--                      {{dateConvert(qprops.row.End)}}-->
<!--                      {{qprops.row.Vnc}}-->
<!--                      {{"dateConvert(qprops.row.End) === '1754' && qprops.row.Vnc"}}-->
                    </q-tab-panel>
                    <q-tab-panel name="vidos">
                      <div style="width: 100%; max-width: 100%;white-space:normal">
                        <q-btn style="margin-left: 10px;margin-bottom: 20px" @click="openVideoPage(qprops.row.SelenoidId)" class="glossy" round color="deep-orange" icon="local_activity" />
<!--                        <vue-friendly-iframe :src="videoPage(qprops.row.SelenoidId)"></vue-friendly-iframe>-->
                        <vue-friendly-iframe :src="vP(qprops.row.SelenoidId)"></vue-friendly-iframe>

                      </div>
                    </q-tab-panel>
                    <q-tab-panel name="textLetter">
                      <div style="width: 100%; max-width: 1000px">
                        <vnc-screen :Vnc="qprops.row.SelenoidId" :viewOnly="true"/>
                      </div>
                    </q-tab-panel>
                    <q-tab-panel name="logCmd">
<!--                      <div class="text-h5 q-mb-md">Вложения</div>-->
                      <prism-editor v-model="list[qprops.row.Start].Out" language="java" :line-numbers="false" class="my-editor" style="max-height: 32em;"/>
                    </q-tab-panel>
                    <q-tab-panel name="errCmd">
<!--                      <div class="text-h5 q-mb-md">Вложения</div>-->
                      <prism-editor v-model="list[qprops.row.Start].Err" language="java" :line-numbers="false" class="my-editor" style="max-height: 32em;"/>
                    </q-tab-panel>
                    <q-tab-panel name="step">
                      <!--                      <div class="text-h5 q-mb-md">Вложения</div>-->
                     {{qprops.row.Steps}}
                    </q-tab-panel>
                  </q-tab-panels>

                </template>

              </q-splitter>
            </q-card-section>
          </q-card>
        </q-td>
      </q-tr>
    </template>

  </q-table>
</template>

<script>
  import Maximized from './Maximized'
  import vncScreen from './vncScreen'
  import deleterRepo from "pages/workspace/deleterRepo";
  import moment from 'moment'
  import {date} from 'quasar'
  import PrismEditor from 'vue-prism-editor'
  import regTicket from "pages/workspace/regTicket";
  import deleter from "pages/workspace/deleter";

    export default {
      props: ['repoList', 'TestId'],
      data() {
        return {
          list: {},
          tab: 'about',
          reports: [
            {
              Start: moment("2020-08-18T15:04:40.4940889+03:00", 'DD.MM.YYYYTHH:mm:ss').toDate(),
              End: moment("2020-08-18T15:04:40.6708111+03:00", 'DD.MM.YYYYTHH:mm:ss').toDate(),
              Err: "",
              Out: "",
              SelenoidId: "",
              Status: 1,
              StatusText: "Ошибка в шаблоне",
              Steps: null
            },
            {
              Start: moment("2020-08-18T15:04:40.4940889+03:00", 'DD.MM.YYYYTHH:mm:ss').toDate(),
              End: moment("2020-08-18T15:04:40.6708111+03:00", 'DD.MM.YYYYTHH:mm:ss').toDate(),
              Err: "kjhjk",
              Out: "ooooooo",
              SelenoidId: "87556ggr4e54",
              Status: 0,
              StatusText: "Ошибка в шаблоне",
              Steps: null
            }
          ],
          columns: [
            {
              name: 'StatusText',
              required: true,
              label: 'Сообщение',
              align: 'left',
              field: row => row.StatusText,
              format: val => `${val}`,
              sortable: true
            },
            {
              name: 'Start',
              align: 'center',
              label: 'Старт',
              field: 'Start',
              format: val => date.formatDate(val, 'DD-MM-YYYY HH:MM:SS'),
              sortable: true
            },
            {
              name: 'End', align: 'center', label: 'Окончание', field: 'End',
              format: val => date.formatDate(val, 'DD-MM-YYYY HH:MM:SS'),
              sortable: true
            },
            {name: 'Status', align: 'center', label: 'Статус', field: 'Status', sortable: true},
          ]
        }
      },
      name: "ReportTable",
      components: {
        PrismEditor,
        'deleter': deleterRepo,
        vncScreen
      },
      methods: {
        vP(id){
          return `http://${this.$store.state.myStore.videoServer}/id?${id}`
        },
        openVideoPage: function (id) {
          window.open(this.vP(id), "_blank");

        },
        upd(qprops) {
          console.log("upd")
          this.list[qprops.row.Start] = qprops.row
        },
        openExpand(qprops) {
          qprops.expand = !qprops.expand
          this.list[qprops.row.Start] = qprops.row
        },
        dateConvert(val) {
          console.log('dateConvert')
          console.log(val)
          return date.formatDate(val, 'YYYY')
        }
      }
    }
</script>

<style scoped>

</style>
