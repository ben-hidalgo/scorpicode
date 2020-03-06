import superagentPromise from 'superagent-promise'
import _superagent from 'superagent'
// import commonStore from './stores/commonStore';
// import authStore from './stores/authStore';


const superagent = superagentPromise(_superagent, global.Promise)

// const encode = encodeURIComponent

const handleErrors = err => {
  if (err && err.response && err.response.status === 401) {
    // authStore.logout()
  }
  return err
}

const responseBody = res => res.body

const tokenPlugin = req => {
  // if (commonStore.token) {
  //   req.set('authorization', `Token ${commonStore.token}`)
  // }
}

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
  makeHat: (inches, color, style) => requests.post('/hats/MakeHat', {inches: inches, color: color, style: style}),
}


export default {
  Hats,
}
