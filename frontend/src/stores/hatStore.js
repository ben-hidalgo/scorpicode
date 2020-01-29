import { observable, action, computed } from 'mobx';

export class HatStore {

  @observable hatRegistry = observable.map();

  @computed get hats() {
    return this.hatRegistry.values();
  };

  @action createHat(hat) {
    this.hatRegistry.set(hat.id, hat);
  }

}

export default new HatStore();
