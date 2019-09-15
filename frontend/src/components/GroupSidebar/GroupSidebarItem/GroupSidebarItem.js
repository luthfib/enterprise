import React, { Component } from 'react';
// import { NavLink } from 'react-router-dom';
import classes from './GroupSidebarItem.module.css'

class groupSidebarItem extends Component {
    render() {
        // let style = classNames(classes.SidebarItem, this.props.classOverride, { SidebarItemOverride: false });
        return (
            <a className={ classes.GroupSidebarItem }>
                {this.props.children}
            </a>
        )
    }
}

export default groupSidebarItem;