const express = require("express");
const app = express();
const expressApp = require("./express-app");
const startApp = async () => {
  app.use(express.json());

  await expressApp(app);
  app.listen(3000, () => {
    console.log(`client is running on 3000`);
  });
};
startApp();
