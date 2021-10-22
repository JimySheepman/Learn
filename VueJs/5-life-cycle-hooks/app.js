const app = Vue.createApp({
  data() {
    return {
      title: "Test Başlığı",
      itemList: [],
    };
  },
  // Life cycle
  beforeCreate() {
    console.log("beforeCreate");
  },
  created() {
    console.log("created");
    setTimeout(() => {
      this.itemList = [1, 2, 3, 4, 5, 6, 7, 8, 9];
    }, 2000);

    setTimeout(() => {
      this.title = "Bu yeni Başlık";
    }, 3000);
  },
  beforeMount() {
    console.log("beforeMount");
  },
  mounted() {
    console.log("mounted");
  },
  beforeUpdate() {
    console.log("beforeUpdate");
  },
  updated() {
    console.log("updated");
  },
  beforeUnmount() {
    console.log("beforeUnmount");
  },
  unmounted() {
    console.log("unmounted");
  },
}).mount("#app");

setTimeout(() => {
  app.unmount();
}, 5000);
