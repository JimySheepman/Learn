const { Kafka } = require("kafkajs");

createTopic();
async function createTopic() {
  try {
    // Admin Stuff..

    const kafka = new Kafka({
      clientID: "kafka_start",
      brokers: ["192.168.1.2:9092"],
    });

    const admin = kafka.admin();
    console.log("Kafka Broker'a bağlanıyor...");
    await admin.connect();
    console.log("Kafka Broker'a bağlantı başarılı...");
    await admin.createTopics({
      topics: [
        {
          topic: "Logs",
          numPartitions: 1,
        },
        {
          topic: "Logs2",
          numPartitions: 2,
        },
      ],
    });
    console.log("topic başaralı bir şekilde oluturulmuştur...");
    await admin.disconnect();
  } catch (error) {
    console.log("Bir hata oluştu.");
  } finally {
    process.exit(0);
  }
}
