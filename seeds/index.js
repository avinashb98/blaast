const quotes = require('./quotes')
const axios = require('axios')

const baseUrl = 'http://localhost:8080'
const createBlaastEndpoint = '/api/v1/blaast'

const sender_id = 'avinash'
const receiver_id = 'bhardwaj'

async function createBlaast(text) {
    return axios.post(baseUrl+createBlaastEndpoint, {
        sender_id,
        receiver_id,
        text
      })
      .then(function (response) {
        console.log(response.data);
      })
      .catch(function (error) {
        console.log(error);
      });
}

async function init() {
    const prs = quotes.map(createBlaast)
    await Promise.all(prs)
}

init()
