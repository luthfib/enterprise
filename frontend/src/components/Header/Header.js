import React from 'react';
import classes from './Header.module.css'
import UserIcon from '../UserIcon/UserIcon'
import WaffleIcon from '../WaffleIcon/WaffleIcon'

const header = () => (
    <header className={classes.Header}>
        <WaffleIcon></WaffleIcon>
        <UserIcon>L</UserIcon>
    </header>
)

export default header;