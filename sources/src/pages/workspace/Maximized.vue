<template>
  <div class="q-pa-md q-gutter-sm">
    <vnc-screen style="max-height: 290px" class="vnc-card" v-if="Vnc.Online && !dialog" :Vnc="Vnc" :viewOnly="true"/>
    <q-btn label="Развернуть" color="primary" @click="dialog = true" />

    <q-dialog
      v-model="dialog"
      persistent
      :maximized="maximizedToggle"
      transition-show="slide-up"
      transition-hide="slide-down"
    >
      <q-card class="bg-primary text-white">
        <q-bar>
          <q-space />

          <q-btn dense flat icon="minimize" @click="maximizedToggle = false" :disable="!maximizedToggle">
            <q-tooltip v-if="maximizedToggle" content-class="bg-white text-primary">Minimize</q-tooltip>
          </q-btn>
          <q-btn dense flat icon="crop_square" @click="maximizedToggle = true" :disable="maximizedToggle">
            <q-tooltip v-if="!maximizedToggle" content-class="bg-white text-primary">Maximize</q-tooltip>
          </q-btn>
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip content-class="bg-white text-primary">Close</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section class="q-pt-none">
          <vnc-screen v-if="Vnc.Online" :Vnc="Vnc" :viewOnly="false"/>
        </q-card-section>
      </q-card>
    </q-dialog>
  </div>
</template>

<script>
import vncScreen from './vncScreen'

export default {

  props: ['Vnc'],
  data () {
    return {
      dialog: false,
      maximizedToggle: true
    }
  },
  components: {
    vncScreen
  }
}
</script>
