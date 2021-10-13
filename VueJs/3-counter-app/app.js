const app = Vue.createApp({
    data() {
        return {
            counter:0,
            counter2:0,
        }
    },
    methods: {
        getCounterResult(){
            console.log("Counter 1 Çalıştı");
           return this.counter > 5 ? 'BÜYÜK': 'KÜÇÜK';
        },
        getCounter2Result(){
            console.log("Counter 2 Çalıştı");
            return this.counter2 > 5 ? 'BÜYÜK': 'KÜÇÜK';
         }
    },
/*     methods: {
        inc(){
            this.counter++;
        },
        dec(){
            this.counter--;
        }
    },
 */
}).mount("#app")