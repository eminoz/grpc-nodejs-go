const { user, todo } = require("./api");
const cors = require("cors");
module.exports = (app) => {
  app.use(cors());
  user(app);
  todo(app)

};
