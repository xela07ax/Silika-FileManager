<template>
  <div>
    <main class="q-page_ doc-page_" style="min-height: 751px;">
<!--    <bread-crop></bread-crop>-->
    <h3 class="title-kz"></h3>
<!--    <filters></filters>-->
    <div class="centered_">
      <div  class="fi-header_">
        <!--      <loading></loading>-->
        <my-table></my-table>
<!--        <div v-if="visibleif">-->
<!--          <p class="data-out">Результаты отсутствуют</p>-->
<!--        </div>-->
<!--        <div v-else>-->
<!--&lt;!&ndash;          <chart-bar  v-if="displayed"></chart-bar>&ndash;&gt;-->
<!--        </div>-->
      </div>
    </div>
    </main>
  </div>
</template>

<script>
// import Filters from './Filters'
// import ChartBar from './body/ChartBar'
// import BreadCrop from './BreadCrop'
import myTable from './qTable'
// import Loading from './components/loading'

export default {
  name: 'App',
  data () {
    return {
      doSomethingOnHide: false,
      items: [
        { title: 'Click Me' },
        { title: 'Click Me' },
        { title: 'Click Me' },
        { title: 'Click Me 2' }
      ],
      lorem: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.'
    }
  },
  components: {
    // Filters,
    // BreadCrop,
    myTable
    // Loading
  },
  computed: {
    visibleif () {
      let checkt = true
      // if (this.$store.getters.getMesIds.length > 0) {
      //   checkt = false
      // }
      return checkt
    },
    displayed () {
      return {}
      // return ((this.$store.getters.getChartVal.length > 0) && (this.$store.getters.getChartMonths.length > 0))
    }
  },
  created () {
  }
}
</script>

<style scoped>
  .my-card {
    width: 100%;
    max-width: 250px;
  }

</style>

<style>
  .doc-page0 {
    padding: 16px 46px;
    font-weight: 300;
    /*max-width: 900px;*/
    margin-left: auto;
    margin-right: auto;
    min-height: 751px;
  }
  .doc-h1 {
    font-size: 2rem;
    line-height: 2rem;
    padding: 1rem 0;
    font-weight: 500;
    /*margin: 0 0 2rem;*/
  }
  .doc-heading {
    color: #027be3;
    cursor: pointer;
  }
  .doc-h2 {
    font-size: 1.5rem;
    line-height: 1.5rem;
    padding: 0.5rem 0;
    font-weight: 500;
    border-bottom: 1px solid #ccc;
    margin: 1rem 0 1.5rem;
  }
</style>
