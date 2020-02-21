import Hat from './Hat';
import React from 'react';

const HatList = props => {

  if (props.hats.length === 0) {
    return (
      <div className="article-preview">
        No hats
      </div>
    );
  }

  return (
    <div>
      <ul>Hats:</ul>
      {
        props.hats.map(hat => {
          return (
            <Hat hat={hat} key={hat.id} />
          );
        })
      }
    </div>
  );
};

export default HatList;
