import React, { Component } from 'react';
// import { NavLink } from 'react-router-dom';
import classes from './SidebarItem.module.css'
import classNames from 'classnames/bind';

class sidebarItem extends Component {
    render() {
        let style = classNames(classes.SidebarItem, this.props.classOverride, { SidebarItemOverride: false });
        return (
            <a className={ style }>
                {this.props.children}
            </a>
        )
    }
}

export default sidebarItem;