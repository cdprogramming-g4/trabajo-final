import React, { useEffect, useState } from 'react';

const wallTypes = {
    CORNER_LEFT_UP: 'corner-UL',
    CORNER_LEFT_DOWN: 'corner-DL',
    CORNER_RIGHT_UP: 'corner-UR',
    CORNER_RIGHT_DOWN: 'corner-DR',
    SIDE_UP: 'side-H',
    SIDE_DOWN: 'side-H',
    SIDE_LEFT: 'side-V',
    SIDE_RIGHT: 'side-V',
    MIDDLE: 'middle',
};

// w = length[0]-1, h = length-1
const Wall = ({w, h, x=0, y=0}) => {
    const [wallType, setWallType] = useState(wallTypes.MIDDLE);

    useEffect(()=>{
        if (y === 0) {
            if (x === 0) {
                setWallType(wallTypes.CORNER_LEFT_UP);
            }
            else if (x === w) {
                setWallType(wallTypes.CORNER_RIGHT_UP);
            }
            else {
                setWallType(wallTypes.SIDE_UP);
            }
        }
        else if (y === h) {
            if (x === 0) {
                setWallType(wallTypes.CORNER_LEFT_DOWN);
            }
            else if (x === w) {
                setWallType(wallTypes.CORNER_RIGHT_DOWN);
            }
            else {
                setWallType(wallTypes.SIDE_DOWN);
            }
        }
        else {
            if (x === 0) {
                setWallType(wallTypes.SIDE_LEFT);
            }
            else if (x === w) {
                setWallType(wallTypes.SIDE_RIGHT);
            }
            else {
                setWallType(wallTypes.MIDDLE);
            }
        }
    }, [w, h, x, y]);

    return (
        <img className='wall'
            src={`images/board/wall-${wallType}.png`}
            alt='wall'
        />
    )
};

export default Wall;