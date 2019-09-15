import * as actionTypes from './actionTypes';

export const toggleSidebarItemHoverState = ( name ) => {
    return {
        type: actionTypes.TOGGLE_SIDEBAR_ITEM_HOVER_STATE,
        sidebarItemIsHovered: name
    };
};

export const toggleGroupSidebarHoverState = ( name ) => {
    return {
        type: actionTypes.TOGGLE_GROUP_SIDEBAR_HOVER_STATE,
        groupSidebarIsHovered: name
    };
};

export const toggleGroupSidebarDisplayStatus= ( name ) => {
    return {
        type: actionTypes.TOGGLE_GROUP_SIDEBAR_DISPLAY_STATUS,
        groupSidebarIsShown: name
    };
};