<template>
  <app-header></app-header>
  <div class="container">
    <h3>Bu benim ilk Vue CLI Uygulamam</h3>
    <p>ilk deneme</p>
  </div>
  <CounterItem></CounterItem>
  <!--
      <CounterItem></CounterItem>
      <CounterItem />
      <counter-item />
      hepsi çalışır
  -->
</template>
<script>
import CounterItem from "@/components/CounterItem";
export default {
  components: { CounterItem },
  component: {
    CounterItem: {
      // "<vermek istediğin isim> : modul ismi"
      CounterItem,
    },
  },
};
</script>