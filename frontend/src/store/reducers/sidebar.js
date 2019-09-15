import { updateObject } from '../utility';
import * as actionTypes from '../actions/actionTypes';

const initialState = {
    sidebarItemIsHovered: false,
    groupSidebarIsHovered: false,
    groupSidebarIsShown: false
};

const toggleSidebarItemHoverState = ( state, action ) => {
    return updateObject( state, { sidebarItemIsHovered: !state.sidebarItemIsHovered } );
};

const toggleGroupSidebarHoverState = ( state, action ) => {
    return updateObject( state, { groupSidebarIsHovered: !state.groupSidebarIsHovered } );
};

const toggleGroupSidebarDisplayStatus = ( state, action ) => {
    return updateObject( state, { groupSidebarIsShown: !state.groupSidebarIsShown } );
};


const reducer = ( state = initialState, action ) => {
    switch ( action.type ) {
        case actionTypes.TOGGLE_SIDEBAR_ITEM_HOVER_STATE: return toggleSidebarItemHoverState( state, action );
        case actionTypes.TOGGLE_GROUP_SIDEBAR_HOVER_STATE: return toggleGroupSidebarHoverState( state, action );
        case actionTypes.TOGGLE_GROUP_SIDEBAR_DISPLAY_STATUS: return toggleGroupSidebarDisplayStatus( state, action );
        default: return state;
    }
};

export default reducer;