const app = Vue.createApp({
    data() {
        return {
            fullName: null,
        };
    },
  methods: {
    updateValue(event) {
      this.fullName = event.target.value
    },
  },
}).mount("#app");
