<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="leftDrawerOpen = !leftDrawerOpen"
        />

        <q-avatar size="42px">
          <img src="~assets/otr-logo-circle2.png" />
        </q-avatar>
        <q-toolbar-title>
          ОТР "Система автоматизированного тестирования"
        </q-toolbar-title>

        <div>Quasar v{{ $q.version }}</div>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
      content-class="bg-grey-1"
    >
      <q-list>
        <q-item-label
          header
          class="text-grey-8"
        >
          Меню
        </q-item-label>
        <div class="q-gutter-md" style="margin-left: 10px">
        <q-btn align="between" style="margin-left: 10px" class="btn-fixed-width" color="teal" label="Список" icon="map"  @click="push('')"/>
        <br>
        <q-btn align="between" style="margin-left: 10px" class="btn-fixed-width" color="accent" label="Создать" icon="flight_takeoff"  @click="push('create')"/>
        </div>
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import EssentialLink from 'components/EssentialLink.vue'

const linksData = [
  {
    title: 'Crearte',
    caption: 'github.com/quasarframework',
    icon: 'code',
    link: '/#/create'
  }, {
    title: 'Docs',
    caption: 'quasar.dev',
    icon: 'school',
    link: 'https://quasar.dev'
  },
  {
    title: 'Discord Chat Channel',
    caption: 'chat.quasar.dev',
    icon: 'chat',
    link: 'https://chat.quasar.dev'
  },
  {
    title: 'Forum',
    caption: 'forum.quasar.dev',
    icon: 'record_voice_over',
    link: 'https://forum.quasar.dev'
  },
  {
    title: 'Twitter',
    caption: '@quasarframework',
    icon: 'rss_feed',
    link: 'https://twitter.quasar.dev'
  },
  {
    title: 'Facebook',
    caption: '@QuasarFramework',
    icon: 'public',
    link: 'https://facebook.quasar.dev'
  },
  {
    title: 'Quasar Awesome',
    caption: 'Community Quasar projects',
    icon: 'favorite',
    link: 'https://awesome.quasar.dev'
  }
];

export default {
  name: 'MainLayout',
  components: { EssentialLink },
  data () {
    return {
      leftDrawerOpen: false,
      essentialLinks: linksData
    }
  },
  created () {
    //console.clear()
    // Начинаем пулл
    this.$store.dispatch('myStore/PoolInit', this)
    // Все пользователи при открытии окна, активируют пулл
    console.log(`Регистрируем пулл`)
    this.$store.dispatch('myStore/ExecFunc', { 'Name': 'PoolStart' })
  },
  methods: {
    push(route){
      this.$router.push({ path: `/${route}` })
    }
  }
}
</script>
