import React from 'react';
import SidebarItem from './SidebarItem/SidebarItem'
import classes from './Sidebar.module.css'
import override from './SidebarItem/SidebarOverride.module.css'
import SidebarItemCustom from './SidebarItem/SidebarItemCustom'

const sidebar = () => (
    <div className={classes.Sidebar}>
        <section></section>
        <nav className={classes.NavigationLinks}>
            <SidebarItem link="/" exact>Projects</SidebarItem>
            <SidebarItemCustom>Tasks</SidebarItemCustom>
            <SidebarItem link="/">Calendar</SidebarItem>
            <SidebarItem link="/" classOverride={override.SidebarItemOverride}>Timeline</SidebarItem>
        </nav>
        
    </div>
)

export default sidebar;