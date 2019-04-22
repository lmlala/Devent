<style>
@import '@fullcalendar/core/main.css';
@import '@fullcalendar/daygrid/main.css';
</style>

<script>
// // import moment from 'moment'
import FullCalendar from '@fullcalendar/vue'
import dayGridPlugin from '@fullcalendar/daygrid'
// import timeGridPlugin from  '@fullcalendar/timeline'
import axios from 'axios'
import eventIcon from '../../config/global'
import Vue from 'vue'

function getEventStyle (e) {
  // console.log(eventIcon)
  // console.log("e is: ", e)
  var color
  switch (e.event_type) {
    case 'Incident':
      switch (e.priority) {
        case 'P5':
          color = eventIcon.eventIcon.incident.critical
          break
        case 'P4':
          color = eventIcon.eventIcon.incident.normal
          break
        default:
          color = eventIcon.eventIcon.incident.default
          break
      }
      break
    case 'Hotfix':
      color = eventIcon.eventIcon.deployment[e.event_type]
      break
    default:
      color = '#A8A8A8'
  }

  // console.log('outside color is: ', color)
  return color.color
}

export default {
  name: 'eventCalendar',
  components: {
    FullCalendar,

  },
  template: `
  <FullCalendar
    ref="eventCalendar"
    defaultView="dayGridMonth"
    nowIndicator="true"
    timezone='local'
    :header="{
        right: 'prev,next today',
        center: 'title',
        left: 'dayGridMonth,dayGridWeek,dayGridDay,listWeek'
      }"
    :plugins="calendarPlugins"
    :weekends="true"
    :events="events"
    :height=800
    :contentHeight=750
    :eventLimit=10
    :editable="true"
    slotDuration='01:00:00'
    @eventRender="eventRender"
    @eventMouseEnter="mouseOver"
    @eventClick="eventClick"
  />`,
  data () {
    return {
      calendarPlugins: [ dayGridPlugin ],
      // events: [{id: 1, title: 'e1', start: '2019-04-17T16:17:00Z', end: '2019-04-17T17:19:00Z'},
      //           {id: 2, title: 'e2', start: '2019-04-18T16:18:00Z', end: '2019-04-20T00:19:00Z'},
      //           {id: 3, title: 'e3', start: '2019-04-19 20:19:00', end: '2019-04-20 03:19:00', color: 'red', allDay: false}
      // ]
      events: []
    }
  },
  mounted () {
    axios.get('http://127.0.0.1:8000/v1/list').then(response => {
      console.log(response.data)
      response.data.forEach(function (e) {
        // console.log(e.is_down)
        // console.log(e.is_service_down)
        // console.log(e)
        // console.log(eventIcon.eventIcon.incident.default.color)
        console.log("each e is: ", e)
        e.color = getEventStyle(e)
      })
      this.events = response.data
    })
  },
  methods: {
    eventRender (info, el) {
    // Vue.set(info.el, 'testattr', 'sdf')
    console.log(info)
    
    },
    mouseOver (info) {
      // console.log(info.el)
      console.log('info.el mouse over...')
    },
    eventClick (clickInfo) {
      console.log('event click...')
      console.log(clickInfo.el)
    }
  }
}
</script>
