
const setCookie = (name, value, days) => {
  var d = new Date();
  d.setTime(d.getTime() + 24*60*60*1000*days);
  // not using string interpolation because we're inside a Go string literal
  document.cookie = name + "=" + value + ";path=/;expires=" + d.toGMTString();
}

const getCookie = (name) => {
  var v = document.cookie.match('(^|;) ?' + name + '=([^;]*)(;|$)')
  return v ? v[2] : null
}

const deleteCookie = (name) => { 
  setCookie(name, '', -1)
}


export {
  setCookie,
  getCookie,
  deleteCookie,
}
