const amqp = require("amqplib");
const queueName = process.argv[2] || "jobsQueue";
const data = require("./data.json");
connetRabbitmq();

async function connetRabbitmq() {
  try {
    const connection = await amqp.connect("amqp://localhost:5672"); // Bağlatı kuruyoruz.
    const channel = await connection.createChannel(); // channel oluşturuyoruz.
    const assertion = await channel.assertQueue(queueName); // hangi job'a ait olduğunu yazıyoruz.
    // mesajın alınması..
    console.log("Mesaj bekleniyor..");
    channel.consume(queueName, (message) => {
      const messageIngo = JSON.parse(message.content.toString());
      const userInfo = data.find((u) => u.id == message.description);
      if (userInfo) {
        console.log("İşenen kayıt", userInfo);
        channel.ack(message); // okunduğunu belirtiyor.
      }
    });
  } catch (error) {
    console.log("Error ", error);
  }
}
