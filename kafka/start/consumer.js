const { Kafka } = require("kafkajs");

const topic_name = process.argv[2] || "Logs2";

createConsumer();
async function createConsumer() {
  try {
    const kafka = new Kafka({
      clientID: "kafka_start",
      brokers: ["192.168.1.2:9092"],
    });

    const consumer = kafka.consumer({
      groupId: "consumer_group_start",
    });
    console.log("Consumer bağlanıyor...");
    await consumer.connect();
    console.log("Consumer'a bağlantı başarılı...");
    // consumer Subscribe...
    await consumer.subscribe({
      topic: topic_name,
      fromBeginning: true,
    });

    await consumer.run({
      eachMessage: async (result) => {
        console.log(
          `Gelen Mesaj ${result.message.value}, Par=> ${result.partition}`
        );
      },
    });
  } catch (error) {
    console.log("Bir hata oluştu.");
  }
}
