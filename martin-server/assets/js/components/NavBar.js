import React from 'react';

export default class NavBar extends React.Component {

    constructor(props) {
        super(props);
    }

    signOut() {
        window.location = '/logout';
    }

    render() {
        const menuItems = this.props.menuItems.map((item, index) => {
            return <li key={item} onClick={() =>
                        this.props.changeViewHandler(index)}><a>{item}</a>
                   </li>
        });

        return <div className="top-bar">
            <div className="top-bar-left">
                <ul className="menu">
                    <li key='title' className="menu-text">{this.props.title}</li>
                    {menuItems}
                </ul>
            </div>
            <div className="top-bar-right">
                <ul className="menu"><li onClick={this.signOut}><a>Sign Out</a></li></ul>
            </div>
        </div>
    }
}
