import React from 'react';
import Header from '../components/Header/Header';
import Body from '../components/Body/Body';
const HomePage = (props) => {
// const {username} = props;


    return (
        <div>
        <Header />
        <Body />
        <div>
        <h2> Hello {props.username}</h2>

        </div>
        </div>
    );
};

export default HomePage;