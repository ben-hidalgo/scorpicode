import { observable, action, computed } from 'mobx';
import agent from '../agent';

export class HatStore {

  @observable hatRegistry = observable.map();
  @observable isLoading = false;

  @computed get hats() {
    return this.hatRegistry.values();
  };

  getHat(name) {
    return this.hatRegistry.get(name);
  }

  @action makeHat(inches) {
    return agent.Hats.makeHat(inches)
      .catch(action(err => {
        throw err;
      }));
  }


  $req() {
    return agent.Hats.listHats();
  }

  @action loadHats() {
    this.isLoading = true;
    return this.$req()
      .then(action(({ hats }) => {
        this.hatRegistry.clear();
        hats.forEach(hat => this.hatRegistry.set(hat.name, hat));
      }))
      .finally(action(() => { this.isLoading = false; }));
  }


}

export default new HatStore();
