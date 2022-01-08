const Info = require("../models/Info");

exports.getPage = (req, res) => {
  res.status(200).json({
    status: "succes",
    req,
  });
};

exports.getAddPage = (req, res) => {
  res.status(200).json({
    status: "succes",
    req,
  });
};

exports.getEditPage = async (req, res) => {
  const info = await Info.findOne({ _id: req.params.id });
  res.status(200).json({
    status: "succes",
    msg: "edit",
    info,
  });
};

exports.sendEmail = async (req, res) => {
  try {
    const outputMessage = `
    
    <h1>Mail Details </h1>
    <ul>
      <li>Name: ${req.body.name}</li>
      <li>Email: ${req.body.email}</li>
      <li>Email: ${req.body.phone}</li>
    </ul>
    <h1>Message</h1>
    <p>${req.body.message}</p>
    `;

    let transporter = nodemailer.createTransport({
      host: "smtp.gmail.com",
      port: 465,
      secure: true, // true for 465, false for other ports
      auth: {
        user: "arinyazilim@gmail.com", // gmail account
        pass: "bpwrtssmqdsdjdjw", // gmail password
      },
    });

    // send mail with defined transport object
    let info = await transporter.sendMail({
      from: '"Smart EDU Contact Form" <arinyazilim@gmail.com>', // sender address
      to: "gcekic@gmail.com", // list of receivers
      subject: "Smart EDU Contact Form New Message ✔", // Subject line
      html: outputMessage, // html body
    });

    console.log("Message sent: %s", info.messageId);
    // Message sent: <b658f8ca-6296-ccf4-8306-87d57a0b4321@example.com>

    // Preview only available when sending through an Ethereal account
    console.log("Preview URL: %s", nodemailer.getTestMessageUrl(info));
    // Preview URL: https://ethereal.email/message/WaQKMgKddxQDoou...

    req.flash("success", "We Received your message succesfully");

    res.status(200).redirect("contact");
  } catch (err) {
    //req.flash("error", `Something happened! ${err}`);
    req.flash("error", `Something happened!`);
    res.status(200).redirect("contact");
  }
};
