import React from 'react';

const trapTypes = ['hole', 'spikes',];

const Trap = () => {
  const trapTypeIdx = Math.floor(Math.random() * trapTypes.length);

  return (
    <img className='trap'
      src={`images/traps/trap-${trapTypes[trapTypeIdx]}.png`}
      alt={trapTypes[trapTypeIdx]}
    />
  )
};

export default Trap;