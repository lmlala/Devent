/**
 * @author [Hilton Li]
 * @email [liminghilton@gmail.com]
 * @create date 2019-04-18 14:01:31
 * @modify date 2019-04-18 14:01:31
 * @desc [description]
 */

<template>
    <div  v-b-tooltip.hover title="Tip from title attribute" id="title">
        {{ message }}
    </div>
</template>

<script>
export default{
  el: '#app',
  data: {
    message: 'Hello Vue!'
  }
}
</script>
