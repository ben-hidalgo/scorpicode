import { extendObservable } from 'mobx';

class HatStore {

  constructor() {
    extendObservable(this, {
      counter: 0,
      hats: [
        {id: "1", inches: 7, color: "RED", style: "DERBY", version: 0},
        {id: "2", inches: 7, color: "RED", style: "DERBY", version: 0}
      ],
    })
  }

}

export default new HatStore();
