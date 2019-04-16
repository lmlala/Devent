<template>
  <full-calendar :config="config" :events="events"/>
</template>

<script>
// // import moment from 'moment'
import axios from 'axios'
import eventIcon from '../../config/global'


function getEventStyle(e) {
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
      color = eventIcon.eventIcon.deployment[e.event_type];
      break

    default:
      color = '#A8A8A8'
  }

  // console.log('outside color is: ', color)
  return color.color
}


export default {
  name: 'hello',
  data () {
    return {
      events: [],
      config: {
        defaultView: 'month',
        height: 800,
        contentHeight: 750,
        eventRender: function (event, element) {
          // console.log(event)
        }
      }
    }
  },
  mounted () {
    axios.get('http://127.0.0.1:8000/v1/list').then(response => {
      response.data.forEach(function (e) {
        // console.log(e.is_down)
        // console.log(e.is_service_down)
        // console.log(e)
        // console.log(eventIcon.eventIcon.incident.default.color)
        // console.log("each e is: ", e)
        e.color = getEventStyle(e)
      })
      this.events = response.data
    })
  }
}
</script>
