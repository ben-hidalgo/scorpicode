// import { observable, action, reaction } from 'mobx';
// import { observable } from 'mobx';
// import agent from '../agent';
import { extendObservable } from 'mobx';

class CommonStore {

  constructor() {
    extendObservable(this, {
      counter: 0,
    })
  }

}

export default new CommonStore();
