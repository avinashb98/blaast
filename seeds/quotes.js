const quotes = require('./quotes.json').map(q => q.content+' - '+q.author)

module.exports = quotes