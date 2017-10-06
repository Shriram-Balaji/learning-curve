const express = require("express");
const graphql = require("express-graphql");
const port = 5000;

const app = new express();
const schema = require("./schema");
app.use(
  "/graphql",
  graphql({
    schema: schema,
    graphiql: true
  })
);
app.listen(port, function() {
  console.log("Listening at " + port);
});
