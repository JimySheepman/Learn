const app = Vue.createApp({
  data() {
    return {
      title: "Vue.js Bootcamp day #1",
      content: "Loren ipsum dolar sit amet..",
      eduflow: {
        title: "website",
        target: "_blank",
        url: "http://merlins.rf.gd",
        alt: "website",
      },
      coords: {
        x: 0,
        y: 0,
      },
    };
  },
  methods: {
    changeTitle(pTitle) {
      this.title = pTitle;
    },
    updateCoords(message, event) {
      console.log(message, event.x, event.y);
      this.changeTitle(`${event.x},${event.y}`)
      this.coords = {
        x: event.x,
        y: event.y,
      };
    },
  },
}).mount("#app");
