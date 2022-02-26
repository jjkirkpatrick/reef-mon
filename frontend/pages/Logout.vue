<template>
    <div id="logout"></div>
</template>

<script>
export default {
  props: {
  },
    created () {
        if (this.$auth.loggedIn) {
            this.$auth.logout()
            this.$router.push('login')
        }else{
            this.$router.push('login')
        }
    },
}
</script>

