import React, { Component } from 'react';
import { connect } from 'react-redux';
// import { NavLink } from 'react-router-dom';
import classes from './SidebarItem.module.css'
import classNames from 'classnames/bind';
import * as actions from '../../../store/actions/index'

class sidebarItemCustom extends Component {
    state = {
        isMouseInside: false
    }

    render() {
        let style = classNames(classes.SidebarItem, this.props.classOverride, { SidebarItemOverride: false });
        return (
            <a onMouseEnter={this.props.toggleSidebarItemHoverState}
               onMouseLeave={this.props.toggleSidebarItemHoverState}
               className={ style }>
              {this.props.children}
            </a>
        )
    }
}

const mapStateToProps = state => {
    return {
        sidebarItemIsHovered: state.sidebar.sidebarItemIsHovered
    };
}

const mapDispatchToProps = dispatch => {
    return {
        toggleSidebarItemHoverState: (hoverAction) => dispatch(actions.toggleSidebarItemHoverState(hoverAction)),
    }
}


export default connect(mapStateToProps, mapDispatchToProps)(sidebarItemCustom);