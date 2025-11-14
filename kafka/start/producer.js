const { Kafka } = require("kafkajs");

const topic_name = process.argv[2] || "Logs2";
const partition = process.argv[3] || 0;

createProducer();
async function createProducer() {
  try {
    const kafka = new Kafka({
      clientID: "kafka_start",
      brokers: ["192.168.1.2:9092"],
    });

    const producer = kafka.producer();
    console.log("Producer bağlanıyor...");
    await producer.connect();
    console.log("Producer'a bağlantı başarılı...");
    const message_result = await producer.send({
      topic: topic_name,
      messages: [
        {
          value: "Bu bir test Log Mesajıdır...",
          partition: partition,
        },
      ],
    });
    console.log("Gonderim işemi başarılıdır", JSON.stringify(message_result));
    await producer.disconnect();
  } catch (error) {
    console.log("Bir hata oluştu.");
  } finally {
    process.exit(0);
  }
}
