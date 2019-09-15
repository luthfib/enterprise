import React, { Component } from 'react';
// import { NavLink } from 'react-router-dom';
import classes from './GroupSidebar.module.css'
import classNames from 'classnames/bind';
import GroupSidebarItem from './GroupSidebarItem/GroupSidebarItem'
import * as actions from '../../store/actions/index'
import { connect } from 'react-redux';

let cx = classNames.bind(classes)

class groupSidebar extends Component {
    
    render() {


        let className = cx({
            GroupSidebar: true,
            GroupSidebarDisplay: true,
            GroupSidebarDisplayNone: !this.props.sidebarItemIsHovered & !this.props.groupSidebarIsHovered
        })

        return (
            <div className={className}
                 onMouseEnter={this.props.toggleGroupSidebarHoverState} 
                 onMouseLeave={this.props.toggleGroupSidebarHoverState}
                // style={{display: "none"}}
                 >
            <section></section>
            <nav className={classes.NavigationLinks}>
                <GroupSidebarItem link="/" exact>Group A</GroupSidebarItem>
                <GroupSidebarItem link="/">Group B</GroupSidebarItem>
                <GroupSidebarItem link="/">Group C</GroupSidebarItem>
                <GroupSidebarItem link="/" >Group D</GroupSidebarItem>
            </nav>
        </div>
        )
    }
}

const mapStateToProps = state => {
    return {
        groupSidebarIsHovered: state.sidebar.groupSidebarIsHovered,
        sidebarItemIsHovered: state.sidebar.sidebarItemIsHovered
    };
}

const mapDispatchToProps = dispatch => {
    return {
        toggleGroupSidebarHoverState: (hoverAction) => dispatch(actions.toggleGroupSidebarHoverState(hoverAction)),
    }
}


export default connect(mapStateToProps, mapDispatchToProps)(groupSidebar);