const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
var PROTO_PATH = __dirname + "../../../todo.proto";
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {});
const grpcObject = grpc.loadPackageDefinition(packageDefinition).userproto;

const todoService = new grpcObject.TodoService(
  "localhost:4040",
  grpc.credentials.createInsecure()
);

module.exports = (app) => {
  app.post("/createnewtodo/:id", (req, res, next) => {
    const CustomerId = req.params.id;
    const { Title } = req.body;

    todoService.CreateTodo({ Title, CustomerId }, (err, response) => {
      const data = {
        todos: response.Title,
        CustomerId: response.CustomerId,
      };
      res.json({ data });
    });
  });
  app.post("/addnewtodo/:id", (req, res, next) => {
    const CustomerId = req.params.id;
    const { Title } = req.body;
    todoService.AddNewToto({ Title, CustomerId }, (err, response) => {
      const data = {
        todos: response.Title,
        CustomerId: response.CustomerId,
      };
      res.json({ data });
    });
  });
  app.get("/gettodosbycustomerid/:id", (req, res, next) => {
    const CustomerId = req.params.id;
    todoService.GetTodo({ Id: CustomerId }, (err, response) => {
      console.log(response);
      const data = {
        todos: response.Title,
        CustomerId: response.CustomerId,
      };
      res.json({ data });
    });
  });
  app.post("/deletetodo/:id", (req, res, next) => {
    const CustomerId = req.params.id;
    const { Title } = req.body;
    todoService.DeleteTodo({ Title, CustomerId }, (err, response) => {
      const data = {
        title: response.Title,
        message: response.Message,
      };
      res.json({ data });
    });
  });
  app.get("/deletealltodos/:id", (req, res, next) => {
    const CustomerId = req.params.id;
    todoService.DeleteAllTodo({ Id: CustomerId }, (err, response) => {
      const data = {
        title: response.Title,
        message: response.Message,
      };
      res.json({
        data,
      });
    });
  });
};
