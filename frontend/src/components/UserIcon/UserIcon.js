import React from 'react';
import classes from './UserIcon.module.css'

const userIcon = ( props ) => (
    <div className={classes.UserIconDiv}>
        <span className={classes.UserIcon}>{props.children}</span>
    </div>
)

export default userIcon;