import superagentPromise from 'superagent-promise'
import _superagent from 'superagent'
import authStore from './stores/authStore';


const superagent = superagentPromise(_superagent, global.Promise)

const handleErrors = err => {
  if (err && err.response && err.response.status === 401) {
    authStore.logout()
  }
  return err
}

const responseBody = res => res.body

const tokenPlugin = req => {
  if (authStore.token) {
    req.set('authorization', `Bearer ${authStore.token}`)
  }
}

// all requests should be POST due to Twirp
const requests = {
  del: url =>
    superagent
      .del(`${url}`)
      .use(tokenPlugin)
      .end(handleErrors)
      .then(responseBody),
  get: url =>
    superagent
      .get(`${url}`)
      .use(tokenPlugin)
      .end(handleErrors)
      .then(responseBody),
  put: (url, body) =>
    superagent
      .put(`${url}`, body)
      .use(tokenPlugin)
      .end(handleErrors)
      .then(responseBody),
  post: (url, body) =>
    superagent
      .post(`${url}`, body)
      .use(tokenPlugin)
      .end(handleErrors)
      .then(responseBody),
}

const Hats = {
  listHats: () => requests.post('/hats/ListHats', {}),
  // TODO: rename makeHats()
  makeHat: (body) => requests.post('/hats/MakeHats', body),
  deleteHat: (id, version) => requests.post('/hats/DeleteHat', {id: id, version: version}),
  fetchHat: (id) => requests.post('/hats/FetchHat', {id: id}),
}


export default {
  Hats,
}
