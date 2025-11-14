const { Kafka } = require("kafkajs");

createProducer();

async function createProducer() {
  try {
    const kafka = new Kafka({
      clientId: "kafka_pub_sub_client",
      brokers: ["192.168.1.2:9092"],
    });

    const producer = kafka.producer();
    console.log("Producer'a bağlanılıyor..");
    await producer.connect();
    console.log("Bağlantı başarılı.");

    const message_result = await producer.send({
      topic: "raw_videos_topic",
      messages: [
        {
          value: "Yeni video içeriği üretildi...",
          partition: 0,
        },
      ],
    });
    console.log("Gonderim işlemi başarılıdır", JSON.stringify(message_result));
    await producer.disconnect();
  } catch (error) {
    console.log("Bir Hata Oluştu", error);
  } finally {
    process.exit(0);
  }
}
