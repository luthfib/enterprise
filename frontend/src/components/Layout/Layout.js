import React, { Component } from 'react';
import Sidebar from '../Sidebar/Sidebar'
import Header from '../Header/Header'
import classes from './Layout.module.css'
import Main from '../Main/Main'
import GroupSidebar from '../GroupSidebar/GroupSidebar'

class Layout extends Component {

    render() {
        return (
            <div className={classes.Layout}>
                
                <div>
                </div>
                <Sidebar />
                <GroupSidebar/>
                <Main>
                </Main>
                <Header></Header>
            </div>
        )
    }
}

export default Layout;