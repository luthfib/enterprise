import React from 'react';
import classes from './GridView.module.css'

const gridIcon = () => (
    
    <a className={classes.GridIcon + ' ' + classes.GridIconLine2}>
		<span className={classes.Layer + ' ' + classes.LayerPrimary}>
			<span></span><span></span>
		</span>
		<span className={classes.Layer + ' ' + classes.LayerSecondary}>
			<span></span><span></span>
		</span>
	</a>
)

export default gridIcon;