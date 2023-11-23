import React from 'react';

const creatureTypes = ['fenix', 'kitsune'];

const Creature = () => {
  const creatureTypeIdx = Math.floor(Math.random() * creatureTypes.length);

  return (
    <img className='creature'
      src={`images/creatures/creature-${creatureTypes[creatureTypeIdx]}.png`}
      alt={creatureTypes[creatureTypeIdx]}
    />
  )
};

export default Creature;