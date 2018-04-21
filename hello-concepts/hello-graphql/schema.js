/* Reasons behind using the Good Reads API
  1. Requires XML Parsing
  2. Messy API, returns more excessive information than needed - a perfect usecase for GraphQL
*/
const request = require("request-promise");

//Used for promisifying a function that generally takes callbacks
const pify = require("pify");
const parseXML = pify(require("xml2js").parseString);
const config = require("config");

const {
  GraphQl,
  GraphQLSchema,
  GraphQLInt,
  GraphQLObjectType,
  GraphQLString
} = require("graphql");

const AuthorType = new GraphQLObjectType({
  name: "Author",
  description: "...",
  fields: () => ({
    name: {
      type: GraphQLString,
      resolve: xml => xml.GoodreadsResponse.author[0].name[0]
    }
  })
});
module.exports = new GraphQLSchema({
  query: new GraphQLObjectType({
    name: "Query",
    description: "...",
    fields: () => ({
      author: {
        type: AuthorType,
        args: {
          id: {
            type: GraphQLInt
          }
        },
        resolve: (root, args) =>
          request(
            `https://www.goodreads.com/author/show.xml?id=${args.id}&key=${config.get(
              "goodReads.key"
            )}`
          ).then(parseXML)
      }
    })
  })
});
