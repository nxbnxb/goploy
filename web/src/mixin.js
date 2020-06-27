// mixin.js
const mixin = {
  methods: {
    hasAdminPermission() {
      return [this.$global.Admin].indexOf(this.$store.getters.role) !== -1
    },
    hasManagerPermission() {
      return [this.$global.Admin, this.$global.Manager].indexOf(this.$store.getters.role) !== -1
    },
    hasGroupManagerPermission() {
      return [this.$global.Admin, this.$global.Manager, this.$global.GroupManager].indexOf(this.$store.getters.role) !== -1
    }
  }
}
export default mixin
