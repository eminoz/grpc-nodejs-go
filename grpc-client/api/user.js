const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
var PROTO_PATH = __dirname + "../../../todo.proto";
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {});
const grpcObject = grpc.loadPackageDefinition(packageDefinition).userproto;

const customerService = new grpcObject.CustomerService(
  "localhost:4040",
  grpc.credentials.createInsecure()
);
module.exports = (app) => {
  app.post("/createuser", (req, res, next) => {
    const { name, email, password } = req.body;
    customerService.CreateUser(
      {
        name,
        email,
        password,
      },
      async (err, response) => {
        const data = {
          name: response.name,
          email: response.email,
          todo: [],
        };
        res.json({ data: data });
      }
    );
  });
  app.get("/getuserbyid/:id", (req, res, next) => {
    const Id = req.params.id;

    customerService.GetUser({ Id }, (err, response) => {
      const data = {
        name: response.name,
        email: response.email,
        todo: [],
      };
      res.json({ data: data });
    });
  });
  app.delete("/deleteuserbyid/:id", (req, res, next) => {
    const Id = req.params.id;
    customerService.DeleteUSer({ Id }, (err, response) => {
      const data = {
        email: response.email,
        id: response.id,
        message: response.message,
      };
      res.json({ data: data });
    });
  });
};
