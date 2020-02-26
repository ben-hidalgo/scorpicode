import Hat from './Hat';
import React from 'react';


const HatList = props => {

  if (props.hats.length === 0 || !props.hats ) {
    return (
      <div className="article-preview">
        No hats
      </div>
    );
  }

  let hats = []
  let hat = props.hats.next()
  // console.log(hat)
  while (!hat.done) {
    //hats.push(props.hats.next())
    hat = props.hats.next()
    console.log(hat)
    //hats.push(<Hat hat={hat.value.target} key={hat.value.target.id} />)
  }

  return (
    <div>
      <ul>Hats:</ul>
      {hats}
    </div>
  );
};

export default HatList;
