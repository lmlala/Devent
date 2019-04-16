/**
 * @author [Hilton Li]
 * @email [liminghilton@gmail.com]
 * @create date 2019-04-15 19:01:20
 * @modify date 2019-04-15 19:01:20
 * @desc [description]
 */

 'use strict';

module.exports = {
    eventIcon: {
        "incident": {
            "critical": { "color": "#FF0000" },
            "normal": { "color": "#CD3333" },
            "default": { "color": "#C67171" }
        },
        "deployment": {
            "Hotfix": { "color": "#32CD32" },
            "default": { "color": "#32CD32" }
        },
        "damo": {
            "start": { "color": "#8A8A8A"},
            "code_freeze": {"color": "#CD00CD"},  
            "depoly_to_sandbox": {"color": "#7FFFD4"},  
            "depoly_to_production": {"color": "#66CD00"},
            "end": {"color": "#383838"},
            "default": { "color": "#8A8A8A"},
        },
        "monitor": {
            "datadog": { "color": "#EEC900"},
            "sentry": { "color": "#EEC900"},
            "splunk": { "color": "#EEC900"},
            "default": { "color": "#EEC900"}
        },
        "component": {
            "create": { "color": "#436EEE"},
            "remove": { "color": "#C6E2FF"},
            "default": {"color": "#C6E2FF"}
        }
    }
}