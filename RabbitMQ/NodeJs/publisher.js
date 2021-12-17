const amqp = require("amqplib");

const message = {
  description: "Bu bir test mesajidir..",
};
const data = require("./data.json");
const queueName = process.argv[2] || "jobsQueue";

connetRabbitmq();

async function connetRabbitmq() {
  try {
    const connection = await amqp.connect("amqp://localhost:5672"); // Bağlatı kuruyoruz.
    const channel = await connection.createChannel(); // channel oluşturuyoruz.
    const assertion = await channel.assertQueue(queueName); // hangi job'a ait olduğunu yazıyoruz.

    data.forEach((element) => {
      message.description = element.id;
      channel.sendToQueue(queueName, Buffer.from(JSON.stringify(message))); // mesaj yolluyoryuz.
      console.log("Gönderilen mesaj ", element.id);
    });

    // Interval Value
    /*     setInterval(() => {
      message.description = new Date().getTime();
      channel.sendToQueue(queueName, Buffer.from(JSON.stringify(message))); // mesaj yolluyoryuz.
      console.log("Gönderilen mesaj ", message);
    }, 1000); */
  } catch (error) {
    console.log("Error ", error);
  }
}
