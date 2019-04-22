
<template>
<div>
  <div id="header">
    <!-- <hr id="headerline"> -->
    <!-- <div id="logo"><img src="../assets/logo.png"></div> -->
    <div id="headerline"></div>
    <div id="logo"><img src="./assets/logo.png"></div>
    <span id="title">Operation Event Monitor System</span>
  </div>
  <div>
    <div id="calendarzone">
      <eventCalendar/>
    </div>
    <div id="infozone">
      <div id="summary"><eventSummary/></div>
      <div id="detail"><eventDetail/></div>
    </div>
  </div>
</div>

</template>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 600px;
}
</style>

<script>
import '../static/main.css'
import eventCalendar from './components/calendar'
import eventDetail from './components/event_detail'
import eventSummary from './components/event_summary'
export default {
  name: 'app',
  components: {
    eventDetail,
    eventCalendar,
    eventSummary
  }
}
</script>
