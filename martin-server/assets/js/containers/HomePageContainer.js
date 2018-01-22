import React from 'react';
import NavBar from 'components/NavBar';
import SettingsView from 'components/views/SettingsView';
import DashboardView from 'components/views/DashboardView';
import DevicesView from 'components/views/DevicesView';
import NotFoundView from 'components/views/NotFoundView';

const VIEWS = {
    dashboard: {
        component: DashboardView,
        display: 'Dashboard',
        path: '/dashboard'
    },
    devices: {
        component: DevicesView,
        display: 'Devices',
        path: '/devices'
    },
    settings: {
        component: SettingsView,
        display: 'Settings',
        path: '/settings'
    },
    notfound: {
        component: NotFoundView,
        display: 'Not Found',
        path: '/notfound'
    }
};

const MENU_ITEMS = [VIEWS.dashboard, VIEWS.devices, VIEWS.settings].map((view) => {
    return view.display;
});

export default class HomePageContainer extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            currentView: this.getViewForPath()
        };

        this.navbarItemHandler = this.navbarItemHandler.bind(this);
    }

    getViewForPath() {
        var path = document.location.pathname;
        if (path == '/') {
            return VIEWS.dashboard;
        }

        for (var view in VIEWS) {
            if (VIEWS[view].path == path) {
                return VIEWS[view];
            }
        }

        return VIEWS.notfound;
    }

    changeView(view) {
        window.history.pushState("", "MARTIN - " + view.display, view.path);
        document.title = 'MARTIN - ' + view.display;
        this.setState({currentView: view});
    }

    changeViewByName(viewName) {
        if (VIEWS[viewName]){
            this.changeView(VIEWS[viewName]);
        }
    }

    navbarItemHandler(menuItemInd) {
        switch (menuItemInd) {
            case 0:
                this.changeView(VIEWS.dashboard);
                break;
            case 1:
                this.changeView(VIEWS.devices);
                break;
            case 2:
                this.changeView(VIEWS.settings);
                break;
        };
    }

    render() {
        const ViewComponent = this.state.currentView.component;
        return <div>
            <NavBar title='MARTIN' changeViewHandler={this.navbarItemHandler}
                    menuItems={MENU_ITEMS} />
            <div className='grid-container view-container'>
                <ViewComponent changeViewHandler={this.changeViewByName}/>
            </div>
        </div>
    }
}
